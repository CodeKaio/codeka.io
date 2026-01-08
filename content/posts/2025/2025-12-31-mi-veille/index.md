---
date: 2025-12-31
language: fr
title: La veille de Wittouck - Fin décembre 2025
series: La veille de Wittouck
tags:
- internet
- java
- kubernetes
- linux
- security
- spring-boot
---

Pour la dernière édition de l'année de "la Veille de Wittouck", pas mal d'articles intéressants qui m'ont donné envie d'expérimenter.
J'ai déjà pas mal bidouillé Spring Boot 4 avec Open Rewrite et j'ai deux articles en cours de rédaction, sur Kubernetes 1.35 et le redimensionnement des ressources d'un pod à chaud, et un sur Pop!_OS @24.04 avec COSMIC.


<!--more-->

## 🛜 Internet

* [Split Keyboards Are Superior And The Reason I'm The Writer I Am Today](https://aftermath.site/best-split-keyboards-diy-qmk-zmk-corne/)  _via_ [Molly White (@molly.wiki)](https://bsky.app/profile/molly.wiki) sur [Bluesky](https://bsky.app/profile/molly.wiki/post/3mb7v3fhcxc2b)

> Un article complet sur les split keyboards. L'auteur y présente le concept, revient sur les keymaps, la gestion de layers, et l aspect DIY (Do It Yourself) avec le choix des switches et des keycaps.
> J'ai d'ailleurs monté mon propre split keyboard, il y a quelques mois, l'article est ici : [Montage d'un clavier mécanique](/2025/07/25/montage-dun-clavier-mécanique)

