---
date: 2025-06-19
language: fr
title: Leeloo Dallas Multipass - Répondre aux 5 élements d'un CFP
tags:
  - internet
draft: true
---

Les conférences sont un lieu important de partage et de veille, y participer en tant que speaker ou assister à des conférences permet d'enrichir votre réseau et de partager vos expériences.

Mais comment devient-on speaker, comment se passe le processus de sélection ?
On m'a récemment demandé : Julien, comment ça marche la réponse à un CFP (_Call For Papers_) ?

C'est une discussion que j'ai déjà eu plusieurs fois avec différentes personnes, j'ai donc décidé de regrouper tous les éléments dans cet article, et de vous donner ma vision et mes astuces issues de mes deux expériences, à la fois en tant speaker à différentes conférences en France, et en tant qu'organisateur de la conférence Cloud Nord.
Je vais vous donner mes attentes en tant qu'orga et mes astuces en tant que speaker dans le but de vous aider à bien rédiger votre réponse à un CFP et (espérons) être sélectionné.

> Le terme conférence peut parfois désigner l'événement auquel vous souhaitez participer (DevLille, Devoxx ou Cloud Nord par exemple), ou l'intervention d'une personne sur scène pour présenter son sujet. Par souci de claretté, j'utilise le terme _conférence_ quand je parle de l'événement, et le terme _talk_ lorsque je parle de l'intervention.

## Les 5 éléments importants lors d'une réponse à un CFP

Lors de la réponse à un CFP on vous demande cinq éléments, qui sont à destination des relecteurs de votre CFP ou des participants à la conférence.

Ces cinq éléments sont :

* le titre de votre _talk_
* l'abstract de votre _talk_
* votre biographie
* vos références
* le track et le format de votre _talk_

Nous allons voir dans la suite de cet article comment bien répondre à ces 5 élements !

## Le titre de votre talk

> Il faut que ce soit green ok ?
> 
> _- Ruby Rhod_

Le titre de votre talk doit donner une idée précise de votre contenu tout en étant accrocheur.
C'est le premier élément que remarquent les relecteurs du CFP ou les participants à une conférence. Ce titre sera publié sur l'agenda de la conférence si votre talk est sélectionné !

Le titre peut éventuellement contenir une note d'humour, ou une référence culturelle (comme le titre de cet article). Il doit néanmoins être suffisamment clair au premier coup d'œil.

Voici les titres que j'ai utilisés sur mes derniers talks ou BBLs :

* Spring Boot & Containers - Do's & Don'ts (BBL)
* Laissez tomber vos Dockerfile, adoptez un buildpack ! (Sunny Tech - Cloud Nord)
* Mais au fait, ça marche comment les service-accounts ? (Devoxx France)
* Rebase d'image Docker/OCI avec crane (Devoxx France - DevLille)

Chacun de ces titres a un style différent. Ils peuvent être percutants, amener au changement, ou à la conversation. Ils mentionnent chacun le cœur du sujet, et évoquent l'approche qu'aura le talk : un ton technique qui amènera des bonnes pratiques, un ton pédagogique qui va aider à comprendre une mécanique peu connue, ou une présentation d'un REX ou d'un cas d'usage précis.

Pour le choix de votre titre, réfléchissez au ton que vous souhaitez donner à votre talk, ou à l'impact que vous espérez provoquer.
C'est un équilibre difficile à trouver, qui peut vraiment faire la différence.

Voici quelques exemples de titres d'autres talks (pas les miens cette fois-ci), sélectionnées au DevLille 2025, qui jouent sur cet équilibre :

* Devenez un Astronaute : Découvrez le Framework Astro
* BullShit IT Awards : Les absurdités tech qui sabotent vos équipes
* La programmation fonctionnelle en Java sans grimacer
* Je malmène ton LLM en direct avec 10 failles de sécurité
* Au corps à CORS - L'art subtil des requêtes Cross Origin
* direnv : petit mais costaud

Par des jeux de mots ou métaphores, certains de ces titres amènent à la curiosité. Dans chacun de ces exemples, on sait à quoi s'attendre, à la fois sur le sujet et sur le ton du talk.
C'est ce qu'on cherche le plus souvent pour une bonne réponse à un CFP.

