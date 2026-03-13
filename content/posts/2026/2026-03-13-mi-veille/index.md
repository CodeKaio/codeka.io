---
date: 2026-03-13
language: fr
title: La veille de Wittouck - Début mars 2026
series: La veille de Wittouck
tags:
  - cloud
  - events
  - ia
  - tools
  - linux
  - docker
---

En ce début de mois de mars, les sujets de souveraineté ont encore fait le buzz. Au-delà des annonces (parfois décevantes), il faut prendre un peu de recul. L'excellent article de Katia donne des pistes concrètes pour aider à élaborer une souveraineté stratégique.

J'ai aussi beaucoup consommé de vidéos et de podcasts les dernières semaines, donc je vous ai listé celles qui m'ont plu, en particulier le nouveau podcast "À la french", très bien réalisé, avec un sujet cool sur "La Suite" développée par la DINUM.

<!--more-->

## ☁️ Cloud

* [Bye bye GAFAM? Les techs de l'Etat développent la Suite Numérique!](https://www.youtube.com/watch?v=aWCk_zhNeQ0) par [À la french](https://www.youtube.com/@alafrench)

> Cet épisode du podcast "À la french" parle du projet "La Suite", l'offre de service numériques à destination des agents du service public. Samuel, le CTO aux commandes, explique la genèse du projet, les différents éléments de La Suite: ProConnect, Visio, Docs, Tchap.
> 
> On y découvre aussi le fonctionnement des services publics, et l'organisation entre les différentes entités.
> 
> C'est parfois technique, mais toujours pédagogique, à ne pas manquer.

* [La réversibilité pour garantir sa véritable souveraineté cloud](https://blog.cockpitio.com/cloud/reversibilite-souverainete-cloud/) par Katia HIMEUR _sur_ [blog.cockpitio.com](https://blog.cockpitio.com)

> Alors que tout le monde s'intéresse à la souveraineté (et c'est pas pour me déplaire), Katia pose les bases de la réversibilité et de la portabilité.
> Sans surprise, les clés principales sont de s'appuyer sur des standards open-source pour éviter le vendor lock-in, s'appuyer sur des architectures hexagonales avec des abstractions, et documenter les stratégies de sorties des offres propriétaires.

## 🎫 Events

* [Dans les coulisses de la sélection des talks à Devoxx France](https://www.youtube.com/watch?v=sO9zuxfz5AU)

> Une vidéo dans laquelle Arnaud, "Chef du CFP" de DevOxx France, répond aux questions d'Emmanuel Bernard. Il donne des infos intéressantes sur les coulisses du CFP et sur le processus de sélection.
> Avec 1000 propositions à relire, le travail est colossal, et 20 à 25 personnes participent à la relecture.
> 
> Arnaud donne aussi beaucoup de conseils sur ce qui fait une bonne réponse à un CFP.

* ["Avoir des nouveaux, ça apporte un air frais" avec Camille Dagbert et Maxime Derroullers - LA VOIX DE L'INNOVATION](https://www.youtube.com/watch?v=oKYWxKuBooM) par [Ekité](https://www.youtube.com/@EKIT3)

> Mon pote Romain a reçu en interview Camille et Maxime, qui co-organisent avec moi l'événement Cloud Nord.
> Ça parle des dessous de l'orga et des changements qu'on a opérés cette année.

## 🧠 IA

* [La Fabrique à Idiots](https://www.youtube.com/watch?v=4xq6bVbS-Pw&t=11s) par [Micode](https://www.youtube.com/@Micode)

> L´IA nous rend-elle bêtes ? Copier/coller du code (ou une rédaction de Français) ne nous rend pas plus intelligents, ça on est plutôt d'accords.
> 
> Sans être alarmiste (bon, en fait si, un peu), Micode explique comment fonctionne le processus d'apprentissage, et le comportement de notre cerveau. Il propose ensuite une façon d'utiliser l'IA pour en faire un tuteur personnel, et de conserver un processus d'apprentissage complet. Une approche plutôt intéressante.

## 🛠️ Tools

* [Slashgear/gdpr-cookie-scanner: CLI tool to scan websites for GDPR cookie consent compliance](https://github.com/Slashgear/gdpr-cookie-scanner)

> Le pote Antoine Caron nous gratifie un petit outil, très bien conçu, pour valider la bonne conformité d'un site en regard du RGPD.
> L'outil génère un rapport avec des points de contrôle, principalement orientés autour du bon respect des banières de cookie, du consentement, etc.
> 
> Je vais devoir creuser la compliance de mon site, car même si j'obtiens la note de 💯, il me reste des questions par rapport au fait que Plausible, que j'utilise, reste un third-party.

* [IT Tools - Collection of handy online tools for developers, with great UX](https://it-tools.tech/) par [Corentin Thomasset](https://corentin.tech/) _via_ [The New Stack (@thenewstack.io)](https://thenewstack.io/it-tools-brings-many-useful-developer-tools-into-one-convenient-location/)

> IT-Tools est une collection de plein de petits outils qu'on utilise au quotidien, regroupés dans une unique page web.
> 
> On y retrouve en vrac, de quoi générer des tokens, mots de passe et UUIDs, de quoi hasher du texte avec différents algos, chiffrer et déchiffrer du texte, convertir du JSON en TOML, XML ou  YAML (et inversement), parser un JWT, prettyfier ces formats, etc.
> En tout, il y a presque 90 petits outils mis à dispo, et on peut self-host tout ça avec un simple container docker, le code est dispo sur [Github](github.com/CorentinTh/it-tools).

## 🐧 Linux

* [Open Source Masterclass : un MOOC pour se lancer dans la contribution au logiciel libre](https://linuxfr.org/news/open-source-masterclass-un-mooc-pour-se-lancer-dans-la-contribution-au-logiciel-libre) _via_ [LinuxFr](https://linuxfr.org/)

> Contribuer à un projet open-source, ce n'est pas juste ouvrir Claude Code, taper un prompt, et ouvrir une PR.
> Ce MOOC gratuit propose de découvrir la philosophie des projets FLOSS (Free/Libre and Open Source Software).
> 
> Je ne l'ai pas encore suivi, mais ça a l'air très bien fait, c'est poussé par Framasoft et Télécom Paris entre autres.

## 🐋 Docker

* [Migration d'un conteneur docker entre deux hôtes Linux](https://domopi.eu/migration-dun-conteneur-docker-entre-deux-hotes-linux/) _via_ [Le Journal du Hacker](https://www.journalduhacker.net/)

> Utiliser des containers ne rend pas une application portable.
> Déplacer un container d'une machine à une autre peut être une tanée, surtout si on a des volumes. Guillaume nous explique la procédure qu'il a menée. C'est très complet, avec tout le détail des commandes exécutéess
> 
> Je me demande quand même s'il n'y a pas des outils qui sont capables de faire ce genre de migration (avec un bout de swarm, par exemple).

---

La prochaine publication est prévue autour du 27 mars 🗓️

Photo de couverture par [Christian Lue](https://unsplash.com/@christianlue?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/white-red-and-green-map-7dEyTJ7-8os?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)

