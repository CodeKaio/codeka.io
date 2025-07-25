---
date: 2025-07-25
language: fr
title: Montage d'un clavier mécanique
tags:
  - internet
---

Ça faisait un moment que j'avais envie de tester un clavier _split_.
J'ai franchi le pas.
Je voulais le monter moi-même, mais ne pas avoir une difficulté trop élevée, ni risquer de cramer un composant en le soudant (mal).

J'ai donc cherché des revendeurs qui proposent des kits à monter.
C'est l'assurance pour moi d'avoir tous les composants compatibles, sans me tromper, et d'avoir aussi le côté personnalisation que je voulais.

Je voulais un clavier _split_ avec au moins quatre lignes de touches (parce que oui, il existe des modèles très petits, sur trois lignes).
J'ai donc choisi le modèle [_Elora_](https://splitkb.com/collections/keyboard-kits/products/halcyon-elora) chez [splitkb.com](https://splitkb.com).
Ce modèle est le plus grand modèle qu'ils ont à disposition. Il dispose de 62 touches, dont plusieurs situées au niveau des pouces. Il est aussi possible de positionner deux modules au-dessus de chaque partie, au choix parmi un petit écran, un bouton rotatif ou un trackpad. 
J'ai choisi l'écran pour la partie gauche, pour afficher les infos du clavier (layer, caps lock, etc.)  (et peut-être coder un peu dessus si j'ai le temps), ainsi qu'un bouton rotatif pour pouvoir régler le volume de ma musique que je mettrai sur la partie droite du clavier.

Le kit est arrivé bien emballé.

{{< img-resized src=20250710_131320.jpg >}}

Les instructions de montage sont sur une doc en ligne et ont l'air plutôt claires.

Je vais devoir monter le clavier (les deux parties séparément), puis flasher le firmware de chaque contrôleur. Ensuite, ce sera le moment de paramétrer le mapping des touches, et aussi de (ré)apprendre à taper au clavier (ouch).

C'est parti !

## Le kit

Une fois le carton déballé, on se rend compte de l'ensemble des pièces qui composent le kit.

{{< img-resized src=20250711_085344.jpg >}}

Le kit contient donc les deux cartes du clavier, les switches (compatibles MX), un kit de touches, les plaques de protection du clavier avec pour chaque côté une mousse, une plaque décorative en aluminium et une base en plexi. On a aussi les câbles USB-C, les deux modules, et un peu de visserie.

La carte de chaque partie du clavier contient le contrôleur déjà soudé, avec une puce RP2040 et deux ports USB-C.
Ici la partie gauche du clavier :

{{< img-resized src=20250711_090103.jpg >}}

On distingue bien les encoches dans lesquelles viendront se plugger les switches, ainsi que les LED pour chacune des touches.

{{< img-resized src=20250711_090114.jpg >}}

Chaque partie du clavier est associée à une plaque métallique (en aluminium propre) et à une petite plaque de mousse.
C'est dans ces plaques que viendront se loger les touches du clavier.

{{< img-resized src=20250711_091051.jpg >}}

## Monter les plaques

La première étape consiste à visser la plaque en alu ainsi que la mousse sur la carte du clavier.
C'est plutôt facile, ça se fait avec un petit tournevis plat, tout s'emboite facilement.

{{< img-resized src=20250711_091424.jpg >}}

Cette première étape était plutôt facile, ça commence à prendre forme !

## Les switches

J'ai choisi les switches proposés par le site web, à savoir des _Kailh Pro Purple_.
Ils ont un toucher "tactile", et devraient être plus silencieux que mon clavier actuel (qui est en _Cherry MX Brown_).

{{< img-resized src=20250711_091846.jpg >}}

Ces switches sont aussi compatibles avec les touches "MX", ce qui me permettra de réutiliser des touches de mon clavier actuel, et de pouvoir acheter quelques touches customisées si le cœur m'en dit.

Les switches viennent se clipser dans la plaque en alu, et sont ensuite être insérés dans la carte en poussant légèrement.
C'est assez simple à faire.

{{< img-resized src=20250711_092639.jpg >}}

Au montage, je me rends compte que changer les switches risque d'être un peu compliqué. J'ai pris une pince pour ça, mais c'est une opération assez délicate, car la plaque métallique accroche fortement.

Bien que le kit soit vendu comme étant _hot swap_ pour les switches, car ils ne sont pas soudés, l'opération n'est clairement pas pensée pour être faite tous les jours.

Poser les switches prend quelques minutes, une fois l'opération finie, ça prend une forme plutôt sympa, on devine bien le futur clavier.

{{< img-resized src=20250711_093125.jpg >}}

L'étape suivante, c'est le petit module avec l'écran !

## Le module

Pour limiter les soudures, les modules sont déjà assemblés, et s'installent avec un petit cable de type _nappe_.
C'est un peu plus technique ici.

Un tout petit écran !

{{< img-resized src=20250711_093432.jpg >}}
{{< img-resized src=20250711_093441.jpg >}}
{class=images-grid-2}

L'installation d'un module est un peu plus technique, il faut réussir à glisser la première nappe dans le connecteur, puis visser le module, et connecter l'autre côté.
C'est plus fin, mais ce n'est pas non plus très compliqué.

{{< img-resized src=20250711_094328.jpg >}}
{{< img-resized src=20250711_094335.jpg >}}
{class=images-grid-2}

Il faudra ensuite flasher tout ça, quand le montage sera fini.

Pour protéger l'écran, j'ai aussi pris une petite _cover_, qui s'installe en trois coups de vis.
Le kit contient des covers pour les deux côtés du clavier, ainsi que pour tous les types de boutons.
Avec la cover, mon petit écran est entièrement couvert, et protégé raisonnable des rayures (et des éclaboussures de café ahaha)/

{{< img-resized src=20250711_095218.jpg >}}

## La plaque de socle

La plaque de socle va supporter le carte du clavier.
Je les ai choisies en plexiglass transparent, pour pouvoir admirer la carte, il existe aussi un modèle "fumé", mais je trouvais ça moins cool.

Elle se visse sur la carte, et on ajoute ensuite les petits "bumps" plastiques qui vont venir dans chaque coin.

{{< img-resized src=20250711_100023.jpg >}}

## Les touches

Une fois tout ça monté, on peut installer les touches du clavier.
J'ai installé les touches suivant le pattern azerty classique pour l'instant, je les changerai probablement plus tard.

{{< img-resized src=20250711_100644.jpg >}}

## La deuxième moitié du clavier

La deuxième partie du clavier s'assemble exactement comme la première, j'ai été plus efficace sur cette deuxième moitié.
Sur cette partie droite, j'ai posé le bouton rotatif, il se pose de la même manière que l'écran.

{{< img-resized src=20250711_111608.jpg >}}

Je trouve que le rendu est plutôt sympa et propre.

## Flasher le firmware

Maintenant que tout est monté, il faut flasher chaque moitié du clavier.

{{< img-resized src=20250711_112134.jpg >}}

Pour ce faire, il suffit de brancher le clavier sur un port usb 
Le firmware est un simple fichier à déposer sur le device USB qui apparait alors.
Il faut exécuter l'opération pour les deux moitiés du clavier, cela prend quelques secondes.

Les firmwares sont dispos sur le site du fabricant, et sont directement paramétré avec le support de l'écran ou du bouton. C'est très facile à récupérer, rien à compiler.

Une fois l'opération faite, le clavier s'illumine, et est prêt â être utilisé.

{{< img-resized src=20250711_112832.jpg >}}

## Configurer le mapping

Le mapping des touches se configure avec un outil nommé 'Vial'.
L'outil est disponible sous une forme binaire, ou peut aussi fonctionner sous Chrome (avec la magie de l'API WebUSB / HID).

Il est plutôt pratique d'utilisation. Il suffit de sélectionner la touche à configurer et d'y associer le caractère souhaité. La touche est alors immédiatement paramétrée.

{{< img-resized src=vial-rocks.png >}}

On reste contraint par les différents mappings supportés et implémentés dans les OS.

## La suite

Il faut maintenant que je m'habitue à deux choses : le QWERTY, qui semble quand même plus pratique pour ce type de clavier, et la disposition des touches en colonne. Mon cerveau va avoir du taf dans les prochains mois.

Je vais passer un peu de temps à m'entrainer sur https://www.keybr.com/.

J'ai aussi un peu galéré à trouver le bon mapping qui me permet d'avoir toutes les touches que je veux, en particulier le support des accents.

J'en suis à une semaine d'utilisation et pour le moment, j'ai environ atteint 20% de mes capacités de frappe sur mon clavier précédent sur lequel je tape à 70 wpm avec 95% de précision. La marge de progression est énorme !

## Liens et références

* https://splitkb.com/ : le site sur lequel j'ai acheté mon clavier
* https://docs.splitkb.com/product-guides/halcyon-series/build-guide/introduction : la doc de montage de mon clavier
* https://www.keybr.com/ : pour s'entraîner