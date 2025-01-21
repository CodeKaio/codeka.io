---
title: Talk - Mais au fait, ça marche comment les service accounts ?
language: fr
date: 2024-04-19
tags:
  - DevOps
events: 
  - name: DevOxx 2024
    date: 2024-04-19
    youtube: 7p9NTgkOy_w
    slides: slides.pdf
---
L'utilisation des comptes de service (service accounts) sur les machines virtuelles et dans les conteneurs est un basique de la sécurité Cloud. Ces comptes ouvrent un accès vital aux multiples API offertes par les fournisseurs de services Cloud.

Quelle magie noire se cache derrière leur fonctionnement ? Comment une application, démarrée dans un conteneur et déployée sur un cluster Kubernetes, parvient-elle à exploiter les API Cloud en utilisant les comptes de service, sans nécessiter de configuration particulière ?

L'objectif de cette présentation est de démystifier le fonctionnement des comptes de service. En analysant le code source des SDK de Google et d'AWS, et en s'appuyant sur des démonstrations pratiques sur des machines virtuelles et des conteneurs, nous verrons comment sont implémentées ces mécanismes, qui n'ont rien de magique !

Nous dissiperons également la magie qui se dissimule derrière des concepts tels que "Workload Identity" dans Google Kubernetes Engine (GKE) et "IAM Roles for Service Accounts" chez AWS.
Rejoignez-nous pour une exploration passionnante des coulisses du Cloud !