## L'abstract de votre talk

> Korben Dallas ! L’heureux vainqueur du grand concours des Gemini Croquette. Ce garçon carbure d’enfer. Il a tous les atouts ! Et en plus, il a des trucs à dire aux quelques 50 milliards de paires d’oreilles qui nous écoutent ! Balance tout mec !
> 
> Euh... Salut.
> 
> _- Ruby Rhod_ & _Korben Dallas_


Avec le titre, l'abstract est un des éléments qui sera publié sur l'agenda de la conférence.
Les titres de talks peuvent utiliser l'anglais sans trop de problème, mais en général, si vous allez parler en français, privilégiez d'écrire votre abstract en français également.

L'abstract de votre talk doit transmettre plusieurs informations (pas forcément dans cet ordre) :

* le cœur de votre sujet, et le contexte dans lequel il s'applique
* le ton du talk
* sa structure principale
* le public cible
* les _Take-Aways_

Pour ces deux premiers éléments, on s'attend à être aligné avec le titre de la proposition. Le sujet doit être présenté, son contexte, et le ton du talk doit transparaître dans les termes utilisés. 

On peut aussi y mentionner la structure du talk, sous la forme d'un plan macro par exemple : "dans un premier temps nous verrons blablabla, puis ensuite nous verrons blablabla". Cette structure permet au lecteur de se projeter dans le talk, de savoir à quoi s'attendre. C'est à la fois très intéressant pour les relecteurs du CFP, mais aussi pour les participants à la conférence qui pourront lire l'abstract sur l'agenda de la conférence.

Il faut également bien préciser le public cible du talk. Est-ce que ce talk est destiné aux Devs ? Aux Ops ? Quel niveau d'expérience ou quel pré-requis ?
Être clair sur le public visé permet d'éviter les erreurs de casting de votre public. Vous souhaitez probablement que les gens qui viennent vous écouter sont intéressés par votre sujet et seront en capacité de le comprendre. Un talk sur des fonctionnalités d'un langage cible probablement les devs qui maîtrisent déjà le langage. Ça peut parfois être évident, mais ça peut aussi être intéressant de le repréciser.

Un élément souvent oublié, mais qui a une importance majeure à mon sens est ce qu'on appellera les _Take Aways_. Je parle aussi parfois de la "Promesse" d'un talk.
Ces éléments répondent à la question : Avec quoi je repars après avoir assisté au talk ? Est-ce que j'aurai simplement passé un bon moment, si le talk est plus humoristique ou théâtral, ou bien est-ce que je repartirai avec des astuces concrètes que je pourrais utiliser dès la semaine prochaine ? Ou encore, est-ce que je vais pouvoir récupérer des éléments de code ou le contenu d'une démo pour pouvoir tester le sujet à la maison ?
Les _Take Aways_ peuvent être simplement évoqués, ou mentionnés de manière explicite, mais c'est une promesse l'on fait aux personnes qui viendront assister au talk : avec quoi on repart en sortant du talk.

Un bon abstract (à mon sens) fait en général quelques paragraphes, 2 ou 3 c'est largement suffisant.
Un abstract d'une ligne "Je présente les nouveautés du logiciel X" est souvent très insuffisant, surtout pour un talk de 45 minutes !

Voici un exemple d'un de mes abstracts qui a été accepté à plusieurs conférences :

