---
created: "2023-05-25"
date: "2023-05-25T00:00:00Z"
language: fr
lastmod: "2023-07-01"
tags:
- basics
title: Anatomy of an HTTP request
slug: anatomie-requete-http
---

HTTP, for _Hypertext Transfer Protocol_, is the main protocol for Internet exchanges. It's used by the browser you're using to read this article, as well as by other applications.
It's based on an exchange of requests and responses, between a client and a server, in text format. The advantage of text format is that it's easy to implement in any programming language.
The HTTP protocol is specified in [RFC 2616](https://www.rfc-editor.org/rfc/rfc2616), the very first version of which dates back to 1990.

This article deals with version 1.1 of the protocol, which was standardized in 1997. A version 2.0 was standardized in 2015, and extends version 1.1 with additional capabilities. The basics of version 1.1 are therefore not obsolete!

![](client-server.png)

In this article, we'll look at the structure of an HTTP request and response.

## The structure of a request

A request is structured in 3 parts:

* the mandatory first line of the request
* the optional headers
* the request body, also optional

```
REQUEST-LINE
[HEADERS]

[BODY]
```

Example of an HTTP _GET_ request without a request body:

```
GET /articles/http-anatomy?version=1#headers HTTP/1.1
Host: codeka.io
Accept: text/plain
Accept-Charset: utf-8
Accept-Encoding: gzip
Accept-Language: fr-fr
```

Example with a request body&nbsp;:

```
POST /articles/ HTTP/1.1
Host: codeka.io
Content-Type: application/json
Content-Length: 31
Authorization: Basic UmljayBBc3RsZXk6TmV2ZXIgR29ubmEgR2l2ZSBZb3UgVXAK

{"title":"my beautiful article"}
```

### The _Request Line_

The _HEADERS_ and _BODY_ are optional.

The _REQUEST-LINE_ is defined in this way:

```
METHOD REQUEST-URI HTTP-VERSION
```

_METHOD_ is the HTTP verb. It can be `GET`, `POST`, `PUT`, `DELETE`, which are the most common HTTP verbs.
RFC 2616 also defines additional methods, such as `HEAD`, `OPTIONS`, `TRACE`, `CONNECT`.

_REQUEST-URI_ is the resource to be called. It can be an absolute URI, like `https://codeka.io/articles/http-anatomy`, or a relative URI, like `/articles/http-anatomy`. In practice, relative URIs are used in combination with a _Header_ `Host`.

The `HTTP-VERSION` field is set to `HTTP/1.1`.

A minimal HTTP request is&nbsp;:

```
GET /articles/http-anatomy HTTP/1.1
```

### _URL_ and _URI_

A _URI_, for _Uniform Resource Identifier_, defines an Internet resource. This resource can be a web page, a document in PDF format, or a video.

A _URL_, for _Uniform Resource Locator_, is a particular _URI_, which contains the network information needed to access a resource.
_URIs_ are defined by a text format made up of _ASCII_ characters.

[RFC 3986](https://www.rfc-editor.org/rfc/rfc3986) specifies _URI_ formats.

The structure of an absolute _URI_ is as follows:

```
SCHEME://HOST/PATH[?QUERY][#FRAGMENT]
```
The _SCHEME_ is often associated with the protocol, _http_, _https_, or _file_ when local files are designated.

The _HOST_ is the name of the server hosting the resource.

The _PATH_ identifies the path to the resource.

The _QUERY_, also often called _Query Strings_, defines additional parameters that the resource can process, in the form `name=value`.

The _FRAGMENT_ defines the part of the resource consulted.

_QUERY_ and _FRAGMENT_ are optional.

A simple _URI_ for our article is&nbsp;:

```
https://codeka.io/articles/http-anatomy
\___/   \_______/ \___________________/
  |         |              |
scheme    host           path
```

A _URI_ with additional elements&nbsp;:

```
https://codeka.io/articles/http-anatomy?version=1#headers
\___/   \_______/ \___________________/ \_______/ \_____/
  |         |              |                |        |
scheme     host           path            query    fragment
```

A _URI_ can also be relative, and contain only part of the _PATH_.
For instance&nbsp;:

```
/http-anatomy?version=1#headers
```

### Request _Headers

_Headers_ define additional information about the request, such as the data format expected in response, or information about data currently cached on the client side.

This simple structure defines the _Headers_:

```
NAME: VALUE
```

[RFC 2616](https://www.rfc-editor.org/rfc/rfc2616#page-38), defines standard _Headers_, but applications can also define their own _Headers_.
The main standard _Headers_ are `Accept`, `Accept-Charset`, `Accept-Encoding` and `Accept-Language` for content negotiation.
This RFC also defines the _Header_ `Authorization` for client authentication. The _Header_ `Host` specifies the host to which the request is sent if it is not specified in the request URI.

An HTTP request containing _Headers_ is&nbsp;:

```
GET /articles/http-anatomy HTTP/1.1
Host: codeka.io
Accept: text/plain
Accept-Charset: utf-8
Accept-Encoding: gzip
Accept-Language: fr-fr
```

### The request _Body_

When a request includes a body, the _Headers_ `Content-Type` and `Content-Length` must be specified, defining its type and size in bytes.
A line break must also separate the _Headers_ from the _Body_.

Example of a query with a _Body_ in `JSON` format:

```
POST /articles/ HTTP/1.1
Host: codeka.io
Content-Type: application/json
Content-Length: 31

{"title":"my beautiful article"}
```
## The structure of a response

HTTP responses have a format quite similar to that of requests.

```
STATUS-LINE
[HEADERS]

[BODY]
```

### _Status Line_

The _Status Line_ represents the server's coded response to the request.

This is how _STATUS-LINE_ is defined:

```
HTTP-VERSION CODE PHRASE 
```

The first element of the response is the HTTP version `HTTP/1.1`.

The codes used are represented by 3 digits.

* `1xx` codes are informational and rarely used in practice.
* `2xx` codes represent a successful operation.
* `3xx` codes indicate a redirection.
* `4xx` codes indicate a client error (malformed request, non-existent resource).
* `5xx` codes represent a server error in processing the request.

Codes are followed by a descriptive _Phrase_.

The most common codes and associated phrases are defined in [RFC 2616](https://www.rfc-editor.org/rfc/rfc2616#section-6.1.1).

Example of an error response:

```
HTTP/1.1 404 Not Found
```

An example of a successful response:

```
HTTP/1.1 200 OK
```

### Response _Headers

Like a request, a response can contain _Headers_. They define the type and size of the response _Body_, or parameters to indicate to the client that the response can be cached. Other _Headers_ define the resource to be called in the event of a redirection (code `3xx`).

Response _Headers_ are defined in the same way as requests:

```
NAME: VALUE
```

Example of an HTTP response including a _Header_ `Location` for a redirect to an `index.html` page:

```
HTTP/1.1 301 Moved Permanently
Location: index.html
```

### The response _Body_

The _Body_ is defined in the same way as requests, and also imposes the same _Headers_ `Content-Type` and `Content-Length` rules.

Example of a successful response with the content of our article&nbsp;:

```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 31

{"title":"my beautiful article"}
```

Example of a server error, with an HTML page representing it:

```
HTTP/1.1 500 Internal Server Error
Content-Type: text/html
Content-Length: 56

<html><body><h1>Server Error</h1></body></html>
```

## Conclusion

The HTTP protocol describes a text format, in the form of a request and response, which is easy to implement in any programming language.
Its simplicity and expressiveness have been the key to its success since the 90s, with new versions continuing to use the same format.

## Références

* [RFC 2616](https://www.rfc-editor.org/rfc/rfc2616): Hypertext Transfer Protocol -- HTTP/1.1
* [RFC 3986](https://www.rfc-editor.org/rfc/rfc3986): Uniform Resource Identifier (URI): Generic Syntax
* [RFC 7540](https://www.rfc-editor.org/rfc/rfc7540): Hypertext Transfer Protocol Version 2 (HTTP/2)
* Cover photo by [Jordan Harrison](https://unsplash.com/@jouwdan?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) on [Unsplash](https://unsplash.com/photos/blue-utp-cord-40XgDxBfYXM?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
