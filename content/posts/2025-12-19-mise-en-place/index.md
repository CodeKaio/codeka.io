---
date: 2025-12-19
language: fr
title: Mise en place 
tags:
  - linux
  - tools
draft: true
---

[//]: # (TODO image de couverture)

J'ai découvert Mise par deux canaux :

* le support par Clever Cloud (plutôt en tant que package manager donc), avec en particulier leur type d'image "Linux"

[//]: # (TODO remettre le lien)
* le calendrier de l'avent de Siegfried Erhet, avec une approche de DevTool

Étant utilisateur de `direnv` depuis plusieurs années, je suis habitué à travailler dans mon shell.
Direnv souffre par contre de plusieurs défauts : il ne peut pas gérer d'alias et se configure uniquement avec du scripting shell.

Aujourd'hui, je teste l'outil "Mise", pour voir dans quelle mesure cet outil peut venir enrichir mon workflow de dev.

<!--more-->

## Mise en place

Mise en place est un outil développé en Rust, ce qui promet de bonnes performances et de la stabilité.

Il propose 3 types d'usage :

* gestionnaire de packages (un peu comme `asdf`)
* gestionnaire de variables d'environnement (exactement comme `direnv` donc)
* exécution de tâches (comme on pourrait le faire avec `make`, `task` ou `just`)

Il se veut aussi extensible avec la possibilité d'utiliser ou d'implémenter des plugins.

## L'installation

L'installation est bien documentée sur le site web : https://mise.jdx.dev/installing-mise.html

Sur mes postes Linux, j'ai suivi la procédure d'installation avec `apt` :

```shell
curl -fSs https://mise.jdx.dev/gpg-key.pub | sudo tee /etc/apt/keyrings/mise-archive-keyring.pub 1> /dev/null
echo "deb [signed-by=/etc/apt/keyrings/mise-archive-keyring.pub arch=amd64] https://mise.jdx.dev/deb stable main" | sudo tee /etc/apt/sources.list.d/mise.list
sudo apt update -y
sudo apt install -y mise
```

Une fois l'installation terminée, la commande `mise` est disponible dans mon shell :

```shell
mise --version
              _                                        __              
   ____ ___  (_)_______        ___  ____        ____  / /___ _________
  / __ `__ \/ / ___/ _ \______/ _ \/ __ \______/ __ \/ / __ `/ ___/ _ \
 / / / / / / (__  )  __/_____/  __/ / / /_____/ /_/ / / /_/ / /__/  __/
/_/ /_/ /_/_/____/\___/      \___/_/ /_/     / .___/_/\__,_/\___/\___/
                                            /_/                 by @jdx
2025.12.10 linux-x64 (2025-12-16)
```

Une dernière étape dans la configuration est importante : la configuration du shell, afin que le shell appelle la commande `mise` lors du changement de répertoires.

Cette opération se fait en éditant le fichier de configuration de votre shell : `~/.bashrc` ou `~/.zshrc` par exemple.

Pour Zsh que j'utilise, il faut alors simplement ajouter la ligne suivante dans le fichier `~/.zshrc` :

```shell
eval "$(mise activate zsh)"
```

Une fois ces opérations faites, il suffit donc de redémarrer votre shell pour que la configuration soit prise en compte.

## Les devtools

Quand je switche entre mes projets, je switche aussi souvent de stack d'outillage, que ce soit la version de Java, de Node, ou d'outils divers (comme Hugo que j'utilise sur ce site).

Pour les outils les plus courants, les langages de programmation principalement, mise dispose d'un plugin _core tools_ qui supporte leur configuration et activation automatique.
Pour les autres outils ou packages, _mise_ utilise un ou plusieurs backends (_asdf_ en fait partie) qui permettent l'installation sans avoir rien à configurer en plus. Les outils disponibles peuvent être listés avec la commande `mise registry` :

```shell
$ mise registry
```


---

La prochaine publication est prévue autour du  🗓️

Photo de couverture par sur 