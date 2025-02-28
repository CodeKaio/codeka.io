---
created: "2025-02-18"
date: "2025-02-18T09:30:00+01:00"
language: fr
tags:
- IA
title: "AI Manifesto"
---

# AI Manifesto

**Aucun** post, **aucune** page de ce blog n'est écrit avec une IA.

Les raisons en sont nombreuses. Je ne suis pas un grand utilisateur d'IA. Cette page, en plus de clamer haut et fort ce postulat, a aussi pour but d'en expliquer les raisons.

## J'aime écrire et être lu

C'est peut-être prétentieux, mais j'aime à penser que les personnes qui viennent lire mes articles le font en partie parce qu'ils espèrent avoir mon point de vue, mon approche, et pas celle d'un LLM.
Mes biais sont suffisants pour ne pas avoir à ajouter ou gérer ceux d'un LLM américain, chinois, ou même français par ailleurs.

Je pense avoir pris goût à l'écriture. Cet exercice (car je le vois comme tel) me force à structurer mes idées, en allant parfois à l'essentiel, en m'assurant que le contenu est clair, précis, et correspond à l'approche que je veux en avoir. Cette approche, principalement autour de la pédagogie et du partage, me définit sur beaucoup de mes activités.

C'est d'ailleurs essentiellement pour cela que j'ai pris le temps d'écrire un livre (spoiler : pas d'IA non plus). L'accomplissement de cet exercice (qui a probablement été l'un des plus difficiles que j'ai eu à mener) n'aurait eu aucune saveur s'il avait été généré par une IA.

Dans mon travail de rédaction, j'aime aussi faire corriger mon contenu textuel par quelqu'un dont c'est le métier : Laura Dhalluin https://lombredelaplume.com/.
Au-delà de la simple correction orthographique, elle s'assure de la cohérence et de la bonne utilisation des termes, supprime les anglicismes, propose des reformulations. C'est aussi une amie et c'est une forme de soutien dans son activité que de solliciter ses services. Je préfère de loin payer quelques euros à Laura pour une correction de qualité, que de payer un abonnement à une IA qui fera un travail (est-ce qu'on peut vraiment utiliser ce terme) de bien moins bonne qualité.

## Mon usage (faible) des IA

### La création de contenu

Je n'ai jamais trop aimé les images rendues par IA. Leur effet "fake" ne me plaît pas et je suis d'autant plus inquiet des photos générées ayant l'air réelles (deepfakes et autres usages malveillants). Je préfère de loin sélectionner une photo libre de droit sur un site comme https://unsplash.com/ ou https://pixabay.com et créditer l'auteur ou l'autrice de la photo _via_ une mention ou un lien. C'est ce que j'essaye de faire pour les photos de couverture de mes articles.

Concernant le contenu textuel, j'utilise parfois un LLM pour affiner un plan que j'ai initialisé à la main ou compléter mes idées. Mais jamais pour la rédaction. Le plus souvent, les LLM ne font que reprendre le plan que je leur ai proposé, en le réorganisant légèrement, et en ajoutant quelques points que j'aurai pu oublier.

### LLM-Driven-Development

Les personnes ayant travaillé avec moi savent que j'utilise aussi assez peu l'IA pour écrire mon code.

Je fais partie de la génération qui a appris à coder sur papier, avec des livres. J'ai souvent combattu ce que j'appelais le "StackOverflow-Driven-Development", qui consiste à copier/coller les réponses marquées d'une coche verte de SO.
Cette pratique présente selon moi plein de défauts : elle n'aide pas à apprendre, elle introduit potentiellement des failles de sécurité, du code de qualité variable avec des pratiques qui peuvent être obsolètes.

Le LLM-Driven-Development suit selon moi le même chemin, mais avec un aspect supplémentaire "rigolo".
Les LLM qui génèrent du code ont des modèles qui ont "appris" sur du code provenant principalement de GitHub. Générer du code en utilisant un outil ayant appris sur du code de qualité non-maîtrisée ne me semble déjà pas forcément une bonne idée.
Ce code généré par ces IA servira ensuite à faire apprendre la génération suivante du modèle de LLM.
Je n'arrive pas à imaginer dans quelle mesure ce processus ne peut pas mener à une dégénérescence globale du code.

On observe déjà aujourd'hui (28 février 2025) des phénomènes inquiétants liés à ce processus. GitClear a publié un rapport sur la qualité observée du code qui semble se dégrader depuis l'adoption massive des IA : https://www.gitclear.com/ai_assistant_code_quality_2025_research.

Autre phénomène, quand on demande à ChatGPT de générer du code front, il génère en premier du code React. Cela peut sembler logique, mais ce phénomène risque d'accélérer l'utilisation de ce framework, et d'enfermer les CTO dans ce "choix".
Est-ce que ChatGPT sait générer du code compatible avec la toute dernière version de Spring-Security, ou génère-t-il du code pour Spring Security 5, qui est le plus présent sur GitHub et donc forcément plus ancré dans son modèle ?
Tout ça pour dire que les personnes ne s'appuyant que sur leur LLM pour écrire le code risquent d'être incapables de l'adapter ou de le maintenir si elles ne savent pas ce qu'elles font.

J'en viens aussi à penser que "On va remplacer les développeurs par des IA" est le take le plus pêté que j'ai pu entendre ces dernières années. Réduire mon métier à la simple écriture du code est ignorant, voire malhonnête. L'IA peut s'avérer utile comme outil pour aider les développeurs oui, pour préparer des revues de code, proposer du refactoring (mais là, un simple LLM ne me semble pas le plus pertinent), expliquer du code, générer des tests.

## Conclusion

Vous l'aurez compris en lisant ces lignes, je ne suis pas un grand utilisateur ou admirateur des IA actuelles.

Je pense que l'IA est un outil formidable, mais que les usages que j'observe massivement aujourd'hui, à savoir générer des images (moches mais c'est mon avis), des deepfakes et du code bancal (encore une fois mon avis), ne sont pas à la hauteur de ce qu'on pourrait espérer.

Je n'aime pas voir apparaître des IA dans tous les outils que j'utilise, sans en avoir besoin.
Leur côté non-déterministe, les "hallucinations" des LLM, et leur dégénérescence qui ne peut que se produire sont pour moi clairement les problèmes.

Je pense que vous méritez mieux que ça.

J'écris moi-même mes articles.

## Inspiration et ressources

Cette page est directement inspirée par ce que j'ai pu voir chez d'autres bloggers, en particulier Denis Germain https://blog.zwindler.fr/ai-manifesto/ et Cassidy Williams https://cassidoo.co/ai/, qui ont eux-même été inspirés par Damola Morenikeji
https://www.bydamo.la/p/ai-manifesto.