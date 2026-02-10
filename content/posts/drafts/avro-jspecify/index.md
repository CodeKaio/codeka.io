---
date: 2026-02-10
title: Apache Avro et JSpecify
draft: true
---

La majorité des projets Java utilisant des communications Kafka s'appuie sur le format Avro.

Ce format est très pratique, puisqu'il permet de bien définir la structure des échanges avec un schéma de données partagé, peut assurer la compatibilité (ascendante ou descendante) des consommateurs, et est plutôt efficace sur les échanges avec la sérialisation en binaire.

Le plugin Maven Avro (qui n'est pas très bien documenté) permet de générer les classes Java à partir d'un schéma.

Les classes peuvent ensuite être utilisées dans le code applicatif.

## Configurer le plugin Maven Avro pour JSpecify

Le guide [Getting Started (Java)](https://avro.apache.org/docs/1.12.0/getting-started-java/) propose la configuration minimale suivante :

```xml
<plugin>
  <groupId>org.apache.avro</groupId>
  <artifactId>avro-maven-plugin</artifactId>
  <version>1.12.0</version>
  <configuration>
    <sourceDirectory>${project.basedir}/src/main/avro/</sourceDirectory>
    <outputDirectory>${project.basedir}/src/main/java/</outputDirectory>
  </configuration>
  <executions>
    <execution>
      <phase>generate-sources</phase>
      <goals>
        <goal>schema</goal>
      </goals>
    </execution>
  </executions>
</plugin>
```

Le goal `schema` est donc invoqué lors de la phase `generate-sources` de Maven (avant la compilation du code).

Pour activer le support des annotations JSpecify, il faut ajouter 3 paramètres de configuration :

* `createNullSafeAnnotations` pour activer la génération des annotations (annotations JetBrains par défaut)
* `nullSafeAnnotationNullable` pour indiquer quelle annotation `@Nullable` doit être utilisée
* `nullSafeAnnotationNotNull` pour indiquer quelle annotation `@NonNull` doit être utilisée

Ce qui donne le paramétrage complet suivant :

```xml
<plugin>
  <groupId>org.apache.avro</groupId>
  <artifactId>avro-maven-plugin</artifactId>
  <version>1.12.0</version>
  <configuration>
    <sourceDirectory>${project.basedir}/src/main/avro/</sourceDirectory>
    <outputDirectory>${project.basedir}/src/main/java/</outputDirectory>
    <!-- JSpecify support -->  
    <createNullSafeAnnotations>true</createNullSafeAnnotations>
    <nullSafeAnnotationNullable>org.jspecify.annotations.Nullable</nullSafeAnnotationNullable>
    <nullSafeAnnotationNotNull>org.jspecify.annotations.NonNull</nullSafeAnnotationNotNull>  
  </configuration>
  <executions>
    <execution>
      <phase>generate-sources</phase>
      <goals>
        <goal>schema</goal>
      </goals>
    </execution>
  </executions>
</plugin>
```



## Liens et références

* Guide officiel de démarrage Avro pour Java : [Getting Started (Java)](https://avro.apache.org/docs/1.12.0/getting-started-java/)
* Le [commit sur GitHub](https://github.com/apache/avro/commit/067c44024c3db832e731c051ba0178f0c4f18a04) qui a apporté le support des annotations null-safe
* La javadoc du plugin Maven Avro : https://avro.apache.org/docs/1.12.0/api/java/org/apache/avro/mojo/AbstractAvroMojo.html