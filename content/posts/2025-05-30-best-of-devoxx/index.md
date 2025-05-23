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
    watchedVideos: 7
---

J'ai regardé pas loin de {{% param watchedVideos %}} videos de la [playlist de Devoxx France 2025](https://www.youtube.com/playlist?list=PLTbQvx84FrATiYy0se8yoHJHicXtmDbB-).
Mon astuce (pour ne pas y passer 300 heures) : je regarde les vidéos en x2, et je ralentis sur les morceaux intéressants.

Cette édition spéciale de "La veille de Wittouck" liste donc les vidéos qui m'ont le plus intéressé. 

<!--more-->

* [Aerospike comme Key/Value Store à l'échelle - Retour d’expérience de Criteo](https://www.youtube.com/watch?v=eeaFt3qnoHk) : Présentation par : Nicolas Wlodarczyk (Aerospike), Peter Goron (Criteo)

> Le use case est d'un autre level. Temps de réponse moyen attendu à 1 ms. 70 milliard de records. 18M req/secondes.
> Utilisation optimisée du page cache linux. Utilisation de disques en raw (sans filesystem).
> On y découvre aussi les challenges rencontrés en production dans le volet REX.

## 🛜 Internet

## 🔒 Sécurité

## ☕ Java

* [AssertJ-DB : Validez vos opérations en bases de données avec style](https://www.youtube.com/watch?v=XILu4r3rIEc) : Présentation par : Julien Roy (Perspecteev)

> Une lib pour tester le contenu d'une base de données en Java. Je connaissais déjà DbUnit sur ce use-case, qui était un peu fastidieux à utiliser. AssertJ-DB s'intègre avec AssertJ, et semble plutôt pratique d'utilisation, avec un apprentissage rapide.
> La killer feature est la possibilité de capturer les changements effectués sur la base de données par le code.

## ☕ IA

## 👷 DevOps

* [Comment allonger notre build nous a fait gagner du temps ?](https://www.youtube.com/watch?v=Fzd45CaQBdk) : Présentation par : Éric LE MERDY (QuickSign), Vincent Galloy (QuickSign)

> Les speakers détaillent l'ensemble des règles et étapes qu'ils ont intégrés dans leur build gradle, avec une estimation du temps gagné par rapport au temps de build ajouté. Outre les règles de formatage, et de validation d'architecture qu'on retrouve souvent, je retiens 2 bonnes idées : le contrôle de l'OpenApi généré (avec un diff automatisé) avec `openapidiff`, la suppression des dépendances inutiles avec du code inspiré de `maven-dependency-analyser`. En définitive, ajouter ces étapes au build plutôt qu'à la CI suit les approches de type _shift-left_, comme on peut aussi les observer avec des outils comme Dagger.

* [Flakiness : Quand tester, c’est vraiment douter](https://www.youtube.com/watch?v=CFXMYZmXfAc) : Présentation par : Maxime CLEMENT (Docker)

> Les _flaky_ tests, qui échouent parfois sont la bête noire des pipelines. On relance le pipeline en croisant les doigts 🤞. Docker a développé un outil interne pour détecter et marquer les tests _flaky_, afin de pouvoir les ignorer dans les pipelines. L'impact observé est de 300 PR mergées du premier coup en plus / mois et 1 an de compute économisé / mois. 

* [{Cloud || Container} Development Environment](https://www.youtube.com/watch?v=Tl9JplAAVB8) : Présentation par : Jean-Philippe Baconnais (Zenika), Benjamin Bourgeois (Zenika)

> Un tour d'horizon des différents moyens pour avoir un environnement de travail déporté : DevContainers, GitPod, JetBrains Gateway, et Cloud Workstation de Google.

## ☸️ Kubernetes

* [J'ai perdu du poids sur Kubernetes avec SlimFaas](https://www.youtube.com/watch?v=sBvBvoB-Cbs) : Présentation par : Guillaume Chervet (AXA France)

> SlimFaas est un système de FaaS pour Kubernetes. La démo est plutôt impressionnante, en quelques minutes, le système est installé et utilisable. Les intérêts principaux sont la capacité à _scale_ les applications à 0 et la possibilité de scheduler le démarrage et l'extinction des pods.

## 🐋 Docker

## 🐧 Linux

## 🏷️ Releases

## 🎫 Évènements

---

La prochaine publication est prévue autour du  🗓️ 13 juin 2025

Photo de couverture par sur 