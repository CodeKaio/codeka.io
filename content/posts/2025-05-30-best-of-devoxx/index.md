---
date: 2025-05-30
language: fr
title: La veille de Wittouck - Fin mai 2025 - Best-Of Devoxx 2025
series: La veille de Wittouck
tags:
  - DevOps
  - Docker
  - Git
  - IA
  - Internet
  - Java
  - Kubernetes
  - Linux
  - Security
  - Tools
draft: true
params:
    watchedVideos: 25
---

J'ai regardé pas loin de {{% param watchedVideos %}} videos de la [playlist de Devoxx France 2025](https://www.youtube.com/playlist?list=PLTbQvx84FrATiYy0se8yoHJHicXtmDbB-).
Mon astuce (pour ne pas y passer 300 heures) : je regarde les vidéos en x2, et je ralentis sur les morceaux intéressants 😅

Cette édition spéciale de "La veille de Wittouck" liste donc les vidéos qui m'ont le plus intéressé.
Je les ai classées selon les tags que j'utilise d'habitude. Il ne sont pas triés dans un ordre précis.

Je n'ai aussi pas listé les talks que j'avais déjà mentionné dans mon article précédent, je vous propose de les retrouver dans l'article [DevOxx 2025 - Bilan]({{< relref "/posts/2025-04-23-devoxx" >}}), que j'ai mis à jour avec les liens des vidéos.

<!--more-->

## 🛜 Internet

