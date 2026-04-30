---
date: 2026-04-30
title: "Tangled - une alternative à GitHub sur AT Protocol"
slug: tangled
tags:
  - git
  - souveraineté
cover_anchor: top
draft: true
---

Avec la chute progressive de GitHub, on voit de plus en plus de messages d'entreprises ou de projets qui cherchent à en sortir.

Au-delà de l'hébergement d'un simple serveur Git, GitHub c'était aussi une espèce de réseau social.

Tangled se pose comme une alternative : un hébergement Git (presque classique), les interactions sociales qui l'accompagnent (stars, issues, etc.), l'hébergement de sites statiques (comme les GitHub Pages), et l'exécution de pipelines (similaire à GitHub Actions).

Et tout ça avec des notions d'auto-hébergement, liées à l'AT Protocol, mais aussi pour le code et les pipelines.

Ça a l'air plutôt complet, et c'est encore en version _alpha_, donc la stabilité et les features vont continuer à évoluer dans les prochains mois.

J'ai testé ça pour vous.

<!--more-->

## Tangled c'est quoi ?

Tangled, c'est donc une plateforme "sociale" d'hébergement de code.

Voyez donc ça comme une alternative à GitHub, ou même GitLab.

Pour fonctionner, Tangled s'appuie sur trois éléments distincts :

* l'AT Protocol pour l'authentification des utilisateurs, et le stockage des données sociales (stars, issues, PRs, etc.)
* des serveurs Git, nommés _knots_ pour le stockage des données du code
* des runners de CI nommées _spindle_

Par défaut, Tangled propose l'utilisation de son propre PDS AT Protocol (`tngl.sh`).
Si vous n'avez pas encore de compte AT Protocol (comme un compte Bluesky), vous pourrez en créer un qui sera hébergé sur `tngl.sh`. Si vous avez déjà un compte Bluesky, sur un PDS appartenant à Bluesky ou un autre (comme Eurosky), vous pouvez utiliser ce compte.

Le PDS de Tangled est hébergé en Finlande.

Et c'est là où est l'idée de génie. Toutes les données relatives à Tangled sont stockées sur votre PDS, dans des records préfixés `sh.tangled`. En cas de migration de compte sur un autre PDS, les données vous suivent.

![Les records AT Proto de mon compte](at-proto-records.png)

Les _knots_ sont des simples serveurs Git, auto-hébergeables. Encore une fois, Tangled a son propre _knot_, qui permet d'héberger le code sans avoir besoin de démarrer sa propre instance.
Mais si vous souhaitez héberger votre propre _knot_, pour conserver la maitrise de vos données, c'est aussi possible.

Enfin, le même principe s'applique pour les _spindle_, qui sont les runners de CI.
Ici aussi, Tangled propose son propre _spindle_, mais il est possible d'en auto-héberger un.

## Le setup du compte

Si vous avez déjà un compte AT Protocol (Bluesky principalement), la création de votre compte sur Tangled se fait simplement en se connectant avec votre compte existant.

Les données relatives à Tangled sont alors stockées sur votre PDS.

> Je n'ai pas encore migré mon compte sur un PDS de type Eurosky, donc je ne sais pas concrètement si ça fonctionne, mais je me doute que c'est bien le cas.

Une fois cette première étape franchie, on a accès à la plateforme.

La page d'accueil présente une timeline avec les activités d'autres personnes, et quelques repos _Trending_.

![Page principale de Tangled](tangled-main-page.png)

Comme pour tous les hébergements Git, il y a un peu de setup à faire : configurer les clés SSH qui permettront de pousser le code et configurer les adresses mails qui permettent de rattacher les commits au compte.

![Configuration des adresses mail de commit](tangled-config-emails.png)

![Configuration de la clé SSH](tangled-config-ssh.png)

## Créer un nouveau repo

La création d'un repo se fait en quelques clics.

![Formulaire de création d'un repo](tangled-repo-creation.png)

Une subtilité dans la création du repo est la sélection du _knot_, qui est le serveur qui héberge le repo. Je reviendrai sur ce point plus loin, en détaillant la partie liée à l'auto hébergement.

Une fois le repo crée, on nous propose simplement d'y pousser le code.

![La page d'un repo tout neuf!](tangled-empty-repo.png)

J'ajoute le repo à mes remote Git avec la commande `git remote`

```shell
$ git remote add tangled git@tangled.org:codeka.io/website

$ git push tangled
The authenticity of host 'tangled.org (2a04:3541:8000:1000:24de:d2ff:fe7c:6eaf)' can't be established.
ED25519 key fingerprint is SHA256:fLyp6ivr5HqmGI8yJiPYstTiJa2AXF/RAa9kF/ur1xo.
This key is not known by any other names.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added 'tangled.org' (ED25519) to the list of known hosts.

Welcome to Tangled's hosted knot! 🧶
Enumerating objects: 4145, done.
Counting objects: 100% (4145/4145), done.
Delta compression using up to 22 threads
Compressing objects: 100% (3858/3858), done.
Writing objects: 100% (4145/4145), 615.41 MiB | 9.50 MiB/s, done.
Total 4145 (delta 1993), reused 309 (delta 72), pack-reused 0
remote: Resolving deltas: 100% (1993/1993), done.
To tangled.org:codeka.io/website
 * [new branch]      main -> main
```

Le code apparait bien dans le repo, première étape franchie !

![Mon repo une fois le code pushé](tangled-repo-pushed.png)

