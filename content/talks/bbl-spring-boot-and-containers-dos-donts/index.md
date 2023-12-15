---
title: BBL - Spring Boot & Containers - Do's & Don'ts
language: fr
date: 2023-04-27
tags:
  - DevOps
  - Spring Boot
  - Docker
  - Kubernetes
events: 
  - name: Axa Web Center Lille
    date: 2023-04-27
  - name: Ippon Lille
    date: 2023-05-15
  - name: ADEO Tech Bubble
    date: 2023-06-07
slides: bbl-spring-boot-and-containers-dos-donts.pdf
---
Bien que Docker soit facile d'utilisation, construire une image reste un exercice compliqué.

Optimisation, layers, configuration, haute disponibilité et sécurité nécessitent de s'abstraire des Dockerfile basiques qu'on peut trouver sur internet.

Dans ce talk, nous verrons comment bien packager une application Spring Boot dans une image Docker/OCI.

À travers 42 bonnes pratiques, nous allons voir :

* les bonnes pratiques préconisées par Spring 🍃
* les pièges à éviter 👿
* comment bien préparer votre application Spring Boot pour un exécution dans un container 📦
* comment bien écrire un Dockerfile pour optimiser la construction et l'image finale 📝
* des moyens alternatifs de construction d'images Spring Boot 🏗
* et l'outillage d'analyse à ajouter à votre toolbox 🔧