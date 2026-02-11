---
date: 2026-02-06
title: Optimiser un site Hugo
draft: true
---

[//]: # (TODO link vers le blog d'antoine)
Sur les bons conseils du pote Antoine Caron, j'ai pris temps cette semaine d'optimiser un peu mon site.

Ce site que vous êtes en train de lire est un site statique, buildé avec Hugo.

J'ai déjà un peu travaillé la compression des différentes ressources, principalement les illustrations, mais je m'étais arrêté à ça.
Dans cet article, je détaille comment j'ai optimisé le build de ce site, pour minimiser les temps de chargement, et comment j'ai amélioré sa sécurité en suivant les bonnes pratiques poussées par MDN.

## Le score Lighthouse

Pour faire un premier travail sur les performances de ce site, j'ai utilisé [une analyse LightHouse](https://pagespeed.web.dev/analysis/https-codeka-io/we5dukzmku?form_factor=desktop).

Lighthouse permet en quelques minutes d'avoir une vue des performances d'une application ou d'un site web, à la fois pour une cible _Desktop_ et _Mobile_.
Il permet aussi de valider certaines propriétés d'accessibilité, comme des contrastes, la présence de texte alternatif pour les lecteurs d'écran, etc.

C'est, je pense, un bon point de départ.

Voici les scores de mon site à l'heure actuelle :


![Score Lighthouse pour un mobile](lighthouse-mobile.png)
![Score Lighthouse pour un desktop](lighthouse-desktop.png)


> J'ai clairement une marge d'amélioration sur l'accessibilité et les performances.

## Minification

Une première étape consiste à minifier les ressources statiques, HTML, CSS et JS.

Cette étape est très simple à mettre en place, car elle est déjà supportée par Hugo.
Il suffit lors du build d'ajouter le flag `--minify` pour demander à Hugo de minifier toutes les ressources.

Ma commande de build est la suivante dans mon `mise.toml` :

```toml
[tasks.build]
description = "Build le site avec Hugo"
run = "hugo --gc --minify --destination public"
```

Ce qui produit des fichiers HTML minifiés de ce type :

```html
<!doctype html><html xmlns=http://www.w3.org/1999/xhtml xml:lang=fr-FR lang=fr-FR><head><script defer language=javascript type=text/javascript src=/js/bundle.min.39a1898ad60dcb3b845d8dc359b7c996c10aa0da902f0d461da32348b1bc5f02.js></script><script defer data-domain=codeka.io src=https://plausible.io/js/script.js></script><script type=text/javascript src=https://app.affilizz.com/affilizz.js async></script><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=/favicon.png><meta property="og:image" content="/pp_ekite_itvw.png"><meta name=twitter:image content="/pp_ekite_itvw.png"><meta name=twitter:card content="summary_large_image"><meta property="og:image:width" content="639"><meta property="og:image:height" content="708"><meta property="og:image:type" content="image/png"><title itemprop=name>Julien Wittouck</title><meta property="og:title" content="Julien Wittouck"><meta name=twitter:title content="Julien Wittouck"><meta itemprop=name content="Julien Wittouck"><meta name=application-name content="Julien Wittouck"><meta property="og:site_name" content="Julien Wittouck">
```

Hop, on peut passer rapidement à autre chose 🚶

## Conversion des images en webp et redimensionnement

Une des actions que j'ai mis en place il y a un moment, est l'utilisation du format _webp_ pour compresser les illustrations que j'utilise dans mes articles.

J'utilise souvent des photos que j'ai capturées avec mon smartphone (pour les articles de conférence), des captures d'écran ou des schémas (produit sur draw.io le plus souvent), ou des photos _stock_ que je vais chercher pour illustrer mes articles de veille.

Ces photos sont souvent lourdes (plusieurs mégaoctets) et en haute résolution, et la première action simple consiste à redimensionner ces photo et les recompresser au format _webp_.

Hugo supporte la recompression des images dans différents formats à la volée, mais pas leur redimensionnement automatique, il faut implémenter soi-même la mécanique.
Pour pouvoir redimensionner les images à la volée, la meilleure solution semble d'utiliser un hook "img" Hugo, qui permet de surcharger la traduction du markdown et d'y mettre le code qu'on souhaite.

Le hook utilisé par défaut est le suivant :

```go
<img src="{{ .Destination | safeURL }}"
  {{- with .PlainText }} alt="{{ . }}"{{ end -}}
  {{- with .Title }} title="{{ . }}"{{ end -}}
>
{{- /* chomp trailing newline */ -}}
```

Pour redimensionner les images à une taille maximale de 820px (la taille utilisée sur la colonne de contenu de ce site), j'utilise le hook suivant :

```go
{{- $image := .Page.Resources.GetMatch .Destination -}}
{{- $width := math.Min 820 $image.Width -}}
{{- $resizeOpts := printf "%dx webp lossless q100 lanczos" (int $width) -}}
{{- with $image.Resize $resizeOpts -}}
<img src="{{ .RelPermalink }}" width="{{ .Width }}" height="{{ .Height }}"
    {{- with $.PlainText }} alt="{{ . }}"{{ end -}}
    {{ with $.Title }}title="{{ . }}"{{ end }}>
{{- end -}}
{{- /* chomp trailing newline */ -}}
```

Je force l'utilisation de `lossless` avec la qualité maximale `q100` pour éviter une perte de données qui rendrait les illustations peu lisible, ce qui serait surtout problématique pour les schémas.

## Pré-compression des ressources statiques

Les images étant maintenant compressées au build par Hugo, je peux m'atteler à la compression des ressources déjà minifiées (HTML, CSS et JS donc).

Avant de passer à la pré-compression en elle-même, il faut regarder comment les ressources seront servies.

Mon site est hébergé chez Clever Cloud, dans une instance de type _static_.
J'avais écrit un article à ce sujet l'année dernière : [Déployer des applications statiques sur Clever Cloud](/2025/06//2025-06-05-static-apps-clever).

Clever Cloud permet d'utiliser Caddy pour servir les fichiers statiques en surcharge de `static-web-server`, simplement en ajoutant un `Caddyfile` à la racine du projet.

Cette option va me permettre de pouvoir configurer Caddy pour servir le répertoire `public` du site :

```Caddyfile
# Clever Cloud needs us to listen on port 8080
:8080

file_server {
	# Clever Cloud serves the public directory in a cc_static_autobuilt directory
    root public
}

# Ask Caddy to compress static files 
encode
```

Lors de l'exécution d'une requête, Caddy va servir les fichiers statiques, et potentiellement compresser les réponses HTTP en alimentant le headers `Content-Encoding`. Les formats utilisés par défaut par Caddy sont `zstd` et `gzip`, et seules les ressources pertinentes sont compressées (les formats déjà compressés comme `jpg` ne sont pas re-compressés).

Cette compression permet d'économiser de la bande passante et accélère le temps de chargement des pages.

Cependant, la compression se fait en utilisant un peu de CPU à la volée.
Il est alors intéressant de pré-compresser les ressources statiques à la phase de build pour économiser un peu de CPU.

Une directive Caddy permet de servir des fichiers statiques pré-compressés : `precompressed`.
Caddy va alors rechercher des variantes compressées des fichiers, sous la forme de fichiers sidecar.
À côté de chaque fichier statique, il faut donc générer les variantes compressées et les nommer en utilisant les extensions `.gz`, `.br` et `.zst` par exemple.

Hugo ne permet pas de générer ces variantes compressées de lui-même, donc je dois utiliser un petit script qui s'exécutera en fin de la phase de build.

J'ai donc créé ce script dans mon fichier `mise.toml` :

```toml
[tasks.build]
description = "Build le site avec Hugo"
run = "hugo --gc --minify --destination public"

[tasks.post-build]
description = "Post build hooks"
depends_post = ["precompress"]

[tasks.precompress]
description = "Precompress static resources"
run = '''
COMPRESSREGEX=".*(html|css|js|xml|ico|svg|md|pdf|woff2)$"
find public/ -type f -regextype egrep -regex $COMPRESSREGEX | xargs zstd --keep --force -19
find public/ -type f -regextype egrep -regex $COMPRESSREGEX | xargs gzip --keep  --force --best
'''
```

J'ai implémenté la compression avec `gzip` en utilisant le plus haut niveau de compression possible (`--best`), et avec `zstd` avec la plus forte compression également (`-19`). Le niveau de compression a surtout un impact à la compression, mais peu à la décompression, donc autant maximiser les différents niveaux.
J'ai fait l'impasse sur le format `br` parce qu'il nécessite d'installer un binaire supplémentaire sur mes instances Clever Cloud, et que `gz` et `zst` sont déjà bien suffisants : `zst` sera supporté par les navigateurs modernes dans les versions les plus récentes, `gzip` fera office de format par défaut raisonnable.

Par défaut, Clever Cloud exécute une tâche `mise run build` si elle existe, donc l'ajouter dans mon fichier permet de pouvoir préciser mes options de build.

Pour la phase de compression, il suffit d'indiquer à Clever Cloud d'exécuter `mise run post-build`, cela se fait avec un hook sur 

Le script `precompress` est inspiré d'un [article de blog de Scott Laird](https://scottstuff.net/posts/2025/03/09/precompressing-content-with-hugo-and-caddy/) sur lequel je suis tombé en faisant quelques recherches.
Il recherche l'ensemble des fichiers matchant la regex donnée, et utilise `zstd` pour compresser ces fichiers.

L'exécution de ces scripts produit la sortie suivante : 

```bash
[build] $ hugo build hugo --gc --minify --destination public
Start building sites …
hugo v0.155.2-d8c0dfccf72ab43db2b2bca1483a61c8660021d9+extended linux/amd64 BuildDate=2026-02-02T10:04:51Z VendorInfo=gohugoio

                  │ EN │ FR
──────────────────┼────┼─────
 Pages            │ 75 │ 139
 Paginator pages  │  0 │   4
 Non-page files   │ 14 │ 222
 Static files     │ 36 │  36
 Processed images │  3 │ 275
 Aliases          │  1 │   8
 Cleaned          │  0 │   0

Total in 272 ms
[precompress] $ COMPRESSREGEX=".*(html|css|js|xml|ico|svg|md|pdf|woff2)$"
245 files compressed : 80.99% (  83.3 MiB =>   67.4 MiB)                       B ==> 98%^T
Finished in 7.77s
```

On peut valider que les fichiers buildés sont précompressés comme souhaité, avec les extensions `.gz` et `.zst` :

```bash
$ ls public/
2020         404.html.zst  fonts           index.xml.zst                           projects
2021         ai            fr              js                                      robots.txt
2022         ai-manifesto  icons           logo_blue.png                           series
2023         books         images          logo_transparent_background.png         sitemap.xml
2024         credentials   index.html      now                                     sitemap.xml.gz
2025         css           index.html.gz   page                                    sitemap.xml.zst
2026         ekite         index.html.zst  posts                                   stats
404.html     en            index.xml       pp_ekite_itvw.png                       tags
404.html.gz  favicon.png   index.xml.gz    pp_ekite_itvw_hu_41404e93ad715bdf.webp  talks
```

et vérifier la taille des fichiers compressés :

```bash
$ ls -al public/index.*
.rw-rw-r-- jwittouck jwittouck  33 KB Wed Feb 11 12:15:21 2026 index.html
.rw-rw-r-- jwittouck jwittouck 9.4 KB Wed Feb 11 12:15:21 2026 index.html.gz
.rw-rw-r-- jwittouck jwittouck 9.0 KB Wed Feb 11 12:15:21 2026 index.html.zst
.rw-rw-r-- jwittouck jwittouck  67 KB Wed Feb 11 12:15:22 2026 index.xml
.rw-rw-r-- jwittouck jwittouck  18 KB Wed Feb 11 12:15:22 2026 index.xml.gz
.rw-rw-r-- jwittouck jwittouck  17 KB Wed Feb 11 12:15:22 2026 index.xml.zst
```

Pour ensuite servir les fichiers précompressés, il faut ajouter la [directive `precompressed`](https://caddyserver.com/docs/caddyfile/directives/file_server#precompressed) dans le `Caddyfile` :

```Caddyfile
# Clever Cloud needs us to listen on port 8080
:8080

file_server {
	# Clever Cloud serves the public directory
    root public
    # serve precompressed files
    precompressed
}

# Ask Caddy to compress static files 
encode
```

La directive `precompressed` recherchera dans l'ordre les fichiers `.zst` et `.gz` pour les servir en priorité, et utilisera comme _fallback_ une compression à la volée.

On peut ensuite simplement vérifier que les fichiers compressés sont bien servis compressés avec une commande `curl`.

Voici ce qui était renvoyé _avant_ la compression :

```bash
$ curl --head https://codeka.io

Content-Length: 81157
Content-Type: text/html; charset=utf-8
Server: Caddy
```

et la même commande après la compression :

```bash
$ curl --compressed --head https://codeka.io

HTTP/1.1 200 OK
Content-Encoding: zstd
Content-Type: text/html; charset=utf-8
Server: Caddy
Content-Length: 9
```

[//]: # (TODO) revérifier après le déploiement

On passe d'une page HTML de 81ko à une donnée compressée de 13ko, sans impacter le CPU du serveur puisque la compression se fait au build !

## Headers de sécurité

La dernière étape de cette configuration consiste à moderniser les headers servis pour impléments un peu de sécurité supplémentaire.

Maintenant que Caddy sert le site et que j'ai un Caddyfile sur lequel j'ai la main, je peux contrôler les headers HTTP renvoyés.

Pour savoir quoi faire, sur les conseils d'Antoine, j'ai utilisé l'analyseur de MDN :

https://developer.mozilla.org/en-US/observatory/analyze?host=codeka.io#scoring

![Résultat de l'analyse de MDN](mdn-analysis.png "Résultat de l'analyse de MDN")

### HSTS

Le premier header intéressant à utiliser est le `Strict-Transport-Security`.

Ce header a pour effet de forcer les navigateurs à utiliser HTTPS.
Bien que j'ai déjà configuré une redirection HTTP vers HTTPS sur mon domaine avec Clever Cloud, c'est une mesure de sécurité supplémentaire.

La recommandation de MDN est de positionner cette valeur :

```HTTP
Strict-Transport-Security: max-age=63072000
```

Dans mon Caddyfile, rien de plus simple, j'ajoute le header `Strict-Transport-Security` :

```Caddyfile
# Clever Cloud needs us to listen on port 8080
:8080

file_server {
	# Clever Cloud serves the public directory
    root public
    precompressed
}

# Custom headers for security
header {
	Strict-Transport-Security "max-age=63072000"
}

# Ask Caddy to compress static files 
encode
```

### Content-Security-Policy

Le premier header intéressant à utiliser est le `Content-Security-Policy`.
Ce header indique au navigateur quelle politique de sécurité appliquer à l'exécution des scripts provenant de sources externes au site web.
C'est une mesure de sécurité permettant de se prémunir des injections de type XSS (Cross-Site Scripting).

Le header doit déclarer l'ensemble des sources (domaines) acceptés pour le chargement des scripts, styles, images et autres ressources.
Utiliser ce header a aussi pour effet de désactiver le CSS et le JS "inline", ce qui est plutôt une bonne pratique.

Après avoir supprimé tous les styles inlines de mon site, j'ai configuré le header dans mon Caddyfile :

```Caddyfile
# Clever Cloud needs us to listen on port 8080
:8080

file_server {
	# Clever Cloud serves the public directory
    root public
    precompressed
}

# Custom headers for security
header {
	Strict-Transport-Security "max-age=63072000"
	
    Content-Security-Policy "
	    script-src 'self' codeka.io plausible.io;
	    frame-src 'self' plausible.io www.youtube-nocookie.com openfeedback.io;
        img-src 'self' img.shields.io;
	    default-src 'self';
	"
}

# Ask Caddy to compress static files 
encode
```

J'utilise plausible.io pour suivre les visites de mes articles, donc son script doit pouvoir être chargé. De la même manière, j'ai des iframes (bouh) sur les pages de mes talks qui référencent les videos Youtube ainsi que les feedbacks OpenFeedback.io. Je dois donc aussi autoriser ces ressources.

La directive `default-src` sert de fallback pour toutes les directives possibles, et indique que seul mon site est autorisé. 

## Conclusion

Ça m'a pris une bonne demi-journée pour mettre en place tous ces mécanismes, mais j'en ressort avec une meilleure compréhension de la sécurité et de la compression en HTTP.
J'ai aussi découvert Caddy, et amélioré mon fichier `mise.toml`.

Pour la plupart de mes lecteurs, l'impact de la compression sera probablement minime, car sur des réseaux performants, la différence de temps de chargement ne se ressentira peut-être pas beaucoup.
Mais avec une compression effectuée uniquement au build, c'est aussi une du CPU de moins de consommé, ce qui devrait pouvoir m'assurer de rester sur des instances les plus petites pour mon site le plus longtemps possible.

## Liens et références

* Configuration de l'optimisation des images avec Hugo : https://gohugo.io/configuration/imaging/#quality
* La méthode [Resize de Hugo](https://gohugo.io/methods/resource/resize/)
* [Le hookimage de Hugo](https://gohugo.io/render-hooks/images/#article)

* Documentation de Caddy :
  * [La directive `encode`](https://caddyserver.com/docs/caddyfile/directives/encode#syntax)
  * [La directive `precompressed`](https://caddyserver.com/docs/caddyfile/directives/file_server#precompressed)

* Documentation MDN :
  * [Content-Security-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Content-Security-Policy)

* [Precompressing Content With Hugo and Caddy](https://scottstuff.net/posts/2025/03/09/precompressing-content-with-hugo-and-caddy/)

* L'excellent talk de Antoine Caron et Hubert Sablonière : [La compression Web : comment (re)prendre le contrôle ?](https://www.youtube.com/watch?v=LWd0hr6ljZk)