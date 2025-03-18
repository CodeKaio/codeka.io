---
created: "2025-03-07"
date: "2025-03-18"
language: fr
aliases:
  - ai
  - ai-manifesto
type: post
tags:
- IA
title: "AI Manifesto"
---

**Aucun** _post_, **aucune** page de ce blog n'est écrit avec une IA.

Les raisons en sont nombreuses. Je ne suis pas un grand utilisateur des IA en général. Cette page, en plus de clamer haut et fort ce postulat, a aussi pour but d'en expliquer les raisons.

## J'aime écrire et être lu

C'est peut-être prétentieux, mais j'aime à penser que les personnes qui viennent lire mes articles le font en partie parce qu'elles espèrent avoir mon point de vue, mon approche, et pas celle d'un LLM (_Large Language Model_).
**Mes biais sont suffisants** pour ne pas avoir à ajouter ou gérer ceux d'un LLM américain, chinois, ou même français d'ailleurs.

Je pense aussi avoir pris goût à l'écriture. Cet exercice (car je le vois comme tel) me force à structurer mes idées, en allant parfois à l'essentiel, en m'assurant que le contenu est clair, précis, et correspond à l'approche que je veux en avoir. Cette approche, principalement autour de la pédagogie et du partage, me définit sur beaucoup de mes activités.