> Depuis plusieurs années maintenant, _Docker_ est utilisé par toute l'industrie de l'IT pour packager et déployer des applications.
> 
> Bien que l'écriture d'un `Dockerfile` soit facile, la construction d'images _OCI_/_Docker_ reste un exercice compliqué :
> * optimisation des layers de l'image
> * bonne gestion des processus Linux
> * séparation des phases de build et de run des images
> * bonnes pratiques de sécurité
> 
> Pire, lorsqu'une faille de sécurité est détecté dans une layer basse (distribution ou runtime) d'une image applicative, il faut alors potentiellement reconstruire plusieurs dizaines ou centaines d'images pour y intégrer les versions patchées.
> 
> Dans ce talk, nous apprendrons comment les **buildpacks** permettent de construire des images OCI/Docker sans _Dockerfile_ et bénéficier des bonnes pratiques issues de la communauté open-source.
> 
> Nous verrons :
> * ce qu'est une image _OCI_, une layer, et comment _Docker_ les construit
> * comment analyser le contenu des layers d'une image _OCI_, et ce qui ne va pas dans les images que nous construisons au quotidien
> * ce qu'est un **buildpack** et comment un **buildpack** construit une image OCI
> * avec une démo, comment utiliser un **buildpack** proposé par la communauté open-source pour construire une image _OCI_ contenant une application _Java_ optimisée
> * enfin, nous verrons comment les **buildpacks** proposent de _rebaser_ des image, et nous permettre de patcher en masse des images applicatives pour corriger des failles de sécurité, sans reconstruire complètement nos images !
> 
> Ce talk est donc à destination des _Ops_ et des _Devs_ qui manipulent _Docker_ au quotidien.
> 
> À la sortie de ce talk, je devrais vous avoir convaincu d'abandonner vos Dockerfile et d'expérimenter les buildpacks !

Cet abstract (un peu long), comprend tous les éléments de réponse dont j'ai parlé plus haut : une remise en contexte, l'exposition d'un enjeu ou d'un problème, la solution et les take-aways (démo et bonnes pratiques). Le plan est aussi décrit, et le public cible. Le ton est pédagogue, avec pour but d'inviter l'auditeur à tester la technologie en sortie.

Voici un autre abstract, tiré de la conférence d'Audrey Wech : "Développer une application mobile avec React Native : pourquoi Expo change la donne ?" :

> Les technologies liées au développement mobile cross-platform ont considérablement évolué. Avec elles, des outils comme Expo ont émergé, simplifiant le cycle de vie des applications mobiles, de l’initialisation à la mise en production.
>
> Chez Karnott, nous avons récemment entrepris la migration de notre application React Native vers Expo, dans le but d’améliorer notre expérience développeur et d’assurer la pérennisation de notre application.
>
> Lors de cette conférence, je vous partagerai mon retour d’expérience sur cette transition. J’explorerai les raisons qui ont motivé notre choix, les défis rencontrés et les bénéfices concrets que nous avons observés. Aucune connaissance préalable de React Native n’est requise.

Encore une fois, on retrouve dans cet abstract tous les éléments : le contexte, l'enjeu autour de l'expérience développeur, la structure du talk, les pré-requis, et la promesse : avoir le retour d'expérience d'une entreprise qui a effectué une migration.
C'est à mon sens un bon abstract.

Je donne enfin un exemple d'abstract qui est clairement insuffisant de mon point de vue, pour un talk nommé "REX IAC Service PaaS Azure" :

> Le talk aura pour objectif d‘expliquer en quoi cette approche permet d’unifier la gestion des services PAAS, les défis à relever mais surtout l’intérêt général de cette approche IAC.
> 
> Technos : Azure, Terraform, Ansible, git

Bien que le fond puisse être intéressant, et que la liste des technos permet de cibler le public, la lecture de cet abstract ne permet pas de savoir à quoi s'attendre : pas de contexte ni d'enjeu, pas de promesse. Une structure a l'air évoquée, mais n'est pas très claire.

## La biographie

> Et toi, comment tu t’appelles ? Moi Korben, Toi ?
> 
> Leeloo Mina Lekatariba Laminatcha Ekbat D Sebat.
> 
> Bien, d’accord. Tout ça c’est ton nom.
> 
> _- Korben Dallas_ et _Leeloo_

Comme le titre et l'abstract du talk, la bio de la speakeuse ou du speaker sera publiée sur l'agenda de la conférence.
Certaines conférences relisent les CFP de manière "anonyme", sans connaître l'identité des personnes, d'autres conférences utilisent la bio dans le process de relecture. Pour les CFP pratiquant l'anonymat lors de la relecture, les bios sont relues après la délibération.

> En général, à la relecture du CFP de Cloud Nord, on ne note pas les talks des personnes qu'on connaît personnellement. À titre personnel, je joue une première relecture anonyme, puis je "supprime" la note que j'aurais attribué à une connaissance proche.

