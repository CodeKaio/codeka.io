---
date: 2026-07-03
language: fr
title: La veille de Wittouck - juin 2026
series: La veille de Wittouck
tags:
  - git
  - ia
  - internet
  - sovereignty
  - events
writing_time: 2h30
---

Ce mois de juin, j'étais en marathon de conférences. DevLille, Tech'Work, Breizhcamp, et je termine la saison avec Riviera Dev.

J'ai donc eu peu de temps pour rédiger cette veille (j'étais tout le temps sur Factorio, parce que _The Factory Must Grow_ 🏭⚙️, vous savez bien),

Je reprends donc le rythme tout doucement, au début de l'été. Bonne lecture !

<!--more-->

## 🔀 Git

* [git-history Documentation](https://git-scm.com/docs/git-history#Documentation/git-history.txt-fixupcommit) _via_ [git-scm.com](https://git-scm.com)

> Je sais pas vous, mais je suis bien hypé par cette commande, livrée avec la version 2.55 de Git, j'ai hâte que cette version soit dans les dépôts pour pouvoir mettre à jour mes scripts 🤓 
> 
> `git history fixup` remplace donc le combo `git commit --fixup` suivi du `git rebase --autosquash`.
> 
> Une belle amélioration de DevEx.

* [Stop Using Conventional Commits](https://sumnerevans.com/posts/software-engineering/stop-using-conventional-commits/) par [Sumner Evans](https://sumnerevans.com)

> L'auteur relève le problème de Conventional Commit : mettre en avant le type de modification plutôt que son sujet.
> 
> Un bon message de commit doit véhiculer l'intention et le périmètre. Le type est souvent implicite dans le message de commit, et n'aide pas particulièrement au debugging lorsqu'il y a un incident.
> 
> L'auteur va aussi plus loin : générer des changelog à partir de l'historique des commits est un non-sens (le changelog est à destination des utilisateurs, pas des développeurs), utiliser semver en fonction des types de commits conduit à générer des breaking-change non souhaités, et surtout, conditionner les pipelines de CI en fonction des messages de commit est la pire des idées (un commit nommé `docs: fix typos` qui introduit une version vérolée d'un package).
> 
> Je suis assez d'accord en fait, même si j'aime bien gitmoji parce que je trouve ça joli ahah.
> 
> L'auteur propose quelques alternatives issus de projets Open Source, qui mettent en avant le scope, et suppriment le type, comme les conventions de Linux, Git, ou Go.
> 
> Et comme [14 standards ne suffisent pas](https://xkcd.com/927/), l'auteur propose aussi le sien nommé [Scoped Commits](https://scopedcommits.com/)

## 🧠 IA

* [Incident Report: CVE-2026-LGTM](https://nesbitt.io/2026/06/26/incident-report-cve-2026-lgtm.html) par [Andrew Nesbitt](https://nesbitt.io)

> Une fiction qui prend la forme d'un post-mortem de traitement d'une CVE.
> La base de l'attaque illustrée: du texte blanc sur fond blanc indiquant à l'agent IA qu'un package vérolé est sécure. L'agent inclut alors la version de package ce qui déclenche l'attaque.
> 
> Les conclusions de cette réflexion sont intéressantes, en particulier : "The phrase “human in the loop” appears in four vendor contracts; in each case they forgot to loop the humans in"

* [There are no lossless transformations of natural-language text](https://sophiebits.com/2026/06/25/there-are-no-lossless-transformations-of-natural-language-text) par [Sophie Alpert](https://sophiebits.com)

> Un court article dans l'esprit des [AI Manifesto](/ai) qu'on voit toujours fleurir sur internet.
> 
> Le propos est ici de remettre en avant que le processus d'écriture est important : il cadre la réflexion, pousse à l'efficacité dans le choix des mots, et à la responsabilité du discours.
> 
> L'aspect de la transaction temporelle est aussi remis en avant : espérer que des gens passent du temps à lire du texte qui a été généré à partir d'un prompt est un manque de respect de leur temps.
> 
> Une des phrases de la conclusion a plutôt bien résonné en moi (raisonné même) : _Readers will appreciate hearing your thoughts, even at the expense of supposed “polish”_.

## 🛜 Internet

* [cartes.gouv.fr](https://cartes.gouv.fr/explorer-les-cartes/)

> L'IGN (Institut national de l'information géographique et forestière) a mis a disposition cartes.gouv.fr.
> 
> Les cartes de bases sont sympa, ça fonctionne plutôt bien. Mais le vrai truc fou, c'est les différentes couches disponibles, cartes et photos satellites, lidar, parcelles cadastrales et mon coup de cœur : la carte Cassini, construite entre 1756 et 1815. On a l'impression de remonter le temps 🚗⚡

* [RDAP.ORG](https://about.rdap.org/)

> RDAP (pour _Registration Data Access Protocol_) est le protocole remplaçant de WHOIS (déprécié depuis début 2025).
> Il a la particularité d'être structuré en JSON, rendant la consommation par des programmes plus facile que le format WHOIS qui était du texte simple.
> 
> Tous les TLDs (_Top Level Domain_) ne le supportent pas encore (le `.io` n'est pas supporté par exemple), mais c'est le cas pour les plus importants (`.fr`, `.com`...).
> 
> J'ai découvert ce protocole lors de l'excellent conférence "[Noms de domaines : la grande histoire des petites extensions](https://youtu.be/FZj6VEnxsRI?si=Akhz2iCfvEo8FSbd)" par Benoît Masson et Théo Bougé tous deux de chez OVH, conférence que j'ai vue à BreizhCamp 2026.

* [Publishing on the Atmosphere with Standard.site](https://piccalil.li/blog/publishing-on-the-atmosphere-with-standardsite/) par [Declan Chidlow](https://vale.rocks/) _via_ [piccalil.li](https://piccalil.li)

> Je m'intéresse de plus en plus à AT Protocol ces derniers temps. Bluesky a récemment annoncé intégrer directement les lexicons de `standard.site` pour pouvoir afficher les contenus avec une jolie carte dans leur interface.
> 
> Cet article explique comment enregistrer un site web dans `standard.site`, avec du code NodeJS, qui crée les enregistrements AT Protocol.
> 
> Ça a l'air très simple, il est certain que je vais tenter l'expérience pour le site que vous êtes en train de lire 😆

* [There Are No Instances in atproto](https://overreacted.io/there-are-no-instances-in-atproto/) par [Dan Abramov](https://overreacted.io/)

> Dan, qui est décidément très très fan de AT Protocol, a publié un nouvel article de vulgarisation du protocole, en appuyant cette fois-ci les concepts liés au réseau et à l'agrégation des donnés du protocole.
> Il compare très simplement AT Protocol et Bluesky à RSS et Google Reader.
> La comparaison marche bien.

* [The Steam Machine Is An Iconoclastic Computer Born In Unforgiving Times](https://aftermath.site/steam-machine-review-price/) par [Chris Person](https://aftermath.site/author/chrisperson/) _via_ [aftermath.site](https://aftermath.site)

> Valve a enfin mis en vente la Steam Machine (ainsi que le Steam Deck, de nouveau en stock).
> Mais le tarif pique : les composant électroniques sont accaparés par les fournisseurs d'IA, et la Steam Machine arrive alors sur le marché plus chère qu'une PS5 Pro.
>
> Le test va assez loin, et appuie sur l'aspect "Open Source": pour son propre développement, Valve a poussé ses contributions à l'écosystème Linux, la machine se veut réparable, on peut imprimer ses propres facades... Ces améliorations vont aussi bénéficier au monde du gaming Linux en général, Valve a peut-être relancé le monde du gaming PC de salon. Mais au pire moment possible.

* [`human.json` ou comment créer un réseau de contenu certifié humain !](https://anthonypena.fr/2026/05/12/humanjson-ou-comment-creer-un-reseau-de-contenu-certifie-humain/) par [Anthony PENA](https://anthonypena.fr)

> Anthony nous présente le protocole `human.json`, qui est un protocole de "vouching" de sites web.
> Le but non-masqué est de créer un réseau de confiance pour s'assurer que la création de contenu des sites web concernés est bien faite par un humain, et pas une IA.
>
> Le protocole vient avec une extension Firefox que j'ai immédiatement installée, et un fichier .json à positionner sur son site avec les contenus que l'on valide de son côté.
>
> Et devinez-quoi, le pote Anthony m'a déjà ajouté ☺️

## 🇪🇺 Souveraineté

* [SecNumCloud en (pas si) bref](https://www.linkedin.com/pulse/secnumcloud-en-pas-si-bref-vincent-strubel-505ge?utm_source=share&utm_medium=member_android&utm_campaign=share_via) par Vincent Strubel _via_ [linkedin.com](https://linkedin.com)

> Cet article écrit par Vincent Strubel, directeur général de l'ANSSI (_Agence Nationale de la Sécurité des Systèmes d'Information_) a pour but de redéfinir le label SecNumCloud et le remettre dans son contexte.
> SecNumCloud est bien un label de sécurité, qui prémunit également des lois extraterritoriales, en particulier les lois étasuniennes, et des kill-switch des entreprises étrangères.
> 
> Vincent Strubel précise que du point de vue de l'ANSSI, le label est donc un label de souveraineté, mais qu'il ne traite pas le sujet de la dépendance (qui est le plus important de mon point de vue).

* [Let’s Encrypt c’est cool, mais il est seul au monde](https://blog.seboss666.info/2026/06/lets-encrypt-cest-cool-mais-il-est-seul-au-monde/) par [Seboss666](https://blog.seboss666.info)

> Seboss666 explique dans cet article l'histoire de Let's Encrypt, et remet en avant les faits : Let's Encrypt, c'est avant tout une entité étasunienne.
> L'impact d'une coupure serait catastrophique.
> 
> Il a donc cherché les alternatives européennes. Certaines implémentent le protocole ACME (_Automated Certificate Management Environment_), qui permet la gestion automatisée des certificats (émission, renouvellement), mais souvent partiellement, et la gratuité des certificats ou la possibilité d'obtenir des wildcards est parfois cachée.
> 
> On manque donc d'alternatives sérieuses et complètes en Europe sur ce sujet. Le risque lié à la dépendance est bien réel.

* [independance.numericatous.fr - Souveraineté Numérique - Analyse de site](https://independance.numericatous.fr/)

> Un petit outil qui évalue la souveraineté numérique (un bien grand concept) en récupérant les données DNS, de certificats, d'hébergement et les contenus des pages web (trackers, SEO et liens sortants).
> 
> C'est pratique et rapide.
> J'ai un plutôt bon score sur mon site ([92/100](https://independance.numericatous.fr/?r=5U1rvTNn)).
> 
> Intéressant de tester avec des quelques sites publics, lemonde.fr ([74/100](https://independance.numericatous.fr/?r=qXGIf4ke)), next.ink ([97/100](https://independance.numericatous.fr/?r=9TcjOLL4))  a un super score, bravo.
> 
> Je ne vous ai pas dit de tester avec [s3ns.io]() ni avec [bleucloud.fr]() hein (le score ne va pas vous surprendre).

## 🎫 Events

* [What to say](https://einarwh.no/blog/2026/04/29/what-to-say/) par [Einar W. Høst](https://einarwh.no)

> Un court article qui aide à répondre à la question qu'on se pose quand on veut écrire un article ou soumettre un talk : _que dire ?_, sous-entendu, comment formuler son sujet.
>
> Le point important est bien de transmettre _son_ message, et pas celui copié depuis une IA (ce qui n'intéresse personne)

---

La prochaine publication est prévue le 17 juillet 2026 🗓️

Photo de couverture par [Upgraded Points](https://unsplash.com/@upgradedpoints?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/brown-wooden-bench-on-beach-during-daytime-KVym2PAn1gA?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)
