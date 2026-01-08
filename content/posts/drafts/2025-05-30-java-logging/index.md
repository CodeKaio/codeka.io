---
date: 2025-05-30
language: fr
title: Introduction au logging en Java 
series: Logging en Java
tags:
  - java
  - observability
draft: true
---

Cette série d'articles, en trois parties, a un objectif simple : faire l'état des lieux du logging en Java.
Je vais donc me concentrer sur mon langage préféré, mais les principes et les bonnes pratiques que je vais illustrer pourront certainement être appliquées à d'autres langages, ou à des domaines connexes au développement, administration système ou réseau (dont je ne suis pas du tout un expert).

Dans ce premier article, je vais poser quelques définitions, expliquer les différents types de logs et leurs usages, et lister quelques standards qu'il faut connaître. 

Dans un deuxième article, je ferai le tour des différents frameworks de log qui existent en Java, avec leurs forces, faiblesses et des exemples concrets.
Ce deuxième article vous permettra de pouvoir faire des choix éclairés en fonction de votre contexte.

Enfin, dans un troisième article, je donnerai des bonnes pratiques, des exemples d'optimisation, et des techniques de logging qu'on pourrait qualifier d'avancées (même si elles restent simples).

## C'est quoi un _log_ ?

> Journal du capitaine, date stellaire 7413.4. 
> 
> Avec l'arrivée et l'aide de Mr. Spock, nous avons pu reparamétrer les moteurs à pleine capacité. Nous sommes maintenant à moins d'une heure de voyage de notre destination : la Terre.
>
> James T. Kirk

Un log est une trace d'un évènement dans un système informatique. Ces traces ont plusieurs buts : détecter les erreurs et pouvoir remonter à leur source, lister les actions qui ont été exécutées dans un système, ou aider au débogage pendant le développement (oui je vous vois avec vos `console.log()`).

Lorsque le capitaine Kirk enregistre une information dans son journal (_Captain's log_), il souhaite tracer une information importante. Il mentionne plusieurs éléments importants : la date d'enregistrement 7413.4, qui correspond au 11 novembre 2279 à 17h05, l'évènement, l'aide apportée par Spock, et la conséquence de l'évènement : l'USS Enterprise n'est plus qu'à une heure de trajet de la Terre. 

Un log informatique contient en général des informations similaires. Voici un log issu de cette session d'écriture dans l'outil que j'utilise pour ce blog (Hugo) :

```text
Change detected, rebuilding site (#4).
2025-05-12 19:02:01.560 +0200
Source changed /posts/2025-05-30-java-logging/index.md
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Total in 149 ms
```

Ce log informatique, sur plusieurs lignes, contient plusieurs informations : l'évènement qui a eu lieu, un changement de fichier a été détecté, on voit également que c'est le quatrième évènement de ce type qui est tracé, la date de l'évènement est aussi tracée (dans un format lisible qui ressemble au format ISO-8601, pas en date stellaire cette fois-ci), et la conséquence de ce changement, le site a été re-construit en 149 millisecondes.

Si on prend un log tel qu'on le retrouve le plus souvent dans des applications Java, on observe des lignes de ce type :

```text
2025-05-12T19:18:17.106+02:00  INFO 53185 --- [  restartedMain] f.u.g.c.GitlabClassroomsApplication      : Started GitlabClassroomsApplication in 4.155 seconds (process running for 4.49)
```

Cette ligne contient encore une fois le même type d'informations, une date (au format ISO-8601, les fameux "TZ") et l'évènement qui a eu lieu (l'application est démarrée).
En complément, nous avons aussi une sévérité du message : `INFO`, et des informations sur le composant qui a émis le message, la classe Java `GitlabClassroomsApplication`.

Un log trace donc au minimum un évènement et sa date d'occurrence. D'autres informations peuvent venir s'y ajouter.

## Les types de log

Maintenant que l'on sait ce qu'est un log, on peut essayer de les  regrouper dans trois grandes familles : les logs d'accès, les logs d'audit et les logs applicatifs.

Les logs d'accès portent bien leur nom : on souhaite tracer l'accès à une ressource. Les logs d'accès sont souvent liés à des évènements réseau : une requête émise par la machine d'adresse IP 172.17.0.34 a contacté la machine d'adresse IP 172.17.0.22 sur le port 5432. Les logs d'accès sont alors produits par des équipements réseau (firewall), ou des systèmes intermédiaires comme des reverse-proxy, ou des API managers.

## Liens et références

* Convertir une date stellaire : https://stardatecalculator.com/