* [Firefox se dote d’une vue partagée pour afficher deux sites côte à côte - Next](https://next.ink/brief_article/firefox-se-dote-dune-vue-partagee-pour-afficher-deux-sites-cote-a-cote/?utm_source=dlvr.it&utm_medium=bluesky) _via_ [next.ink](https://next.ink)

> Une feature très cool. Vu que beaucoup de sites sont optimisés pour mobile, on se retrouve souvent avec deux bandes inutilisées sur les côtés, splitter en deux dans une même fenêtre peu être très pratique donc.

* [Write that blog!](https://writethatblog.substack.com/)

> Une newsletter qui donne des tips pour écrire son blog tech. On y trouve aussi des interviews de bloggers qui expliquent pourquoi ils/elles bloggent, et ce que ça leur a apporté.

* [Backing up Spotify](https://annas-archive.li/blog/backing-up-spotify.html) _via_ [Marme & lade (@marmelade.bsky.social)](https://bsky.app/profile/marmelade.bsky.social) sur [Bluesky](https://bsky.app/profile/marmelade.bsky.social/post/3mahhuoqp5223)

> J'ai vu passer cette info et je ne sais pas quoi en penser. Au-delà de l'acte "pirate" en lui-même et de l'usage qui pourrait être fait de ces données (faire apprendre une IA ?), il semble y avoir une réelle démarche de conservation.
Quoi qu'il en soit, l'article détaille des informations plutôt intéressantes et des graphes sur l'écoute des morceaux et l'algo supposé "aléatoire" de Spotify.

## 🌱 Spring Boot

* [Spring Boot: Chapitre 4](https://www.youtube.com/watch?v=oGZQF4qERFU) _via_ [le Ch´ti JUG](https://youtube.com/@chtijug)

> Le Java User Group des _Ch'tis d'un Ch'nord_ a accueilli Stéphane Nicoll, qui a présenté avec une longue démo les nouveautés de Spring Boot 4. J'étais présent, et j'ai vraiment trouvé cette session intéressante.
> Il prend le temps d'expliquer les nouveautés, la raison derrière certains changements, et comment migrer.

## ☕ Java

* [Considering raising the Java baseline to 25 for Micronaut 5](https://github.com/micronaut-projects/micronaut-core/discussions/12288) _via_ [Micronaut Framework (@micronautfw.bsky.social)](https://bsky.app/profile/micronautfw.bsky.social/post/3malpktlqd22t)

> Micronaut envisage une baseline en Java 25, contrairement à Spring qui reste sur Java 17. J'aime bien l'idée, ça force les entreprises à se mettre à jour lors des migrations.
> Ils expliquent vouloir se servir de certaines nouveautés du langage dans le framework, ce que ne fait pas encore Spring.

## ☸️ Kubernetes

* [Le premier intérêt de Kubernetes n'est pas le scaling](https://mcorbin.fr/posts/2025-12-29-kubernetes-scale/) par [mcorbin (@mcorbin.bsky.social)](https://bsky.app/profile/mcorbin.bsky.social/post/3mb4ftrjp522g)

> Un rappel essentiel de ce qu'est Kubernetes. On parle souvent de scalabilité, mais ça met de côté tous les autres aspects, qui sont ici bien rappelés.
> C'est bien tout ces aspects qui font de Kubernetes une infrastructure essentielle, et pas forcément overkill pour les applications.

* [Kubernetes 1.35: In-Place Pod Resize Graduates to Stable](https://kubernetes.io/blog/2025/12/19/kubernetes-v1-35-in-place-pod-resize-ga/)

> Redimensionner les ressources d'un pod ou d'un container à chaud (sans recréation du pod, mais avec redémarrage éventuel) est maintenant possible.
> C'était le disponible depuis un moment, mais c'est passé en GA (_General Availability_) en version 1.35. C'est une évolution importante de k8s. J'ai commencé à rédiger un article pour en étudier l'impact sur des applications Java.
> La prochaine feature intéressante sera l'intégration de ces fonctionnalités avec les VPA, c'est en cours de réflexion avec les [CPU Startup Boost](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/enhancements/7862-cpu-startup-boost) notamment.

## 🐧 Linux

* [The Linux Kernel is just a program](https://serversfor.dev/linux-inside-out/the-linux-kernel-is-just-a-program/) par [Zsolt Kacsandi](https://serversfor.dev) _via_ [Steve Klabnik (@steveklabnik.com)](https://bsky.app/profile/steveklabnik.com/post/3mapj6kf76k2d)

> Un court article qui revient sur le démarrage du Kernel Linux et une partie de son fonctionnement. Il est intéressant de voir comment on peut très facilement démarrer un Kernel dans une VM QEMU, et lancer un simple programme.
> Ça permet aussi de réfléchir au concept de distribution Linux.

* [GNOME et KDE ont pris un coup de vieux : pourquoi tout le monde parle de COSMIC ?](https://goodtech.info/cosmic-desktop-rust-pop-os-system76-linux-revolution/) _via_ [Le journal du Hacker](https://www.journalduhacker.net/)

> Je teste COSMIC depuis quelques jours, et je suis déjà séduit. Ça marche bien, c'est frais, efficace (tiling auto💙), et c'est joli. À voir sur la durée, j'aurai un avis plus poussé d'ici à quelques semaines.

## 🔒 Security

* [How I Used TPM for Key Encryption in Rust on Linux](https://dev.to/tsuruko12/how-i-used-tpm-for-key-encryption-in-rust-on-linux-hardware-tpm-vtpm-4ond) _via_ [ytroncal (@ytroncal.bsky.social)](https://bsky.app/profile/ytroncal.bsky.social/post/3matp2xjavc2x)

> Je ne connaissais pas bien la fonction des TPM, l'article détaille l'utilisation des API permettant de manipuler un TPM (hardware ou virtuel), créer une clé et la stocker dans le TPM, puis utiliser cette clé pour chiffrer et déchiffrer des données.
Je me demande dans quelle mesure des gestionnaires de mots de passe utilisent ces features.

---

La prochaine publication est prévue autour du 16 janvier 2026 🗓️

Photo de couverture par [Marisol Benitez](https://unsplash.com/@marisolbenitez?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/a-person-is-holding-a-sparkler-in-their-hand-Os8IuLEsSJU?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)