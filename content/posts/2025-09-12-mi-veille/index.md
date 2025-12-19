---
date: 2025-09-12
language: fr
title: La veille de Wittouck - Début Septembre 2025
series: La veille de Wittouck
tags:
  - Docker
  - Git
  - IA
  - Internet
  - Java
  - Kubernetes
  - Linux
---

C'est la rentrée de la Veille de Wittouck !

J'espère que vous avez passé une bonne rentrée.
J'ai une belle liste de liens à dépiler que j'ai accumulé pendant l'été.

Comme pour l'année dernière, je vais essayer de garder le rythme d'une publication toutes les deux semaines.

J'ai aussi passé une partie de l'été à travailler sur l'organisation de CloudNord, en particulier sur le CFP.
L'agenda est d'ailleurs disponible sur [cloudnord.fr/schedule](https://cloudnord.fr/schedule/).

<!--more-->

## 🛜 Internet

* [Investing in what moves the internet forward | The Mozilla Blog](https://blog.mozilla.org/en/mozilla/building-whats-next/)

> Mozilla a annoncé en début d'été la fin de vie de son service Pocket, qui était utilisé encore par beaucoup de bloggers.
C'est assez soudain, le service sera complètement éteint d'ici au 8 octobre.
Par chance, j'avais déjà migré vers raindrop.io début mai.

* [Quelques livres que je recommanderais, pour des ingénieur(e)s logiciels](https://blog.pascal-martin.fr/post/quelques-livres-que-je-recommande-pour-des-ingenieurs-logiciels/) par [Pascal Martin](https://blog.pascal-martin.fr/)

> Pascal partage une liste plutôt exhaustive des livres qu'il considère comme devant faire partie de la bibliothèque idéale d'un ou une dev.
> La liste est longue, et contient beaucoup d'ouvrages considérés comme des références, la majeure partie des livres sont en anglais.

* [Building Bluesky Comments for My Blog](https://natalie.sh/posts/bluesky-comments/) par [natalie.sh](https://natalie.sh/) _via_ [Nicolas Fränkel](https://bsky.app/profile/frankel.ch) sur [Bluesky](https://bsky.app/profile/frankel.ch/post/3lw57yjcn4k2p)

> Une idée de génie, tout simplement. Je pense que je vais tenter d'implémenter la même chose, mais je sens déjà que je vais galérer avec Hugo 😅

## ☕ Java

* [Upgrading to JUnit 6.0](https://github.com/junit-team/junit-framework/wiki/Upgrading-to-JUnit-6.0)

> Je suis tombé là-dessus un peu par hasard en préparant mes cours.
> JUnit 6 amène une baseline en Java 17 et l'utilisation des annotations JSpecify pour la null-safety (à combiner à la compilation avec NullAway par exemple).
> C'est pour l'instant en version 6.0.0-RC2, donc je pense que la version 6.0.0 ne devrait pas tarder.
> Par contre, je pense que l'intégration dans Spring Boot ne sera peut-être pas pour la 4.0 (_a priori_, ils sont encore sur la 5.X sur la dernière milestone).

## 🧠 IA

* [Function calling using LLMs](https://martinfowler.com/articles/function-call-LLM.html) par [Kiran Prakash](https://www.linkedin.com/in/kiran-prakash/) _via_ [Martin Fowler](https://www.linkedin.com/in/martin-fowler-com/) on [LinkedIn](https://www.linkedin.com/posts/martin-fowler-com_function-calling-using-llms-activity-7325506940615221248-tJo-?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)

> Un excellent article qui décrit pas-à-pas la construction d'un agent IA capable d'appeler des fonctions externes au LLM.
L'article conclu par l'implémentation d'un serveur MCP. C'est une super introduction pour découvrir le fonctionnement de ce protocole.

* [Problèmes et dangers de l'IA générative sous l'angle du développement informatique](https://bsky.app/profile/nathe.xyz/post/3luue6mgwcc2w) par [Anathexyz](https://bsky.app/profile/nathe.xyz/) sur [Bluesky](https://bsky.app/)

> Cet excellent thread porte des sujets de réflexion sur l'usage de l'IA dans nos métiers. Les grands thèmes abordés sont les problématiques de sécurité, les coûts, les gains de productivité, et le sujet qui me tient le plus à cœur : le déterminisme.
> C'est très bien écrit, et je pense que c'est compréhensible par des gens qui bossent dans notre secteur, sans être des experts en IA ni même des techs.

* [Titles matter](https://joshcollinsworth.com/blog/titles-matter) par [Josh Collinsworth](https://joshcollinsworth.com) sur [Bluesky](https://bsky.app/profile/collinsworth.dev/post/3lxcwd4lzvc2v)

> Une réflexion qu'on pourrait résumer en "quelqu'un qui a vibe-codé une application peut-il se dire développeur ?".
> L'article se veut plutôt constructif et honnête, sans chercher à faire du _gatekeeping_ ni à dénigrer les utilisateurs des IA Gen.

* [Charte du bon usage des IA génératives à l’Université de Toulouse](https://www.univ-tlse3.fr/medias/fichier/charte-du-bon-usage-des-ia-generatives-a-l-universite_1755517998722-pdf)

> Ce document a servi de base à la rédaction d'une mini-charte par les enseignants de l'Université de Lille.
> Cette charte vise à cadrer l'usage des IA Gen par les étudiants, sans l'interdire, mais plutôt en proposants un usage d'aide à la compréhension.
> "L'attente de l'enseignant ne peut jamais être que vous rendiez intégralement ou partiellement une solution rédigée par une IA Gen."

## ☸️ Kubernetes

* [Resize CPU and Memory Resources assigned to Containers](https://kubernetes.io/docs/tasks/configure-pod-container/resize-container-resources/)

> Cette feature de Kubernetes me fait de l'œil depuis un bon moment. Elle permet de modifier les ressources attribuées à un pod en cours d'exécution.
> Elle pourrait permettre d'aider au démarrage des applications et aussi sur le scale-up.
> Je pense qu'elle fera l'objet d'un article détaillé ici dans les prochaines semaines, je dois tester tout ça avec des runtime Java.

## 🐋 Docker

* [Fork Yeah: We’re Bringing Kaniko Back](https://www.chainguard.dev/unchained/fork-yeah-were-bringing-kaniko-back) _via_ [Idriss Neumann](https://bsky.app/profile/ineumann.fr) sur [Bluesky](https://bsky.app/profile/ineumann.fr/post/3lsetrdcomc2i)

> J'étais passé au travers de l'annonce de la fin du support de Kaniko par Google.
> C'est dommage, car c'est un des rares outils qui fonctionnait sans avoir besoin de Docker ou de Docker-in-Docker (🤢).
> Je l'ai beaucoup utilisé dans des pipelines de CI/CD.
> L'outil est maintenant développé et supporté par ChainGuard sur leur GitHub : [chainguard-dev/kaniko](https://github.com/chainguard-dev/kaniko)

## 🔀 Git

* [Git clean local branches](https://alsohelp.com/blog/git-clean-local-branches) par [David Boureau](https://alsohelp.com/) _via_ [Yannick TRONCAL](https://about.me/ytroncal) sur [Bluesky](https://bsky.app/profile/ytroncal.bsky.social/post/3lq7i7pyn322q)

> Un joli petit bout de script pour faire le ménage dans les branches locales Git.
> Peut-être à regrouper avec un alias ?

* [Git Notes: git's coolest, most unloved feature](https://tylercipriani.com/blog/2022/11/19/git-notes-gits-coolest-most-unloved-feature/) par [Tyler Cipriani](https://tylercipriani.com/) _via_ [Lea Linux](https://bsky.app/profile/lea-linux.org/post/3lsa2ayugck2b)

> Une feature de Git que je ne connaissais pas. Les cas d'usage présentés (review, commentaires, attachement de résultat de tests) sont hyper cools, mais sans outillage ça doit quand même être difficile à manipuler.
> La doc officielle de [Git Notes](https://git-scm.com/docs/git-notes#_examples) liste un exemple d'utilisation pour annoter un commit avec le nom de la personne qui a testé le code.

---

La prochaine publication est prévue autour du  🗓️ 26 septembre 2025.

Photo de couverture par [kyo azuma](https://unsplash.com/@tokyo_boy?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) sur [Unsplash](https://unsplash.com/photos/empty-building-hallway-x_TJKVU1FJA?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)