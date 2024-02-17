<!-- customize-category:IDEA -->

# How to run multiple Tomcat projects in Intellij IDEA Community

- [How to run multiple Tomcat projects in Intellij IDEA Community](#how-to-run-multiple-tomcat-projects-in-intellij-idea-community)
  - [Prerequisites](#prerequisites)
  - [Tomcat Configuration](#tomcat-configuration)
  - [Start Tomcat](#start-tomcat)
  - [Auto deploy project to Tomcat](#auto-deploy-project-to-tomcat)
  - [DEBUG](#debug)

## Prerequisites

1. Install Intellij IDEA Community.
2. Install Tomcat
3. Install Maven

## Tomcat Configuration

Open `File -> Settings` and expand `Tools` and select `External Tools`

<img width=400 src='/assets/image/1706625961.png'/>

## Start Tomcat

<img width=400 src='/assets/image/1706626004.png'/>

## Auto deploy project to Tomcat

```text
war:war org.codehaus.mojo:wagon-maven-plugin:upload-single -Dwagon.fromFile=C:\Users\lance\Code\personal\test\test2\target\test2-1.0-SNAPSHOT.war -Dwagon.url=file://D:\App\apache-tomcat-8.5.98\webapps\ -Dwagon.toFile=a#b.war
```

<img width=400 src='/assets/image/1706626049.png'/>

## DEBUG

<img width=400 src='/assets/image/1706626082.png'/>

> <https://stefancosma.xyz/2018/10/01/how-to-use-tomcat-intellij-idea-community/>
