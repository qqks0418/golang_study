FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD password
ENV MYSQL_DATABASE test_db
ENV MYSQL_USER test_user
ENV MYSQL_PASSWORD test_pass

COPY ./config/my.conf /etc/mysql/conf.d/my.cnf

# docker build -t docker-mysql:1 -f Dockerfile .
# docker image ls
# docker run --name docker-mysql -d -v $PWD/mysql_db:/var/lib/mysql -p 23306:3306 docker-mysql:1
# docker exec -it docker-mysql bash