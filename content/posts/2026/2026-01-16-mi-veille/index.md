---
date: 2026-01-16
language: fr
title: La veille de Wittouck - Début janvier 2026
series: La veille de Wittouck
tags:
  - devops
  - docker
  - events
  - shell
  - ia
  - internet
  - java
  - linux
---

Après des vacances bien méritées, [la Veille de Wittouck](/series/la-veille-de-wittouck) est de retour avec la même formule : mon analyse rapide des articles, vidéos ou podcasts qui m'ont intéressés sur les deux dernières semaines.
Je ferai également un feedback de mes participations aux différentes conférences de l'année en _éditions spéciales_, comme je l'ai fait en 2025.

<!--more-->

## 👷 DevOps

* [Zéro lien cassé : automatiser la détection avec Lychee](https://blog.stephane-robert.info/post/lychee-verification-liens-casses/) par [Stéphane Robert](https://blog.stephane-robert.info) _via_ [linkedin](https://www.linkedin.com/posts/stephanerobert1_devops-devsecops-documentation-activity-7416715399678914560-CSv1/)

> Stéphane Robert a encore frappé.
Il nous présente ici un outil pour vérifier les liens d'un site web, blog ou documentation : Lychee.
C'est déjà indispensable pour moi, et ça m'a permi de découvrir des liens cassés après une migration, donc ça fait bien le job.

## 🐋 Docker

* [Shutdown Signals with Docker Entrypoint Scripts](https://itnext.io/shutdown-signals-with-docker-entry-point-scripts-5e560f4e2d45) _via_ [Nicolas Fränkel](https://bsky.app/profile/frankel.ch/post/3mccakjqpra24)

> Un article simple, mais qui explique un point important dans les systèmes de containers : le passage des signaux OS (SIGTERM en particulier).
>
> _TLDR_ : Utilisez `exec` dans vos scripts shells pour muter le process de votre entrypoint et passer les signaux automatiquement.

## 🎫 Events

* [Conférences : 10 conseils pour être sélectionné·e à un CFP](https://blog.humancoders.com/conferences-10-conseils-pour-etre-selectionne%c2%b7e-a-un-cfp-3993/) _via_ [linkedin](https://www.linkedin.com/posts/human-coders_prendre-la-parole-en-conf%C3%A9rence-cest-un-activity-7406343149708820481-hzJX)

> Encore un bon article plein de bons conseils pour aborder la saison des conférences, en particulier la rédaction et la soumission d'un CFP.

* [Et si on parlait](https://www.youtube.com/@et-si-on-parlait)

> Une chaîne Youtube animée par Jean-François Garreau, speaker Nantais bien connu.
>
> C'est plein de bon petits tips pour la prise de parole, avec des vidéos courtes, claires, et bien animées. Une bonne source pour préparer ses conférences, pour les speakeuses et speakers de tout niveau.

* [🎂 13 years: Migrating my speaker page with AI](https://david.pilato.fr/posts/2026-01-10-13-years-migrating-to-hugo-with-cursor/) par [David Pilato](https://david.pilato.fr) _via_ [Bluesky](https://bsky.app/profile/pilato.fr/post/3mc2rpbl6w72h)

> David nous explique comment il a migré son site sur Hugo, comment il a créé un type de page pour ses talks (j'ai fait la même chose mais en moins aboutit), et partage quelques stats sur son expérience de speaker.

## 🧠 IA

* [The rise and fall of robots.txt](https://www.theverge.com/24067997/robots-txt-ai-text-file-web-crawlers-spiders) _via_ [theverge.com](https://theverge.com)

> Cet article relate l'histoire du fameux fichier `robots.txt`.
Aujourd'hui, à l'aire des IA qui crawlent le web, c'est bon de se rappeler cette époque de respect tacite.

* [Programmeur, un métier qui tend à disparaître depuis 40 ans](https://contretemps.azeau.com/post/programmeur-un-metier-qui-tend-a-disparaitre-depuis-40-ans/) _Par_ [Olivier Azeau](https://contretemps.azeau.com)

* [L'alignement de l'esprit importe plus que celui du code](https://www.emaxilde.net/posts/2020/08/01/l-alignement-de-l-esprit-importe-plus-que-celui-du-code.html) _par_ [Olivier Poncet](https://emaxilde.net)

> Olivier et Olivier (!) partagent à quelques jours d'écart des réflexions sur nos métiers et sur leur avenir de nos métiers, à l'ère de l'IA.
> Après avoir vécu les transformations des langages, du low-code, du no-code et maintenant de l'IA-Gen, on peut être convaincu que notre métier ne va pas disparaître cette fois-ci non plus, mais encore une fois être transformé.
>
> Ce que je résume souvent en "Écrire le code ce n'est pas le problème le plus dur de mon métier".

## 🛜 Internet

* [IFTTD 341 - Bilan 2025 - IA, productivité et rupture : nouvelle donne pour les développeurs ?](https://www.ifttd.io/episodes/bilan-2025) _via_ [Quentin Adam](https://bsky.app/profile/waxzce.org/post/3mbwkwgtlc625)

> Comme tous les ans, Quentin donne son avis sur l'année écoulée, et se projette sur l'année suivante.
L'IA est forcément au cœur de la conversation, mais aussi la géopolitique actuelle et la souveraineté. 

* [Comment j'ai viré Algolia et recréé le Google de 1998 sur mon site](https://korben.info/pagefind-recherche-statique-hugo-algolia-alternative.html) _par_ [Korben](https://korben.info)

> Un des problèmes casse-pied avec les sites statiques, c'est l'implémentation d'un moteur de recherche interne.
J'avais expérimenté un peu avec [fusejs](https://www.fusejs.io/), mais je l'ai décâblé, car je n'étais pas satisfait des résultats.
Korben présente PageFind et comment il l'a mis en place sur son site.
À tester quand j'aurai envie de remettre un moteur de recherche (quand j'aurai plus de contenu).

## 🖥️ Shell

* [Scripts I wrote that I use all the time](https://evanhahn.com/scripts-i-wrote-that-i-use-all-the-time/) par [Evan Hahn](https://evanhahn.com/) _via_ 
  [Anthony Pena](https://anthonypena.fr/2025/12/02/revue-de-presse-decembre/)

> Une collection incroyable de scripts shells. C'est une bonne source d'inspiration pour automatiser de petites tâches courantes (comme le `mkdir && cd`). L'auteur partage tous ses scripts dans un [repo Git](https://codeberg.org/EvanHahn/dotfiles). Pépite.

## ☕ Java

* [Command completion: No more shortcuts!](https://www.youtube.com/watch?v=waY6HAmyHOw)

> Un nouvel accès aux commandes de l'IDE IntelliJIDEA. Il suffit de taper ".." n'importe où dans l'éditeur.
Je suis plutôt hypé par cette feature, ça va être bien pratique, pour éviter de retenir que pour refactorer un bout de code il faut faire CTRL+ALT+SHIFT+T (c'est le vrai raccourci)

* [Flaky Tests: a journey to beat them all](https://foojay.io/today/flaky-tests-a-journey-to-beat-them-all/) par [Loic Mathieu](https://loicmathieu.fr) _via_ [Bluesky](https://bsky.app/profile/foojay.io/post/3mccoqlcmw223)

> Le pote Loïc nous parle des tests "flaky" (qui ne passent pas tout le temps), et nous donne des pistes concrètes pour améliorer les choses.
>
> J'avais vu une conférence à DevOxx France 2025 sur le même sujet, à voir aussi [Flakiness : Quand tester, c’est vraiment douter](https://www.youtube.com/watch?v=CFXMYZmXfAc) : de Maxime CLEMENT

* [Stepping down as Mockito maintainer after 10 years](https://github.com/mockito/mockito/issues/3777) _via_ [Erik C. Thauvin (@erik.thauvin.net)](https://bsky.app/profile/erik.thauvin.net/post/3mb3ab4qoss2h)

> Une news qui remet encore en avant le statut des maintainers de projets Open Source. Ici un contributeur de Mockito déclare arrêter de travailler sur la librairie et passer la main.

## 🐧 Linux

* [Fingerprint reader support](https://fprint.freedesktop.org/)

> Sur mon nouveau laptop, j'ai un lecteur d'empreintes digitales, et je ne savais pas dans quelle mesure j'allais pouvoir l'utiliser sous Linux, et ça marche plutôt pas mal en fait. `fprint` fait bien le job, couplé au `pam_fprintd` pour coupler l'outil à l'authentification. J'en reparlerai dans mon article sur Pop!_OS 24.04.

---

La prochaine publication est prévue autour du 16 janvier 2026 🗓️

Photo de couverture par [Filip Bunkens](https://unsplash.com/@thebeardbe?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/road-covered-by-snow-near-vehicle-traveling-at-daytime-R5SrmZPoO40?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)