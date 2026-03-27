---
date: 2026-03-27
language: fr
title: La veille de Wittouck - Fin mars 2026
series: La veille de Wittouck
tags:
  - souverainete
  - open source
  - git
  - ia
  - linux
  - events
  - java
  - tools
writing_time: 2h30
---

Comme depuis plusieurs semaines, je continue à m'intéresser aux sujets autour de la souveraineté numérique.
J'ai aussi lu plusieurs articles et rapports sur l'Open Source, écrits par la CNCF et la Linux Foundation (les sujets sont connexes).

Ces sujets ont d'ailleurs maintenant le droit à leur propre tag 😆

Je vous conseille aussi deux chouettes talks que j'ai regardé cette semaine.

<!--more-->

## 🇪🇺 Souveraineté

* [La Suite numérique de l’État : critique des critiques – Framablog](https://framablog.org/2026/03/19/la-suite-numerique-de-letat-critique-des-critiques/) sur [framablog.org](https://framablog.org)

> Framasoft commente l'initiative de l'état concernant La Suite dans un long article.
> Je suis plutôt d'accord avec leur vision. La Suite s'appuie sur un ensemble de logiciels libres, et la DINUM re-contribue du code, c'est sain. Je ne trouve pas déconnant que l'état développe ses outils et "wrappe" des logiciels pour les adapters à son besoin (ce que plein de boites font), plutôt que d'acheter un logiciel sur étagère.
>
> Après, je pense aussi que l'état devrait soutenir les entreprises Françaises en privilégiant un achat local. S'ils développent tout eux-même, ça va aussi à l'encontre de cette approche.
>
> En tout cas, je ne vais pas pleurer pour Microsoft ou Google, ou des ESN qui se seraient gavées sur un tel projet (ou qui l'auraient vautré).

* [SUSE Cloud Sovereignty Framework Self Assessment](https://www.suse.com/fr-fr/cloud-sovereignty-framework-assessment/) _via_ [suse.com](https://suse.com)

> Un questionnaire rapide (32 questions) permettant à une organisation d'estimer son niveau de souveraineté.
>
> C'est assez difficile à répondre en fait, beaucoup de questions sont sur le plan de la juridiction, mais je pense que c'est un outil qui permet déjà de se poser les bonnes questions.

* [Ponos - Alternative europeenne à LinkedIn](https://ponos-job.eu/)

> Une alternative européenne à LinkedIn.
> Intéressant, l'authentification peut se faire avec un compte Mastodon ou Bluesky.
>
> Je me suis créé un compte, je vais voir ce qu'il s'y passe

## 🧩💻 Open Source

* [CNCF Annual Report 2025](https://www.cncf.io/reports/cncf-annual-report-2025/) _via_ [cncf.io](https://cncf.io)

> La CNCF a publié son rapport 2025. On y retrouve plein d'infos intéressantes sur les communautés, les events (Kubecon et CloudNativeCon), et les différents projets.
> 
> Sans surprise, OpenTelemetry et Kubernetes sont les projets avec le plus de contributions.

* [ROI for Open Source Software Contribution](https://www.linuxfoundation.org/research/contribution-roi?hsLang=en) sur [linuxfoundation.org](https://linuxfoundation.org)

> Un autre rapport, cette fois-ci émis par la Linux Foundation sur le ROI lié aux contributions Open Source.
>
> Les coûts liés à des absences de contribution sont vastes : maintient de forks privés, redévelopper des features ou trouver des solutions alternatives, reporter des features parce qu'un produit opensource ne permet pas de la développer, subir des changements sur les produits opensource.

* [When Projects Fail: Why Companies Should Treat Open Source as Infrastructure](https://dev.to/katcosgrove/when-projects-fail-why-companies-should-treat-open-source-as-infrastructure-32c0) par [Kat Cosgrove](https://bsky.app/profile/kat.lol)

> "If You Rely on Open Source, Help Maintain It", c´est l'appel de Kat Cosgrove.
> Nous dépendons tous de projets Open Source qui sont maintenus par des contributeurs souvent peu aidés ([xkcd.com/2347/](https://xkcd.com/2347/)).
> L'arrêt de ses projets peut avoir des impacts importants sur toute l'industrie qui en dépend.


* [Meet Kit, your companion for a new internet era](https://blog.mozilla.org/en/firefox/meet-kit/) sur [blog.mozilla.org](https://blog.mozilla.org)

> Ok, la nouvelle mascotte de Firefox, nommée Kit, est beaucoup trop mignonne.
>
> Sur la page du branding de Kit, il y a des exemples d'utilisation de la mascotte sur des tote bags ou t-shirts. Trop choupi.

## 🔀 Git

* [Git Worktree Comme Un Chef](https://www.metal3d.org/blog/2026/git-worktree-comme-un-chef/) par [Patrice Ferlet](https://www.metal3d.org/) _via_ la newsletter de [RudeOps](https://www.rudeops.com/newsletters/2026-03-23)

> J'utilise les worktree depuis un moment pour travailler efficacement sur plusieurs branches en paralèlle, et éviter de manipuler des stash.
> 
> Cet article explique comment ça fonctionne, et donne aussi quelques astuces, et unn script utilisable immédiatement. Plus d'excuses.

## 🧠 IA

* [GitHub va utiliser vos données pour entraîner Copilot](https://www.programmez.com/actualites/github-va-utiliser-vos-donnees-pour-entrainer-copilot-39223) sur [programmez.com](https://programmez.com)

> GitHub continue à aspirer nos données pour alimenter son Copilot. Maintenant, ce sont les données de chat, et des intéractions avec leur modèle.
> 
> Bon, bah il est peut-être temps de quitter ce truc.

* [Le plus grand danger de l’IA à l’université n’est pas la triche, c’est l’érosion de l’apprentissage lui‑même](https://www.rtbf.be/article/le-plus-grand-danger-de-l-ia-a-l-universite-n-est-pas-la-triche-c-est-l-erosion-de-l-apprentissage-lui-meme-11696887?utm_source=bskyrtbfinfo) sur [rtbf.be](https://rtbf.be)

> Une réflexion sur l'utilisation de l'IA Gen pour tricher aux exams.
> 
> Si les profs utilisents des IA pour rédiger des sujets et que les étudiants utilisent des IA pour y répondre, à quoi tout ça sert ?
> 
> Ça peut faire sourir, mais peut-être que l'exam sur feuille va redevenir à la mode.

* [You Might Debate It — If You Could See It](https://blog.jim-nielsen.com/2026/opacity-of-generative-tools/) par [Jim Nielsen](https://blog.jim-nielsen.com)

> L'auteur a une réflexion intéressante. Quand on génère du code avec l'IA Gen, des choix implicites sont présentés et admis sans aucun débat.
> 
> La dernière phrase de l'article pique un peu : "When you offload your thinking, you might be on-loading someone else’s you’d never agree to — personally or collectively"
> 
> L'auteur parle aussi de "Cheval de Troie" du craft.

## 🐧 Linux

* [Your Linux OS has a Chaos Engine - and Nobody Told You](https://hackernoon.com/your-linux-os-has-a-chaos-engine-and-nobody-told-you) _via_ [hackernoon.com](https://hackernoon.com)

> Une commande `tc` (pour _traffic control_) permet d'indiquer au noyau Linux de modifier le comportement du réseau : ajouter du délai avant l'envoi des paquets ou en supprimer un pourcentage.
> 
> C'est intéressant, parce que c'est déjà intégré dans les distribs et dans le noyau. Pas besoin de proxy supplémentaire, une commande suffit.
> Par contre ça fait un peu flipper, comme les règles iptables, on risque facilement de se retrouver coincé si on manipule mal. Les règles sont effacées au reboot, donc dans le doute...

* [Now playing](https://lkdjiin.github.io/blog/2026/03/18/now-playing/) par [Xavier Nayrac](https://lkdjiin.github.io/)

> J'avais déjà regardé un peu le protocole D-Bus, pour permettre à mon StreamDeck de contrôler Spotify (comme les boutons media sur un clavier).
> 
> Là c'est aussi intéressant, la commande playerctl permet de manipuler un sous-protocole pour les lecteurs multimédia, et même de récupérer la cover de l'album en cours de lecture.
> 
> via https://bsky.app/profile/news.humancoders.com/post/3mhe2xgxwfh2z

## 🎫 Events

* [DevLille 2025 - Au secours ! Mon manager me demande des KPIs !](https://www.youtube.com/watch?v=z0hi-gmZmw8) par [Geoffrey Graveaud](https://www.linkedin.com/in/geoffrey-graveaud-033319b0)

> Ça faisait un moment que je voulais voir cette conférence de Geoffrey Graveaud, j'ai pu la regarder cette semaine dans le train.
> Sa technique de speaker est parfaite, des temps de pause, le discours est calme et appuyé.
> Sur le fond, il donne des astuces pour comprendre pourquoi les managers demandes des KPIs, et pour aller chercher derrière les indicateurs qui n'ont pas de sens.

## ☕ Java

* [G1, ZGC, Shenandoah, ... Chez Java, ils sont gentils avec tous leurs GCs mais moi je choisis lequel](https://www.youtube.com/watch?v=gvKXB8VQZXw) par [Antoine Dessaigne](https://fr.linkedin.com/in/antoine-dessaigne-3618aa37)

> À Touraine Tech, je n'avais pas vu cette conf du pote Antoine Dessaigne, qui passait juste avant mon créneau.
> Je l'ai regardée cette semaine, et j'ai fait un feedback à Antoine à sa demande.
> Si vous êtes curieux du fonctionnement des GC, tout est bien illustré. Antoine donne aussi quelques astuces avec des flags intéressants à activer.

## 🛠️ Tools

* [This Is Not The Computer For You](https://samhenri.gold/blog/20260312-this-is-not-the-computer-for-you/) par [Sam Henri Gold](https://samhenri.gold)

> Une review du MacBook Neo. Pas sur la machine en elle-même. Pas sur ses limites en elle-même. Mais sur ce qu'elle va apporter aux personnes qui vont l'utiliser.
> 
> Ce n'est pas la machine parfaite, mais une machine pour débuter, et apprendre, comprendre, expérimenter. Comme on l'a tous plus ou moins fait avec nos premières machines, achetées avec l'argent économisé pendant de longs mois ou années.
> 
> Chromebook prend un taquet au passage : "The kid who tries to run Blender on a Chromebook doesn’t learn that his machine can’t handle it. He learns that Google decided he’s not allowed to."

---

La prochaine publication est prévue autour du 16 avril 🗓️

Photo de couverture par [Markus Winkler](https://unsplash.com/@markuswinkler?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/a-box-with-a-key-chain-and-a-key-chain-on-it-Z8yWSsx8OWE?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)
