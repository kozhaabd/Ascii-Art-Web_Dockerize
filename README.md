# Ascii-art-web-dockerize
## Description
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of ascii-art project,which are work with terminal .

Ascii-art-web-dockerize here we try cofigurate our project by using technology Docker. Docker is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.


## Links

 - [What is a docker](https://selectel.ru/blog/what-is-docker/#whatis)
 - [Dockerfile good practices.](https://github.com/matiassingers/awesome-readme)
 - [How to write a Good readme](https://readme.so/editor)

## Usage:

1.Create image for Dockerfile

  
    docker build -t ascii-art-web . 

2.Enter this command to start the program

    docker run --name cont -dp 3333:3000 ascii-art-web
3.Open the web browser and go to

    http://127.0.0.1:3333/
    