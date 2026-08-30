---
date: "2025-03-27"
language: fr
tags:
  - Java
  - Kubernetes
  - IA
  - Internet
  - Docker
title: La veille de Wittouck - Fin Mars 2025
slug: la-veille-de-wittouck-fin-mars-2025
series: La veille de Wittouck
atUri: "at://did:plc:a27wdjlmq3ebx4v5f2jpzvsk/site.standard.document/3mucgxanwyv2j"
---

Avec la sortie récente de Java 24, cette deuxième édition de _La veille de Wittouck_ rassemble tout d'abord quelques lectures indispensables autour de la toute nouvelle version de mon langage préféré.
Quelques articles généralistes également, dont un très bien illustré sur les technologies de stockage. Enfin, parmi les nouveautés, l'outil _kaniuse_ référence la compatibilité des API Kubernetes (très pratique), et Docker Compose supporte maintenant la publication des fichiers `compose.yml` dans un registry OCI.

Bonne lecture !

## ☕ Java

* [Java 24 : Quoi de neuf ?](https://www.loicmathieu.fr/wordpress/informatique/java-24-quoi-de-neuf/?utm_source=pocket_saves) par [Loïc Mathieu](https://www.loicmathieu.fr)
> Java 24 est sorti ce mercredi 18 mars. L'occasion de revoir les fonctionnalités de cette version sur le blog du pote Loïc. Les 2 nouveautés les plus importantes à mon avis sont la _JEP 483: Ahead-of-Time Class Loading & Linking_ pour optimiser le démarrage d'une application et _JEP 491: Synchronize Virtual Threads without Pinning_ pour booster les perfs des _threads_ virtuels. Les _Stream Gathereres_ passent aussi en version standard.

* [Welcome, GraalVM for JDK 24!🚀](https://medium.com/graalvm/welcome-graalvm-for-jdk-24-7c829fe98ea1) par [Alina Yurenko](https://www.linkedin.com/in/alinayurenko/)
> Avec la sortie de Java 24, s'accompagne la sortie de GraalVM pour JDK 24. Quelques nouveautés importantes et attendues : des exécutables natives un peu plus petits (~5%), de meilleures performances attendues pour les workloads de type IA, l'export facile de SBOM (Software Bill Of Material), et la génération de rapports HTML qui fournissent des détails sur le build des images.

* Mon pote Guillaume Dufrêne retrace l'histoire de Java sur LinkedIn, avec le contexte historique IT associé :
  * [Java 9](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_springboot-kubernetes-alphago-activity-7305879062458126338-mTDe)
  * [Java 11](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_webmethods-springboot-honey-activity-7306241440496472064-lpP7)
  * [Java 17](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_springio-honey-java-activity-7307328618362294272-8rO9)
  * [Java 24](https://www.linkedin.com/posts/guillaume-dufr%C3%AAne-90179410_springboot-opentelemetry-javaone-activity-7307690998195802112-dnwB)

> Deuxième partie de l'histoire avec les versions "modernes" de Java. Guillaume se concentre sur les versions LTS (avec l'impasse de la version 21, on ne lui en tiendra pas rigueur 😅)

* [Programmatic Bean Registration](https://docs.spring.io/spring-framework/reference/7.0-SNAPSHOT/core/beans/java/programmatic-bean-registration.html?utm_source=pocket_shared) _via_ [Josh Long](https://bsky.app/profile/starbuxman.joshlong.com) sur [Bluesky](https://bsky.app/profile/starbuxman.joshlong.com/post/3ljrthfuch22r)

> Une 3ᵉ façon d'enregistrer des beans dans le contexte Spring (en plus de `@Bean` dans une classe `@Configuration` et `@Component`). L'approche va permettre de simplifier des cas d'usage de `@Conditional` et se veut compatible avec les approches d'AOT.

## 🛜 Internet

* [How Core Git Developers Configure Git](https://blog.gitbutler.com/how-git-core-devs-configure-git/) par [Scott Chacon](https://bsky.app/profile/scottchacon.com) _via_ [RudeOps](https://www.rudeops.com/)

> J'avais déjà lu cet article il y a quelque temps, et je suis retombé dessus via l'excellente newsletter de RudOps. Pour faire simple, le TLDR est un bon point de départ. Tous les paramètres sont expliqués et détaillées, c'est un article très complet. 

* [IO devices and latency](https://planetscale.com/blog/io-devices-and-latency?utm_source=pocket_shared) _via_ [La veille des ours](https://www.linkedin.com/newsletters/la-veille-des-ours-7100088441966575616/)

> Un article qui reprend en détail les différentes technologies de stockage, depuis le stockage sur bande, jusqu'aux disques réseau, en passant par les disques rotatifs et SSD. Les technologies sont illustrées par des animations interactives. C'est très bien écrit et très intéressant pour se rendre compte de l'évolution des technologies au fil des années.

* [LinkedIn Post Inspector](https://www.linkedin.com/post-inspector/)

> J'ai découvert cet outil qui permet de voir à quoi va ressembler la "carte" générée par LinkedIn pour les liens qui y sont postés. Ça va bien me servir 😅

## 🧠 IA

* [ChatGPT : le mythe de la productivité](https://framablog.org/2025/03/09/chatgpt-le-mythe-de-la-productivite/?utm_source=pocket_shared) par [Hubert Guillaud](https://hubertguillaud.wordpress.com/) _via_ [Framablog](https://framablog.org)

> Ce _post_ fait écho avec le sentiment que j'ai aujourd'hui autour de l'usage des IA génératives et l'impact ressenti, que j'ai un peu développé dans mon [AI Manifesto](/ai). C'est une lecture indispensable !

## ☸️ Kubernetes

* [kaniuse](https://kaniuse.gerome.dev/) par [Gérôme Grignon](https://www.linkedin.com/in/gerome-grignon/)

> Un super outil qui permet de visualiser le cycle de vie de fonctionnalités de Kubernetes. Indispensable pour celles et ceux qui doivent manager un cluster.

## 🐋 Docker

* [Using Docker Compose with OCI artifacts](https://docs.docker.com/compose/how-tos/oci-artifact/) _via_ [Guillaume Lours](https://www.linkedin.com/in/guillaumelours/) sur [LinkedIn](https://www.linkedin.com/posts/guillaumelours_docker-compose-activity-7306423870838747136-SvCl?utm_source=share&utm_medium=member_desktop&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc)

> Une nouvelle feature de `docker compose` qui permet de partager facilement des compositions en tant qu'artifacts OCI. C'est une belle avancée pour pouvoir partager des stacks prêtes à l'emploi.

---

La prochaine publication est prévue autour du 18 avril (après DevOxx donc) 🗓️

* Photo de couverture par [Alexandre Boucey](https://unsplash.com/@thisisareku?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) sur [Unsplash](https://unsplash.com/photos/black-framed-eyeglasses-on-white-book-page-FuhXMEU8LNw?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
      