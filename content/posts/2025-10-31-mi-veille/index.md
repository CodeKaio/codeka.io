---
date: 2025-10-31
language: fr
title: La veille de Wittouck - Fin octobre 2025
series: La veille de Wittouck
tags:
  - IA
  - Databases
  - Internet
  - Java
  - Linux
---

Après deux semaines bien chargées, principalement autour de l'organisation de [Cloud Nord](https://cloudnord.fr), et pour ne pas finir comme une citrouille enfumée, on lit ["La veille de Wittouck"]({{< relref "/series/la-veille-de-wittouck">}}), en se gavant de bonbecs (non).

<!--more-->

## 🛜 Internet

* [Plotting my Factorio Base](https://www.reddit.com/r/factorio/comments/1oh5f3y/plotting_my_factorio_base/) sur [r/factorio](https://www.reddit.com/r/factorio).

> En scrollant Reddit, je suis tombé sur cette vidéo folle. Une usine Factorio dessinée comme blueprint. Le rendu est chouette.
> Drawscape dessine des schémas avec un bras robotisé et des feutres (un Pen Plotter). Ils font de chouettes blueprints, principalement d'avions ou de voitures.
> Le truc dingue, ils ont mis sur [Github](https://github.com/drawscape-labs/drawscape-factorio) leur script qui convertit une map du jeu en SVG pour leur outil.

* [is it really FOSS ?](https://isitreallyfoss.com/) _via_ [Rémi Verchère](https://bsky.app/profile/r.verchere.fr) sur [Bluesky](https://bsky.app/profile/r.verchere.fr/post/3m4ak6lpkhc2x).

> Il est parfois difficile de s'y retrouver dans les licenses open-source, et surtout de savoir si le logiciel qu'on utilise est réellement open-source ou pas. Ce site répond à cette question. C'est loin d'être exhaustif, mais il y a déjà pas mal de projets référencés.

* [Twelve years of blogging](https://vladmihalcea.com/twelve-years-of-blogging/) par [Vlad Mihalcea](https://vladmihalcea.com/)

> On ne présente plus Vlad Mihalcea, très connu pour son expertise sur Hibernate. Il revient sur ses 12 ans de blogging et partage quelques stats impressionnantes. Un modèle et une inspiration sur l'aspect consistant et régulier de son travail.

## ☕ Java

* [Try Out JEP 401 Value Classes and Objects](https://inside.java/2025/10/27/try-jep-401-value-classes/) sur [Inside Java](https://inside.java/)

> Ce court article présente le principe des _Value Classes_ en Java, leur utilisation, et l'impact sur les performances du passage de certaines classes du JDK en _Value Classes_.
> C'est assez impressionnant. Vu qu'un build du projet Valhalla est dispo sur jdk.java.net, j'ai bien envie de tester tout ça.

* [Why Bloomberg Chose Vendor-Neutral Java Over Big Tech](https://thenewstack.io/why-bloomberg-chose-vendor-neutral-java-over-big-tech/) sur [The New Stack](https://thenewstack.io/)

> Je pense depuis longtemps que le choix d'une implémentation de Java ne doit pas être pris à la légère. Les implémentations sont nombreuses et certaines ont leurs spécificités sur le plan des performances (Azul Zulu), ou du support LTS.
> Bloomberg a adopté Eclipse Temurin pour son côté "Vendor Neutral", et l'aspect sécurité. Ils communiquent aussi contribuer au développement, ce qui est une bonne initiative, bravo.

## 💾 Databases

* [Build Your Own Database](https://www.nan.fyi/database) _via_ [Nicolas Fränkel](https://bsky.app/profile/frankel.ch) sur [Bluesky](https://bsky.app/profile/frankel.ch/post/3m3zrvi4nl22n)

> Oh my god. Cet article est complètement dingue. Il explique pas à pas, les étapes qui ont amené à l'écriture d'une base de données. C'est hyper intéressant pour comprendre ou se rappeler tout ce qui se cache derrière les BDD qu'on manipule tous les jours.
> L'article est intéractif, avec des exemples à dérouler, c'est probablement un des meilleurs articles que j'ai lu ce mois-ci.

## 🧠 IA

* [It's insulting to read your AI-generated blog post](https://enocc.com/ai/2025/10/24/insulting-ai-writing.html) par [Pablo Enoc](https://enocc.com/about) _via_ [Emile Heitor](https://bsky.app/profile/imil.net) sur [Bluesky](https://bsky.app/profile/imil.net/post/3m46u5i7pbs2l).

> Ce coup de gueule me parle bien. Si vous n'écrivez pas vous-même votre contenu, pourquoi je devrai prendre le temps de le lire ?
> Ça résonne avec ce que j'avais moi même écrit sur mon [AI Manifesto]({{< relref "/ai">}}).

* [Microsoft force un bouton de partage Copilot directement dans la barre des tâches !](https://www.generation-nt.com/actualites/windows-11-copilot-microsoft-ia-taskbar-vision-2062756) sur [Generation NT](https://www.generation-nt.com/)

> Une des raisons qui fait que je *déteste* Microsoft et Windows : on ne possède pas et on ne contrôle pas le système d'exploitation de nos machines sous Windows.
> 
> #WittouckPasContent On a maintenant le droit à du Copilot partout. En plus des boutons physiques sur les claviers (!), maintenant un système Windows qui vient avec un compte Microsoft en ligne obligatoire, probablement pour vendre encore plus d'IA que personne ne veut ou n'a demandé.

## 🐧 Linux

* [A Word on Omarchy](https://マリウス.com/a-word-on-omarchy/)

> Je ne sais plus retrouver où j'ai découvert cet article.
> J'ai dû récemment réinstaller 2 machines sous Linux, et j'ai résisté à la tentation de tester la "distribution" Omarchy. Cet article très pointilleux analyse le contenu de cette distrib.
> La conclusion (TLDR) est qu'Omarchy est surtout un bundle d'applications préinstallées, mais avec des scripts qui semblent plutôt artisanaux, là où des paquets d'une distribution auraient été plus standards.
> L'auteur de l'article n'est donc clairement pas convaincu par l'intérêt de cette distribution, qu'il refuse de qualifier en tant que telle.
> C'est important parfois de démonter une hype, et là c'est très bien écrit.

* [Terminal Trove](https://terminaltrove.com/explore/) _via_ [Benji Le Gnard](https://bsky.app/profile/benjilegnard.bsky.social) sur [Bluesky](https://bsky.app/profile/benjilegnard.bsky.social/post/3lzssminryq2q)

> "oh wow, la mine d'or" cite Benji, et il a raison. Ce site référence 727 outils en ligne de commande. On peut les filtrer par langage, license, OS, etc. Il y a aussi une jolie liste de Terminaux, qu'on peut facilement comparer comme mon [Tilix préféré _Vs_ Ghostty](https://terminaltrove.com/compare/terminals/ghostty-vs-tilix/).

## 🎫 Évènements

* [Cloud Nord](https://cloudnord.fr)
> Je ne reviens pas en détails sur Cloud Nord, j'écris un article à part sur les dessous de l'organisation que je publierai la semaine prochaine.
J'ai encore un peu de travail pour finaliser les vidéos qui seront sur [notre chaîne Youtube](https://www.youtube.com/playlist?list=PLVQhat0Bx0WB-fhbbQ0bQkhfTLAZIU2IU).
> 
> J'ai l'impression que cette édition 2025 était un succès, malgré quelques couacs de projection dans une des salles.
Il est vraissemblable que l'édition 2026 suivra le même format.

---

La prochaine publication est prévue autour du 14 novembre 🗓️

Photo de couverture par [Colton Sturgeon](https://unsplash.com/@coltonsturgeon?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/close-up-of-pumpkin-near-wall-EFQlS6SL9uw?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText).
