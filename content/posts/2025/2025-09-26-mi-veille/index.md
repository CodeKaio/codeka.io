---
date: 2025-09-26
language: fr
title: La veille de Wittouck - Fin Septembre 2025
slug: la-veille-de-wittouck-fin-septembre-2025
series: La veille de Wittouck
tags:
  - Databases
  - Git
  - Internet
  - Java
atUri: "at://did:plc:a27wdjlmq3ebx4v5f2jpzvsk/site.standard.document/3mucgx7hf2a2o"
---

Enfin ! La version 25 de Java, LTS tant attendue par les fans (moi le premier 😅), est arrivée.
PostgreSQL 18 est également sorti cette semaine, donc [La veille de Wittouck]({{< relref "/series/la-veille-de-wittouck">}}) s'attarde particulièrement sur ces deux événements.

<!--more-->

## 🛜 Internet

* [I love UUID, I hate UUID](https://blog.epsiolabs.com/i-love-uuid-i-hate-uuid) par [Maor Kern](https://hashnode.com/@maor10) _via_ [Ludovic Borie](https://bsky.app/profile/b0dul.bsky.social/post/3lykf4k2jne2u)

> Une excellente lecture sur l'utilisation des UUID, en particulier comme identifiants dans les bases de données, et les gains apportés par UUIDv7.
> Cette version des UUID n'est pas encore supportée par Java, mais la librairie [java-uuid-generator](https://github.com/cowtowncoder/java-uuid-generator) l'implémente.

* [you do not have to use generative ai "art" in your blogs because there are websites where you can get real, nice images for free](https://livelaugh.blog/posts/non-ai-images-for-blogs-websites/) par [Jenn Schiffer](https://livelaugh.blog/) _via_ [Cassidy Williams](https://buttondown.com/cassidoo/archive/its-important-to-have-something-to-walk-towards/)

> Un article qui liste 14 sites mettant à disposition des images gratuitement. J'utilise [Unsplash](https://unsplash.com/) depuis longtemps pour illustrer mes articles, c'est pratique d'avoir d'autres sources sous la main.

## ☕ Java

* [Java 25 : Quoi de neuf ?](https://www.loicmathieu.fr/wordpress/fr/informatique/java-25-quoi-de-neuf/)

> Mon pote Loïc Mathieu publie un excellent billet de blog accompagnant chaque release de Java.
> C'est donc l'occasion cette semaine d'aller (re)lire son analyse sur les fonctionnalités de cette nouvelle version, avec ce qui va particulièrement intéresser les développeurs.

* [OpenJDK JDK 25 General-Availability Release](https://jdk.java.net/25/)

> Java 25 est officiellement disponible depuis le 16 septembre !
> La [release note](https://jdk.java.net/25/release-notes) liste l'ensemble des JEP embarquées, mais allez plutôt lire l'article de Loïc 😊

* [GraalVM 25](https://www.graalvm.org/downloads/)

> La release de l'OpenJDK est accompagnée de celle de GraalVM. Les images de container embarquant l'outil _native-image_ sont disponibles sur [Oracle Container Registry](https://container-registry.oracle.com/ords/ocr/ba/graalvm).
> L'équipe de GraalVM n'a pas encore publié d'article pour accompagner cette release, cela risque d'arriver dans les prochains jours.

* [Temurin JDK 25](https://adoptium.net/)

> Oracle fourni le build officiel d'OpenJDK, mais uniquement pour un nombre limité d'architectures (et pas de JRE).
> La distribution Temurin de Java, buildée par la fondation Eclipse avec le projet Adoptium, a produit les builds de Java 25 pour tous les OS et architectures possibles, y compris Alpine (essentiel aujourd'hui dans les containers), et des architectures RISC et PowerPC (oui ça existe encore).
> Les images OCI (JRE ou JDK) sont disponibles sur [hub.docker.com](https://hub.docker.com/_/eclipse-temurin), avec des base image Alpine, Ubuntu, ou RHEL Ubi.
> Temurin est aussi disponible sous la forme de packages pour les distributions Linux (rpm, apk ou deb), ou en tant que package MSI pour Windows.

* [9 most-watched Java conference talks of 2025 (so far)](https://www.reddit.com/r/java/comments/1nk4wqu/9_mostwatched_java_conference_talks_of_2025_so_far/) par [Tech Talks Weekly](https://www.techtalksweekly.io/) sur [Reddit](https://www.reddit.com/user/TechTalksWeekly/)

> Je suis tombé un peu par hasard sur ce post Reddit, qui liste les conférences intéressantes à regarder sur Java depuis le début de l'année.
> Leur newsletter est intéressante, je m'y suis abonné (en espérant y voir apparaître un de mes talks un jour 😅).

## 🔀 Git

* [Git worktrees for fun and profit](https://blog.safia.rocks/2025/09/03/git-worktrees/) par [Safia Abdalla](https://blog.safia.rocks/) _via_ [Bluesky](https://bsky.app/profile/captainsafia.com/post/3lxxc2jdjlc2z)

> Une approche originale permettant de travailler sur de multiples branches en parallèle. On en apprend tous les jours sur Git !

## 💾 Databases

* [PostgreSQL 18](https://www.postgresql.org/about/news/postgresql-18-released-3142/)

> La version 18 de ma base de données préf est disponible. Je suis particulièrement hypé par le [support de OAuth 2.0](https://www.postgresql.org/docs/18/auth-oauth.html). J'ai hâte de tester ça avec un petit serveur Keycloak, et de voir ce que les fournisseurs de Cloud vont faire de cette feature.

* [Get Excited About Postgres 18](https://www.crunchydata.com/blog/get-excited-about-postgres-18) par [Elizabeth Christensen](https://x.com/sqlliz) _via_ [Nicolas Fränkel](https://www.linkedin.com/posts/nicolasfrankel_get-excited-about-postgres-18-crunchy-data-activity-7372909485234876417-Cozi)

> Un article qui rentre un peu dans le détail de cette nouvelle version de Postgres.

* [Postgres for Everything](https://postgresforeverything.com/) _via_ [Ludovic Borie](https://bsky.app/profile/b0dul.bsky.social/post/3ly5srjh3lx24)

> Caching, OLAP, JSON, on savait déjà que PostgreSQL avait de multiples capacités.
> Cette page liste tout ce que peut faire notre base de données préf. Il y a des cas d'usage très surprenants, comme l'indexation de documents PDF ou la manipulation de SVG.
> J'ai aussi découvert un build WASM nommé [PGlite](https://github.com/electric-sql/pglite), qui peut donc s'exécuter dans un browser !

## 🎫 Évènements

* [DevFest Nantes 2025](https://devfest.gdgnantes.com/)

> Le plus grand DevFest de France revient cette année avec 2 jours de confs, pas moins de 108 speakers pour 76 talks et ateliers.
> Mon préf : [Let's play Factorio](https://devfest.gdgnantes.com/sessions/let_s_play_factorio) (🫣)

---

La prochaine publication est prévue autour du 🗓️ 17 octobre (je parlerai de mon expérience Nantaise 🐘)

Photo de couverture par [Matt](https://unsplash.com/@matty10?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) sur [Unsplash](https://unsplash.com/photos/white-ceramic-mug-with-brown-liquid-2mpcN1dYPbQ?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)