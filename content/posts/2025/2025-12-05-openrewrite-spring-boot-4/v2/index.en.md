---
date: 2025-12-12
lastmod: 2025-12-19
language: en
title: Spring Boot 4 Upgrade with OpenRewrite
tags:
  - java
  - spring-boot
params:
  original: "v1.md"
---

One of the projects I actively maintain is [GitLab Classrooms](/projects/gitlab-classrooms).

The code for this project is written in Spring Boot 3 and Java 25.
With the recent release of Spring Boot 4, I wanted to upgrade this project quickly.

To do that, I have two possibilities: either I do the upgrade manually, or I use a tool to do it automatically.

I took the opportunity to test OpenRewrite.

<!--more-->

## OpenRewrite

[OpenRewrite](https://docs.openrewrite.org/) is a tool that allows performing refactoring operations on Java code.
It relies on a concept of recipes, which implement transformations on the code.

I discovered this tool during [Jérôme Tama's talk at Devoxx France 2025](https://www.youtube.com/watch?v=aYHb7sLhsoQ).

## The Spring Boot 4 Recipe

A [Spring Boot 4 migration recipe](https://docs.openrewrite.org/recipes/java/spring/boot4/upgradespringboot_4_0-community-edition) is available for the community edition.

By browsing the recipe code on [Github](https://github.com/openrewrite/rewrite-spring/blob/main/src/main/resources/META-INF/rewrite/spring-boot-40.yml), it seems that the recipe does a large part of what is indicated in the Spring migration guide:

* upgrade of the `spring-boot-starter-parent` pom
* modifications related to coordinate changes of certain maven artifacts
* migration to Spring Framework and Spring Security 7
* update of deprecated properties
* update to testcontainers 2
* migration to modular starters

> [!INFO]
> The recipe evolves regularly, so maybe it does even more things by the time you read this article.

The recipe takes the form of a YAML file, and it is accompanied by code that implements the various transformations:

> I'm not going into the details of how OpenRewrite works, check out Jérôme Tama's talk mentioned above for more information.

```yaml
type: specs.openrewrite.org/v1beta/recipe
name: org.openrewrite.java.spring.boot4.UpgradeSpringBoot_4_0
displayName: Migrate to Spring Boot 4.0
description: >-
  Migrate applications to the latest Spring Boot 4.0 release. This recipe will modify an application's build files,
  make changes to deprecated/preferred APIs.
tags:
  - spring
  - boot
recipeList:
  - org.openrewrite.java.spring.boot3.UpgradeSpringBoot_3_5
  - org.openrewrite.java.spring.framework.UpgradeSpringFramework_7_0
  - org.openrewrite.java.spring.security7.UpgradeSpringSecurity_7_0
  - org.openrewrite.java.spring.batch.SpringBatch5To6Migration
  - org.openrewrite.java.spring.boot4.SpringBootProperties_4_0
  - org.openrewrite.java.spring.boot4.ReplaceMockBeanAndSpyBean
  - org.openrewrite.hibernate.MigrateToHibernate71
  - org.openrewrite.java.testing.testcontainers.Testcontainers2Migration
  - org.openrewrite.java.spring.boot4.MigrateToModularStarters
  - org.openrewrite.java.dependencies.UpgradeDependencyVersion:
      groupId: org.springframework.boot
      artifactId: "*"
      newVersion: 4.0.x
      overrideManagedVersion: false
  - org.openrewrite.java.dependencies.UpgradeDependencyVersion:
      groupId: org.springframework.boot
      artifactId: spring-boot-dependencies
      newVersion: 4.0.x
      overrideManagedVersion: true
  - org.openrewrite.maven.UpgradePluginVersion:
      groupId: org.springframework.boot
      artifactId: spring-boot-maven-plugin
      newVersion: 4.0.x
  - org.openrewrite.maven.UpgradeParentVersion:
      groupId: org.springframework.boot
      artifactId: spring-boot-starter-parent
      newVersion: 4.0.x
  - org.openrewrite.gradle.plugins.UpgradePluginVersion:
      pluginIdPattern: org.springframework.boot
      newVersion: 4.0.x

  # Replace deprecated starters with their new names
  # https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Migration-Guide#deprecated-starters
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-oauth2-authorization-server
      newArtifactId: spring-boot-starter-security-oauth2-authorization-server
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-oauth2-client
      newArtifactId: spring-boot-starter-security-oauth2-client
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-oauth2-resource-server
      newArtifactId: spring-boot-starter-security-oauth2-resource-server
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-web
      newArtifactId: spring-boot-starter-webmvc
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-web-services
      newArtifactId: spring-boot-starter-webservices
  # https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Migration-Guide#aop-starter-pom
  - org.openrewrite.java.dependencies.RemoveDependency:
      groupId: org.springframework.boot
      artifactId: spring-boot-starter-aop
      unlessUsing: org.aspectj.lang.annotation.*
  - org.openrewrite.java.dependencies.ChangeDependency:
      oldGroupId: org.springframework.boot
      oldArtifactId: spring-boot-starter-aop
      newArtifactId: spring-boot-starter-aspectj
```

The OpenRewrite documentation indicates that you can use a simple _Maven_ command to perform the migration:

```shell
mvn -U org.openrewrite.maven:rewrite-maven-plugin:run \
  -Drewrite.recipeArtifactCoordinates=org.openrewrite.recipe:rewrite-spring:RELEASE \
  -Drewrite.activeRecipes=org.openrewrite.java.spring.boot4.UpgradeSpringBoot_4_0 \
  -Drewrite.exportDatatables=true
```

> Quite convenient, because I won't have to modify my `pom.xml`, nor add a configuration file to my project to be able to perform this one-shot migration.

Running the command takes a few seconds and displays the operations performed (I've cleaned up the logs a lot to make it more readable):

```shell
[INFO] --- rewrite:6.25.0:run (default-cli) @ gitlab-classrooms ---
[INFO] Using active recipe(s) [org.openrewrite.java.spring.boot4.UpgradeSpringBoot_4_0]
[INFO] Using active styles(s) []
[INFO] Validating active recipes...
[INFO] Project [gitlab-classrooms] Resolving Poms...
[INFO] Project [gitlab-classrooms] Parsing source files
[INFO] Running recipe(s)...
[INFO] Printing available datatables to: target/rewrite/datatables/2025-12-12_17-33-51-247

[WARNING] Changes have been made to pom.xml by:
[WARNING]     org.openrewrite.java.spring.boot4.MigrateToModularStarters
[WARNING]     org.openrewrite.maven.UpgradeParentVersion: {groupId=org.springframework.boot, artifactId=spring-boot-starter-parent, newVersion=4.0.x}
[WARNING]     org.openrewrite.java.testing.testcontainers.Testcontainers2Migration

[WARNING] Changes have been made to src/main/resources/application-local.properties by:
[WARNING]     org.openrewrite.text.FindAndReplace: {find=javax., replace=jakarta., filePattern=**/*.js;**/*.ts;**/*.properties}

[WARNING] Changes have been made to src/test/java/fr/univ_lille/gitlab/classrooms/adapters/jpa/PostgresqlJPAAdaptersTest.java by:
[WARNING]     org.openrewrite.java.testing.testcontainers.Testcontainers2Migration
[WARNING]     org.openrewrite.java.spring.boot4.MigrateToModularStarters

[WARNING] Changes have been made to src/test/java/fr/univ_lille/gitlab/classrooms/mvc/ExportControllerMVCTest.java by:
[WARNING]     org.openrewrite.java.spring.boot4.ReplaceMockBeanAndSpyBean
[WARNING]         org.openrewrite.java.ChangeType: {oldFullyQualifiedTypeName=org.springframework.boot.test.mock.mockito.MockBean, newFullyQualifiedTypeName=org.springframework.test.context.bean.override.mockito.MockitoBean}

[...]

[WARNING] Please review and commit the results.
[WARNING] Estimate time saved: 1h 16m
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  20.769 s
[INFO] Finished at: 2025-12-12T17:33:51+01:00
[INFO] ------------------------------------------------------------------------
```

OpenRewrite seems to have run correctly and indicates that several files have been modified. A `git status` allows to see what has been impacted:

```shell
git status
On branch feature/migration-spring-boot-4
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   pom.xml
	modified:   src/main/resources/application-local.properties
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/adapters/jpa/PostgresqlJPAAdaptersTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/api/UploadJunitGradingRestControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/assignments/AssignmentScoreServiceImplTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/assignments/AssignmentServiceImplTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/domain/classrooms/ClassroomStudentControllerTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/mvc/ClassroomControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/mvc/ExportControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/mvc/assignments/AssignmentMVCControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/mvc/assignments/StudentAssignmentResetGradeMVCControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/mvc/dashboard/DashboardControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/quiz/QuizAnswerControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/quiz/QuizEditionControllerMVCTest.java
	modified:   src/test/java/fr/univ_lille/gitlab/classrooms/quiz/QuizServiceImplTest.java

no changes added to commit (use "git add" and/or "git commit -a")
```

* the `pom.xml` (which was expected first)
* the properties configuration files (some properties were renamed)
* the test files (mainly for the deprecation of `@MockBean` and `@SpyBean`)

A `git diff` allows to check all that:

```shell
git diff pom.xml

@@ -6,7 +6,7 @@
     <parent>
         <groupId>org.springframework.boot</groupId>
         <artifactId>spring-boot-starter-parent</artifactId>
-        <version>3.5.6</version>
+        <version>4.0.0</version>
         <relativePath/> <!-- lookup parent from repository -->
     </parent>
 
@@ -44,7 +44,7 @@
-        <jacoco-maven-plugin.version>0.8.13</jacoco-maven-plugin.version>
+        <jacoco-maven-plugin.version>0.8.14</jacoco-maven-plugin.version>

@@ -57,7 +57,7 @@
         <dependency>
             <groupId>org.springframework.boot</groupId>
-            <artifactId>spring-boot-starter-web</artifactId>
+            <artifactId>spring-boot-starter-webmvc</artifactId>
         </dependency>

@@ -67,12 +67,12 @@
         <dependency>
             <groupId>org.springframework.boot</groupId>
-            <artifactId>spring-boot-starter-oauth2-client</artifactId>
+            <artifactId>spring-boot-starter-security-oauth2-client</artifactId>
         </dependency>
 
         <dependency>
             <groupId>org.springframework.boot</groupId>
-            <artifactId>spring-boot-starter-oauth2-resource-server</artifactId>
+            <artifactId>spring-boot-starter-security-oauth2-resource-server</artifactId>
         </dependency>
         
@@ -91,8 +91,8 @@
         <dependency>
-            <groupId>org.flywaydb</groupId>
-            <artifactId>flyway-core</artifactId>
+            <groupId>org.springframework.boot</groupId>
+            <artifactId>spring-boot-starter-flyway</artifactId>
         </dependency>
         
@@ -127,22 +127,32 @@
+    <dependency>
+      <groupId>org.springframework.boot</groupId>
+      <artifactId>spring-boot-starter-webmvc-test</artifactId>
+      <scope>test</scope>
+    </dependency>
+
+    <dependency>
+      <groupId>org.springframework.boot</groupId>
+      <artifactId>spring-boot-starter-data-jpa-test</artifactId>
+      <scope>test</scope>
+    </dependency>
 
@@ -136,7 +136,7 @@
 
         <dependency>
             <groupId>org.testcontainers</groupId>
-            <artifactId>postgresql</artifactId>
+            <artifactId>testcontainers-postgresql</artifactId>
             <scope>test</scope>
         </dependency>
         
         <dependency>
-            <groupId>org.springframework.security</groupId>
-            <artifactId>spring-security-test</artifactId>
+            <groupId>org.springframework.boot</groupId>
+            <artifactId>spring-boot-starter-security-test</artifactId>
             <scope>test</scope>
         </dependency>
```

At the `pom.xml` level, everything went well, all the expected modifications have been applied.

The new modular architecture of Spring Boot 4 was correctly handled.

The test code was also cleaned of the old deprecated `@MockBean`:

```text
@@ -37,10 +37,10 @@ class ClassroomControllerMVCTest {
     @Autowired
     private MockMvc mockMvc;
 
-    @MockBean
+    @MockitoBean
     private ClassroomService classroomService;
 
-    @MockBean
+    @MockitoBean
     private Gitlab gitlab;
```

and some properties (which were in comments) were correctly renamed.

```text
 # generate full creation sql script if needed
-spring.jpa.properties.javax.persistence.schema-generation.create-source=metadata
-spring.jpa.properties.javax.persistence.schema-generation.scripts.action=update
-spring.jpa.properties.javax.persistence.schema-generation.scripts.create-target=update.sql
+spring.jpa.properties.jakarta.persistence.schema-generation.create-source=metadata
+spring.jpa.properties.jakarta.persistence.schema-generation.scripts.action=update
+spring.jpa.properties.jakarta.persistence.schema-generation.scripts.create-target=update.sql
```

## The adjustments I had to do manually

After this execution, my code doesn't compile.

An import in my Spring Security configuration is not resolved

```java
import org.springframework.boot.actuate.autoconfigure.security.servlet.EndpointRequest;
```

because this class was moved to another package:

```java
import org.springframework.boot.security.autoconfigure.actuate.web.servlet.EndpointRequest;
```

Once these small adjustments were made, I ran my unit tests.

This time, I got an error message related to Spring Security at startup:

```text
Caused by: java.lang.IllegalArgumentException: pattern must start with a /
```

I hadn't paid attention to this change in the migration guides, so I might have missed it. Regardless, it's not a very complicated change, I easily applied it.

Once these last adjustments were made, the tests pass correctly 🎉:

![Screenshot of my unit tests passing!](tests.webp)

## Conclusion

It took me about 1 hour to migrate my project from Spring Boot 3.5 to Spring Boot 4.0.

OpenRewrite clearly made the work easier, it modified all my dependencies, and migrated deprecated annotations (which would have been tedious).
I still had to finalize the migration manually, and I couldn't skip reading the [Spring Boot 4.0 Migration Guide](https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-4.0-Migration-Guide).

I think that Spring Boot 4 support in OpenRewrite is still in its early stages (the version introducing support was published on December 5, 2025, and the update for the modular architecture was published on December 16), so it's not impossible that the operations I had to do manually will be automated in the future.

Anyway, 1 hour of work to migrate a project of about 3,000 lines of code, I think it's quite efficient.
