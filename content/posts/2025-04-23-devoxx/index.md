---
date: "2025-04-23"
language: fr
title: DevOxx 2025 - Bilan
draft: true
tags:
  - Internet
  - Kubernetes
---

Pour la deuxième année consécutive, j'ai la chance d'être speaker à DevOxx France.
Le pass de speaker me permet d'assister aux trois jours de la conférence (dont les billets partent plus vite que ceux d'un concert d'AC/DC).

Ce post fait le bilan de ma participation à cette édition 2025, sur les deux plans, en tant que speaker et en tant que participant.

## L'orga

Être speaker permet de découvrir les coulisses de l'organisation. En observant un peu les fameux _gilets rouges_ (tenue officielle des organisateurs), on peut se rendre compte du travail colossal de l'organisation d'une conférence comme DevOxx France (rien à voir avec ma conf pref : [Cloud Nord](https://cloudnord.fr/)). 15 salles, ça nécessite 15 _gilets rouges_ pour briefer les speakers, et 15 ingés pour la partie technique (micro, équipement de captation). Sans compter l'accueil et la remise des badges, la bagagerie, la logistique pour la distribution des repas, le point info, etc.
Les orgas sont aux petits soins.

L'aspect technique est parfait. Sur chaque pupitre, 2 câbles HDMI permettent de brancher des ordinateurs. Un commutateur permet de switcher facilement d'une source à une autre, ce qui facilite les présentations avec 2 speakers. Un compteur de temps est aussi mis dans le champ de vision du speaker. Le temps restant pour la présentation défile (toujours trop vite), une lumière jaune s'allume quand il reste quelques minutes (3 à 5 en fonction des formats), une lumière rouge signale quand le temps est écoulé.

> Petit bémol, dans ma salle, ce compteur était pile dans l'alignement du pupitre, il fallait que je me décale légèrement pour pouvoir le voir. Rien de très gênant.

## Mercredi

Mercredi, je suis arrivé un peu tard, j'ai donc loupé la keynote. En arrivant, ça m'a permis de faire un peu le tour des stands, de prendre un café tranquillement pour m'accoutumer à l'ambiance si particulière du lieu.

### Bring the Action : Using GraalVM in Production - Alina Yurenko

Cela fait plusieurs années qu'Alina nous présente l'utilisation de GraalVM. Cette année, elle se concentre sur la facilité d'utilisation de la compilation native (quelques lignes dans notre `pom.xml`), sur les options à ajouter pour accélérer la compilation lorsqu'on développe, et sur la compatibilité des librairies. Elle a également fait un focus sur les performances des binaires compilés en présentant quelques tirs de perfs (similaires à la JVM, mais sans le temps de warmup). Autre nouveauté intéressante, GraalVM commence à supporter WebAssembly, ce qui permettra à terme de pouvoir compiler du Java pour l'exécuter dans un navigateur !

> Je suis un grand fan du travail d'Alina, et cette présentation tout en démo montre bien les évolutions de GraalVM native, orientées expérience développeur. Cela donne clairement envie de tester sur un projet !

### Kestra : un orchestrateur open source, event driven et déclaratif, codé en Java - Loïc Mathieu

Dans ce lunch talk (15 minutes, ça va vite !), mon pote Loïc a présenté Kestra, son architecture, et a fait une démo d'écriture et d'exécution d'un workflow simple dans l'interface Kestra, et a présenté l'écosystème des plugins Kestra. Il a aussi expliqué les avantages d'avoir choisi Java pour le développement de cet outil, en particulier l'utilisation de Nashorn (qui sera remplacé à terme par GraalVM polyglot), pour exécuter les scripts écrits dans le langage du choix du développeur.

> J'ai apprécié re-découvrir ces éléments, même si 15 minutes sur ce sujet, c'est un peu court ! J'ai eu l'occasion de rediscuter du fonctionnement de Kestra avec Loïc, au détour d'un couloir, donc ma curiosité a été satisfaite :D

### BullShit IT Awards : Célébrons les erreurs des équipes Tech !

Une salle comble pour mon pote Romain ! Romain nous présente les meilleures pépites qu'il a pu voir ou entendre sur des projets IT ! Le public a été mis à contribution pour voter pour la meilleure pépite.

> J'ai voté pour "PR ouverte depuis 72 jours, 19 commentaires, 0 merge". Bonne ambiance dans la salle, le public semble avoir apprécié le show

### OpenRewrite : Refactoring as code - Jérôme Tama

Jérôme nous explique comment fonctionne OpenRewrite, et comment écrire notre propre code de refactoring. Le cas d'usage présenté est assez concret (migration JUnit 4 à 5) et illustre bien l'intérêt de l'outil.

> Un outil intéressant. Par contre, écrire les migrations soit même semble quand même assez compliqué. Mais à tester en utilisant les recette contribuées par la communauté.

### Rebase d'image Docker/OCI avec crane - Julien Wittouck

Je ne pouvais pas manquer mon propre talk :). Cela s'est plutôt bien passé de mon point de vue. J'ai même eu le temps de jouer une démo que je m'étais gardé sur le côté au cas où.

