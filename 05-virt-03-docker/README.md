# Домашнее задание к занятию "5.3. Введение. Экосистема. Архитектура. Жизненный цикл Docker контейнера"


## Задача 1

Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

https://hub.docker.com/r/romanovyp/mynginx

`$ sudo docker run -d nginx`

```
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
461246efe0a7: Pull complete 
060bfa6be22e: Pull complete 
b34d5ba6fa9e: Pull complete 
8128ac56c745: Pull complete 
44d36245a8c9: Pull complete 
ebcc2cc821e6: Pull complete 
Digest: sha256:1761fb5661e4d77e107427d8012ad3a5955007d997e0f4a3d41acc9ff20467c7
Status: Downloaded newer image for nginx:latest
92d8295004eb978bf324642bd8f07e8dbebbf549bf9f843312be28c7d9789aff
```

`$ sudo docker ps`

```
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS     NAMES
92d8295004eb   nginx     "/docker-entrypoint.…"   2 minutes ago   Up 4 minutes   80/tcp    modest_kirch

```

`$ sudo docker stop modest_kirch`

*Создаем файл index.html*

```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I'm DevOps Engineer!</h1>
</body>
</html>
```

*Создаем Dockerfile*

```
FROM nginx:latest
COPY ./index.html /usr/share/nginx/html/index.html
```

`$ sudo docker build -t rom/mynginx:v1 .`

*Проверяем работоспособность*

`$ sudo docker run -it --rm -d -p 80:80 rom/mynginx:v1`

`$ sudo docker push romanovyp/mynginx`

## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Сценарий:

- Высоконагруженное монолитное java веб-приложение;

    *Лучше использовать виртуализацию.  Докер лучше использовать где много небольших сервисов, там же где речь о монолитном приложении, лучше использовать виртуализацию или физическую машину. Но по возможностям развернуть приложениее, переехать, физическая проигрывает виртуализации.*

- Nodejs веб-приложение;

    *Докер в этом случае хорошо справится с распределенной структурой.*

- Мобильное приложение c версиями для Android и iOS;

    *Если приложени создано как куча микросервисов, то Докер подойдет хорошо. Если как монолит с БД, то лучше виртуалки не найти.*

- Шина данных на базе Apache Kafka;

    *Спроектирована как распределенная, горизонтально масштабируемая система. Докер должен подойти хорошо.*

- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;

    *Для высокомаштабируемой распределенной системы докер подойдет очень хорошо*

- Мониторинг-стек на базе Prometheus и Grafana;

    *Хорошо подойдет докер. Докер как нельзя лучше подходит для распределенных систсем и инфраструктур.*

- MongoDB, как основное хранилище данных для java-приложения;

    *Виртуализация скорее подходить лучше, но Докер так же можно применить.*

- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.

    *Докер должен подойти под эти задачи. Масштабируемвя система с распределением нагрузки, внутренними сетями*

## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;

`$ sudo docker run -it --rm -d --name centos -v ~/devops-netology/05-virt-03-docker/data:/data centos`

- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;

`$ sudo docker run -it --rm -d --name debian -v ~/devops-netology/05-virt-03-docker/data:/data debian`

- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;

```
[root@8cab317a85ba data]# touch from_centos    
[root@8cab317a85ba data]# ls
from_centos
[root@8cab317a85ba data]# echo test1 > from_centos 
[root@8cab317a85ba data]# cat from_centos 
test1
[root@8cab317a85ba data]# exit

```

- Добавьте еще один файл в папку ```/data``` на хостовой машине;

```
echo test2 > from_host
yuri@uflash:~/devops-netology/05-virt-03-docker/data$ cat from_host 
test2

```

- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.

```
$ sudo docker exec -it debian bash
root@b4531eee79a5:/# cd data
root@b4531eee79a5:/data# ls
from_centos  from_host
root@b4531eee79a5:/data# cat *
test1
test2

```



