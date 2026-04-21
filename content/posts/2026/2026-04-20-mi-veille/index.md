---
date: 2026-04-20
language: fr
title: La veille de Wittouck - Début avril 2026
series: La veille de Wittouck
tags:
  - sovereignty
  - cloud
  - ia
  - internet
  - java
  - kubernetes
  - linux
  - security
---

En ce début avril, toutes les têtes étaient tournées vers la lune, et la mission Artemis II, mais aussi vers la DINUM, qui a encore une fois capté l'attention avec l'annonce de son OS à base de Nix.

Je profite de la découverte du compte Flickr de la NASA (oui Flickr existe encore ahah) pour utiliser une des jolies photos de la mission pour illustrer cet article.

<!--more-->

## 🇪🇺 Souveraineté

* [Sécurix et Bureautix : le Linux de l’État pour remplacer Windows ](https://www.it-connect.fr/securix-et-bureautix-le-linux-de-etat-pour-remplacer-windows/)

> Après l'annonce, la DINUM commence à dévoiler des éléments de ses distributions Linux destinées à ses agents.
> Les 2 distributions sont nommées Sécurix et Bureautix, et s'appuient sur NixOS.
>
> Le code est dispo sur GitHub (faute de mieux ?) : https://github.com/cloud-gouv/securix
>
> On s'amuse du nom qui fait penser à Astérix (et semble eclipser Nix), c'est bien trouvé.
>
> Je trouve le choix d'une base NixOS intéressante. Les avantages en termes administration, de stabilité et de fork pour les différentes entités de l'État sont cool. La montée en compétence pour les profils non-tech risque de piquer par contre.

* [EU Cloud Sovereignty Framework Explained - Emiel Brok, SUSE / DOSBA ](https://www.youtube.com/watch?v=9Xky-bfqCss)
* [Lightning Talk: Designing for Exit: Why Portability Is a Sovereignty Requirement - Alireza Rahmani](https://www.youtube.com/watch?v=h7sD08_3Do4)

> Deux talks issus de la journée _Open Souvereign Cloud Day_, organisée par la CNCF en marge de la Cloud Native Con.
> 
> Le premier est un talk animé par Emiel Brok de SUSE. Il nous présente quelques enjeux de souveraineté, et le "Cloud Sovereignty Framework¨, un document de la Commission Européenne qui cadre huit axes de souveraineté, ainsi que l'outil développé par SUSE pour effectuer le self-assessment (que j'avais déjà listé lors de ma dernière veille).
>
> Le second est un talk court, animé par Alireza Rahmani avec un angle clair : la souveraineté a pour pré-requis la portabilité.
> Sans surprise, on parle de standards, de containers, de limiter le vendor lock-in.

## ☁️ Cloud

* [Datacenter map : DCWatch data](https://datacenters.hubblo.org/) _via_ [linkedin.com](https://www.linkedin.com/posts/bepetit_datacenters-dcwatch-ai-ugcPost-7449867949538119680-MiJV)

> Une carte interactive des datacenters en France, Belgique, Luxembourg et Suisse, pour 433 datacenters référencés.
> Apparaissent également les projets pour 89 nouveaux datacenters.
> 
> Je ne sais pas dans quelle mesure les données sont exhaustives, mais j'y retrouve bien les principaux DC de ma région.

* [Dropping Cloudflare for bunny.net](https://jola.dev/posts/dropping-cloudflare) _via_ [Bearstech](https://www.linkedin.com/posts/un-bon-retour-dexp%C3%A9rience-sur-la-migration-share-7447727797416058881-DHtZ)

> Bunny.net est un CDN développé par une entreprise Slovénienne. Il se veut une alternative sérieuse à CloudFlare, avec 119 PoPs (Point of Presence) répartis dans le monde.
> Il propose du CDN, un WAF, mais aussi des services edge, comme du scripting ou des containers.
> 
> L'article semble encourageant, l'auteur ne semble pas avoir eu de difficultés à migrer. C'est à tester donc.

* [s3fs-fuse/s3fs-fuse: FUSE-based file system backed by Amazon S3](https://github.com/s3fs-fuse/s3fs-fuse)

> Alors que tout le monde s'extasie devant la nouveauté de AWS : S3 Files, je voulais rappeler que si vos applications ont ce genre de besoin, peut-être que S3 n'était pas le bon choix (NFS ?).
> 
> Sinon, il existe depuis un moment ce driver FUSE qui permet de monter un bucket en filesystem depuis n'importe quelle machine, et que ça supporte tous les providers compatibles S3.
>
> Un bon moyen d'esquiver un lock-in supplémentaire.

## 🧠 IA

* [Cybersecurity Looks Like Proof of Work Now](https://www.dbreunig.com/2026/04/14/cybersecurity-is-proof-of-work-now.html) _via_ [HumanCoders](https://bsky.app/profile/news.humancoders.com/post/3mjoe2bifbc2z)

> Payer des tokens pour analyser une application, alors que des hackers payent des tokens pour essayer de la hacker.
> 
> La citation clé de cet article : _"to harden a system we need to spend more tokens discovering exploits than attackers spend exploiting them"_.
> 
> Qui gagne à la fin ? OpenAI et Anthropic (avec le modèle Mythos illustré dans l'article).
> 
> Ça me fait penser au modèle économique de Youtube : les publicitaires payent pour vous afficher de la pub, vous payez pour ne plus avoir de pub, un seul acteur gagne dans cette histoire.

* [Vibe Coding Failures: Documented AI Code Incidents](https://crackr.dev/vibe-coding-failures)

> Une liste (amusante) d'incidents imputables au "Vibe Coding".
> Il serait intéressant de pouvoir comparer cette liste avec une liste d'incidents plus généralistes pour pouvoir comparer les impacts, mais la démarche est amusante.

## 🛜 Internet

* [Powering NASA's real-time data with Lightstreamer](https://lightstreamer.com/customer-story/how-nasa-is-using-lightstreamer/) _via_ [lightstreamer.com](https://lightstreamer.com)

> En surfant sur la hype de la mission Artemis, je suis tombé un peu par hasard sur LightStreamer, qui est utilisé pour la télémétrie temps-réel de l'ISS, pour environ 100&nbsp;000 points de données, allant de la puissance captée par les panneaux solaires, la température des différents modules, aux différents niveaux du système de survie.
> 
> Il y a un client Javascript dispo sur GitHub pour se connecter au flux de données et y brancher une app.
> 
> Quelques cas d'usage sont rigolos, comme [Mimic](https://github.com/ISS-Mimic/Mimic), une maquette animée de l'ISS qui reproduit les différents mouvements de la station, et [ISS Piss Tracker](https://bsky.app/profile/iss-piss-tracker.bsky.social), un bot Bluesky qui poste le niveau du système de recyclage de l'urine.

* [NASA Johnson](https://www.flickr.com/photos/nasa2explore/)

> Le compte Flickr (oui ça existe encore) du centre spatial Johnson de la NASA.
> 
> On y retrouve des photos publiées presque quotidiennement, que ce soit des missions Artemis, de la mission en cours sur l'ISS avec Sophie Adenot, et des photos du centre de contrôle.
> 
> Les photos sont d'une qualité affolante, et c'est très exhaustif, on peut y retrouver des photos originales du programme Apollo !

* [Why Your Engineering Team Is Slow (It's the Codebase, Not the People)](https://piechowski.io/post/codebase-drag-audit/) _via_ [Hugo Lassiège sur Bluesky](https://bsky.app/profile/eventuallycoding.com/post/3miove77y7c2z)

> L'auteur a identifié 5 pain-points (qu'on retrouve souvent) qui permettent d'expliquer la lenteur d'une équipe : les estimations hors-sols (et dépassées), la peur du déploiement, les fichiers maudits, la coverage qui ne rassure personne et le Time to First Commit des nouveaux contributeurs.
> 
> Des bons angles d'analyse.

* [siddharthvaddem/openscreen: Create stunning demos for free. Open-source, no subscriptions, no watermarks, and free for commercial use. An alternative to Screen Studio.](https://github.com/siddharthvaddem/openscreen)

> Un outil qui permet de créer des démos vidéo et de les éditer.
> 
> Capture, édition, zooms, capture de l'audio, annotations. Ça a l'air plutôt complet.
> 
> Je pense que je vais essayer de m'en servir pour mes prochains supports de formation ou tutoriels.

* [Am I a Tech Bro?](https://amiatechbro.com/)

> Un quiz rapide qui permet de vérifier votre "alignement" avec les idées des Tech-Bros.
> 
> Les questions sont directes, les intentions sont claires.
> Toutes les questions sont sourcées, c'est assez intéressant d'aller les consulter.
> 
> Bon, pas de surprise, je ne suis pas un Tech Bro (dommage ?).

## ☕ Java

* [Valkey: Announcing Spring Data Valkey: A New Season for High-Performance Spring Applications](https://valkey.io/blog/spring-data-valkey/) _via_ [valkey.io](https://valkey.io)

> Valkey publie sa propre implémentation de Spring Data : Spring Data Valkey, ainsi que le starter Spring Boot qui l'accompagne, et l'implémentation de Spring Cache.
> 
> Ça reste compatible avec Redis (quel intérêt d'ailleurs ?). L'intérêt principal est de s'appuyer sur le client GLIDE, qui propose une meilleure gestion des clusters et une intégration native d'OpenTelemetry.
> 
> La doc officielle est disponible sur [spring.valkey.io](https://spring.valkey.io).

* [Best Oracle Java Alternatives in 2026 Comparison of OpenJDK Distributions](https://www.youtube.com/watch?v=Ytdo8OGEYFI)

> Un vidéo intéressante qui liste différentes distributions de l'OpenJDK : OpenJDK, Temurin, Zulu, Liberica, etc.
> 
> Les avantages et inconvénients de chaque distribution sont listées, et surtout les éléments qui peuvent aider au choix : le support (LTS et +), le patching, etc.
> 
> Un article un peu plus détaillé reprend les différents points et même un tableau comparatif bien utile en fin de page.
> 
> Attention, la vidéo est produite par Bellsoft, donc peut-être un peu biaisée (même si le comparatif m'a semblé honnête).

## ☸️ Kubernetes

* [First Contact with Tetragon](https://blog.littlejo.link/en/tetragon/started/) par [Joseph Ligier](https://blog.littlejo.link)

> Joseph nous présente dans ce premier article le fonctionnement de Tetragon, qui utilise eBPF pour tracer les exécutions de processus au niveau du noyau Linux, et en particulier pour des clusters Kubernetes.
> 
> Ce premier article présente rapidement l'outil, j'ai hâte de voir l'intégration avec Kubernetes et l'utilisation de policies pour surveiller des comportements anormaux.

## 🐧 Linux

* [Dotfiles + Claude Code = my tiny config workshop](https://www.hsablonniere.com/dotfiles-claude-code-my-tiny-config-workshop--95d5fr/) par [Hubert Sablonnière](https://www.hsablonniere.com/)

> Hubert présente dans cet article son setup shell. Au menu : Fish, Starship, Stow pour la gestion des dotfiles, et quelques raccourcis intelligents.
> 
> Il nous met aussi à dispo son repo GitHub avec ses dotfiles, si jamais on veut s'en inspirer.

## 🔒 Security

* [L'affaire Trivy — Acte IV : Aqua a parlé, voici ce que ça change](https://blog.stephane-robert.info/post/trivy-aqua-reponse-supply-chain/) par [Stéphane Robert](https://blog.stephane-robert.info)

> Stéphane Robert a résumé pour nous l'affaire Trivy, avec le déroulé des événements et l'explication technique détaillée et pédagogique.
> C'est probablement une des meilleures sources d'information sur le sujet, tant on voit qu'il a suivi l'affaire en détails.
> 
> Cette affaire a des conséquences dures : la confiance est perdue, il est probable que Trivy ne soit plus l'outil d'analyse privilégié (d'ailleurs Stéphane ne le recommande plus).

* [YubiKey SSH Authentication: Stop Trusting Key Files on Disk](https://dev.to/orthogonalinfo/yubikey-ssh-authentication-stop-trusting-key-files-on-disk-157o)

> Un setup de connexion SSH avec une clé YubiKey. La clé privée étant stockée sur la YubiKey. C'est plutôt simple, le tuto liste toutes les commandes nécessaires.
> 
> Bien penser à configurer deux clés pour éviter d'être coincé en cas de perte !

---

La prochaine publication est prévue autour du 1 mai (avec une édition spéciale DevOxx France) 🗓️

Photo de couverture par [NASA Johnson](https://www.flickr.com/photos/nasa2explore/55208327975/) sur [Flickr](https://www.flickr.com/photos/nasa2explore/)