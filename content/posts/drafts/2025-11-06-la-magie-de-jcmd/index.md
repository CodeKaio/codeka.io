---
date: 2025-10-31
language: fr
title: "Abracajava : La magie de `jcmd`"
tags:
  - java
  - tools
draft: true
lastmod: 2025-11-14
---

[//]: # (TODO intro)

Cet article présente `jcmd`, quelques cas d'usage bien pratiques, pour comprendre le comportement d'une JVM en prod.

# `jcmd`

`jcmd` est l'un des binaires fournis avec les différentes distributions de JDK. Il existe depuis au moins la version 7 du JDK.

Il est inclus par défaut dans les JDK Eclipse Temurin, OpenJDK et GraalVM.

Attention, il n'est pas forcément inclus dans les JRE, donc pensez à vérifier s'il est bien présent dans votre version de Java.

Il permet d'exécuter des commandes de diagnostic sur une JVM _en cours d'exécution_.
C'est donc un outil bien pratique pour débugger une JVM qui tourne dans un container par exemple.

Certaines de ces commandes ont un impact sur l'application. Il est par exemple possible de forcer l'exécution d'un _Garbage Collector_, ou d'extraire un _Heap Dump_.

Les commandes disponibles évoluent avec les différentes versions de Java.

Je liste ici les plus pratiques pour diagnostiquer une application.

# Exécuter `jcmd`

`jcmd` vient avec une aide en ligne de commande très minimale, si on lui passe le paramètre `--help` :

```shell
$ jcmd --help
Usage: jcmd <pid | main class> <command ...|PerfCounter.print|-f file>
   or: jcmd -l                                                    
   or: jcmd -h                                                    
                                                                  
  command must be a valid jcmd command for the selected jvm.      
  Use the command "help" to see which commands are available.   
  If the pid is 0, commands will be sent to all Java processes.   
  The main class argument will be used to match (either partially 
  or fully) the class used to start Java.                         
  If no options are given, lists Java processes (same as -l).     
                                                                  
  PerfCounter.print display the counters exposed by this process  
  -f  read and execute commands from the file                     
  -l  list JVM processes on the local machine                     
  -? -h --help print this help message  
```

Lorsqu'on exécute `jcmd` sans paramètre, il liste l'ensemble des JVM en cours d'exécution, ainsi que leur PID (_Process ID_) sur le système d'exploitation.
Cette exécution est alors très similaire à ce que l'on peut faire avec `jps`.
On obtient le même résultat en exécutant `jcmd` avec le paramètre `-l` (pour _list_) :

```shell
$ jcmd
308099 jdk.jcmd/sun.tools.jcmd.JCmd
2909 fr.univ_lille.gitlab.classrooms.GitlabClassroomsApplication

$ jps
308100 Jps
2909 GitlabClassroomsApplication
```

On observe que `jcmd` et `jps` s'auto-découvrent (comme le fait la commande `ps` d'ailleurs).

À noter que `jcmd` ne liste pas les process qui tournent dans des containers séparés. Il faut donc parfois utiliser `jcmd` directement dans les containers (ce qui est mon _use-case_ principal).

Une fois qu'on a identifié le PID de la JVM concernée (2909 dans mon exemple), on peut invoquer la sous-commande `help` pour obtenir la liste des commandes disponibles d'une JVM
(j'ai tronqué le résultat par souci de lisibilité) :

```shell
$ jcmd 2909 help
2909:
The following commands are available:
Compiler.CodeHeap_Analytics
Compiler.codecache
[...]
VM.uptime
VM.version
help

For more information about a specific command use 'help <command>'.
```

Et de la même manière, on peut alors récupérer l'aide d'une sous-commande particulière.
Un premier exemple avec la commande `VM.uptime` : 

```shell
 $ jcmd 2909 help VM.uptime
2909:
VM.uptime
Print VM uptime.

Impact: Low

Syntax : VM.uptime [options]

Options: (options must be specified using the <key> or <key>=<value> syntax)
        -date : [optional] Add a prefix with current date (BOOLEAN, false)
```

Et un autre exemple avec la commande `VM.flags` :
```shell
 $ jcmd 2909 help VM.flags
2909:
VM.flags
Print VM flag options and their current values.

Impact: Low

Syntax : VM.flags [options]

Options: (options must be specified using the <key> or <key>=<value> syntax)
        -all : [optional] Print all flags supported by the VM (BOOLEAN, false)
```

En plus de la description de la commande et de ses options supportées, `jcmd` indique l'impact de l'exécution de la commande sur la JVM.
L'impact est classé en 3 niveaux :

* _Low_, avec potentiellement la mention _no impact_ : ces commandes ont donc un impact réduit sur la JVM. Il s'agit principalement des commandes qui listent les paramètres ou activent des options. À utiliser sans modération ;
* _Medium_, avec parfois une indication de ce qui fait varier cet impact, comme _depends on number of threads_ ou _depends on number of classes_. Utiliser ces commandes sur une JVM en production est donc acceptable et sera le plus souvent invisible ;
* _High_, avec souvent l'indication _depends on heap size_. Utiliser ces commandes occasionnera un impact fort sur les performances de la JVM et le code en cours d'exécution, potentiellement même une pause dans l'exécution pour l'exécution d'un _garbage collector_ ou la création d'un _heap dump_.

> On retrouve tous ces éléments dans la doc de `jcmd` sur [docs.oracle.com](https://docs.oracle.com/en/java/javase/25/docs/specs/man/jcmd.html)

Maintenant qu'on a compris comment fonctionnait `jcmd`, on va pouvoir voir quelles sont les commandes les plus utiles pour diagnostiquer une JVM.

# Les commandes les plus utiles

> Par "utile" je veux surtout dire "utile pour moi", ne soyez pas vexés si votre commande préf n'est pas dans cette liste.

## Lister les informations de la JVM

* `VM.flags` : liste les options de la JVM
* `VM.version` : affiche la version de la JVM
* `VM.uptime` : affiche l'uptime de la JVM

# Liens et références

* La doc de `jcmd` sur [docs.oracle.com](https://docs.oracle.com/en/java/javase/25/docs/specs/man/jcmd.html)

---

Photo de couverture par [Aron Visuals](https://unsplash.com/@aronvisuals?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) sur [Unsplash](https://unsplash.com/photos/multicolored-plasma-ball-in-dim-light-room-_l4yffWjgt4?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)