Pas besoin de rédiger votre auto-biographie ou vos mémoires, on attend ici simplement une photo, une courte bio et les liens vers vos réseaux sociaux si vous souhaitez les partager.

Concernant la photo, un avatar est aussi souvent accepté. Évitez simplement de mettre une photo qui représente quelqu'un autre (pas de photo de Philippe Etchebest ou de Lionel Messi svp 😅).

> Aknot, c'est vous ? Quelle horrible tête, ça ne vous va pas. Enlevez la !
>
> _- Zorg_

La bio peut présenter vos passions, votre métier, l'entreprise dans laquelle vous travaillez. J'aime bien présenter les choses dans ce sens, puisque ça met plus en avant ma personne, plutôt que mes activités pro. Si vous avez déjà lu ma bio, vous savez quel est mon animé préféré (en même temps, vu le nom de mon site, vous aviez peut-être déjà un indice aussi 🐉).

Le nom de votre entreprise peut aussi parfois vous être demandé à part, ainsi que vos liens vers les différents réseaux.
Il n'est pas rare que des spectateurs de conférence décide de vous suivre sur un réseau ou un autre et vous notifie dans un post pour vous remercier de votre talk ou vous féliciter pour votre prestation. C'est toujours sympa de recevoir ce type de feedback 😁

## Vos références

> Où...Où...Où est-ce qu'il a appris à négocier ?!
> 
> _-Matelot anonyme_

La section "Références" n'est visible que des relecteurs du CFP, elle n'est jamais publiée sur les sites des conférences. C'est donc un espace sur lequel vous pouvez directement parler aux relecteurs !
Vous pouvez indiquer pourquoi vous souhaitez parler de ce sujet, votre expérience sur le sujet ou la techno, et votre expérience en prise de parole.

Deux cas se présentent alors, soit vous postulez pour votre premier talk (dans ce cas bravo !), soit vous avez déjà une première expérience.

Si c'est votre premier talk, n'hésitez pas à le préciser. Beaucoup de conférences aiment mettre en avant de nouvelles personnes, et laisser leur chance à des nouvelles speakeuses ou des nouveaux speakers.
Si c'est le cas, n'hésitez pas également à préciser si vous avez déjà donné des présentations ou des formations dans le cadre de votre travail, cela permet aux relecteurs d'évaluer votre capacité à prendre la parole en public. Si le talk a déjà été donné en interne, c'est aussi l'occasion de partager un lien vers les slides, ou détailler plus finement le contenu, plan ou démos.

Si vous avez déjà une première expérience de talk en conférence, c'est le bon endroit pour en parler ! Si le talk que vous proposez a déjà été donné dans une autre conférence, meetup ou tremplin, précisez-le, donnez les liens vers les slides ou la vidéo si elle existe.
Si c'est un nouveau talk, indiquez le et expliquez aux relecteurs pourquoi ils devraient vous choisir vous !

L'idée est vraiment de donner tous les éléments qui permettront aux relecteurs de faire leur choix. Il n'est aussi pas rare que les relecteurs ou les orgas contactent la personne pour obtenir plus de détail. Mais autant être le plus clair dès le départ.

## le track et le format du talk

> Vous êtes classé dans la catégorie humain ?
> 
> Négatif, je suis une mitte en pull over.
> 
> _- Policier et Korben Dallas_

Lors de l'envoi de votre proposition, on vous demandera la plupart du temps de choisir un ou plusieurs track et un ou plusieurs formats.

Les tracks correspondent aux thèmes ou aux catégories des conférences. En fonction de leur ligne éditoriale, les conférences vont privilégier un ou plusieurs tracks.

Un exemple concret pour illustrer : Devoxx France est une conférence principalement orientée autour de Java. Il y a donc un track Java, mais aussi un track Architecture, Data & IA, Déploiement, Front-end, etc... 
Pour Cloud Nord, nous avons des tracks Cloud, DevOps, Sécurité, mais aussi Data & IA, Architecture et Containers.
Les conférences ont aussi presque toutes un track "Découverte", "Alien" ou "Hors piste" qui permet de proposer des talks en dehors des thèmes principaux des conférences.

