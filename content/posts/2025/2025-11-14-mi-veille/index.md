---
date: 2025-11-14
lastmod: 2025-11-14
language: fr
title: La veille de Wittouck - Début novembre 2025
series: La veille de Wittouck
tags:
  - internet
  - cloud
  - git
  - security
  - java
  - docker
  - linux
  - events
  - clevercloud
  - certifications
---

En cette mi-novembre, la sortie de Spring Boot 4 se profile et la saison des conférences continue.
Voici ma sélection d'articles de veille en attendant le prochain DevFest (Lyon le 28 novembre pour moi).

<!--more-->

## 🛜 Internet

* [Pourquoi je passe des certifs ?](https://blog.antoinemayer.fr/2025/11/10/pourquoi-je-passe-des-certifs/) par [Antoine Mayer](https://blog.antoinemayer.fr/)
* [À quoi ça sert, des certifs en 2025 ?](https://blog.zwindler.fr/2025/11/04/a-quoi-ca-sert-des-certifs-en-2025/) par [Denis Germain](https://blog.zwindler.fr/)

> Deux points de vue sur l'exercice des certifications. Ce sujet fait toujours débat dans notre milieu, donc c'est intéressant d'avoir des avis un peu éclairés.
> Je pense que toutes les certifications ne se valent pas, et qu'il est quand même assez rare qu'une certif soit nécessaire (sauf pour bosser chez les Clouds Américains).
> C'est surtout un bon exercice, et ça peut valider un apprentissage le plus souvent théorique. Bref, je ne vais pas refaire un article ici, je vous laisse vous forger votre propre avis.

## ☁️ Cloud

* [Clever Cloud Academy](https://academy.clever.cloud/)

> Ça y est, Clever Cloud lance son programme de formations certifiantes. La première certification est nommée "Cloud Concepts 101" et parle des concepts généraux du Cloud et de la plateforme Clever Cloud (console et CLI). Et c'est gratuit ! Ce serait dommage de s'en priver.

* [Comment Bedrock Streaming optimise ses coûts AWS](https://aws.amazon.com/fr/blogs/france/comment-bedrock-streaming-optimise-ses-couts-aws/)

> Un article intéressant qui présente les principaux leviers FinOps sur AWS. Pas de surprise sur les optimisations compute, stockage et réseau. Par contre, le volet d'optimisation sur les appels aux APIs AWS est plutôt inédit et intéressant.

## 🔀 Git

* [Git: avoid reset --hard, use reset --keep instead](https://adamj.eu/tech/2024/09/02/git-avoid-reset-hard-use-keep/) par [Adam Johnson](https://adamj.eu/) _via_ [Dave Rupert](https://bsky.app/profile/davatron5000.bsky.social) sur [Bluesky](https://bsky.app/profile/davatron5000.bsky.social/post/3m2wltqd5ns2c)
> Une astuce à laquelle je n'avais réellement jamais prêté attention. Je fais rarement des hard reset, les keep sont une super alternative.

## 🔒 Sécurité

* [OWASP Kubernetes Top Ten](https://owasp.org/www-project-kubernetes-top-ten/)

> Je ne savais pas qu'il existait un Top Ten OWASP consacré à Kubernetes. C'est à parcourir et à surveiller pour la version 2025 qui est en cours d'étude (la version actuelle est de 2022).

## ☕ Java

* [Spring Boot 4.0 Migration Guide](https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Migration-Guide)

> Spring Boot 4 arrive, et le guide de migration est déjà disponible.
> On m'a aussi signalé qu'il y avait [un chantier en cours chez Open Rewrite](https://github.com/openrewrite/rewrite-spring/issues/754) pour proposer une recette toute faite, à tester.

## 🐋 Docker

* [Drastically Reduce and Improve the CI/CD Feedback Loop by Going Local with Docker Compose](https://k33g.hashnode.dev/drastically-reduce-and-improve-the-cicd-feedback-loop-by-going-local-with-docker-compose) par [Philippe Charrière](https://k33g.hashnode.dev/)
> Oh. My. God.
> C'est un des use-case les plus originaux que j'ai pu voir ces derniers temps. Philippe implémente des pipelines complets avec juste docker-compose. Ça shift-left sévèrement, et c'est bluffant de simplicité.

## 🐧 Linux

* [System76 CEO Carl Richell Reveals COSMIC Desktop Launch Date](https://linuxiac.com/system76-ceo-carl-richell-reveals-cosmic-desktop-launch-date/)
> J'utilise Pop!_OS sur mes machines depuis assez longtemps (5 ou 6 ans, je pense). Leur environnement de bureau Cosmic est très attendu, je suis content de voir qu'il sera bientôt disponible, je vais tester ça dès que c'est disponible en version finale (je n'ai pas le courage de gérer des bugs ou de l'instabilité).

## 🎫 Évènements

* [Comment être refusé à coup sûr à un CFP ?](https://www.webofmars.com/blog/2025-10-29_comment-etre-refuser-a-un-cfp/) par [Frédéric Léger](https://bsky.app/profile/webofmars.com) sur [Bluesky](https://bsky.app/profile/webofmars.com/post/3m4g65oilpq22)
> Encore un article qui parle des conférences et des CFP, avec cette fois-ci les erreurs à éviter.
> Cette liste complète bien ce que j'avais déjà écrit dans mon article [« Leeloo Dallas Multipass - Répondre aux 5 éléments d'un CFP »]({{< relref "/posts/2025/2025-06-19-cfp-tips-and-tricks" >}}), je l'ai ajouté comme lien complémentaire.

* [200 heures de préparation pour 45 minutes de présentation](https://contretemps.azeau.com/post/200-heures-de-preparation-pour-45-minutes-de-presentation/) par [Olivier Azeau](https://contretemps.azeau.com) _via_ [Bluesky](https://bsky.app/profile/oaz.bsky.social/post/3m5ejv5augs2l)
> Je parle souvent des conférences et de ma propre implication en tant qu'orga ou speaker. Habituellement, je sais que je passe à peu près une heure de préparation pour une heure de présentation en conf (donc plutôt 45 à 50 heures). Mais les formats originaux sont plus lourds à préparer, et je ne suis pas surpris du travail qu'Olivier a investi dans son talk. Je l'ai déjà dit, mais j'ai environ 100 heures de préparation sur mon talk "[Let's play Factorio]({{<relref "talks/talk-lets-play-factorio">}})". Hâte de voir le replay de cette conf.

---

La prochaine publication est prévue autour du 28 novembre 2025 🗓️.
Je serai au DevFest de Lyon ce jour-là, ce sera surtout un récap de cette nouvelle conférence.

Photo de couverture par [Brigitte Tohm](https://unsplash.com/@brigittetohm?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/white-ceramic-cup-with-coffee-UnACLA4mhLQ?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)