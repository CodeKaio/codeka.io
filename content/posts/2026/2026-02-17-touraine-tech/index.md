---
date: 2026-02-17
language: fr
title: Touraine Tech 2026 - Premier voyage au pays de la rillette
---

Dans ma tournée des conférences tech de France pour jouer à Factorio, je suis passé par Touraine Tech, la semaine dernière.
Comme toujours, c'est l'occasion de faire de nouvelles rencontres et de croiser les autres speakeuses et speakers.

Voici donc le récap de ma visite à Touraine Tech, pour cette 8ᵉ édition, mais la première pour moi.

<!--more-->

## Back to school, les bancs de la fac

Touraine Tech 2026 a eu lieu à la Faculté des Sciences et Techniques de l'Université de Tours.

[20260213_152718.webp](20260213_152718.webp)

C'est amusant de se balader dans les couloirs de la fac et de s'assoir dans les amphis.
On retrouve sur les tables les âmes d'artistes des étudiants de Tours.

[20260212_123033.webp](20260212_123033.webp)

Les amphis sont très beaux. En tant que speaker, c'est un plaisir d'être présent devant le public dans de si belles salles. Le public de Touraine Tech est très sympa, souriant, et généreux en feedbacks. J'ai ressenti un bon _mood_ sur ces deux jours.

## Les talks du jeudi

La journée du jeudi était déjà bien chargée, j'ai assisté à six talks.

### Keynote - Tech Vs Bots - Debuggons la démocratie - par Clément Hammel

Dans cette première keynote, Clément Hammel explique le principe de _social listening_ et montre quelques exemples d'analyse des réseaux sociaux.

Un des exemples mis en avant est celui de la mobilisation du monde agricole par rapport à la crise de la dermatose nodulaire contagieuse récente.
Il nous explique alors la méthode d'analyse qui consiste à sélectionner un ensemble de tweets (en Français, et dans un temps dédié), leur classification en fonction de mots clés dans les bios des auteurs, la détection de communautés _via_ les interactions entre les auteurs et les tweets, ainsi que la duplication de contenus.

Le résultat met en avant une cartographie, et les polarités des différentes communautés.
Sont aussi analysées les dates de création des comptes, avec des pics de création pendant certains événements, comme l'investiture de Donald Trump.

Le constat est sans appel, la manipulation politique de cette crise est évidente, puisque les agriculteurs, premiers concernés par cet évènement, sont très peu présents dans le discours face aux mouvements politiques (plutôt situés à droite).

Pour conclure, Clément pousse la réflexion au monde de l'IA Gen, qui est entraînée sur ces contenus (en particulier Grok, entraîné sur le contenu de Twitter donc) : la propagande russe infiltre l'IA Gen.

Il nous a également présenté l'initiative Française VigiNum (unique au monde), qui vise à surveiller les tendances sur les réseaux sociaux, et détecter les tentatives d'ingérences qui pourraient être menées par certains états, ainsi que quelques axes permettant de riposter face à ces attaques.

[20260212_093554.webp](20260212_093554.webp)

## L'architecture hexagonale au pays des irréductibles développeurs - par Nathan Castelein et Ambre Person

Ordralfabétix (Ambre avec ses couettes et muni de son poisson pas frais) et Cétautomatix (Nathan avec son tablier et son marteau), présentent chacun leur vision de l'architecture hexagonale, pour implémenter Jarvix, l'Intelligence Armoricaine.

[20260212_101047.webp](20260212_101047.webp)

Le ton est vite posé. Ce talks est un des plus drôles que j'ai vus récemment. Tout est prétexte à un jeu de mots en _ix_ : une partie de _Chifoumix_ : Menhir/Parchemin/Serpe, la commande `curlix` pour exécuter des requêtes HTTP, le serveur de cache _Redix_.

Le jeu entre les deux speakers fonctionne bien, sur fond de rivalité éternelle entre ces deux personnages.

Les règles de l'architecture hexagonale sont bien présentées et illustrées avec le village d'Astérix et ses portes gardées, et par des implémentations en Java et en Go du métier de druide pour préparer la recette de la potion magique. Et c'est là tout l'intérêt du talk, d'y voir illustrés ces deux implémentations très différentes, et comment les concepts de ports et adapters peuvent être repris dans ces langages aux concepts différents.

Et on y découvre enfin la recette de la potion magique (j'en ai une photo, mais c'est un secret).

> C'est un talk que je conseillerai à mes étudiants, et qui a beaucoup plus au public. Bravo aux speakers.

> PS : Ce talk m'a donné envie d'intégrer une section sur l'architecture hexagonale dans Factorio. Ça traînait dans ma tête depuis un moment, mais ça s'est débloqué en voyant le village illustré, donc ça m'aura bien servi.

## Le hasard fait bien les tests - par David Pilato

Les tests un peu _flakky_, on connait. Parfois, certains éléments aléatoires et difficiles à reproduire font échouer des tests.

David présente l'utilisation de [_Randomized Testing_](https://labs.carrotsearch.com/randomizedtesting.html), une librairie de tests Java développée par Carrot Search, qui s'intègre avec JUnit.

[20260212_123700.webp](20260212_123700.webp)

Randomized Testing permet de facilement générer des données aléatoires de tout type (nombres, chaînes de caractères, dates, locales, etc), pour les utiliser dans les tests.
Chaque exécution est associée à un _seed_ qui va rendre le test reproductible.
Un axe est aussi d'exécuter les tests dans des répétitions pour maximiser les cas couverts.

Un seul point négatif : la librairie est conçue pour être utilisée avec JUnit 4.

## Les talks du vendredi

Le vendredi matin, j'étais très fatigué (j'ai passé une mauvaise nuit). J'ai donc passé la matinée à peaufiner mon talk (j'ai ajouté une partie sur l'architecture hexagonale) et à discuter avec les speakers dans notre salle réservée.

J'ai profité du regain d'énergie de l'heure du midi pour aller voir quelques talks _Lightning_ de 15 minutes.



## Une belle démarche de transparence

La keynote de cloture du vendredi soir (juste après ma game de Factorio), a été l'occasion pour les orgas de remercier tout le monde : sponsors, speakeuses et speakers, et le public présent.

[20260213_174521.webp](20260213_174521.webp)

L'orga en a profité pour afficher la situation financière de l'association derrière l'évènement.
Pas de surprise, comme partout, les dernières années sont difficiles. Touraine Tech 2026 est quand même à l'équilibre, mais sans le soutien des sponsors, on voit que ça va être difficile.

Je souhaite à toute l'équipe que les sponsors se remobilisent les prochaines années, pour que ce bel event puisse perdurer.

[20260213_174817.webp](20260213_174817.webp)

## En conclusion

J'ai passé une super conférence. L'accueil des orgas aux petits soins, les soirées autour d'un verre et d'une planche qui passent beaucoup trop vite, la qualité des interventions, et le public définitivement très agréable. Tout ça fait que je reviendrai avec plaisir à Touraine Tech 2027.
