# module-test

## Network
<pre>
$ docker network create --driver bridge mybridge
</pre>

## MYSQL Container
<pre>
$ docker volume create mysql-vol

$ docker run --name=mysql-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=giri -d -v mysql-vol:/var/lib/mysql -—network mybridge mysql

$ docker inspect mysql-db --format '{{json .Mounts}}' | jq

$ docker container logs mysql-db

$ docker stop mysql-db

$ docker start mysql-db
</pre>

## Go Application
<pre>
$ docker build -t demo-api . 

$ docker run -d -p 8080:8080 --name=demo-api --network mybridge demo-api

$ docker logs demo-api
</pre>
