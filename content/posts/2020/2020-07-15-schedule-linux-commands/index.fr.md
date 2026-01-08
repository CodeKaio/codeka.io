---
created: "2020-07-15"
date: "2020-07-15T00:00:00Z"
language: en
lastmod: "2022-06-17"
tags:
- devops
- linux
- shell
title: Planifier des commandes Linux avec `at`
slug: schedule-linux-commands
---

Comme je prépare et exécute beaucoup de scripts, j'ai parfois besoin d'exécuter un script à une heure précise de la journée.

Lorsqu'un script ne doit être exécuté qu'une seule fois, `cron` n'est pas une solution viable.
J'ai donc découvert le planificateur `at`.

Vous devez d'abord l'installer, en utilisant `apt` comme d'habitude pour les utilisateurs de debian, ubuntu ou autre dérivés :

```shell
$ sudo apt install at
```

## Planifier l'exécution d'une commande

1. utilisez la commande `at` avec une heure / date
2. saisissez les commandes à exécuter dans l'invite
3. tapez CTRL+D pour quitter (^D)

```shell
$ at 9AM       
warning: commands will be executed using /bin/sh
at> cd workspaces/github/dotfiles
at> git pull
at> < EOT >
job 1 at Sat Apr 16 09:00:00 2022
```

Cet exemple va récupérer le contenu d'un dépôt à 9h demain matin !

`at` prend en charge de nombreuses spécifications de temps.

Voici un extrait de sa page de [manuel](https://fr.manpages.org/at) :

> At permet d'indiquer l'heure de lancement de manière assez complexe, en extension de la norme POSIX.2. Il accepte des spécifications de la forme HH:MM pour exécuter un travail à une heure donnée de la journée en cours (ou du lendemain si l'heure mentionnée est déjà dépassée). On peut aussi lui fournir l'un des arguments suivants : midnight (minuit), noon (midi), ou teatime (l'heure du thé, soit 16 heures). Il est également possible de fournir un suffixe du type AM (matin), ou PM (après-midi) avec une heure sur un cadran de 12 heures.

Les commandes sont exécutées avec le compte de l'utilisateur connecté, en utilisant un shell `/bin/sh`.
Il utilisera les variables d'environnement disponibles du shell au moment où la commande `at` est exécutée, et se placera dans le répertoire courant avant d'exécuter votre script.

## Voir les commandes planifiées

```shell
$ atq
1	Sat Apr 16 09:00:00 2022 a jwittouck
```

## Voir les détails d'une tâche

```shell
$ at -c 1


cd /home/jwittouck || {
	 echo 'Répertoire d'exécution inaccessible' >&2
	 exit 1
}
cd workspaces/github/dotfiles
git pull

```

## Supprimer une tâche

```shell
$ atrm 1
```