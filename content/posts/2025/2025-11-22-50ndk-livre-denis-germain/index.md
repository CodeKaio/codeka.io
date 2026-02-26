---
date: 2025-11-22
language: fr
title: 'Kubernetes : 50 solutions pour les postes de dev et les clusters de prod'
tags:
  - kubernetes
---

Il y a quelques semaines, Denis Germain (aka [Zwindler](https://blog.zwindler.fr/) sur les internets), m'a fait parvenir un exemplaire de son livre fraîchement paru : [Kubernetes : 50 solutions pour les postes de développement et les clusters de production](https://www.editions-eyrolles.com/livre/kubernetes).

L'expertise de Denis sur Kubernetes n'est plus à prouver, c'est son sujet de prédilection, et j'assiste systématiquement à ses conférences quand j'en ai l'occasion (fanboy mode 😁).
Je suis content qu'il ait proposé de m'envoyer un exemplaire de son livre, si tu lis ces lignes Denis, merci 😊

Je me suis donc plongé dedans cette semaine.

>_TLDR_ : C'est un état de l'art solide, parfois amusant (on s'amuse comme on peut avec _Kubernetes_), et c'est très bien structuré.

<!--more-->

## Un état de l'art solide

Forcément, avec l'ambition de présenter 50 solutions Kubernetes différentes, on se dit que ça va être dense. 

C'est le cas.

Sans forcément chercher l'exhaustivité, la sélection de Denis couvre tous les usages qu'on pourrait imaginer : depuis le poste de travail, jusqu'à des usages de production, en passant par des solutions "alien" 👽 parfois originales.
Pour nous guider dans la lecture, les 300 pages du livre sont structurées en 7 grandes parties (chapitres) :

* les solutions destinées aux développeurs, installables sur leur machine, que Denis a regroupé sur le nom "Outils de type Desktop"
* les solutions managées par des opérateurs Cloud
* les solutions de déploiement automatisées qui permettent d'installer un cluster
* les solutions de déploiement de type Infrastructure as Code
* les OS Kubernetes (une grande découverte pour moi)
* les plateformes de management multi-clusters
* et les solutions "inclassables"

Dans chaque chapitre, on va donc retrouver une courte intro qui explique le but de ce regroupement de ces solutions, puis un sous-chapitre pour chacune d'entre elles, ainsi qu'une courte conclusion.

Chaque solution est présentée en suivant la même structure : 

* un tableau récapitulatif de l'outil (systèmes supportés, licence, etc.)
* une courte description ainsi que les pré-requis à l'installation ;
* l'installation en elle-même à la manière d'un tutoriel ;
* des pistes, liens et conseils pour aller plus loin ;
* un tableau récapitulatif des avantages et inconvénients de l'outil.

Cette structure permet de rapidement comprendre à quel usage convient chaque solution, et l'écriture sous forme de tutoriel détaillés (ligne de commande, consoles et autres screenshots) permet de rapidement mettre le pied à l'étrier.

## Un format original

Vous l'aurez compris, sa structure fait que ce n'est probablement pas un livre qu'on va lire de couverture à couverture. Chaque section est indépendante. On commence donc par lire l'introduction, et on peut ensuite aller parcourir le chapitre qui présente les solutions dont on a besoin.

Pour chaque solution qui serait compatible sur plusieurs OS, ou qui aurait plusieurs méthodes d'installation, Denis prend pour exemple une méthode particulière, charge au lecteur d'adapter les commandes ou de se référer à la documentation officielle. Ce qui est original, c'est que pour rendre ses exemples les plus diversifiés possibles, Denis ne prend pas toujours la méthode la plus simple. Le livre est donc très varié, que ce soit avec des tutos en mode ligne de commande, ou en console web, il y en a pour tous les goûts.

Pour ma part, j'ai passé pas mal de temps à lire le premier chapitre sur les solutions desktop que j'ai trouvé très complet (14 solutions proposées).
Je me suis aussi beaucoup attardé sur le chapitre portant sur les solutions managées, et sur celui traitant de l'_Infrastructure as Code_. Ce sont deux domaines sur lesquels je travaille régulièrement, donc c'était intéressant de pouvoir le retrouver et compléter ma boite à outils.

Je me suis un peu moins concentré sur les autres chapitres, car ils couvrent des domaines sur lesquels je n'ai pas de besoin, ou en dehors de mon expertise. Mais si l'envie me prend de bidouiller, je sais que j'y trouverai la matière pour pouvoir tester et installer une solution.

## À qui s'est destiné ?

Ce livre ne s'adresse clairement pas aux débutants sur Kubernetes. 
Bien que Denis rappelle rapidement les concepts principaux de kube en intro, le but n'est pas d'apprendre Kubernetes, ses concepts et ses usages.
Il est néanmoins très intéressant pour comprendre les méandres du fonctionnement de Kubernetes.
Chaque solution (certaines étant certifiées, je suppose par la _CNCF_ ou _via_ un _TCK_) a une approche différente de Kubernetes. Et découvrir les façons de les installer permet de démystifier le fonctionnement d'un cluster.

Le livre est donc plutôt adressé aux personnes qui souhaitent découvrir de nouvelles solutions d'installation, que ce soit dans un but expérimental (bidouille ou home-lab) ou pour monter un cluster de production.

On entend souvent parler sur les réseaux des questions de souveraineté, d'auto-hébergement, et de préférer des solutions simples. À ce titre, l'état de l'art écrit par Denis permet de répondre à ces questions. Quel que soit votre enjeu autour de Kubernetes, il y aura une solution qui pourra y répondre.

Les tableaux de présentation de chaque outil, ainsi que les récapitulatifs des avantages et inconvénients ont une valeur incroyable dans l'aide au choix.

## Mes solutions prefs

Forcément, en parcourant les différentes solutions, je suis tombé sur des solutions que j'utilise régulièrement : Minikube (qui est certes un peu lourd, mais fait bien le job), k3s qui tourne sur certaines machines que j'ai chez moi, Docker Desktop et Rancher Desktop que j'utilise parfois chez mes clients, ainsi que les différentes offres managées, et l'incontournable OpenTofu.

Mais j'ai aussi découvert (ou redécouvert) des implémentations auxquelles je n'avais jamais prêté attention et que je vais rapidement tester : _kind_ (Kubernetes IN Docker), que j'ai déjà installé sur mon poste de travail principal (il n'y a plus qu'à jouer avec pour mon prochain projet);

J'adore aussi le principe de KWOK (Kubernetes WithOut Kubelet) qui consiste à simuler le comportement d'un cluster.
Comme Denis, je pense que ça peut être très pratique pour tester le bon fonctionnement de topologies, contraintes de ressources, charts Helms, etc. 
Je testerai probablement mes charts Helms sur cette implémentation à l'avenir (peut-être même directement dans une CI, qui sait ? 🤫)

Mention spéciale pour la dernière solution présentée, qui n'est autre que le contenu du talk de Denis "Démystifions Kubernetes, binaire par binaire".
Il ajoute également en référence la vidéo de cette conférence captée à Cloud Nord 2023 (j'étais dans la salle ce jour là 💙).

Deuxième mention spéciale aussi pour l'installation de Kosmos/Kapsule, la solution managée par Scaleway. Plutôt que d'utiliser la console ou le CLI, Denis propose un tuto entièrement avec la ligne de commande _curl_, ce qui permet de bien comprendre les différentes étapes de l'installation, et de rappeler le fonctionnement des Clouds avec leurs différentes API.

Enfin, j'ai aussi découvert les solutions directement packagées sous forme d'OS. C'est clairement quelque chose que j'essayerai sur une de mes machines. L'administration n'étant pas ma spécialité (je suis surtout un Dev dans l'âme), avoir une solution toute prête peut être une super alternative. J'avais déjà entendu parler de Talos Linux, c'est probablement cette solution que je testerai quand j'aurai un peu les temps.

## Ce qu'il manque à mon avis

> Il ne manque pas grand-chose en fait. 

C'est très complet. On pourrait dire qu'il manque telle ou telle solution, mais le tour d'horizon que nous propose Denis permet déjà de voir l'évential des solutions possibles.
J'imaginais qu'il y aurait eu un chapitre sur l'offre Kubernetes de Clever Cloud, mais je suppose que le calendrier de l'édition ne matchait pas avec la sortie récente de cette offre.

Je pense que j'aurai aimé avoir directement dans le livre un tableau sous forme d'index, qui reprend l'ensemble des solutions avec tous les tableaux récapitulatifs détaillés en en-tête de chapitre (type d'outil, compatibilité, production-ready, etc.).
Cela permettrait de pouvoir trouver rapidement une solution en fonction d'un besoin particulier.
Denis a néanmoins publié ce genre de tableau dans un article de blog il y a quelques semaines : [93 façons de déployer Kubernetes : j'ai recensé (presque) toutes les méthodes existantes](https://blog.zwindler.fr/2025/11/02/93-facons-de-deployer-kubernetes/).

Si vous restez sur votre faim par rapport à une solution particulière, Denis fournit les liens et pointeurs (souvent vers la documentation officielle des outils) pour pouvoir "Aller plus loin" (du nom de la section dans chaque sous-chapitre).
Cette section est parfois un peu légère, mais une fois la phase "tuto" passée, il est courant de se plonger dans la documentation officielle, donc rien de très surprenant en fait.

## Conclusion

Je pense que c'est un livre incontournable pour les bidouilleurs de Kubernetes.

Mais c'est surtout un état de l'art, très bien structuré, qui permet de trouver rapidement une solution pour un besoin particulier.
Le fait d'avoir, pour chaque outil, un tuto d'installation, permet de savoir à quoi s'attendre quand on voudra l'installer ou le tester.

Les tableaux récapitulatifs et d'avantages et d'inconvénients permettent de pouvoir vite déterminer si une solution est adaptée à un besoin.

Ce livre m'a donné envie de tester plusieurs solutions, je pense que c'est mission accomplie donc 😅

Le travail qu'a fait Denis sur ce livre est impressionnant, bravo à lui.

## Où l'acheter ?

C'est dispo sur toutes les bonnes plateformes au prix de 37€, je vous remets ici les liens vers mes préférées : 

* Sur le site de l'éditeur [Eyrolles](https://www.eyrolles.com/Informatique/Livre/kubernetes-9782416022647/), en version papier ou numérique
* [Fnac](https://www.fnac.com/a21789006/Denis-Germain-Kubernetes)
* [Furet du Nord](https://www.furet.com/livres/kubernetes-denis-germain-9782416022647.html)
* [Cultura](https://www.cultura.com/p-kubernetes-50-solutions-pour-les-postes-de-developpement-et-les-clusters-de-production-9782416022647.html) 

Bonne lecture !