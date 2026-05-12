---
title: "Gestion de dotfiles avec `chezmoi`"
slug: dotfiles-chezmoi
tags:
  - tools
  - linux
draft: true
---

Je travaille sur plusieurs machines différentes : mon NUC qui est ma machine principale, et mon laptop qui me sert en mobilité et en déplacement.

Pour harmoniser mes configurations, j'avais depuis un certain temps un repo Git, sur lequel de déposais quelques scripts customisés et des dotfiles que je gérais à la main.

J'ai décidé d'outiller ce workflow de manière plus sérieuse avec l'outil [chezmoi](https://www.chezmoi.io/).

## Installation et configuration basique

J'utilise `mise` pour installer mes outils. Donc l'installation de `chezmoi` est une simple commande :

```shell
$ mise use -g chezmoi

$ chezmoi --version
chezmoi version v2.70.2, commit b1aacd726df533ffd1f2fec7ded8e2ecfdb26e0e, built at 2026-04-17T00:17:32Z, built by goreleaser
```

> Ça pose d'ailleurs une question de l'oeuf ou de la poule, puisque mes dotfiles de `mise` seront également gérés avec `chezmoi`. Lorsque je vais installer une nouvelle machine, je pense que le premier outil que je devrai installer c'est bien `mise`, puis `chezmoi` pour récupérer mes dotfiles, y compris ceux de `mise`.

Une fois l'installation terminée, la commande `chezmoi init` permet d'initialiser une nouvelle machine.

```shell
$ chezmoi init
```

Cette commande crée un nouveau répertoire `.local/share/chezmoi` et l'initialise avec `git` (pour pouvoir partager les configurations entre plusieurs machines).

La commande `chezmoi cd` permet de se déplacer dans le répertoire de configuration de chezmoi pour voir ce qu'il s'y passe, et exécuter les commandes `git`.

```shell
$ chezmoi cd

$ pwd
/home/jwittouck/.local/share/chezmoi

$ ls -a
 .   ..   .git
```

L'installation est maintenant terminée, on peut commencer à ajouter quelques dotfiles.

## Ajout de dotfiles

L'ajout de dotfiles à `chezmoi` se fait avec la commande `chezmoi add`.

```shell
$ chezmoi add .vimrc
```

Une fois cette commande exécutée, le fichier `.vimrc` est copié dans `.local/share/chezmoi`, on le retrouve avec un simple `ls`:

```shell
$ chezmoi cd

$ ls
```