* [Retourner le cerveau de sa Gameboy et les attraper (vraiment) tous 💪](https://www.youtube.com/watch?v=7f-zajwDuNk) : Présentation par : Audren Burlot (Max Digital Services)

{{< youtube 7f-zajwDuNk >}}

> Une conférence sur l'histoire du développement des premiers jeux Pokemon, et du glitch qui permet de capturer Mew. Le speaker rentre dans le détail des modifications de valeurs de variables tout au long de l'exécution du glitch.
La conférence est hyper bien réalisée, avec des démos dans le jeu !

* [Une documentation de qualité pour avoir des logiciels opérationnels 📄⚙️](https://www.youtube.com/watch?v=_BUEJquIKbU) : Présentation par : Geoffrey Graveaud (Inside)

{{< youtube _BUEJquIKbU >}}

> Un talk autour de la gestion de la doc ! Le speaker donne les billes pour avoir une documentation de qualité, à travers trois critères : clarté, pertinence et accessibilité (au sens "facile à trouver"), et l'utilisation de revue de doc et de commentaires avec conventional comments.

* [Dans les coulisses des géants de la Tech !](https://www.youtube.com/watch?v=ulL4RLtJFzo) : Présentation par : Rachel Dubois (Bttrways)

{{< youtube ulL4RLtJFzo >}}

> La speakeuse y déconstruit les mythes des licornes, de la perfection qui est perçue de l'extérieur, l'illusion de l'agilité. Elle présente les approches produit concrètes de quelques licornes comme Spotify ou AirBnb, et l'importance des métriques techniques et produit, des tests et de l'acceptation de l'échec.

* [Les clés de l’architecture pour les devs](https://www.youtube.com/watch?v=ZoYDxF_7LoI) : Présentation par : Christian Sperandio (Arolla), CYRILLE Martraire (Arolla)

{{< youtube ZoYDxF_7LoI >}}

> "L'architecture, c'est le boulot de tout le monde". Les speakers donnent quelques clés d'architecture à destination des devs. Quelques tips simples comme "énoncer clairement les problèmes" à résoudre, et quels sont les quality attributes à mesurer dans une bonne architecture : disponibilité, performance, etc. Ils présentent aussi quelques points autour architecture monolithiques, modulaires, et des tradeoffs associés. Enfin, quelques points comme les ADR, les tests d'architecture et les principes de couplage et de cohésion.

## 🔒 Sécurité

* [Pragmatic OpenID Connect - Guillaume Chervet (AXA France)](https://www.youtube.com/watch?v=6coBmM5GO6o) : Présentation par : Guillaume Chervet (AXA France)

{{< youtube 6coBmM5GO6o >}}

> Le speaker présente le fonctionnement de OIDC, les différents types de tokens, les scopes, et les flow clients et serveur, avec des démos live. C'est à voir pour tous les devs, front ou back !

## ☕ Java

* [AssertJ-DB : Validez vos opérations en bases de données avec style](https://www.youtube.com/watch?v=XILu4r3rIEc) : Présentation par : Julien Roy (Perspecteev)

{{< youtube XILu4r3rIEc >}}

> Une lib pour tester le contenu d'une base de données en Java. Je connaissais déjà DbUnit sur ce use-case, qui était un peu fastidieux à utiliser. AssertJ-DB s'intègre avec AssertJ, et semble plutôt pratique d'utilisation, avec un apprentissage rapide.
La killer feature est la possibilité de capturer les changements effectués sur la base de données par le code.

* [Continuations: The magic behind virtual threads in Java](https://www.youtube.com/watch?v=iwuGgGxD9lA) : Présentation par : Balkrishna Rawool (Netherlands / ING Bank) (en anglais 🇬🇧)

{{< youtube iwuGgGxD9lA >}}

> Un talk qui explique le principe sous-jacent au fonctionnement des threads virtuels disponibles depuis Java 21, en particulier la gestion des scope et contextes d'exécution. L'API Continuation n'est pas sensée être utilisée en dehors des mécanismes internes du JDK, mais découvrir son fonctionnement permet de se rendre compte de la complexité qui se cache derrière les threads virtuels. La démo est incroyable, le speaker réimplémente un Generator, et un VirtualThread !

* [Booster le démarrage des applications Java : optimisations JVM et frameworks](https://www.youtube.com/watch?v=9k2HX9GUZRg) : Présentation par : Olivier Bourgain (Mirakl)

{{< youtube 9k2HX9GUZRg >}}

> Le speaker présente plusieurs axes d'optimisation : AOT et CDS, allocation de threads aux JIT C1 et C2, CRaC, et des éléments autour de l'infrastructure Kubernetes et allocation de ressources. C'est assez basique, mais ce sont toujours des bons rappels. J'aurais néanmoins aimé une démo de CRaC, car on en voit rarement.

* [30 ans d'Hello World en Java avec les JDK 1.0 à 24](https://www.youtube.com/watch?v=kZ8AMRcEBwg) : Présentation par : Jean-Michel Doudoux (Sciam)

{{< youtube kZ8AMRcEBwg >}}

> La légende JM Doudoux présente 34 formes de _Hello World_ en Java, du fameux `System.out.println("Hello World")` dans une méthode `main` à des moyens plus fous comme des blocs statiques, ou l'API Class-File pour générer dynamiquement du code.

## ☕ IA

* [IA Générative, TDD et Architecture Hexagonale : Une Synergie Révolutionnaire ?](https://www.youtube.com/watch?v=nL5L6KqGAxw) : Présentation par : Clément Virieux (Ippon), Florine Chevrier (Ippon)

{{< youtube nL5L6KqGAxw >}}

> Des stratégies concrètes pour imposer à une IA des standards de qualité issus du _craftmanship_. Ça passe par une approche de TDD, de test d'architecture et de maquettage, ces trois éléments étant passés au niveau du prompt de l'IA. L'approche semble efficace. Le plugin _CLine_ qui est présenté affiche aussi le coût financier de chaque requête. 

## 👷 DevOps

* [Comment allonger notre build nous a fait gagner du temps ?](https://www.youtube.com/watch?v=Fzd45CaQBdk) : Présentation par : Éric LE MERDY (QuickSign), Vincent Galloy (QuickSign)

{{< youtube Fzd45CaQBdk >}}

> Les speakers détaillent l'ensemble des règles et étapes qu'ils ont intégrés dans leur build gradle, avec une estimation du temps gagné par rapport au temps de build ajouté. Outre les règles de formatage, et de validation d'architecture qu'on retrouve souvent, je retiens 2 bonnes idées : le contrôle de l'OpenApi généré (avec un diff automatisé) avec `openapidiff`, la suppression des dépendances inutiles avec du code inspiré de `maven-dependency-analyser`. En définitive, ajouter ces étapes au build plutôt qu'à la CI suit les approches de type _shift-left_, comme on peut aussi les observer avec des outils comme Dagger.

* [Flakiness : Quand tester, c’est vraiment douter](https://www.youtube.com/watch?v=CFXMYZmXfAc) : Présentation par : Maxime CLEMENT (Docker)

{{< youtube CFXMYZmXfAc >}}

> Les _flaky_ tests, qui échouent parfois sont la bête noire des pipelines. On relance le pipeline en croisant les doigts 🤞. Docker a développé un outil interne pour détecter et marquer les tests _flaky_, afin de pouvoir les ignorer dans les pipelines. L'impact observé est de 300 PR mergées du premier coup en plus / mois et 1 an de compute économisé / mois. 

* [{Cloud || Container} Development Environment](https://www.youtube.com/watch?v=Tl9JplAAVB8) : Présentation par : Jean-Philippe Baconnais (Zenika), Benjamin Bourgeois (Zenika)

{{< youtube Tl9JplAAVB8 >}}

> Un tour d'horizon des différents moyens pour avoir un environnement de travail déporté : DevContainers, GitPod, JetBrains Gateway, et Cloud Workstation de Google.

* [Et si on faisait du simulation-driven development ?](https://www.youtube.com/watch?v=12LO_l90vDk) : Présentation par : Pierre Zemb (Clever Cloud)

{{< youtube 12LO_l90vDk >}}

> Le speaker présente le _Deterministic Simulation Testing_. Cela passe par du property-based testing, l'injection de chaos avec des mocks qui ajoutent des temps de latence aléatoires. Un test est associé à une seed qui permettra de rejouer les tests. C'est intéressant sur le principe, mais ça manque d'exemples concrets à mon sens.

## ☸️ Kubernetes

* [J'ai perdu du poids sur Kubernetes avec SlimFaas](https://www.youtube.com/watch?v=sBvBvoB-Cbs) : Présentation par : Guillaume Chervet (AXA France)

{{< youtube sBvBvoB-Cbs >}}

> SlimFaas est un système de FaaS pour Kubernetes. La démo est plutôt impressionnante, en quelques minutes, le système est installé et utilisable. Les intérêts principaux sont la capacité à _scale_ les applications à 0 et la possibilité de scheduler le démarrage et l'extinction des pods.

## 🐋 Docker

* [Questions pour un conteneur : Quiz sur les images, les conteneurs, OCI & Docker](https://www.youtube.com/watch?v=ZAucjf1mKOw) : Présentation par : Aurélie Vache (OVHcloud), Sherine Khoury (Red Hat)

{{< youtube ZAucjf1mKOw >}}

> Un talk interactif ! Le format est original ! Après chaque question, les speakeuses donnent une petite explication, avec une démo. On découvre des détails autour du fonctionnement des images et registries, des architectures, des SBOM

## 🐧 Linux

* [Aerospike comme Key/Value Store à l'échelle - Retour d’expérience de Criteo](https://www.youtube.com/watch?v=eeaFt3qnoHk) : Présentation par : Nicolas Wlodarczyk (Aerospike), Peter Goron (Criteo)

{{< youtube eeaFt3qnoHk >}}

> Le use case est d'un autre level. Temps de réponse moyen attendu à 1 ms. 70 milliard de records. 18M req/secondes.
Utilisation optimisée du page cache linux. Utilisation de disques en raw (sans filesystem).
On y découvre aussi les challenges rencontrés en production dans le volet REX.

*  [Quand le Terminal dévore la UI : TUI pour tout le monde !](https://www.youtube.com/watch?v=yEzKbvbOmTI) : Présentation par : Thierry Chantier (Dataiku)

{{< youtube yEzKbvbOmTI >}}

> Un peu d'histoire autour des _Terminal User Interface_, ainsi qu'une démo de développement d'un CLI TUI en Python, Golang ou Rust. Assez inspirant de voir comment il est facile d'implémenter de tels outils avec des librairies faciles d'accès.

---

La prochaine publication est prévue autour du  🗓️ 13 juin 2025

Photo de couverture par sur 