Les slides sont dispo ici : [Rebase d'images Docker/OCI avec crane]({{< relref "/talks/talk-rebase-crane">}})

### Les stands

J'ai pas mal discuté avec les potes de Clever Cloud, des nouveautés à venir, et du sursaut de souveraineté récent qu'on certains de leurs clients. Cela montre une prise de conscience qui sera intéressante à creuser. Sur le stand de Michelin, une animation _Gran Turismo_ était organisée : 1 tour de circuit du Mans, dans un siège baquet, avec volant. J'ai posé le deuxième temps et gagné une casquette Michelin plutôt cool.
J'ai un peu discuté avec Redis et MongoDb également.

### La soirée des speakers

J'ai fait un saut rapide à la soirée des speakers. C'est assez impressionnant de voir tout ce monde. Nous avons un peu discuté, mangé un bout et dégusté un verre de vin.  Comme j'étais quand même assez fatigué par cette journée, je n'ai pas passé beaucoup de temps.

## Jeudi

Jeudi, j'arrive boosté au Palais des congrès. Malheureusement, il y a déjà la queue pour rentrer dans l'amphi bleu pour les keynotes. J'opte donc pour une tactique alternative : café + croissant, et salle Maillot qui est une bonne salle d'overflow.

### Keynote : Langage, IA et propagande : la guerre des récits a déjà commencé - Elodie Mielczareck

Elodie présente les différents étages de la correspondance entre le langage et le monde réel. Elle évoque des termes qu'on observe souvent dans l'actualité, en particulier la notion de "post-vérité" et les thèmes abordés par le film Matrix autour de la notion du réel.

> Assez inspirant, parfois difficile à suivre. Je pense que je re-visionnerai cette keynote pour m'assurer de comprendre ces notions que j'ai eu un peu de mal à suivre (le café met 30 minutes à faire effet, on était un peu juste là :D )

### Keynote : La territorialisation des infrastructures comme levier de pouvoir - Ophélie Coelho

En s'appuyant sur son travail de recherche et sur des cartes géographique des câbles sous-marin et de la localisation des centres de données, Ophélie met en avant le pouvoir et le contrôle que peuvent avoir des entreprise ou des états sur nos communications réseau. Elle met aussi en avant l'importance du logiciel dans ce jeu, tout en haut du modèle OSI.

> Là encore, c'est un sujet qui résonne pas mal avec les sursauts de souveraineté qu'on observe en ce moment. Le lien entre géopolitique et numérique est indiscutable. Je pense que je vais acheter son livre pour me renseigner plus en détail.

### Anatomie d'une faille - Olivier Poncet

Olivier retrace les différentes étapes qui ont mené à l'implémentation de la faille dite 'xz' de l'année dernière. De l'ingénieurie sociale pour "infiltrer" les maintainers du paquet cible, à l'ingénieurie technique pour intégrer le code malveillant dans les paquets, jusqu'a la découverte "accidentelle" de la faille.

> Beaucoup de gens on cité cette conférence comme étant une de leurs préférées de cette édition de DevOxx 2025. C'est aussi le cas pour moi. Le travail de recherche qu'a produit Olivier sur ce talk est impressionnant et on découvre (avec la pédagogie qui le caractérise) les détails de cette faille, qui est complètement folle. C'est bien construit, et c'est effrayant.

### Kubernetes : 5 façons créatives de flinguer sa prod 🔫 - Denis Germain

Denis présent cinq cas issus de ses expériences, qui ont conduit à une prod en PLS. Des erreurs bêtes, des cas d'erreurs en cascade. Au delà des erreurs, Denis présente aussi les actions mises en place pour que cela ne se reproduise plus, backups du cluster, admission controller et policies Kyverno ou OPA.

> Un talk plein d'humour, sous la forme d'un REX. C'est bien expliqué et on repart avec des solutions concrètes pour éviter de reproduire ces cas chez nous (on aura au moins été prévenu). Un de mes talks préférés sur cette édition.

### Comment builder Java depuis ses sources - Antoine Dessaigne

Antoine explique sous forme de démo les étapes pour builder un JDK, en partant du `git clone`, pour finir avec un `./java --version`.
L'environnement de build est construit au fur et à mesure avec des `apt install`. Le process de build est pour finir assez simple, mais contient des dépendances amusantes issues de certains modules de Java : alsa pour la partie "son", "cups" pour le code d'impression. Chose intéressante, le build cross-plateforme a été évoqué, et semble assez simple à mettre en place.

> J'étais curieux de ce talk. Je n'ai jamais pris le temps de builder moi-même un jdk, donc je voulais savoir ce que ça impliquait. C'est beaucoup plus simple que ce que j'imaginais. Je testerai probablement l'image Docker qu'il a mis à disposition pour me faire une idée.

### Communiquer à 36000 km : l'art de l'efficacité avec moins d'un Watt - Paul Pinault

Un talk sur les communications IOT via satellite. On y parle de modulation de fréquences, de consommation électriques, de gain d'une antenne satellite, et de la différence entre les constellations de satellites et les géostationnaires.

> J'ai compris 10% de ce talk. Néanmoins, c'est assez intéressant de voir que le sujet reste accessible au commun des mortels en termes d'implémentation et de coût. Un talk à voir pour s'ouvrir l'esprit.

---

