---
date: "2025-05-02"
language: fr
title: La veille de Wittouck - Fin Avril 2025
series: La veille de Wittouck
tags:
  - Internet
  - Java
  - Security
---

En ce fin avril, il fait beau, donc on n'a pas très envie de passer du temps sur son ordi à scroller pour faire sa veille.
Donc on lit la veille de Wittouck !

> Sur quelques liens, j'ai oublié de noter les sources qui m'ont permis de les découvrir (oups).

La playlist de Devoxx France 2025 se remplit avec des vidéos masquées, donc la publication devrait bientôt être faite 🤞

## 🛜 Internet

* [Professeur Fabrizio Bucella](https://www.youtube.com/@FabrizioBucella/shorts) sur Youtube.

> J'ai découvert les _shorts_ de ce professeur belge par hasard. Il nous explique des principes de physique en quelques minutes, principalement avec un tableau noir et une craie. C'est incroyable de pédagogie, plein d'humour et de bel accent belge 🇧🇪. Salukes.

## 🔒 Sécurité

L'actualité principale de ce mois d'avril a été l'annonce de la fin du financement des programmes CVE et CWE par le gouvernement fédéral Américain, _via_ l'entité MITRE. Le gouvernement est depuis revenu en arrière et a prolongé le financement de 11 mois, mais cette annonce a inquiété toute la sphère sécu. Et l'ombre de la fin du financement plane toujours.
À noter que le financement de _NVD_ (la _National Vulnerability Database_) qui est aussi utilisée par des outils comme _OWASP Dependency Check_ ne serait pas directement impacté par un arrêt du financement de MITRE (car étant financés par une autre entité fédérale). Ces outils continueraient à fonctionner (ouf).

Plusieurs liens autour de ce sujet :

* [La lettre envoyée aux membres du _board_ CVE](https://bsky.app/profile/tib3rius.bsky.social/post/3lmulrbygoe2g) _via_ [Tib3rius](https://bsky.app/profile/tib3rius.bsky.social) sur Bluesky

> Ce document semble être la première information sur l'arrêt du financement

Suite à cette annonce, et pour assurer sa pérennité et son indépendance aux financements du gouvernement américain, la _CVE Foundation_ a été officialisée :

* [CVE Foundation](https://www.thecvefoundation.org/)

> Ce document décrit principalement les enjeux du programme et les buts principaux de la fondation : une gouvernance internationale et la transparence sur la gestion des financements.

Pour comprendre les événements : 

* [CVE Program Funding Reinstated—What It Means And What To Do Next](https://www.forbes.com/sites/kateoflahertyuk/2025/04/16/cve-program-funding-cut-what-it-means-and-what-to-do-next/) _via_ [Gergely Orosz](https://bsky.app/profile/gergely.pragmaticengineer.com) sur [BlueSky](https://bsky.app/profile/gergely.pragmaticengineer.com/post/3lmwpojt55c27)

> Cet excellent article résume ce qu'il s'est passé sur ces quelques jours, depuis l'annonce de la fin du financement, et quelle pourrait être la suite.

Enfin, il serait peut-être intéressant de creuser l'alternative européenne : _EUVD_ (pour _European Union Vulnerability Database_), qui reprend aussi les données de CVE, mais permettrait une indépendance en cas de coupure de CVE, qui semble maintenant partiellement protégé par la fondation, mais reste encore sous gouvernance Américaine.

* [EUVD](https://euvd.enisa.europa.eu/homepage) : La page d'accueil de la base de données _EUVD_, portée par l'_ENISA_ (_European Union Agency For CyberSecurity_)

> Je pense que cette base de données pourrait devenir une alternative fiable à la NVD américaine. À voir si des outils comme OWASP DC ou OWASP DT sont forkés / paramétrables pour se brancher sur ces données. Je n'ai pas trouver comment télécharger le JSON avec les EUVD 😅

## ☕ Java

* [What's new in Maven 4?](https://maven.apache.org/whatsnewinmaven4.html)

> Maven 4 n'est pas encore sorti, mais dévoile une partie de ses fonctionnalités : Java 17 par défaut (en replacement de Java 5 ?), Des modification sur le format du pom avec la séparation du pom de _build_ du pom de consommation. Et aussi l'officialisation du type `bom` comme type de package. `mvnd` reste un projet à part. Un guide de migration est aussi déjà disponible.

## ☕ IA

* [AI 2027](https://ai-2027.com)

> Un article qui imagine l'évolution de l'IA dans les années à venir. L'article prend beaucoup en compte le contexte géopolitique actuel comme inspiration. On frise la science-fiction  (surtout sur la fin, qu'on peut choisir entre 2 scenarios). À lire absolument.

---

La prochaine publication est prévue autour du 16 mai 🗓️

[//]: # (Photo de couverture par [Oğuzhan Akdoğan]&#40;https://unsplash.com/@jeffgry?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash&#41; sur [Unsplash]&#40;https://unsplash.com/photos/man-using-computer-inside-room-qYMkkREOHa4?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash&#41;)