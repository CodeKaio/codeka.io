---
date: "2025-03-14"
language: fr
tags:
- IA
- Java
- Kubernetes
- Internet
- Sécurité
- Linux
title: La veille de Wittouck - Début Mars 2025
---

Je vais essayer dans les prochaines semaines de publier les liens que je découvre lors de ma veille techno courante, dans le cadre de mon boulot ou depuis mon canap' 🛋️

Le rythme de publication sera probablement bi-mensuel : un rythme hebdomadaire me semble trop compliqué à tenir et un rythme mensuel peut-être un peu trop long.

Cette démarche a deux buts principaux : partager ce que je découvre et décharger mon esprit (et mes onglets Firefox 😅)

Les liens seront classés par thèmes avec un avis ou commentaire.
Je vais aussi essayer de citer la source de ma découverte à chaque fois (_via_), afin que vous puissiez aussi vous abonner directement à ces sources.

Parmi les liens que vous allez retrouver (mais pas que) :

* des articles sur l'informatique et internet en général
* les liens vers les releases des outils et produits que j'utilise
* les posts LinkedIn que je trouve intéressants

## 🛜 Internet

* [Publish on your Own Site, Syndicate Elsewhere](https://indieweb.org/POSSE) vu sur [Bluesky](https://bsky.app/profile/christopheml.fr/post/3ljspbggvok2d) _via_ Christophe Michel.

> J'aime bien l'approche et la théorisation. C'est ce que j'essaye de faire sur ce blog. Je vais probablement aller plus loin sur ces concepts dans les mois à venir.

* [Pourquoi les jeunes sont devenus si nuls en informatique](https://www.mac4ever.com/societe/187334-pourquoi-les-jeunes-sont-devenus-si-nuls-en-informatique) par Nicolas Sabatier, _via_ [Bluesky](https://bsky.app/profile/nolotec.bsky.social/post/3litzcyzrbx26).

> Un article qui retrace le lien particulier avec l'informatique des gens de ma génération. Les perspectives, couplées au succès des IA génératives, sont quand même assez inquiétantes.

## 🤖 IA

* [L'extension Firefox & Chrome de Next.ink permet de détecter les médias générés par IA](https://next.ink/171509/alerte-genia-notre-extension-debarque-sur-firefox-android-et-identifie-2-000-sites/) _via_ [next.ink](https://next.ink/)

> Hâte d'en voir le code source. C'est principalement une liste de sites web, mais ça peut bien être utile pour savoir à quoi s'attendre lors d'une navigation.

* [New Junior Developers Can’t Actually Code](https://nmn.gl/blog/ai-and-learning) _via_ Siegfried Ehret sur [ehret.me](https://ehret.me/news-from-last-month/202503-developer).

> Nous avons fait ce même constat avec certains collègues de l'Université de Lille en observant les étudiants. Je pense que nous (enseignants) devons travailler à recentrer l'usage des IA génératives et limiter le LLM-Driver-Development.

## 🐧 Linux

* [You Should Use /tmp/ More](https://atthis.link/blog/2025/58671.html) _via_ Siegfried Ehret sur [ehret.me](https://ehret.me/news-from-last-month/202503-developer).

> Ok, je suis convaincu `/tmp` devient mon nouveau répertoire de `Downloads` par défaut. Pourquoi personne ne l'avait suggéré avant ? J'ai aussi envie d'y rediriger mes `target` et autre ! En plus le répertoire `/tmp` est probablement géré intégralement en RAM, sans écriture disque avec le `tmpfs`, donc ça augmenterait la durée de vie de mon SSD

## ☕ Java

* Mon pote Guillaume Dufrêne retrace l'histoire de Java sur LinkedIn, avec le contexte historique IT associé :
  * [Java 1.2](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_java-activity-7302617575677480961-Z9ZO?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.3](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_javaone-java-daezveloppement-activity-7302979942772822018-uho5?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.4](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_javaone-java-daezveloppement-activity-7302979942772822018-uho5?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.5](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_maven-svn-ubuntu-activity-7303704720794083328-1Lkj?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.6](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_soap-rest-jpa-activity-7304791908348801024-nKXS?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.7](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_android-wikeo-spring-activity-7305154301553766401-4T7q?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)
  * [Java 1.8](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_ruby-rails-prestashop-activity-7305516664895262720-5wM9?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)

> Ça rappelle des souvenirs, et le contexte remet un peu les fonctionnalités du langage dans l'histoire, et explique le succès de Java au fur et à mesure des années !

## 🏷️ Releases

 * k9s en version [0.40.7](https://github.com/derailed/k9s/releases/tag/v0.40.7)

> Une nouveauté sympa, la personnalisation des colonnes affichées disponible depuis la version 0.40.0.

## 🎫 Évènements

* La programmation de DevOxx France 2025 est dispo : https://www.devoxx.fr/agenda-2025/schedule/

> Beaucoup de belles confs, je partagerai mon agenda dans la prochaine édition.

* DevFest Nantes a publié son [pack de sponsoring](https://docs.google.com/presentation/d/14CUEDTIynRDDbDzWFeYzicSSsrOyz4zoJgotskWm0X8/edit#slide=id.g33ca69c931d_0_289)

> Nous allons nous inspirer de ce pack pour rédiger celui de Cloud Nord.

## 🔒 Sécurité

* [Panorama de la cybermenace 2024](https://cyber.gouv.fr/publications/panorama-de-la-cybermenace-2024) par l'ANSSI _via_ [cyber.gouv.fr](https://cyber.gouv.fr)

> Je suis tombé dessus un peu par hasard. Ce doc est très intéressant et il est construit en 3 volets : les opportunités, les moyens et les finalités.

* [Recommandations de sécurité pour l'architecture d'un système de journalisation](https://cyber.gouv.fr/publications/recommandations-de-securite-pour-larchitecture-dun-systeme-de-journalisation) par l'ANSSI _via_ [cyber.gouv.fr](https://cyber.gouv.fr)

> J'ai étudié ce doc dans le cadre de ma mission actuelle et de ma veille. Il est très orienté autour des journaux de sécurité (authentification, accès, évènements systèmes) et moins autour des journaux applicatifs, mais ça reste intéressant à lire.

---

La prochaine publication est prévue autour du 26 mars 🗓️

* Photo de couverture par [Ran Berkovich](https://unsplash.com/@berko?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) sur [Unsplash](https://unsplash.com/photos/close-up-selective-focus-photo-of-black-binoculars-kSLNVacFehs?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
      