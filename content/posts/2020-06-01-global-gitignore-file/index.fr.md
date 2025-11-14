---
created: "2020-04-16"
date: "2020-06-01T00:00:00Z"
language: fr
lastmod: "2022-06-17"
tags:
- devops
- git
title: "Configurer un .gitignore global \U0001F648"
slug: global-gitignore-file
---

Cet article explique comment configurer un fichier `.gitignore` global, pour exclure des fichiers ou des répertoires pour tous vos dépôts git.

Utiliser un fichier global permet d'ignorer des fichiers dans l'ensemble des répertoires Git de votre poste.
C'est très utile pour certains fichiers, comme les fichiers `.env`. Cela empêche surtout les commits accidentels.
J'ignore aussi les répertoires communs pour les développements liés à Java et NodeJS (`target/` et `node_modules`), ainsi que les fichiers IntelliJ IDEA (`*.iml` et `.idea/`).

Cela dit, vous devriez toujours configurer un fichier `.gitignore` dans vos projets, car le fichier global ne fonctionne que pour vous et ne sera pas partagé avec le code de votre projet.

## Créer le fichier `.gitignore`

Sur les systèmes Linux, l'emplacement par défaut pour un fichier `.gitignore` global est `~/.config/git/ignore`.

Créez ce fichier s'il n'existe pas, et mettez-y votre contenu :

```shell
# create the directory if it doesn't exists
$ mkdir -p ~/.config/git

# create the global ignore file
$ cat <<EXCL >> ~/.config/git/ignore
# global gitignore file

# idea settings
.idea/
*.iml

# java
target/

# direnv
.envrc
.env
EXCL
```

Et c'est tout !