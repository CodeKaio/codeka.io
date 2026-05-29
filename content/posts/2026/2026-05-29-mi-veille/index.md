---
date: 2026-05-29
language: fr
title: La veille de Wittouck - Fin mai 2026
series: La veille de Wittouck
tags:
  - docker
  - git
  - ia
  - internet
  - kubernetes
  - sovereignty
  - tools
  - architecture
writing_time: 2h30
---

Entre deux matchs de Roland Garros, et pour prendre un peu l'air, pourquoi ne pas en profiter pour lire quelques articles intéressants (enfin, que _J'AI_ trouvés intéressants 😅).

Ce mois-ci dans "La Veille de Wittouck", plusieurs sujets autours de l'AT Protocol (je m'y intéresse depuis quelques semaines), mais aussi autour de `mise`, mon outil de dev préf du moment, et autour du _drama_ en cours sur GitHub.
Et en toile de fond, toujours aussi quelques articles autour de la souveraineté numérique.

<!--more-->




## 🔀 Git

* [The Missing GitHub Status Page](https://mrshu.github.io/github-statuses/)

> Une simple page de statut, qui collecte les incidents de GitHub et recalcule un temps d'_uptime_ prenant en compte l'intégralité des services de GitHub et les incidents associés.
> Le calcul peut sembler malhonnête, puisque cumuler l'ensemble des services dans un même indicateur ne reflête pas réellement l'usage de tous les utilisateurs.
> Néanmois, le regroupement affiche une info bien plus sombre que ce que communique GitHub sur sa propre [status page](https://www.githubstatus.com/).
> La vérité est probablement au milieu, en tout cas GitHub ne semble pas aller très bien.

* [Ghostty Is Leaving GitHub](https://mitchellh.com/writing/ghostty-leaving-github) par [Mitchell Hashimoto](https://mitchellh.com)

> Mitchell Hashimoto a annoncé souhaiter quitter GitHub pour héberger son projet Ghostty.
> L'annonce a fait pas mal de bruit (j'ai même entendu le terme "caprice"), conséquence des instabilités récentes et répétées de GitHub.
> Mais si un outil ou une plateforme ne marche pas ou ne convient pas, il est naturel d'en changer.
>
> Ici, il n'est pas impossible que son choix (pas encore communiqué) aie une influence sur le destin de GitHub, ou sur le destin de la plateforme qu'il choisira.
> Je l'ai spotté sur Tangled.org, où il a créé quelques repos, et où il a même démarré un nouveau projet (autour de la CI de Tangled).
> La lib Go pour Ghossty est d'ailleurs hébergée sur Tangled, et le repo GitHub est devenu un mirroir.


## 🧠 IA

* [Keep Your Slop to Yourself](https://davidgasquez.com/keep-your-slop-to-yourself) par [David Gasquez](https://davidgasquez.com)

> Produire du contenu est plus long que le consommer. Écrire un article me prend souvent de nombreuses heures, pour seulement quelques minutes de lecture. L'effort le plus important est à la charge de l'auteur.
> 
> David Gasquez lève, dans son court article, un constat simple : aujourd'hui, générer du texte avec une IA est sans-effort. Les efforts de lecture, et de nettoyage du slop (tout le blabla IA, ou les fausses informations) sont souvent laissés à la charge du lecteur.
> 
> Donc il considère malpoli (_rude_) de partager du contenu IA "brut", sans l'avoir relu ou pré-nettoyé.
>  
> Je partage complètement son point de vue. Je ne veux pas avoir à subir la flemme des gens avec qui je suis censé échanger ou travailler (je hais les vocaux, les chats ou mails générés). Si vous ne prenez pas le temps d'écrire ce que vous avez à dire, je considère que je n'ai pas à prendre le temps de le lire.

## 🛜 Internet

* [L'EstamiTech - Décentralisons avec ATProto](https://estamitech.fr/episode/94ec84b2-63a8-4827-a678-854f57004861)

> Le pote Ludovic Borie discute avec Édouard Paris de l'_AT Protocol_.
> Ils parlent de la genèse de BlueSky, du protocol, de quelques projets annexes. C'est hyper intéressant, à découvrir si vous n'avez jamais entendu parler de ce protocole.
> 
> Je m'y intéresse depuis quelques semaines, et j'ai très envie de participer au prochain [meetup AT Proto sur Lille](https://platform.openmeet.net/members/did:web:did.atproto.fr).

* [Webmention.io](https://webmention.io/) et [Webmention sur Indieweb](https://indieweb.org/Webmention)

> En me baladant [sur le web](https://matthiasott.com/notes/lazy-and-prompt#webmentions), je suis tombé sur un site qui utilisait WebMention. Le principe est de collecter les mentions d'un article faites ailleurs, ou les likes et reposts.
> J'aime bien l'idée, j'ai déjà essayé de connecter mes likes BlueSky à certains de mes articles et de récupérer des stats au build de mon site. Je pense que je vais creuser l'idée.

* [Use widgets on your Firefox New Tab page](https://support.mozilla.org/en-US/kb/firefox-new-tab-widgets)

> Une nouvelle fonctionnalité chouette de Firefox est l'intégration de widgets sur la page d'accueil. Deux widgets sympa sont déjà dispo: une todo-list basique et un timer de type pomodoro. Peut-être que d'autres arriveront à terme, ou peut-être même avec un nouveau point d'extension. 

* [A Social Filesystem](https://overreacted.io/a-social-filesystem/) par [Dan Abramov](https://overreacted.io)

> Un long post qui explique en détails le fonctionne d'AT Protocol, en prenant l'analogie du filesystem. C'est extrêmement bien écrit et intéressant. S'il y a une seule ressource à lire pour comprendre comment fonctionne AT Protocol, c'est bien cet article.

* [Taken - Since you arrived](https://sinceyouarrived.world/taken) _via_ [Ross Floate sur Bluesky](https://bsky.app/profile/rossfloate.bsky.social/post/3mlerdtejbs2q)

> Une page web hyper impressionnante, qui liste l'ensemble (l'ensemble) des données qu'il est possible d'extraire d'une simple session de navigation web. 23 points de données ont été extraits de mon browser (localisation _via_ l'IP, user-agent du navigateur, etc.), mais aussi de mon comportement (temps passé sur la page, scrolls, changement d'onglets).
> C'est terrifiant (et j'utilise Firefox, pas le navigateur le moins soucieux de la vie privée).

## 🇪🇺 Souveraineté

* [La gendarmerie a économisé 1/2 milliard d’euros grâce au Libre et Linux en 20 ans](https://next.ink/237345/la-gendarmerie-a-economise-1-2-milliard-deuros-grace-au-libre-et-linux-en-20-ans/) par [next.ink](https://next.ink)

> Utiliser du logiciel libre, ça peut faire économiser de l'argent (même si ça ne doit pas être un but premier, je pense).
> La Gendarmerie Nationale, avec son approche orientée Linux et Logiciels libres en est la preuve. Content que mes impôts ne sont pas dépensés en licences Windows ou Office 365 (en tout cas pour la gendarmerie, le reste de l'état a peut-être encore du travail à ce titre).

* [Commission d'enquête sur les dépendances structurelles et les vulnérabilités systémiques dans le secteur du numérique et les risques pour l’indépendance de la France - Assemblée nationale](https://www.assemblee-nationale.fr/dyn/17/organes/autres-commissions/commissions-enquete/ce-vulnerabilites-numeriques)

> Une commission d'enquête est actuellement menée par les députés Philippe Latombe, et Cyrielle Chatelain. Les comptes rendus sont disponibles au fur et à mesure des auditions. 44 réunions (auditions ou tables rondes) ont été tenues jusqu'à présent. J'attends avec impatience le rapport définitif (qui devrait être vaste), qui devrait être rendu pour le mois d'août.
> 
> Parmi les organisations auditionnées, on retrouve Mistal AI, les acteurs du libre comme Mastodon, Framasoft, l'April, des acteurs du Cloud comme Cloud Temple, Outscale, OVH, Scaleway, et des entités étatiques comme la DINUM, DGFIP, etc.

## 🛠️ Tools

* [Mise-en-place & Fnox : mon setup de gestion multi-projets](https://www.vrchr.fr/posts/2026/04/28/mise-en-place-fnox-setup-multi-projets/) par [Rémi Verchère](https://vrchr.fr)

> Pour faire suite à [son excellent talk à DevOxx France 2026](https://www.youtube.com/watch?v=ZEtc6WnreI0), Rémi livre un article reprenant le contenu qu'il a présenté, avec les exemples issus des démos.

* [Mise : un seul outil pour tes runtimes et tes CLI](https://vmaerten.fr/blog/installer-ses-outils-avec-mise/) par [Valentin Maerten](https://vmaerten.fr)

> Encore un nouvel article sur `mise`. Ici, Valentin présente surtout la gestion des outils et runtimes, mais explique aussi son fonctionnement avec le `PATH` et les `shims`, les lockfiles, et présente comment utiliser `mise` dans une CI pour atteindre le graal de la reproductibilité.

## ☸️ Kubernetes

* [Kloak : injection de secrets en kernel-space via eBPF sur Kubernetes](https://une-tasse-de.cafe/blog/kloak/) par [Quentin Joly](https://une-tasse-de.cafe/)

> La gestion des secrets est toujours compliquée, et leur injection dans les applications peut les exposer à des fuites (_via_ des dumps de mémoire par exemple). Dans ce post, le pote Quentin présent Kloak, qui permet de remplacer les secrets par des placeholders qu'une probe eBPF viendra remplacer par la réelle valeur juste avant d'envoyer les paquets sur le réseau.
> 
> C'est du génie, plus de peur que les secrets fuitent par l'application ou ses logs.
> 
> Ça fonctionne avec tous les outils et langages qui utilisent `libssl.so`, ou la lib Go `crypto/tls`, donc Go, Python, Ruby, NodeJS, Rust, C et même curl. Par contre, pour Java, c'est mort, puisque la JVM implémente sa propre stack.

## 🐋 Docker

* [Docker Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Docker_Security_Cheat_Sheet.html) _via_ [Renaud Lifchitz](https://bsky.app/profile/nono2357.bsky.social) sur [BlueSky](https://bsky.app/profile/nono2357.bsky.social/post/3mlxpqzprdn23)

> Cette Cheat Sheet, proposée par l'OWASP, liste 13 règles pour sécuriser l'exécution de containers Docker sur une machine.
> On y retrouve surtout des classiques, ne pas exposer la socket du `dockerd`, utiliser un `user` ou limiter les capabilities des containers.
> Les règles sont accompagnées d'exemples, pour Docker, Docker Compose, et Kubernetes.
>
> Le site propose beaucoup d'autres Cheat Sheets, toutes pertinentes, à consulter pour les technos qui vous intéressent.

## 🏗️ Architecture

* [The Hive vs. Spring Modulith : deux visions différentes du monolithe modulaire](https://medium.com/@tpierrain/the-hive-vs-spring-modulith-deux-visions-diff%C3%A9rentes-du-monolithe-modulaire-27a6d8f1bdb4) par [Thomas Pierrain](https://github.com/tpierrain/me-myself-and-i)

> Une comparaison entre deux approches pour la construction d'architectures modulaires : le pattern Hive et l'approche Spring Modulith.
> L'auteur fait une comparaison honnête en représentant les deux approches.
> D'ailleurs si vous n'avez pas entendu parler du Hive, je vous invite à le découvrir avec [cette vidéo](https://www.youtube.com/watch?v=TL_nVaOs47g).
> Personnellement, j'aime bien l'approche Hive, qui me semble plus propre et une simple dérivée de l'approche Hexagonale. Spring Modulith reste aussi cantonné à Spring et Spring Boot, alors qu'il est possible d'implémenter Hive dans n'importe quel langage.

---

La prochaine publication est prévue autour du 19 juin 🗓️

Photo de couverture par [John Fornander](https://unsplash.com/@johnfo?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/person-wearing-pair-of-white-low-top-sneakers-while-holding-wilson-tennis-racket-4R9CcBdQTEg?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)