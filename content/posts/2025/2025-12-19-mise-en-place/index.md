---
date: 2025-12-19
lastmod: 2025-12-20
language: fr
bluesky: https://bsky.app/profile/codeka.io/post/3mae6vmuy4225
title: Adieu `direnv`, Bonjour `mise`
tags:
  - linux
  - tools
  - shell
  - devops
---

J'ai découvert `mise` dans le [calendrier de l'avent de Siegfried Ehret](https://sieg.fr/ied/avent-2025/04-mise/), avec une approche autour de l'outillage de dev.

Étant utilisateur de `direnv` depuis plusieurs années, je suis habitué à travailler dans mon shell.
`direnv` souffre par contre de plusieurs défauts : il ne peut pas gérer d'alias et se configure uniquement avec du scripting.
J'ai donc tout un tas de scripts shells qui traînent sur mes machines, plus ou moins adaptées aux package-managers `apt` et `pacman`, qui s'intègrent avec `direnv`. J'avais d'ailleurs écrit un [article sur mon usage de `direnv`](/2022/06/17/direnv-pour-votre-shell) en 2022.

Aujourd'hui, j'ai donc décidé de tester l'outil `mise`, pour voir dans quelle mesure cet outil peut venir enrichir mon workflow de dev, ou (pourquoi pas) remplacer `direnv`.

<!--more-->

## mise-en-place

[mise-en-place](https://mise.jdx.dev/) propose 3 types d'usage :

* gestionnaire de packages (un peu comme `asdf`)
* gestionnaire de variables d'environnement (exactement comme `direnv` donc)
* exécution de tâches (comme on pourrait le faire avec `make`, `task` ou `just`)

Il se veut aussi extensible avec la possibilité d'utiliser ou d'implémenter des plugins.

Il est développé en Rust, ce qui promet de bonnes performances et de la stabilité.

## L'installation

L'installation est bien documentée sur le site web : https://mise.jdx.dev/installing-mise.html

Sur mes postes Linux, j'ai suivi la procédure d'installation avec `apt` :

```shell
$ curl -fSs https://mise.jdx.dev/gpg-key.pub | sudo tee /etc/apt/keyrings/mise-archive-keyring.pub 1> /dev/null
$ echo "deb [signed-by=/etc/apt/keyrings/mise-archive-keyring.pub arch=amd64] https://mise.jdx.dev/deb stable main" | sudo tee /etc/apt/sources.list.d/mise.list
$ sudo apt update -y
$ sudo apt install -y mise
```

Une fois l'installation terminée, la commande `mise` est disponible dans mon shell :

```shell
$ mise --version
              _                                        __              
   ____ ___  (_)_______        ___  ____        ____  / /___ _________
  / __ `__ \/ / ___/ _ \______/ _ \/ __ \______/ __ \/ / __ `/ ___/ _ \
 / / / / / / (__  )  __/_____/  __/ / / /_____/ /_/ / / /_/ / /__/  __/
/_/ /_/ /_/_/____/\___/      \___/_/ /_/     / .___/_/\__,_/\___/\___/
                                            /_/                 by @jdx
2025.12.10 linux-x64 (2025-12-16)
```

Une dernière étape dans la configuration est importante : la configuration du shell, afin que le shell appelle la commande `mise` lors du changement de répertoires.

Cette opération se fait en éditant le fichier de configuration de votre shell : `~/.bashrc` ou `~/.zshrc` par exemple.

Pour Zsh que j'utilise, il faut alors simplement ajouter la ligne suivante dans le fichier `~/.zshrc` :

```shell
$ eval "$(mise activate zsh)"
```

Une fois ces opérations faites, il suffit donc de redémarrer votre shell pour que la configuration soit prise en compte.

## Les devtools

Quand je switche entre mes projets, je switche aussi souvent de stack d'outillage, que ce soit la version de Java, de Node, ou de divers outils (comme Hugo que j'utilise sur ce site).

Pour faciliter ces opérations, j'avais implémenté quelques scripts shell que j'utilisais depuis `direnv`. Ces scripts modifiaient mon $PATH pour pointer vers la bonne version de Java, de Node, etc. J'utilisais aussi parfois des outils comme `nvm`
ou `tfenv`.

`mise` est une alternative plus beaucoup plus simple.

Pour les outils les plus courants, les langages de programmation principalement, `mise` dispose d'un plugin _core tools_ qui supporte leur configuration et activation automatique.
Pour les autres outils ou packages, `mise` utilise un ou plusieurs backends (_asdf_ en fait partie) qui permettent l'installation sans avoir rien à configurer en plus. Les outils disponibles peuvent être listés avec la commande `mise registry` :

```shell
$ mise registry | grep java
java      core:java

$ mise registry | grep graalvm
graalvm   asdf:mise-plugins/mise-graalvm
```

La configuration se fait avec un fichier `mise.toml` ou `.mise.toml`.
La configuration est hiérarchique, mise va lire le fichier de configuration dans le répertoire courant, et merger ou surcharger les valeurs de ce fichier avec les valeurs des répertoires parents.

Il est aussi possible de définir un fichier `mise.toml` à la racine du `$HOME` de l'utilisateur, ou dans le sous-répertoire `$HOME/.config/` pour partager des configurations sur tout le système.

### Une stack java

Déclarer une stack Java est très simple avec un fichier `mise.toml` :

```toml
[tools]
java = "temurin-25"
maven = "3"
```

Une fois le fichier crée, au chargement du répertoire, `mise` affichera une erreur, car il ne souhaite par exécuter les fichiers par défaut (par mesure de sécurité) :

```shell
$ cd workspaces/gitlab-classrooms
mise ERROR error parsing config file: ~/workspaces/gitlab-classrooms/mise.toml
mise ERROR Config files in ~/workspaces/gitlab-classrooms/mise.toml are not trusted.
Trust them with `mise trust`. See https://mise.jdx.dev/cli/trust.html for more information.
mise ERROR Run with --verbose or MISE_VERBOSE=1 for more information
```

La commande `mise trust` permet donc d'activer le fichier de configuration donné :

```shell
$ mise trust
mise trusted /home/jwittouck/workspaces/gitlab-classrooms
mise WARN  missing: java@temurin-25.0.1+8.0.LTS maven@3.9.12
```

Il est aussi possible de configurer `mise` pour truster des répertoires complets dans le fichier `.config/mise/config.toml` :

```toml
[settings]
trusted_config_paths = ["~/workspaces"]
```

Une fois le fichier créé et trusté, `mise` affichera les outils manquants s'il y en a avec un warning :

```shell
mise WARN  missing: java@temurin-25.0.1+8.0.LTS maven@3.9.12
```

Pour installer les outils manquants, il suffit alors d'exécuter la commande `mise install` : 

```shell
$ mise install
mise maven@3.9.12 download apache-maven-3.9.12-bin.tar.gz                                          8.57 MiB/8.81 MiB (0s) ███████████████████░  2s
mise java@temurin-25.0.1+8.0.LTS download OpenJDK25U-jdk_x64_linux_hotspot_25.0.1_8.tar.gz     128.19 MiB/134.21 MiB (0s) ███████████████████░  2s
```

Une fois les outils installés, ils sont ajoutés au `$PATH` et peuvent être utilisés immédiatement :

```shell
$ echo $PATH

/home/jwittouck/.local/share/mise/installs/java/temurin-25.0.1+8.0.LTS/bin:/home/jwittouck/.local/share/mise/installs/maven/3.9.12/apache-maven-3.9.12/bin:/home/jwittouck/.local/bin:/usr/local/bin:/usr/bin:/usr/local/sbin

$ java --version
openjdk 25.0.1 2025-10-21 LTS
OpenJDK Runtime Environment Temurin-25.0.1+8 (build 25.0.1+8-LTS)
OpenJDK 64-Bit Server VM Temurin-25.0.1+8 (build 25.0.1+8-LTS, mixed mode, sharing)

$ mvn -v
Apache Maven 3.9.12 (848fbb4bf2d427b72bdb2471c22fced7ebd9a7a1)
Maven home: /home/jwittouck/.local/share/mise/installs/maven/3.9.12/apache-maven-3.9.12
Java version: 25.0.1, vendor: Eclipse Adoptium, runtime: /home/jwittouck/.local/share/mise/installs/java/temurin-25.0.1+8.0.LTS
Default locale: en_US, platform encoding: UTF-8
OS name: "linux", version: "6.17.12-1-manjaro", arch: "amd64", family: "unix"
```

Si je veux utiliser une autre version de Java, rien de plus simple, je peux simplement modifier le fichier `mise.toml`, ou exécuter la commande `mise use` (qui va aussi modifier le fichier `mise.toml`) :

```shell
$ mise use graalvm@25
mise ~/workspaces/gitlab-classrooms/mise.toml tools: graalvm@25.0.1

$ mise use --remove java
mise ~/workspaces/gitlab-classrooms/mise.toml removed: java

$ java --version
openjdk 25.0.1 2025-10-21
OpenJDK Runtime Environment GraalVM CE 25.0.1+8.1 (build 25.0.1+8-jvmci-b01)
OpenJDK 64-Bit Server VM GraalVM CE 25.0.1+8.1 (build 25.0.1+8-jvmci-b01, mixed mode, sharing)

$ mvn -v
Apache Maven 3.9.12 (848fbb4bf2d427b72bdb2471c22fced7ebd9a7a1)
Maven home: /home/jwittouck/.local/share/mise/installs/maven/3.9.12/apache-maven-3.9.12
Java version: 25.0.1, vendor: GraalVM Community, runtime: /home/jwittouck/.local/share/mise/installs/graalvm/25.0.1
Default locale: en_US, platform encoding: UTF-8
OS name: "linux", version: "6.17.12-1-manjaro", arch: "amd64", family: "unix"
```

C'est pratique, flexible et bluffant.

Les packages sont installés par défaut dans le répertoire `~/.local/share/mise/installs/`, et utilisent des liens symboliques pour relier les différentes version et alias de versions : 

```shell
$ tree -L 2 ~/.local/share/mise/installs/
/home/jwittouck/.local/share/mise/installs/
├── graalvm
│   ├── 25 -> ./25.0.1
│   ├── 25.0 -> ./25.0.1
│   ├── 25.0.1
│   └── latest -> ./25.0.1
├── java
│   ├── temurin-25 -> ./temurin-25.0.1+8.0.LTS
│   ├── temurin-25.0 -> ./temurin-25.0.1+8.0.LTS
│   ├── temurin-25.0.1+8.0.LTS
│   └── temurin-latest -> ./temurin-25.0.1+8.0.LTS
└── maven
    ├── 3 -> ./3.9.12
    ├── 3.9 -> ./3.9.12
    ├── 3.9.12
    └── latest -> ./3.9.12
```

Vu que mon IDE (IntelliJ) télécharge aussi automatiquement des versions de Java, je risque d'avoir des doublons, mais je pense qu'il ne sera pas bien difficile d'utiliser ces versions à la place de celles proposées par IntelliJ.

Pour une stack Java plutôt classique, `mise` fait très bien le job.

### Une stack Hugo pour builder mon site

Mon site (ce site) est buildé avec [_Hugo_](https://gohugo.io).

Bien qu'installer Hugo est plutôt facile, autant essayer de basculer le maximum de choses avec `mise` pour ce test.

Hugo est disponible dans le registry de `mise` :

```shell
$ mise registry | grep hugo
hugo            aqua:gohugoio/hugo ubi:gohugoio/hugo asdf:NeoHsu/asdf-hugo asdf:nklmilojevic/asdf-hugo
hugo-extended   aqua:gohugoio/hugo/hugo-extended
```

Plutôt que d'intialiser directement le fichier `mise.toml` avec la configuration de l'outil,
il est possible d'utiliser la commande `mise use`, qui va générer le fichier pour nous et installer l'outil directement, qui sera immédiatement utilisable :

```shell
$ mise use hugo
mise ~/workspaces/codeka.io/mise.toml tools: hugo@0.152.2

$ hugo version
hugo v0.152.2+extended+withdeploy linux/amd64 BuildDate=unknown
```

Le fichier généré est alors le suivant :

```toml
[tools]
hugo = "latest"
```

Utiliser une version "latest" est bien entendu déconseillé, mais cela permet de tester rapidement la nouvelle version de l'outil.

Et encore une fois, tout fonctionne parfaitement :

```shell
hugo serve
Watching for changes in /home/jwittouck/workspaces/codeka.io/{archetypes,assets,content,data,i18n,layouts,static,themes}
Watching for config changes in /home/jwittouck/workspaces/codeka.io/hugo.toml, /home/jwittouck/workspaces/codeka.io/themes/poison/config.toml
Start building sites … 
hugo v0.152.2-6abdacad3f3fe944ea42177844469139e81feda6 linux/amd64 BuildDate=2025-10-24T15:31:49Z VendorInfo=gohugoio


                  │ FR  │ EN  
──────────────────┼─────┼─────
 Pages            │ 134 │  26 
 Paginator pages  │   3 │   0 
 Non-page files   │ 234 │   0 
 Static files     │ 182 │ 182 
 Processed images │ 160 │   0 
 Aliases          │   7 │   1 
 Cleaned          │   0 │   0 

Built in 299 ms
Environment: "development"
Serving pages from disk
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1) 
Press Ctrl+C to stop
```

## Utiliser des variables d'environnement

`mise` permet de déclarer des variables d'environnement dans le fichier `mise.toml` avec le bloc `[env]` :

```toml
[env]
SPRING_PROFILES_ACTIVE = 'local'
```

Il est aussi possible de charger un fichier `.env` avec [une directive](https://mise.jdx.dev/environments/#env-directives) :

```toml
[env]
_.file = '.env'
```

C'est ici très similaire à `direnv`, donc pas de surprise.

### Utiliser des secrets

Là où `mise` se distingue, c'est sur la gestion des secrets.

`mise` s'intègre avec `fnox` (par le même développeur). Il est alors possible de charger des secrets depuis un Vault ou un secret manager cloud (AWS Secrets Manager, Azure Key Vault, Google Cloud Secret Manager), depuis un Keypass.

`mise` s'intègre aussi avec `age` qui permet d'utiliser des clés SSH pour chiffrer les secrets, ce qui est bien pratique si vous baladez vos clés sur vos différentes machines, puisqu'il n'y a alors aucune configuration à faire la clé `~/.ssh/id_ed25519` ou `~/.ssh/id_rsa` sera utilisée (dans cet ordre de priorité).


```shell
$ mise set --age-encrypt --prompt GITLAB_CLIENT_ID
Enter value for GITLAB_CLIENT_ID ************

$ mise set --age-encrypt --prompt GITLAB_CLIENT_SECRET
Enter value for GITLAB_CLIENT_SECRET ************
```

Les valeurs chiffrées sont stockées dans le fichier `mise.toml`, qui peut alors être poussé sur un repo git en toute sécurité :

```toml
[env]
GITLAB_CLIENT_ID = { age = { value = "KLUv/QBYDR0AukFwDj0ANfCYs5DmZabA2Mi9YcrEQIqnBlqmxxa45qzJ6BIme5SvwM1b8bd87RMmC0YR540QLz2IWr5cNzRj/PEI0ADOANIAqBfCLgps21MFvkmjrMekjza1mjgCw+p8L11IHD38hL0pa5LJTF+jo1AGgjcpTFoprubn359HcX5T458wH/F00eoLinZHXzTatFtdm3l6rwWfMp5Fb6Wb0fa9a3qt53IQaaN2bE4r8x/PWJqUtrROIYthKb91VtSKW/Lk6CrXY87IgenHjEcrm5N87Zo/pep79tdFOz0vxuWG89vnlpJTnIfd+XmzIY2pzaRAunuB0ar/sJi3lbR6PinpumTCpMJ4Cim3I9S2vN9+/tbwBCa1AbNNrhUxnDXLrdl8UoryOpoLM3a9fjoomVX5hnHybBTy3o4crTXKi0D6/qU0D3vGjsZHRe9DgyomNrrXJndAdY+m1phS2rGIyjZbqVcn3n2XRsoi7rb3E6Wx6oWCUyT+E1eRbZrmLMq8XpY+E6/PHVIDdkQdEe7hGJB2ReHvglVm1riPWty/iz6d7fTeDGKpsrY8sbv824jjf2j35IXd2au1Z6g2WUhRt+Hkiphy+32HK9mCYaWSDWt6Y8FgX3tV1HJ7f1P14P40p7iUxsEB9ZflPQbDSfbfZZWFIvO2BRFHUprYBzSEBxcJTc6zKoH1Fs5JE01sQOEy4YJkwqRC3QpH0k/7um0wn3yBvXkS00y88MyU+9JCsThhCE69nh/PKjt+rQt2ycvX/8RXRdkrDmu6CbNjHspqmbv8o74Hc6fsfVVKJ5Dz+O/6rOY58x/SlMlmddqAsxvy2Ke28eLceRzRAnronyS1FNNEb4xlnY+VUKNLYTUd+pPqydJwiCsrnV82yS2r1sznsqnm5mKhvjxH6G3RKC2GU3dfds0tqY8BOgbQuJHBUoGDD8SDDxMmSNQ4kWCIJqCARkaFF0J43FBAZMSHIwFKoMAhZAUBBxv4FlsKDTkW7AABsCFBhIoOP2L2h8kBBxxFQIDAwCCIEZ8DwAJBiNJB66KeAeNADy2GgSE10gpMw9uD+t2+iVrXN0Omy2vqB4+oaZeYO6UBXe2aSyhP/Hwkk8mkQjOVOno3j2vjtCOKpYv7MjhXLReVwzGzVs7Sg+k24m7Ouk1SQD2/cWi+cpl9Di3Z2W3jpbA6ZT0Tp/h5ax0nzFFitrK3mRpZTvHOecZW7VqL9EIWCCU9+2uFrmZzAgA", format = "zstd" } }
GITLAB_CLIENT_SECRET = { age = "YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IHNzaC1yc2EgYWM5WlV3Cmpsa0EweTlISnFLQmJ4NHBRQTRDVGRBd21NM3JWbVo1b3NFZHQ3dXcvZHZaQmFYMkEyWjViWmpJWCtLT0R1aVkKckNsUW1qVHZ1NlBEelYzQm9nNkk1SEp3REJsNzZvdmJuTkdYd2tLV3VCRFBRUXlCYjhhVzNvSnBib0dDcUFXTwpydW1INWIyMWtjaCt2WnlGeEx5dXM5Q205NUJBdmlycjUzdE41ejFFOW9PZ1YzaXczM292bXlsSkpPWW5hVldQCkJOUUZ3ditlbGFJbzRWZU5TeTBaSitoUk1TeDh4RVJqUjRaVkV5MGJrOXhJeitsbjBKcEQ0TkZla1V4bnNybUoKWVBjeSttUkFMVkd2Q2t1dDJVUWxjMk9vcER3cWFZcjNBRHNTZVJqMGZNc01xQ1VZd3pHU2hQazJWczFRbjd6OAowSHRqS0tCeE1HL2MrVG5FaUhaVndvWm9xNzNrME41VjJDRzlzbEFGRzdvQ040elNSeVAxcVdGMVlSK2NsQlM3ClJFbGZtV2pVOUh3VmF2RXlhNE5pekpIcnVpSzF2WWs5YWMvSTRqYnBQMnBMUFpValNTajI2WVp2YXRSTWgxd1AKTGVZektqdWxjRENjMVJaZy9aVEhvQkgzNXdyT01PY1NKT2t0MHB2VW9RK1hCSmxkbk5aWElGVndKeTRZRi90VApXMDhnWFlCVGRZeTFZcHRKSUdRbHkrRXlidXdKY0lCUVNjYy8vR21LbTk5eG5UTFp3L1RUM3JFbldRNldkTkJFCmVsZ2VqWXc3c3Q0eERtb0YxSDhNZjlrOUdJdmp4QjhEVW16UlRGY0xrbGtBZXkwcVkvZ1MrWVRKYXdqTm9QYjkKcUJzNS9MdUR4UENvU05iRHo5cGVEaGpvTDVwdU0xMUsrUGRBQkNaUFZaVQotPiB0LWdyZWFzZSA5KklKaXIgWlcyCjJzUlRqRXo3ZnJ6ZVZndk92UVE5QllNejhNZGZqVjYxMWN5U0JjbzhUbTJhCi0tLSBmRXEwQnh0Ykc4NWtIdk1nd200LzI2aEZtVEJqZ2x6QWp3ZCs0TlJjYThBCrPsAPxQrwsNRNctYosJQ7GQM8+Zc4bdoMoTbSehN6Fyf/5AvfJ2ko3Gm5FJ9L2qnW0ZGW+QTLMsR4vLInmecojm/eu4LKC5EqBB/3oFCAJJNZAdOtYiJnwUCWml5WY" }
```

Et les valeurs sont automatiquement déchiffrées pour être disponibles dans le shell en variable d'environnement : 

```shell
$ env | grep GITLAB
GITLAB_CLIENT_ID=a578cf01-96cf-4d60-bef0-c89240583992
GITLAB_CLIENT_SECRET=9520fa56-6663-4008-888a-cd50e6b575a9
```

C'est clairement une killer-feature. Je n'ai pas pris le temps de tester en détail l'intégration avec `fnox`, mais je pense que je vais m'y atteler prochainement.

## Exécuter des tâches

Le dernier use-case proposé par `mise` est l'exécution de _task_, qui sont autant de petits scripts que l'on veut, un peu à la manière d'un `Makefile`, `justfile`, ou `Taskfile`.

L'avantage est que tout s'intègre parfaitement avec `mise`, il est facile d'utiliser des variables d'environnement et des tools dans des scripts qu'on va pouvoir exécuter directement avec `mise run`.

Et on peut aller assez loin dans la configuration, avec des dépendances entre les tâches ou conditionner l'exécution à la mise à jour d'un fichier par exemple.

Pour mon site, j'avais un `justfile` qui contenait quelques commandes, je l'ai rapidement migré, et en voici le résultat :

```toml
[tools]
hugo = "0.152"

[env]
HUGO_OPTIONS = "--printI18nWarnings --printPathWarnings --cleanDestinationDir --logLevel debug"

[tasks]

[tasks.build]
description = "Build le site avec Hugo"
run = "hugo build $HUGO_OPTIONS"

[tasks.serve]
description = "Lance le serveur de développement Hugo"
run = "hugo server --openBrowser $HUGO_OPTIONS"

[tasks.draft]
description = "Lance le serveur en incluant les brouillons et futures publications"
run = "hugo server --openBrowser $HUGO_OPTIONS --buildDrafts --buildFuture"
```

Maintenant, avec un simple fichier `mise.toml`, j'ai tout au même endroit, et je n'ai besoin que d'un seul outil.
Les tâches peuvent bien entendu en appeler d'autres, ou utiliser des [fichiers de script](https://mise.jdx.dev/tasks/file-tasks.html) pour leur implémentation, cette partie est assez riche, 

Je démarre mon site en test avec la commande `mise run draft` ou même directement avec `mise draft`, et je n'aurai plus à me soucier de l'installation de Hugo et Just sur mes machines (voire même dans ma CI 👀).

```shell
$ mise run build
[build] $ hugo build $HUGO_OPTIONS
Start building sites … 
hugo v0.152.2-6abdacad3f3fe944ea42177844469139e81feda6+extended linux/amd64 BuildDate=2025-10-24T15:31:49Z VendorInfo=gohugoio

                  │ FR  │ EN  
──────────────────┼─────┼─────
 Pages            │ 134 │  26 
 Paginator pages  │   3 │   0 
 Non-page files   │ 234 │   0 
 Static files     │ 182 │ 182 
 Processed images │ 160 │   0 
 Aliases          │   7 │   1 
 Cleaned          │   0 │   0 

Total in 304 ms
Finished in 540.7ms
```

C'est plutôt pratique de pouvoir avoir quelques scripts packagés avec cet outil. Ça permet surtout de ne pas avoir à installer d'autres outils comme `task` ou `just` sur mes machines.

## Déployer sur Clever Cloud avec `mise`

Maintenant qu'on a vu comment fonctionnait `mise`, j'ai envie d'implémenter un cas concret supplémentaire.

Le site que vous lisez est hébergé sur Clever Cloud. J'ai donc sur mon poste des variables d'environenment `CLEVER_TOKEN` et `CLEVER_SECRET` qui traînent, un fichier `.clever.json` avec les liens vers mes apps, ainsi qu'un runtime node avec le CLI `clever-tools`.

Est-ce qu'il est facile de migrer tout ça vers `mise` ?

### L'installation du CLI `clever-tools`

Installer le CLI de Clever Cloud avec `mise` se fait en une commande : `mise use npm:clever-tools`.
Le préfixe `npm:` permet d'indiquer à `mise` où aller chercher le package, cela ouvre tout un tas de cas d'usages intéressants. Fini les `npm install -g` (ou pire `sudo npm install -g`) dans tous les sens, on installe juste l'outil au bon endroit, et la dépendance devient même explicite.

```shell
$ mise use npm:clever-tools
mise ~/workspaces/codekaio/codeka.io/mise.toml tools: npm:clever-tools@4.4.1
```

Une fois la commande exécutée, `clever-tools` apparaît dans le fichier `mise.toml`, dans le bloc `[tools]` :

```toml
[tools]
"npm:clever-tools" = "latest"
```

### `CLEVER_TOKEN` et `CLEVER_SECRET`


Pour s'authentifier aux API, le CLI `clever-tools` utilise 2 variables d'environnement `CLEVER_TOKEN` et `CLEVER_SECRET`, qu'on peut récupérer en exécutant la commande `clever login` :

```shell
$ clever login
Opening https://console.clever-cloud.com/cli-oauth?cli_version=4.4.1 in your browser to log you in…
Login successful as Julien WITTOUCK <julien@codeka.io>
```

Les variables sont alors récupérables depuis la fenêtre de navigateur qui sert à l'authentification, et sont déposées dans le fichier `.config/clever-tools/config.json` :

```json
{"token":"081e54b8-853b-4e8f-a000-8f314238c5c3","secret":"6834da77-9edb-4e4a-bde8-9a6109934311","expirationDate":"2026-12-21T11:35:15.291Z"}
````

Pour stocker les variables d'environnement de manière sécurisée avec `mise`, c'est plutôt facile, j'utilise la commande `mise set` avec les valeurs que je récupère directement depuis le fichier `.config/clever-tools/config.json`:

```shell
$ jq -r '.token' ~/.config/clever-cloud/clever-tools.json | mise set --raw --age-encrypt --prompt CLEVER_TOKEN
$ jq -r '.secret' ~/.config/clever-cloud/clever-tools.json | mise set --raw --age-encrypt --prompt CLEVER_SECRET
```

Les deux variables apparaissent bien dans mon `mise.toml` :

```toml
[env]
CLEVER_TOKEN = { age = "XXX" }
CLEVER_SECRET = { age = "YYY" }
```

J'ai fait la même chose avec mon `APP_ID` afin de pouvoir également le stocker sans souci sur git :

```shell
$ mise set --raw --age-encrypt --prompt PROD_APP_ID
Enter value for PROD_APP_ID ************
```

### La tasks

Écrire les tasks pour `mise` est alors plutôt facile, il suffit d'enchaîner les deux commandes `clever link` et `clever deploy` dans un seul script :

```toml
[tasks."clever:deploy"]
description = "Déploie sur l'alias Clever Cloud de production"
run = '''
clever link $PROD_APP_ID --alias codeka.io
clever deploy --alias codeka.io
'''
```

Il suffit alors ensuite de lancer la commande `mise clever:deploy` et mon site est en production :

```shell
mise clever:deploy
[clever:deploy] $ clever link $PROD_APP_ID --alias codeka.io
✓ Application [redacted] has been successfully linked to local alias codeka.io!
🚀 Deploying codeka.io
   Application ID  [redacted]
   Organisation ID [redacted]

🔀 Git information
   Remote head     366fb84818efb3c9e67a758a99c7fa5c53466b3b (refs/heads/main)
   Deployed commit 366fb84818efb3c9e67a758a99c7fa5c53466b3b
   Local commit    4cf8d3a5d18ea0d12411fc16e362a47224fb57b8 [will be deployed]

🔄 Deployment progress
   → Pushing source code to Clever Cloud…
   ✓ Code pushed to Clever Cloud
   → Waiting for deployment to start…
```

## Conclusion

J'ai testé simplement pendant quelques heures, en migrant les outils de mon site et un de mes projets.
Je suis vraiment bluffé par la puissance de `mise`, sa flexibilité et ses performances. C'est rapide à prendre en main, ça supporte tous les outils et langages que j'utilise, et ça vient en plus avec une solution pour la gestion des secrets et de scripting.

En une aprèm, je suis déjà conquis.

Adieu `direnv`, `asdf`, `task` et `just`, et tous mes scripts custom, vous m'avez bien servi.

## Liens et références

* Le blog post de Siegfried Erhet : https://sieg.fr/ied/avent-2025/04-mise/
* Mon [article sur mon usage de `direnv`](/2022/06/17/direnv-pour-votre-shell)
* La documentation officielle de `mise` : https://mise.jdx.dev
  * L'installation de `mise` : https://mise.jdx.dev/installing-mise.html
  * L'utilisation des DevTools : https://mise.jdx.dev/dev-tools/
  * La liste des outils disponibles : https://mise.jdx.dev/registry.html#tools
  * L'utilisation des variables d'environnement : https://mise.jdx.dev/environments/
  * Utilisation de tâches : https://mise.jdx.dev/tasks/