Le bon choix du track permet de s'assurer que votre talk matche bien avec la ligne éditoriale de la conf, et permet aussi de cibler le bon public.

Les tracks principaux des conférences auront aussi plus de talks alloués. Il y a aussi cet élément à prendre en compte lors de l'envoi de la proposition.
A Cloud Nord, nous avons moins de 20 talks (probablement 16 ou 18 cette année), donc nous accepterons probablement 1 ou 2 talks sur le track "Découverte", pour avoir le plus de contenu possible sur le thème qui nous intéresse le plus : "Cloud". La sélection sera alors rude dans ce track.

Concernant le format, on retrouve en général trois grands types de formats dans les conférences.

Le format le plus classique est le format "conférence", de 45 ou 50 minutes. Il permet d'aller en détail dans un sujet. C'est souvent le format le plus souhaité.
Le deuxième format est un format plus court, souvent appelé "Short" ou "Quickie". Ces formats sont en général d'une durée de 15 ou 20 minutes. Ce sont les formats les plus faciles pour débuter. En effet, les organisateurs prennent moins de "risque" à allouer ce format à une personne qui débute.
Le dernier format est un format qui est souvent appelé "Labs" ou "Workshop". Ce sont des formats de deux ou trois heures, plutôt consacrés à des mises en pratiques d'une techno, sous la forme d'un tutorial ou d'un TP. Ce sont des formats aussi faciles pour débuter, car ils sont souvent fait à plusieurs, et il est moins intimidant d'animer un workshop avec 10 personnes, plutôt que de présenter un talk devant un amphithéâtre plein.

## Autres éléments

> C'est green ?
> 
> Super green !
> 
> _- Korben Dallas & Ruby Rhod_

Beaucoup de conférences acceptent des talks avec plusieurs speakeuses ou speakers.
Un talk en binôme peut aussi être une très bonne façon de se lancer. Un côté théâtral peut alors être présent entre les deux personnes, ou le talk peut être découpé en deux, entre une partie présentation et une partie démo, chacun s'occupant d'une partie.

Envoyez vous-même vos propositions, il est souvent mal vu de voir une proposition envoyée par quelqu'un avec la mention 'Le talk sera présenté par X'. Cela donne l'impression qu'on impose le talk à une personne.

Les conférences acceptent en général un nombre limité de propositions par personne ou binôme. Ce nombre est généralement 2 ou 3. Maximisez vos chances en proposant plusieurs sujets, mais c'est également au risque que tous vos sujets soient acceptés ! Si c'est le cas, n'hésitez pas à en parler avec les relecteurs et organisateurs, car le but n'est pas de vous mettre en difficulté.

N'hésitez pas également à re-proposer un sujet qui a été refusé à une autre conférence, ou une année précédente. Il n'est pas impossible que le sujet soit devenu plus important avec les années.

Enfin, ça peut sembler évident, mais évitez les abstracts générés par IA.
Le but est bien de transmettre VOS expériences, pas celles de ChatGPT (dans tous les cas, ce n'est pas ChatGPT qui sera sur scène le jour J).

> Si tu veux que quelque chose soit fait, fais le toi même !
> 
> _- Zorg_

Si vous m'avez lu jusqu'au bout, bravo et merci !
J'espère que tous ces conseils vous aideront.

Le CFP de Cloud Nord est déjà ouvert, donc n'hésitez pas à soumettre vos sujets sur https://conference-hall.io/cloud-nord-2025.

## Liens et références

* [Préparez et donnez votre première conférence (quand ce n’est pas votre métier)](https://votre-premiere-conference.fr/) : Ce livre de Pascal Martin donne beaucoup de très bon conseils. L'envoi de la proposition à un CFP fait l'objet d'un chapitre complet.
* [Proposer une conférence – la réponse au CFP](https://devfesttoulouse.fr/2025/03/18/proposer-une-conference-le-cfp/) : Un chouette billet de blog écrit par Sylvain Wallez, on y retrouve des conseils similaires
