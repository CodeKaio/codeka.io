---
date: 2025-12-02
language: fr
title: La veille de Wittouck - Fin novembre 2025
series: La veille de Wittouck
tags:
  - internet
  - ia
  - java
  - kubernetes
  - events
---

L'hiver approche et la saison des conférences se termine tranquillement.
Avant de pouvoir ouvrir les premières case de son calendrier de l'avent, on lit "La veille de Wittouck" pour savoir ce j'ai trouvé intéressant en cette fin novembre.

<!--more-->

## 🛜 Internet

* [How Did I Get Here](https://how-did-i-get-here.net/) _via_ [Siegfried Ehret](https://ehret.me/news-from-last-month/202512-developer/)

> Une implémentation amusante de _traceroute_ sur une page web.
> En plus des détails du chemin emprunté entre le navigateur web qui ouvre cette page et le serveur qui la génère, l'article propose des explications de quelques protocoles réseau bas-niveau : ICMP, WHOIS, BGP.
> On a aussi le droit à une introduction aux AS (pour _Autonomous System_) qui constituent le squelette d'internet.

* [Ces logiciels libres qui changent ma vie !](https://nirinarabeson.fr/posts/semaine-du-libre-2025) par [Nirina Rabeson](https://nirinarabeson.fr/) sur [Bluesky](https://bsky.app/profile/nirinarabeson.fr/post/3m6i6lfnjas2l)

> Nirina nous raconte son histoire, en toute simplicité, et en particulier sa découverte du monde du libre et de Linux.
> La lecture de cet article a pas mal résonné avec ma propre histoire.
> Nirina nous donne aussi la liste des outils qu'il utilise au quotidien, ça peut aider à trouver des alternatives libres à certains usages.

## ☕ Java

* [Spring Boot 4](https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Release-Notes)

> La version 4.0 de mon framework pref est enfin sortie. Beaucoup de nouveautés sont listés dans les release notes, la migration sera un peu coûteuse (peut-être facilitée avec un OpenRewrite).

* [Are you really wasting your time in Java without these 10 libraries?](https://blog.frankel.ch/wasting-time-without-10-libraries/) par [Nicolas Fränkel](https://blog.frankel.ch)

> Un article en réponse à un autre. Les opinions de Nicolas sont assez tranchées. Celles et ceux qui me connaissent bien verront que je suis plutôt d'accord.
> Beaucoup de librairies (plutôt _mainstream_, rien d'exotique) règlent des problèmes qui n'existent pas (ou plus) et donc pourraient être dispensables. Mais certaines sont clairement incontournables.

* [I'm working on Electron for Java. Anyone is interested in trying it out?](https://www.reddit.com/r/java/comments/1owdru7/im_working_on_electron_for_java_anyone_is) sur [r/java](https://www.reddit.com/r/java)

> Il y a des développeurs qui ont parfois des idées marrantes. Ici un dev propose une implémentation équivalente à Electron pour Java.
> Le code est disponible sur [GitHub](https://github.com/tanin47/java-electron).
> C'est rigolo, mais l'approche de la communication entre la WeebView et la JVM repose sur HTTP, ce qui ne semble pas le plus efficace.

* [IntelliJ IDEA 2025.3 💚 Spring 7](https://blog.jetbrains.com/idea/2025/11/intellij-idea-2025-3-spring-7/)

> Spring Framework 7 est sorti ce mois-ci.
> L'équipe d'IntelliJ nous explique quelles fonctionnalités ont été ajoutées à l'IDE pour améliorer l'expérience de dev, en particulier sur le versionning des _API REST_, les _HTTP Interfaces_ et les _Dynamic Bean Registration_.

* [jMolecules 2.0](https://odrotbohm.de/2025/11/jmolecules-2.0-stereotypical/) par [Oliver Drotbohm](https://odrotbohm.de/) _via_ [LinkedIn](https://www.linkedin.com/posts/odrotbohm_jmolecules-20-stereotypical-activity-7395134565708664834-Vsgm/?utm_source=social_share_send&utm_medium=android_app&rcm=ACoAAAnJockBYMCZmKvFfK2Ytyqf-fRZDwyzaKc&utm_campaign=share_via)

> Je ne connaissais pas cette lib, sur laquelle je suis tombé au détour d'un post LinkedIn.
> jMolecules ([Github](https://github.com/xmolecules/jmolecules)) propose un ensemble d'annotations "marqueur" pour identifier les composants des différents types d'architecture (CQRS, en couches classiques, Onion, Hexagonale).
> Des jeux de test ArchUnit sont aussi proposés, ainsi qu'une intégration avec Spring Modulith. Ça a l'air bien fait et complet, à tester prochainement.

## 🧠 IA

* [Entre la chaise et le clavier #1 : Éviter la bêtise artificielle avec Anne Alombert](https://next.ink/podcast/entre-la-chaise-et-le-clavier-1-eviter-la-betise-artificielle-avec-anne-alombert/) sur [Next.ink](https://next.ink)

> Le premier podcast d'une nouvelle série portée par le media indépendant Next.ink.
> Pendant plus d'une heure, Anne Alombert explique le fonctionnement des IA génératives, leurs usages observés, et les impacts concrets qu'elles ont sur le fonctionnement de notre cerveau.
> C'est passionnant.

## ☸️ Kubernetes

* [Ingress NGINX Retirement: What You Need to Know](https://www.kubernetes.dev/blog/2025/11/12/ingress-nginx-retirement/)

> Comme annoncé [il y a quelques mois](https://github.com/kubernetes/ingress-nginx/issues/13002) par l'équipe de devs (et relayé [ici]({{< relref "/posts/2025/2025-04-18-mi-veille" >}})), le projet ingress-nginx tire sa révérence.
> On a maintenant une date de fin de vie prévue pour Mars 2026. Bien que les systèmes installés avant cette date continueront à fonctionner, il va falloir migrer vers un autre controller.
> [Cette page](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/) liste les implémentations d'Ingress, en attendant de [migrer vers Gateway API](https://gateway-api.sigs.k8s.io/guides/) qui est la cible à terme.
> 

## 🎫 Évènements

* [DevFest Lyon 2025](https://devfest.gdglyon.com/)

> Ce 28 novembre dernier, j'étais au DevFest Lyon, et c'était bien bien chouette. Mon feedback est [ici]({{< relref "/posts/2025/2025-11-29-devfest-lyon" >}}) en attendant les photos officielles.

---

La prochaine publication est prévue autour du 12 décembre 🗓️.

Photo de couverture par [Elena Mozhvilo](https://unsplash.com/@miracleday?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/white-and-red-love-print-box-wBenc1bRgGI?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText).
      