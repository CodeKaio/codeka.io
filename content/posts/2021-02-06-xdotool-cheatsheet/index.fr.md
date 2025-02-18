---
created: "2021-02-06"
date: "2021-02-06T00:00:00Z"
language: fr
modified: "2024-01-05"
tags:
- Tools
title: xdotool cheatsheet
slug: xdotool-cheatsheet
---

J'ai beaucoup joué ces jours-ci avec xdotool, pour essayer d'automatiser certaines choses pour mon Elgato Streamdeck.

Voici les choses que j'essaie de faire :

* Sélectionner une fenêtre, et envoyer une séquence clavier (comme CTRL+B pour couper ou rétablir le son d'un appel Teams)
* Taper des emojis dans la fenêtre active 😅
* Déplacer une fenêtre ou la redimensionner

Voici quelques liens que j'ai trouvés à propos de `xdotool` :

* [Xdotool - Window Stack](https://www.linux.org/threads/xdotool-%E2%80%93-window-stack.10687/)
* [Xdotool - Examples](https://www.linux.org/threads/xdotool-examples.10705/#post-36275)

Vous trouverez ci-dessous les commandes que j'ai trouvées utiles au cours de mes recherches.

## obtenir les infos (classe) d'une fenêtre

```shell
xprop | grep 'CLASS'
```

Cliquez ensuite sur la fenêtre que vous voulez analyser.

Plus de détails sur les classes dans ce fil de discussion : [xdotool : what are « class » and « classname » for a window ?](https://askubuntu.com/questions/1060170/xdotool-what-are-class-and-classname-for-a-window)

## trouver la fenêtre active

```shell
xdotool getactivewindow
```

## trouver une fenêtre par classe

```shell
xdotool search --onlyvisible --limit 1 --class "Firefox"
```

## positionner le focus sur une fenêtre

```shell
xdotool windowactivate 123456
```

## envoyer une touche de clavier

```shell
xdotool search --onlyvisible --limit 1 --class "Firefox" key ctrl+t
```

## envoyer une touche de clavier à firefox (en changeant de fenêtre active et en revenant)

```shell
ACTIVE_WINDOW=$(xdotool getactivewindow)
FIREFOX_WINDOW=$(xdotool search --onlyvisible --limit 1 --class "Firefox")
xdotool windowactivate $FIREFOX_WINDOW
xdotool key ctrl+s
xdotool windowactivate $ACTIVE_WINDOW
```

## changer la taille d'une fenêtre

```shell
xdotool search --onlyvisible --limit 1 --class "Firefox" windowsize 800 600 
```