C'est d'ailleurs essentiellement pour cela que j'ai pris le temps d'écrire un livre (_spoiler_ : pas d'IA non plus). L'accomplissement de cet exercice, qui a probablement été l'un des plus difficiles que j'aie eu à mener dans ma vie, n'aurait eu aucune saveur s'il avait été généré par une IA.

Dans mon travail d'écriture, j'aime aussi faire corriger mon contenu textuel par quelqu'un dont c'est le métier : [Laura Dhalluin](https://lombredelaplume.com/).
Au-delà de la simple correction orthographique, elle s'assure de la cohérence et de la bonne utilisation des termes, remplace les anglicismes, propose des reformulations. C'est aussi une amie et c'est une forme de soutien dans son activité que de solliciter ses services. Je préfère de loin payer quelques dizaines d'euros à Laura pour une correction de qualité, que de payer un abonnement à une IA qui fera un travail (est-ce qu'on peut vraiment utiliser ce terme) de bien moins bonne qualité.

## Mon usage (faible) des IA

### Dans la création de contenu

Je n'ai jamais trop aimé les images rendues par IA. Leur effet _fake_ ne me plaît pas et je suis d'autant plus inquiet des photos générées ayant l'air réelles (_deepfakes_ et autres usages malveillants). Je préfère de loin sélectionner une photo ou illustration libre de droits sur un site comme https://unsplash.com/ ou https://pixabay.com et créditer l'auteur ou l'autrice _via_ une mention ou un lien. C'est ce que j'essaye de faire pour les photos de couverture de mes articles.

Concernant le contenu textuel, j'utilise parfois un LLM pour affiner un plan que j'ai commencé à la main ou pour compléter mes idées. Mais jamais pour la rédaction. Le plus souvent, les LLM ne font que reprendre le plan que je leur ai proposé, en le réorganisant légèrement, et en ajoutant quelques points que j'aurais pu oublier.

### LLM-Driven-Development

Les personnes ayant travaillé avec moi savent que j'utilise aussi assez peu l'IA pour écrire mon code.

Je fais partie de la génération qui a appris à coder sur papier et avec des livres. J'ai souvent combattu ce que j'appelais le _StackOverflow-Driven-Development_, qui consiste à copier/coller les réponses marquées d'une coche verte de _StackOverflow_.
Cette pratique présente selon moi plein de défauts : elle n'aide pas à apprendre, elle introduit potentiellement des failles de sécurité, du code de qualité variable avec des pratiques qui peuvent être obsolètes.

Le _LLM-Driven-Development_ suit selon moi le même chemin, mais avec un aspect supplémentaire « rigolo ».
Les LLM qui génèrent du code ont des modèles qui ont « appris » sur du code provenant principalement de GitHub. Produire du code en utilisant un outil ayant appris sur du code de qualité non maîtrisée ne me semble déjà pas forcément une bonne idée.
Ce code généré par ces IA servira ensuite à faire apprendre la génération suivante du modèle de LLM.
Je n'arrive pas à imaginer dans quelle mesure ce processus ne pourrait pas mener à une **dégénérescence** globale du code.

On observe déjà aujourd'hui des phénomènes inquiétants liés à ce processus. GitClear a publié un [rapport](https://www.gitclear.com/ai_assistant_code_quality_2025_research) sur la qualité observée du code d'une centaine de _repositories_ _open-source_ (dont ceux de Chromium, VSCode, React, Electron, Elasticsearch, Ansible, et beaucoup d'autres) qui semble se dégrader depuis l'adoption massive des IA dans nos métiers. La quantité de code dupliqué a tendance à augmenter, ce qui aura probablement pour effet l'augmentation du nombre de défauts constatés, ainsi que l'augmentation de la complexité de leur résolution. Le code écrit à l'aide d'un assistant IA semble aussi plus soumis à des bugs.

[Un papier de Microsoft](https://www.microsoft.com/en-us/research/publication/the-impact-of-generative-ai-on-critical-thinking-self-reported-reductions-in-cognitive-effort-and-confidence-effects-from-a-survey-of-knowledge-workers/) (!) cite dans sa conclusion : « _While GenAI can improve worker efficiency, it can inhibit critical engagement with work and can potentially lead to long-term overreliance on the tool and diminished skill for independent problem-solving_ ». Ce constat résonne énormément avec le constat partagé avec des professeurs de l'Université de Lille : certains étudiants sont parfois en difficulté pour réaliser les tâches qu'on leur demande (leur futur travail hein) sans l'aide de leur IA préférée.

Autre phénomène : quand on demande à ChatGPT de générer du code _front_, il génère en premier du code React. Cela peut sembler logique, étant donné que React semble être la librairie la plus utilisée pour développer des _frontends_. Beaucoup de développeurs choisissent alors d'utiliser cette librairie, ainsi que les dépendances que ChatGPT leur suggère. Cet exemple (je n'ai rien contre React), illustre un cas d'enfermement technologique et pourrait même s'avérer dangereux, si jamais une IA, sous le contrôle d'une organisation quelconque, s'amusait à suggérer une dépendance contenant une faille de sécurité par exemple.

Est-ce que ChatGPT sait générer du code compatible avec la toute dernière version de Spring Security, ou génère-t-il du code pour Spring Security 5, qui est le plus présent sur GitHub et donc forcément plus ancré dans son modèle ?
Tout ça pour dire que les personnes ne s'appuyant que sur leur LLM pour écrire le code risquent d'être incapables de l'adapter ou de le maintenir si elles ne savent pas ce qu'elles font. Les approches de type _RAG_ (_Retrieval-Augmented Generation_), permettent de pallier en partie ces biais, mais ne sont pas toujours utilisées par défaut.

J'en viens aussi à penser que les phrases vues parfois qui, en substance, se résument par « On va remplacer les développeurs par des IA  » est le _take_ le plus pété que j'ai pu entendre ces dernières années. Réduire mon métier à la simple écriture du code est ignorant, voire malhonnête. L'IA peut s'avérer utile comme outil pour aider les développeurs, pour préparer des revues de code, proposer du _refactoring_ (mais là, un simple LLM ne me semble pas le plus pertinent), expliquer du code, générer des tests, mais ne remplacera pas les développeurs (tout comme le No-Code ne les a pas remplacés).

## Conclusion

Vous l'aurez compris en lisant ces lignes, je ne suis pas un grand utilisateur ou admirateur des IA actuelles.

Je pense que l'IA est un outil formidable, mais que les usages que j'observe massivement aujourd'hui, à savoir générer des images (moches, mais c'est mon avis), des _deepfakes_ et du code bancal (encore une fois mon avis), ne sont pas à la hauteur de ce qu'on pourrait espérer d'un tel outil.

Je n'aime pas voir apparaître des IA dans tous les outils que j'utilise, sans en avoir besoin.
Leur côté non déterministe (donc non testable), les « hallucinations » des LLM, et leur dégénérescence qui, selon moi, ne peut que se produire, sont pour moi clairement les problèmes de la génération actuelle de ces outils.

Je pense que vous méritez (nous méritons) mieux que ça.

J'écris moi-même mes articles et mon code.

## Inspiration et ressources

Cette page est directement inspirée par ce que j'ai pu voir chez d'autres bloggers, en particulier Denis Germain sur https://blog.zwindler.fr/ai-manifesto/ et Cassidy Williams sur https://cassidoo.co/ai/, qui ont eux-même été inspirés par Damola Morenikeji sur https://www.bydamo.la/p/ai-manifesto.

J'éditerai probablement cette page de temps en temps, avec de nouveaux éléments et de nouveaux liens issus de ma veille. Après tout, il n'est pas impossible que je change d'avis.

* [ChatGPT : le mythe de la productivité](https://framablog.org/2025/03/09/chatgpt-le-mythe-de-la-productivite)
* [AI Copilot Code Quality: 2025 Look Back at 12 Months of Data](https://www.gitclear.com/ai_assistant_code_quality_2025_research)
* [The Impact of Generative AI on Critical Thinking: Self-Reported Reductions in Cognitive Effort and Confidence Effects From a Survey of Knowledge Workers](https://www.microsoft.com/en-us/research/publication/the-impact-of-generative-ai-on-critical-thinking-self-reported-reductions-in-cognitive-effort-and-confidence-effects-from-a-survey-of-knowledge-workers/)
* [StackOverflow Driven-Development (SODD)](https://dzone.com/articles/stack-overflow-driven-development-sodd-its-really)