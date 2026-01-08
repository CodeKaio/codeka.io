---
date: "2025-01-17"
lastmod: "2025-02-12"
language: fr
tags:
- devops
- terraform
- vault
- clevercloud
title: Vault sur Clever Cloud
series: Terraform et Clever Cloud
---

Pour les besoins des cours que je donne à l'Université de Lille, j'ai dû configurer un serveur Vault sur Clever Cloud.

Et [bien entendu](/books/iac-avec-terraform), j'ai fait tout ça avec Terraform.

Cet article décrit comment utiliser le provider Terraform de Clever Cloud pour déployer un serveur Vault. Un article suivant décrira comment le configurer pour l'authentification OIDC avec GitLab et y stocker quelques secrets à titre d'exemple.

Le code de cet article est aussi disponible sur GitHub : https://github.com/juwit/terraform-clevercloud-playground/tree/main/vault.

> Cet article a été écrit avec des commandes Terraform, mais fonctionne également avec les commandes OpenTofu équivalentes.

## Architecture cible

Avant d'entrer dans la mise en pratique, il convient ici d'expliquer quelques choix illustrés par le schéma suivant.

![](cc_vault.png)

Clever Cloud propose de déployer des applications dans de [nombreux langages](https://www.clever-cloud.com/developers/doc/applications/). Pour héberger une instance Vault, le plus simple semblait d'utiliser une instance Docker.

Par défault, Vault propose l'utilisation du _backend_ de stockage _Integrated storage_ pour le stockage des données. Étant donné la nature du déploiement avec une instance Docker sur un seul nœud et le fait que Clever Cloud ne supporte pas le stockage persistant pour ce type d'instance, il m'a semblé judicieux d'utiliser un _backend_ de stockage externalisé. Parmi les options proposées par Vault, 3 options sont envisageables sur Clever Cloud : les bases de données MySQL ou PostgreSQL, ou S3 _via_ l'implémentation _Cellar_ fournie par Clever Cloud.

Le stockage externalisé sur S3 ne supporte pas la haute disponibilité et pourrait s'avérer incompatible avec l'implémentation _Cellar_ (cf. [les adaptations requises par le _backend_ Terraform S3 pour _Cellar_](/posts/2024-12-31-terraform-clever-cloud#configuration-du-backend)), donc je l'ai directement écarté et j'ai privilégié l'implémentation avec PostgreSQL.

L'authentification _via_ GitLab permet à mes étudiants d'utiliser leur compte GitLab existant, en exploitant l'instance GitLab fournie par l'Université de Lille. C'est donc très pratique pour eux (pas besoin d'avoir un compte ailleurs) et pour moi (pas besoin de créer et de fournir des comptes).
J'aurais aussi pu utiliser une instance KeyCloak pour implémenter l'authentification, mais cela aurait complexifié inutilement l'implémentation.

À noter aussi que je ne suis pas expert Vault, donc je ne suis pas à l'abri d'avoir fait une erreur de configuration quelque part, alors attention si vous utilisez cette configuration en production 🙂

## SetUp de Terraform

Dans un article précédent, j'ai déjà expliqué comment configurer [Terraform pour Clever Cloud](/posts/2024-12-31-terraform-clever-cloud#configurer-le-provider-clever-cloud), ainsi que comment [configurer un _backend_ _via_ un _bucket_ _Cellar_](/posts/2024-12-31-terraform-clever-cloud#configuration-du-backend).
Ces étapes ne sont pas décrites ici pour ne pas alourdir cet article, mais sont bien nécessaires.

## Création de la base de données avec Terraform

La première étape consiste à créer une base de données consacrée à Vault.
Avec Terraform, la création de la base de données se fait avec le code suivant :

```terraform
resource "clevercloud_postgresql" "vault_storage" {
  name   = "vault_storage"
  plan   = "dev"
  region = "par"
}
```

Vault nécessite que le schéma de la base de données soit initialisé avant que l'application ne soit démarrée.
Le schéma est fourni dans la [documentation](https://developer.hashicorp.com/vault/docs/configuration/storage/postgresql) du _backend_ :

```sql
CREATE TABLE vault_kv_store (
  parent_path TEXT COLLATE "C" NOT NULL,
  path        TEXT COLLATE "C",
  key         TEXT COLLATE "C",
  value       BYTEA,
  CONSTRAINT pkey PRIMARY KEY (path, key)
);

CREATE INDEX parent_path_idx ON vault_kv_store (parent_path);
```

Ce script peut être passé à la main _via_ `psql`, ou dans la console Clever Cloud.

Il est aussi possible d'utiliser un _provisioner_ Terraform pour exécuter le script après la création de la base de données :

```terraform
resource "clevercloud_postgresql" "vault_storage" {
  name = "vault_storage"
  plan = "dev"
  region = "par"

  provisioner "local-exec" {
    # wait for the database to be up
    command = "sleep 10 && psql -f vault-schema.sql"
    environment = {
      PGHOST = self.host
      PGPORT = self.port
      PGDATABASE = self.database
      PGUSER = self.user
      PGPASSWORD = self.password
    }
  }
}
```
Ici, le _provisioner_ `local-exec` est utilisé pour exécuter la commande `psql` après avoir attendu quelques secondes, le temps que la base de données soit effectivement créée.
Les variables d'environnement nécessaires à l'exécution de `psql` sont également positionnées.

> Je ne suis pas un grand fan de l'exécution de _provisioners_, car ils impliquent une dépendance avec la machine qui exécute Terraform. Ici, c'est le binaire `psql` et la commande `sleep` dans le script _shell_ qui sont nécessaires.

## Création de l'instance Vault avec Terraform

Une fois la base de données créée et le schéma initialisé, on peut créer l'instance Docker pour notre Vault sur Clever Cloud avec le code suivant :

```terraform
resource "clevercloud_docker" "vault_instance" {
  name = "vault_instance"

  # vertical auto-scaling disabled
  smallest_flavor = "XS"
  biggest_flavor = "XS"

  # horizontal auto-scaling disabled
  min_instance_count = 1
  max_instance_count = 1

  # network setup
  additional_vhosts = ["vault-instance.cleverapps.io"]
  redirect_https = true

  # URL for the storage backend
  environment = {
    VAULT_LOCAL_CONFIG = jsonencode(
      {
        "storage": {
          "postgresql": {
            "connection_url": "postgres://${clevercloud_postgresql.vault_storage.user}:${clevercloud_postgresql.vault_storage.password}@${clevercloud_postgresql.vault_storage.host}:${clevercloud_postgresql.vault_storage.port}/${clevercloud_postgresql.vault_storage.database}"
          }
        },
        "listener": [{ "tcp": { "address": "0.0.0.0:8080", "tls_disable": true } }],
        "disable_mlock": true,
        "ui": true
      })
  }
}
```

Parmi les paramètres de configuration intéressants, on retrouve les paramètres principaux de la ressource `clevercloud_docker`, avec les paramètres de scalabilité horizontale et verticale, ainsi que la déclaration d'un nom de domaine customisé.

Les variables d'environnement permettent de passer sa configuration à Vault (plutôt que d'utiliser un fichier).
C'est un des aspects bien pratique de l'image Docker de Vault (documenté sur [dockerhub](https://hub.docker.com/r/hashicorp/vault)).
Ici, on utilise la variable `VAULT_LOCAL_CONFIG`, dans laquelle on donne du contenu formaté en _JSON_, à l'aide de la fonction Terraform `jsonencode()`.

Concernant la configuration de Vault, le stockage sur l'instance PostgreSQL est défini à travers le paramètre `"storage" : { "postgresql" : {} }`. L'URL de connexion est passée en paramètre, elle est reconstruite à partir des attributs de la ressource `clevercloud_postgresql.vault_storage`. Le paramètre `listener` permet de forcer Vault à écouter sur le port `8080`, à la place du port par défaut `8200`, qui est le port d'écoute attendu par Clever Cloud. L'utilisation de l'adresse `0.0.0.0` permet aussi d'écouter sur les connexions provenant d'internet (à la place de l'adresse localhost `127.0.0.1` par défaut). C'est aussi Clever Cloud qui va s'occuper de l'exposition d'un certificat pour l'accès en HTTPS à l'instance, on désactive donc le TLS avec l'option `tls_disable`.
Enfin, on désactive le _lock_ de mémoire en RAM avec `disable_mlock`, car l'exécution de _containers_ Docker sur Clever Cloud ne permet pas, à ma connaissance, l'utilisation de la _capability_ Linux `IPC_LOCK`. Cette [capability](https://www.man7.org/linux/man-pages/man7/capabilities.7.html) de Linux permet de donner les droits à un processus de verrouiller sa mémoire en RAM pour éviter que la mémoire soit écrite sur le _swap_. Le paramètre `ui` permet d'activer la console graphique de Vault, qui sera bien pratique pour les étapes suivantes.

Une fois l'application Docker créée, on peut récupérer son identifiant Clever Cloud avec un _output_ Terraform :

```terraform
output "vault_instance_id" {
  description = "Clever Cloud id for the instance. Use with `clever link` before deploying."
  value = clevercloud_docker.vault_instance.id
}
```

```shell
$ terraform output -raw vault_instance_id        

app_72d4b5a4-1ab8-4653-a825-9be0c62e0fa1
```

Cet _output_ permettra d'exécuter les commande `clever link` et `clever deploy` pour déployer l'instance Vault à l'étape suivante.

## Déploiement de Vault

Le déploiement d'une application Docker sur Clever Cloud passe par l'écriture d'un `Dockerfile` et l'exécution de la commande `clever deploy`.

Le contenu du fichier `Dockerfile` est simpliste :

```dockerfile
FROM hashicorp/vault:1.18

CMD ["server"]
```

On part d'une version fixée de Vault, (la version 1.18 étant la plus récente à l'heure de l'écriture de ces lignes), et on surcharge la commande exécutée par Vault au démarrage de l'application avec la directive `CMD ["server"]`.
Par défaut, Vault démarre en mode « développement », avec la commande `CMD ["server", "-dev"]`. Si vous souhaitez simplifier vos tests, vous pouvez conserver cette directive, mais elle est déconseillée pour de la production. Je l'ai donc désactivée dans cet article.

Après avoir créé un _repository_ Git pour notre fichier et commité celui-ci, le déploiement se fait en 2 commandes, `clever link` pour associer le _repository_ Git courant à l'instance Clever Cloud Docker, puis `clever deploy` pour soumettre le code source à Clever Cloud :

```shell
$ clever link app_72d4b5a4-1ab8-4653-a825-9be0c62e0fa1

Your application has been successfully linked!

$ clever deploy

Remote application is app_id=app_72d4b5a4-1ab8-4653-a825-9be0c62e0fa1, alias=vault_instance, name=vault_instance
Remote application belongs to orga_0331b635-5a61-4786-8f2f-dee81a1b8970
App is brand new, no commits on remote yet
New local commit to push is c6eb36c12ee5ca4a6f0cbcaa2683310856ef7f42 (from refs/heads/main)
Pushing source code to Clever Cloud
Your source code has been pushed to Clever Cloud.
Waiting for deployment to start
Deployment started (deployment_f5deb5ec-e9af-4f19-a5c2-978356632954)
Waiting for application logs
Couldn't start vault with IPC_LOCK. Disabling IPC_LOCK, please use --cap-add IPC_LOCK
==> Vault server configuration:
Administrative Namespace:
                     Cgo: disabled
   Environment Variables: APP_HOME, APP_ID, CC_APP_ID, CC_APP_NAME, CC_COMMIT_ID, CC_DEPLOYMENT_ID, CC_ENVIRON_UPDATE_TOKEN, CC_ENVIRON_UPDATE_URL, CC_INSTANCE_ID, CC_OWNER_ID, CC_PRETTY_INSTANCE_NAME, CC_REVERSE_PROXY_IPS, CC_USE_PULSAR_LOGSCOLLECTION, COMMIT_ID, HOME, HOSTNAME, INSTANCE_ID, INSTANCE_NUMBER, INSTANCE_TYPE, NAME, PATH, PORT, PWD, SHLVL, VAULT_LOCAL_CONFIG, VAULT_PG_CONNECTION_URL, VERSION
              Go Version: go1.23.3
              Listener 1: tcp (addr: "0.0.0.0:8080", cluster address: "0.0.0.0:8081", disable_request_limiter: "false", max_request_duration: "1m30s", max_request_size: "33554432", tls: "disabled")
               Log Level:
                   Mlock: supported: true, enabled: false
           Recovery Mode: false
                 Storage: postgresql (HA disabled)
                 Version: Vault v1.18.3, built 2024-12-16T14:00:53Z
             Version Sha: 7ae4eca5403bf574f142cd8f987b8d83bafcd1de
2025-01-03T14:25:52.301Z [INFO]  proxy environment: http_proxy="" https_proxy="" no_proxy=""
2025-01-03T14:25:52.333Z [INFO]  incrementing seal generation: generation=1
2025-01-03T14:25:52.333Z [WARN]  no `api_addr` value specified in config or in VAULT_API_ADDR; falling back to detection if possible, but this value should be manually set
2025-01-03T14:25:52.336Z [INFO]  core: Initializing version history cache for core
2025-01-03T14:25:52.336Z [INFO]  events: Starting event system
==> Vault server started! Log data will stream in below:
Application start successful
Successfully deployed in 0 minutes and 28 seconds
```

Au démarrage, Vault indique que la configuration est bien chargée, et affiche quelques _warnings_.

> Concernant le _warning_ mentionnant l'IPC_LOCK, il n'est pas possible à ma connaissance de forcer l'option `--cap-add IPC_LOCK` sur Clever Cloud. Néanmoins, ce _warning_ ne pose pas de problème, puisque le _lock_ de la mémoire est désactivé avec le paramètre `disable_mlock`.

Une fois le démarrage terminé, la commande `clever open` permet d'ouvrir un navigateur web sur notre instance de Vault !

```shell
$ clever open

Opening the application in your browser
```

L'URL d'accès à Vault peut aussi être récupérée de deux manières : avec la commande `clever domain`, ou avec un _output_ Terraform qu'on ajoute au code qui crée l'instance Docker.

```shell
$ clever domain
  app_72d4b5a4-1ab8-4653-a825-9be0c62e0fa1.cleverapps.io
* vault-instance.cleverapps.io
```

```terraform
output "vault_url" {
  description = "URL of the Vault instance."
  value = clevercloud_docker.vault_instance.vhost
}
```

```shell
$ terraform output -raw vault_url

app_72d4b5a4-1ab8-4653-a825-9be0c62e0fa1.cleverapps.io
```

## Initialisation du Vault

Lors de sa première ouverture, Vault doit être initialisé, puis déverrouillé. Ces étapes permettent de créer ses clés de déverrouillage (_unseal keys_), ainsi que le _token_ d'accès `root` qui permettra d'utiliser l'API dans un premier temps.

Ces opérations doivent être faites une seule fois à la création du serveur Vault et doivent être faites manuellement _via_ le CLI Vault ou sa console. Dans cet exemple, nous allons effectuer ces manipulations dans la console de Vault :

![img.png](vault-create-root-keys.png)

Une fois le nombre de clés choisi, ainsi que les différentes options de chiffrement, Vault génère les clés et les met à disposition sur l'écran suivant :

![img.png](vault-root-keys.png)

> Ces clés ne doivent être perdues en aucune circonstance ! En cas d'utilisation en production, le nombre de clés souhaité sera probablement différent de 1 !

Après avoir stocké les clés en lieu sûr, l'écran suivant nous invite à déverrouiller Vault en saisissant une clé de déverrouillage.
Lorsque suffisamment de clés auront été entrées, Vault sera déverrouillé et prêt à l'utilisation.

![img.png](vault-unseal.png)

Une fois Vault déverrouillé, l'écran de _login_ apparaît, il est alors possible de se connecter avec le _token_ d'accès `root` obtenu aux étapes précédentes :

![img.png](vault-login.png)

La console de Vault est maintenant disponible :

![img.png](vault-console.png)

Vault est maintenant initialisé, déverrouillé et prêt à être utilisé !

L'article suivant traitera de la configuration de Vault pour utiliser l'authentification OIDC de GitLab, et finaliser cette architecture.

## En conclusion

Cet article a présenté comment mettre en œuvre l'installation et la configuration d'un serveur Vault sur Clever Cloud.

C'est cette infrastructure qui m'a permis de pouvoir mettre à disposition rapidement un serveur Vault pour mes étudiants, afin de les former à la récupération de secrets depuis une application Spring Boot.

Pour exécuter l'infrastructure proposée dans cet article, il vous en coûtera environ 16 €/mois avec les plans utilisés :

| article          | prix/mois |
|------------------|-----------|
| PostgreSQL - Dev | 0&nbsp;€  |
| Docker - Plan XS | 16&nbsp;€ |

Cette architecture n'est pas parfaite, mais permet de facilement déployer un Vault pour des cas d'usage simples ou un environnement de dev. Il faudrait bien entendu la revoir (en particulier les plans utilisés) pour un environnement de production.

## Liens et références

* Exemples de code de cet article sur [GitHub](https://github.com/juwit/terraform-clevercloud-playground/tree/main/vault)
* Page d'accueil de [Clever Cloud](https://www.clever-cloud.com/)
* Installation du [CLI Clever Cloud](https://www.clever-cloud.com/developers/doc/cli/getting_started/)
* Installation du [CLI Terraform](https://developer.hashicorp.com/terraform/install)
* Installation du [CLI OpenTofu](https://opentofu.org/docs/intro/install/)
* Documentation du provider [Terraform Clever Cloud](https://registry.terraform.io/providers/CleverCloud/clevercloud/latest) :
  * Ressource [`clevercloud_postgresql`](https://registry.terraform.io/providers/CleverCloud/clevercloud/latest/docs/resources/postgresql)
  * Ressource [`clevercloud_docker`](https://registry.terraform.io/providers/CleverCloud/clevercloud/latest/docs/resources/docker)
* Documentation de [Vault](https://developer.hashicorp.com/vault/docs) :
  * L'image Docker de Vault sur [dockerhub](https://hub.docker.com/r/hashicorp/vault)
  * Configuration du [storage PostgreSQL](https://developer.hashicorp.com/vault/docs/configuration/storage/postgresql)
* IPC_LOCK et mlock :
  * Manpage des [capabilities](https://www.man7.org/linux/man-pages/man7/capabilities.7.html) Linux (pour l'option `--cap-add IPC_LOCK`)
  * Manpage de l'appel système [mlock](https://www.man7.org/linux/man-pages/man2/mlock.2.html)
  * Un article [Vault and mlock()](https://support.hashicorp.com/hc/en-us/articles/115012787688-Vault-and-mlock) dans le help center de HashiCorp
  * Le paramètre [`disable_mlock` dans la configuration de Vault](https://developer.hashicorp.com/vault/docs/configuration#disable_mlock)
* Photo de couverture par [Jason Dent](https://unsplash.com/@jdent?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) sur [Unsplash](https://unsplash.com/photos/black-and-silver-door-knob-3wPJxh-piRw?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
