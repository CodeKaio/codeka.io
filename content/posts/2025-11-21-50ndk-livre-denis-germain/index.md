---
date: 2025-11-21
language: fr
title: 'Kubernetes : 50 solutions pour les postes de dev et les clusters de prod'
tags:
  - kubernetes
draft: true
---

Il y a quelques semaines, le pote Denis Germain (aka [Zwindler](https://blog.zwindler.fr/) sur les internets), m'a fait parvenir un exemplaire de son livre [Kubernetes : 50 solutions pour les postes de développement et les clusters de production](https://www.editions-eyrolles.com/livre/kubernetes).
L'expertise de Denis sur Kubernetes n'est plus à prouver, c'est son sujet de prédilection en conférence tech.

Je me suis donc plongé dedans cette semaine.

_TLDR_ : C'est un état de l'art, parfois amusant (on s'amuse comme on peut avec _k8s_), et c'est très bien structuré.

<!--more-->

## Un état de l'art

Forcément, avec l'ambition de présenter 50 solutions Kubernetes différentes, on se dit que ça va être dense, et parfois décousu.
Pour nous guider dans la lecture, le livre est structuré en 7 grandes parties (chapitres) :

* les solutions destinées aux développeurs, installables sur leur machine, que Denis a regroupé sur le nom "Outils de type Desktop"
* les solutions managées par des opérateurs Cloud
* les solutions de déploiement automatisées qui permettent d'installer un cluster
* les solutions de déploiement de type Infrastructure as Code
* les OS Kubernetes
* les plateformes de management multi-clusters
* et les solutions "inclassables"

Dans chaque chapitre, on va donc retrouver une courte intro qui explique le regroupement des solutions, puis une description de chacune d'entre elles, ainsi qu'une conclusion.

Chaque solution est présentée en suivant la même structure : 

* un tableau récapitulatif de l'outil (systèmes supportés, licence, etc.)
* une courte description ainsi que les pré-requis à l'installation ;
* l'installation en elle-même à la manière d'un tutoriel ;
* des pistes, liens et conseils pour aller plus loin ;
* un tableau récapitulatif des avantages et inconvénients de l'outil.

## À qui s'est destiné ?

Ce livre ne s'adresse clairement pas aux débutants sur Kubernetes. 
Bien que Denis rappelle rapidement les concepts principaux de kube en intro, le but n'est pas d'apprendre Kubernetes, ses concepts et ses usages.

Le livre est adressé aux personnes qui souhaitent découvrir de nouvelles solutions d'installation, que ce soit dans un but expérimental ou pour monter un cluster de production.

À ce titre, les tableaux de présentation de chaque outil, ainsi que les récapitulatifs des avantages et inconvénient ont une valeur incroyable dans l'aide à un choix.

## Mes solutions prefs

Forcément, en parcourant les différentes solutions, je suis tombé sur des solutions que j'utilise régulièrement : Minikube (qui est certes un peu lourd, mais fait bien le job), k3s qui tourne sur certaines machines que j'ai chez moi, Docker Desktop et Rancher Desktop que j'utilise parfois chez mes clients, ainsi que les différentes offres managées, et l'incontournable OpenTofu.

Mais j'ai aussi découvert des implémentations auxquelles je n'avais jamais prêté attention et que je vais rapidement tester : kind (Kubernetes IN Docker) que j'ai déjà installé sur mon poste de travail principal (il n'y a plus qu'a jouer avec)

J'adore le principe de KWOK (Kubernetes WithOut Kubelet) qui consiste à simuler le comportement d'un cluster.
Comme Denis, je pense que ça peut être très pratique pour tester le bon fonctionnement de topologies, contraintes de ressources, charts Helms, etc. 
Je testerai probablement mes charts Helms sur cette implémentation à l'avenir (peut-être même directement dans une CI, qui sait ? 🤫)

Mention spéciale pour la dernière solution présentée, qui n'est autre que le contenu du talk de Denis "Démystifions Kubernetes, binaire par binaire".
Il ajoute également en référence la vidéo de cette conférence captée à Cloud Nord 2023 (j'étais dans la salle ce jour là 💙)

## Ce qu'il manque à mon avis

Pas grand chose en fait.
Je pense que j'aurai aimé avoir directement dans le livre un tableau sous forme d'index, qui reprend l'ensemble des solutions avec tous les points détaillés.
Cela permettrai de pouvoir trouver rapidement une solution en fonction d'un besoin particulier.
Denis a néanmoins publié ce genre de tableau dans un article de blog il y a quelques semaines : [93 façons de déployer Kubernetes : j'ai recensé (presque) toutes les méthodes existantes](https://blog.zwindler.fr/2025/11/02/93-facons-de-deployer-kubernetes/)

## Conclusion

Bien qu'il soit un ovni par le format, je pense que c'est un livre incontournable pour les bidouilleurs de Kubernetes.

---

La prochaine publication est prévue autour du 28 novembre 2025 🗓️.
Je serai au DevFest de Lyon ce jour-là, ce sera surtout un récap de cette nouvelle conférence.

Photo de couverture par [Brigitte Tohm](https://unsplash.com/@brigittetohm?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/white-ceramic-cup-with-coffee-UnACLA4mhLQ?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)