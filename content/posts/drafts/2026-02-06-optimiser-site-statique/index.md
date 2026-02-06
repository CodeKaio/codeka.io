---
date: 2026-02-06
title: Optimiser un site Hugo
draft: true
---

[//]: # (TODO link vers le blog d'antoine)
Sur les bons conseils du pote Antoine Caron, j'ai pris le temps d'optimiser un peu mon site.

Ce site que vous êtes en train de lire est un site statique, buildé avec Hugo.

J'ai déjà un peu travaillé la compression des différentes ressources, principalement les illustrations, mais je m'étais arrêté à ça.

## Le score Lighthouse

Pour faire un premier travail sur les performances de ce site, j'ai utilisé [une analyse LightHouse](https://pagespeed.web.dev/analysis/https-codeka-io/we5dukzmku?form_factor=desktop).

Lighthouse permet en quelques minutes d'avoir une vue des performances d'une application ou d'un site web, à la fois pour une cible _Desktop_ et _Mobile_.
Il permet aussi de valider certains propriétés d'accessibilité, comme des contrastes, la présence de texte alternatif pour les lecteurs d'écran, etc.

C'est, je pense, un bon point de départ.

Voici les scores de mon site à l'heure actuelle :


![Score Lighthouse pour un mobile](lighthouse-mobile.png)
![Score Lighthouse pour un desktop](lighthouse-desktop.png)


> J'ai clairement une marge d'amélioration sur l'accessibilité et les performances.

## Compression des images en webp

Une des actions que j'ai mis en place il y a un moment, est l'utilisation du format _webp_ pour compression les illustrations que j'utilise dans mes articles.

J'utilise souvent des photos que j'ai capturées avec mon smartphone (pour les articles de conférence), des schémas que je produits sur draw.io le plus souvent et que j'exporte en _png_, ou des photos _stock_ que je vais chercher pour illustrer mes articles de veille.

Ces photos sont souvent lourdes (plusieurs mégaoctets) et en haute résolution, et la première action simple consiste à redimensionner ces photo et les recompresser au format _webp_.

