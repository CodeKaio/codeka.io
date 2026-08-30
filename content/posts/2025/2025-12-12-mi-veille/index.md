---
date: 2025-12-12
language: fr
title: La veille de Wittouck - Début décembre 2025
slug: la-veille-de-wittouck-début-décembre-2025
series: La veille de Wittouck
tags:
  - devops
  - ia
  - internet
  - java
  - linux
  - security
---

Les vacances approchent, et la veille de Wittouck continue.
Plusieurs vidéos et podcasts intéressants en ce début de mois, l'interview de Kelsey Hightower est à ne pas manquer.

<!--more-->

## 👷 DevOps

[AI, DevOps, and Kubernetes: Kelsey Hightower on What’s Next](https://www.youtube.com/watch?v=HdUbTyvrfKo) par [JetBrains](https://www.youtube.com/@JetBrainsTV) _via_ [Bluesky](https://bsky.app/profile/kelseyhightower.com/post/3m7qccwj2qs2f)

> Une interview fleuve (plus d'une heure) de Kelsey Hightower.
> Il revient sur les fondations de la philosophie DevOps et le platform-engineering, le Cloud et les approches serverless, et bien entendu sur Kubernetes et les concepts de Workload APIs.
> 
> S'il y a une seule chose à regarder de tout ce mois de décembre, c'est bien cette vidéo.

## 🧠 IA

* [Devstral2 Mistral Vibe CLI](https://mistral.ai/fr/news/devstral-2-vibe-cli) _via_ [Frédéric Camblor (@fred.camblor.fr)](https://bsky.app/profile/fred.camblor.fr/post/3m7q2tkzefc23)

> Mistral tape très fort avec son nouveau modèle consacré au dev : Devstral2, et avec son CLI Vibe. Le modèle Devstrall Small 2 est conçu pour tourner localement. De quoi pouvoir s'abstraire des abonnements ?
> Si Vibe permet à terme de pouvoir se brancher sur un modèle local, c'est la mort de Claude Code assurée.

* [RSL 1.0 - L'heure de passer à la caisse a sonné pour les IA](https://korben.info/rsl-web-ia-payer-scraping.html) par [Korben](https://korben.info)

> Les bots IA pompent tout le contenu du web pour pouvoir ensuite vendre des tokens.
> Des solutions de Firewall pour protéger le contenu du scrapping sont de plus en plus hype, Cloudflare avait présenté [leur solution Firewall for AI](https://www.cloudflare.com/application-services/products/firewall-for-ai/) cette année, et c'est aussi le propos du firewall [Anubis](https://github.com/TecharoHQ/anubis).
> 
> RSL propose un standard (encore un) pour publier du contenu sous license, pour autoriser ou non le scrapping, et porter les attributions et compensation des auteurs de contenue (paiement par crawl ou par inférence). Vu que plusieurs géants du CDN sont derrière tout ça, ça risque de prendre.

## 🛜 Internet

* [Conventions dans l'industrie logicielle](https://www.reddit.com/r/developpeurs/s/vzcN2X7WKH) _via_ [r/developpeurs](https://www.reddit.com/r/developpeurs)

> Un post reddit qui liste quelques conventions dans notre industrie : commits, versionning, gestions de branches, styles, etc. C'est loin d'être exhaustif mais on retrouve les outils qu'on connait déjà.
> 
> J'ai d'ailleurs découvert le site https://keepachangelog.com (même si j'écrivais déjà des CHANGELOG.md)

* [Advent of Code, Behind the Scenes](https://youtu.be/uZ8DcbhojOw) par [Eric Wastl (@was.tl)](https://bsky.app/profile/was.tl/post/3m73d2wdr6k2r)

> Le créateur de Advent of Code nous explique comment s'est passée sa journée du 1er décembre 2015, comment il est passé de 70 utilisateurs estimés à 4000, puis 15 000 au 2 décembre, pour finir à 52 000 au 1er janvier 2016.
> Il explique aussi ce qui fait un bon puzzle, et comment il les génère, y compris les inputs.
> La partie la plus intéressante est quand il nous montre les visualisations développées par la communauté pour la résolution des puzzles. C'est bluffant de créativité.

* [Fediverse - Le futur sera fédéré et auto-hébergé - Elena Rossini](https://www.projets-libres.org/podcast/s04e4-fediverse-futur-federe-auto-heberge-elena-rossini/) _via_ [Raphaël Semeteys](https://www.linkedin.com/posts/raphaelsemeteys_fediverse-le-futur-sera-f%C3%A9d%C3%A9r%C3%A9-et-auto-h%C3%A9berg%C3%A9-activity-7391003730726363137-fQlf?utm_source=share&utm_medium=member_android&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)

> On entend de plus en plus parler du web libre, avec le [Fediverse](https://fr.wikipedia.org/wiki/Fediverse) ou les approches comme celles du [IndieWeb](https://indieweb.org/) et de l'auto hébergement. Dans ce podcast, Elena nous explique comment elle a quitté Twitter suite au rachat par Elon Musk, et son usage du Fediverse, en particulier sur Mastodon et PeerTube.
> 
> Elena maintient d'ailleurs un blog sur le sujet, avec pas mal d'articles abordables : [blog.elenarossini.com](https://blog.elenarossini.com/tag/the-future-is-federated/)

## ☕ Java

* [What's New in IntelliJ IDEA](https://www.jetbrains.com/idea/whatsnew/) _via_ [jetbrains.com](https://jetbrains.com)

> Jetbrains liste les nouveautés de la dernière mouture de son IDE. Au menu, un support amélioré de Java 25 et Spring Boot 4, un nouveau thème (plutôt joli), et quelques annonces sur les améliorations à venir sur Junie avec le _Bring Your Own Key_ qui est particulièrement attendu.

* [Implémenter un lock en Elasticsearch](https://www.loicmathieu.fr/wordpress/fr/informatique/implementer-un-lock-en-elasticsearch/) par [Loïc Mathieu](https://www.loicmathieu.fr/)

> Loïc propose une implémentation simple d'un lock dans un Elasticsearch, pour essayer de reproduire ce qu'on peut faire avec une base de données relationnelle et un `SELECT FOR UPDATE`. C'est pas parfait (uniquement logiciel), mais ça fait le job.

* [Which Version of JDK Should I Use?](https://whichjdk.com/) _via_ [whichjdk.com](https://whichjdk.com)

> Une page efficace qui liste les distributions de JDK les plus courantes. Les contraintes de license sont rappelés, et il en sort des recommandations sur l'usage de telle ou telle distribution. Eclipse Temurin est selon moi la distribution à toujours privilégier (comme les auteurs de ce site).

## 🐧 Linux

* [Pop!_OS 24.04 LTS Released: A Letter From Our Founder](https://blog.system76.com/post/pop-os-letter-from-our-founder?utm_source=newsletter_all&utm_medium=email&utm_term=pop&utm_content=blog&utm_campaign=pop-cosmic&mc_cid=4710c4bf7f)

> J'utilise Pop!_OS depuis plusieurs années sur une grande partie de mes machines. La sortie de cette version embarquant l'environnement COSMIC était très attendue, j'ai hâte de tester ça lors de ma prochaine réinstallation.

## 🔒 Security

* [CWE - 2025 CWE Top 25 Most Dangerous Software Weaknesses](https://cwe.mitre.org/top25/archive/2025/2025_cwe_top25.html) _via_ [Catalin Cimpanu (@campuscodi.risky.biz)](https://bsky.app/profile/campuscodi.risky.biz/post/3m7qekq3pec2x)

> MITRE (oui ça existe encore cf [La veille de Wittouck - Fin Avril 2025](/posts/2025/2025-05-02-mi-veille/#-sécurité)) a publié la liste des 25 CWE (pour _Common Weakness Enumeration_) les plus dangereuses. On y retrouve les mêmes CWE qui composent l'OWASP Top 10, mais ici, on est à jour sur une vision 2025 inspirée des CVE observées sur l'année.

---

La prochaine publication est prévue autour du 26 décembre 2025 🗓️

Photo de couverture par [Elena Mozhvilo](https://unsplash.com/@miracleday?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/red-and-gold-love-print-gift-boxes-hlHmSFZrATc?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)