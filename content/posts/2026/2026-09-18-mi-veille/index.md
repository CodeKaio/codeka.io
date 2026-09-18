---
date: 2026-09-18
language: fr
title: La veille de Wittouck - Début septembre 2026
slug: la-veille-de-wittouck-debut-septembre-2026
series: La veille de Wittouck
tags:
  - at-proto
  - docker
  - events
  - ia
  - internet
  - linux
writing_time: 2h
atUri: "at://did:plc:a27wdjlmq3ebx4v5f2jpzvsk/site.standard.document/3mvsfyp24kr2a"
---

C'est la rentrée, et donc la reprise de mes articles de veille.

Cet été, pas mal de sujets autour d'AT Protocol, qui a maintenant le droit à sa propre catégorie.

Je pense que j'en parlerai encore dans de prochains articles, le sujet me passionne.

Bonne rentrée !

<!--more-->

## 🦋 AT Proto

* [airspace](https://getair.space/) _via_ [danielroe (@danielroe.dev)](https://bsky.app/profile/danielroe.dev/post/3mviimz36cc2e)

> Encore un cas d'usage intéressant de AT Protocol.  Utiliser AT Protocol comme une base de données JSON typée avec les _Lexicons_.
> 
> Plus besoin de backend ni de BDD, c'est le protocole qui assure ça.
> 
> C'est certainement le prochain sujet d'un article de fond ici.

* [standard.horse](https://standard.horse/)

> Si vous avez lu [mon dernier article](/2026/09/07/mon-site-sur-at-proto-avec-standard-site/), vous avez découvert standard.site.
> 
> standard.horse est une appli qui permet d'éditer directement les _Publications_ et _Documents_ au format standard.site, dans une interface web minimaliste.
> 
> C'est plutôt bien fait.

## 🐋 Docker

* [CNCF Announces Graduation of Cloud Native Buildpacks, Advancing the Standard for Container Builds | CNCF](https://www.cncf.io/announcements/2026/08/11/cncf-announces-graduation-of-cloud-native-buildpacks-advancing-the-standard-for-container-builds/) _via_ [cncf.io](https://cncf.io)

> Le projet [buildpacks.io](https://buildpacks.io), dont j'avais fait [un talk en 2023](/talks/talk-laissez-tomber-vos-dockerfile-adoptez-un-buildpack/), est passé au niveau de maturité "Graduated". Il était en "Incubating" depuis 2020 !
> 
> Belle reconnaissance pour ce projet, je vais peut-être refaire quelques test, ça fait un moment que je ne l'ai pas utilisé.

## 🎫 Events

* [Conférences : 10 conseils pour être sélectionné·e à un CFP - RÉPONDEUR #11](https://www.humancoders.com/podcast/conferences-10-conseils-pour-etre-selectionne-e-a-un-cfp-repondeur-11) _via_ [humancoders.com](https://humancoders.com)

> Dans ce podcast, il est question de la réponse à un CFP, comment bien formuler son sujet, comment bien rédiger son abstract et ses références et sa bio.
> C'est plein de conseils concrets, donnés par cinq orgas de confs qui expriment aussi leurs attentes de leur côté de la barrière.
> Je ne peux qu'être d'accord avec tout ce qui a été dit.

* [La cheatsheet de vos CFP](https://virginie-blog.pageaud.net/articles/tech/2026/cfp/) _via_ [virginie-blog.pageaud.net](https://virginie-blog.pageaud.net)

> 4 sketchnotes simples et efficaces qui expliquent les points importants à la réponse d'un CFP et des petits conseils pratiques.
> 
> À relire avant de soumettre un sujet ! Ça complète bien le podcast précédent.

* [Lettre ouverte pour la survie des événements techniques en France](https://lettreouverte.afup.org/)

> Une lettre que vous avez probablement déjà vu passer sur les différents réseaux.
> Nous l'avons signée avec Cloud Nord, je me devais donc de la repartager ici.

* [Les conférences développeurs appellent les entreprises à ne pas les laisser disparaitre - Next](https://next.ink/255239/les-conferences-developpeurs-appellent-les-entreprises-a-ne-pas-les-laisser-disparaitre/)

> La lettre ouverte proposée par l'AFUP et signée par 66 organisations a aussi été relayée par le media Next.ink.
> 
> L'article est en accès libre.

## 🧠 IA

* [I'm done using AI](https://brettcodes.com/im-done-using-ai/) _par_ [brettcodes.com](https://brettcodes.com)

> J'arrète de coder avec l'IA.
> Brett nous explique son expérience sur les 2 dernières années avec l'IA et explique pourquoi il a décidé d'arrêter de l'utiliser : "The work got done faster than I could have done on my own. And I barely had to think", "I realized, well, I'm basically just a hamster on a wheel", "So what did this do? Well, it made me lazy. It made me stop caring. It made me a worse programmer. It made me depressed. Because I stopped doing the hard work, I stopped learning, I stopped growing, I stopped being the one making the software."
> 
> C'est un sentiment qui revient de plus en plus j'ai l'impression.

* [Who does Anubis actually stop?](https://fzakaria.com/2026/07/09/who-does-anubis-actually-stop) _par_ [Farid Zakaria](https://fzakaria.com)

> Un court article qui traite de Anubis, le proxy HTTP anti-IA, non pas sur son implémentation ou son usage, mais sur les impacts.
> 
> Les bots IA arrivent à contourner le proxy en calculant la proof-of-work demandée par Anubis.
> 
> En définitive, ce sont seuls les utilisateurs humains qui sont impactés par la présence d'Anubis, et tout ça consomme de l'énergie.
> 
> Ça fait un peu penser aux captchas, que les IA ou les bots résolvent plus facilement que les humains parfois.
> 
> On pourrait se dire "à quoi bon tenter de résister alors".

## 🛜 Internet

* [Experience Better Browsing: Introducing Native Containers in Firefox 153](https://blog.mozilla.org/en/firefox/firefox-containers-preview/) _via_ [blog.mozilla.org](https://blog.mozilla.org)

> La version 153 de Firefox propose maintenant nativement les "containers", qui séparent les sessions de différents onglets, en particulier au niveau des cookies.
> 
> Un container "shopping" sera celui qui aura tous les cookies des sites web, et vous évitera de voir des pubs pour des chapeaux melon partout, alors que vous aimez juste John Steed.
> 
> Je l'utilise aussi pour pouvoir ouvrir plusieurs sessions de mail, en séparant le pro du perso, c'est très pratique.

* [The Website Specification](https://specification.website/)

> On sait tous ce qu'un un site web, basiquement une page index.html, avec un `<head>` et un `<body>`.
> 
> Cette spécification liste 170 points qui doivent être implémentés pour respecter toutes les bonnes pratiques.
> La liste des balises de <head> recommandées, les aspects SEO, mais aussi les règles basiques d'accessibilité, de sécurité, etc.
> 
> Ils proposent aussi une checklist pour vérifier un site, mais pas d'outil permet de vérifier les points automatiquement.

## 🐧 Linux

* [Linux sera votre prochain système d’exploitation – franceinfo](https://www.franceinfo.fr/replay-radio/nouveau-monde/linux-sera-votre-prochain-systeme-d-exploitation_8072087.html) _via_ [franceinfo.fr](https://franceinfo.fr)

> C'est une des rares fois où on entend parler de Linux sur un media grand public. Dans cette très courte chronique, le journaliste explique de manière plutôt bien vulgarisée ĺ'intérêt de Linux, et évoque aussi à demi-mot les sujets de souveraineté.

---

La prochaine publication est prévue autour du 2 octobre 2026 🗓️

Photo de couverture par [MChe Lee](https://unsplash.com/@mclee?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/white-table-with-black-chairs-PC91Jm1DlWA?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)

