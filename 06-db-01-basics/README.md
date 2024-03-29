# Домашнее задание к занятию "6.1. Типы и структура СУБД"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Архитектор ПО решил проконсультироваться у вас, какой тип БД 
лучше выбрать для хранения определенных данных.

Он вам предоставил следующие типы сущностей, которые нужно будет хранить в БД:

- Электронные чеки в json виде

*Документо-ориентированная NoSQL. Основная единица логичесокй модели данных является документ- структурированный текст, с определенным снтаксисом.Хорошо подходит под формат json.*

- Склады и автомобильные дороги для логистической компании

*Графовая NoSQL БД.Склады и дороги даже визуально напоминают графы с их узлами и отношениями между узлами. *

- Генеалогические деревья

*Сетевая БД. Имеет иерархическую структуру, а так же можно описать сложные отношения родства.*

- Кэш идентификаторов клиентов с ограниченным временем жизни для движка аутентификации

*БД ключ-значение.Каждая сущность представляет сосбой набор пар ключ-значение.*

- Отношения клиент-покупка для интернет-магазина

*Реляционная БД. Структура БД строится на отношениях между данными, которые хранятся в виде таблиц. Используется SQL для взаимодействия с данными.*

Выберите подходящие типы СУБД для каждой сущности и объясните свой выбор.

## Задача 2

Вы создали распределенное высоконагруженное приложение и хотите классифицировать его согласно 
CAP-теореме. Какой классификации по CAP-теореме соответствует ваша система, если 
(каждый пункт - это отдельная реализация вашей системы и для каждого пункта надо привести классификацию):

- Данные записываются на все узлы с задержкой до часа (асинхронная запись).

*По CAP - AP, по PACELC - PA/EL*

- При сетевых сбоях, система может разделиться на 2 раздельных кластера

*По CAP - AP, по PACELC - PA/EC*

*Возможны два варианта: Это потеря консистентности, (когда кластера после разделения по разному обслуживают клиентов) тогда По CAP - AP, по PACELC - PA/EC.

*Второй вариант - потеря доступности одного из кластеров - тогда по CAP - CP, по PACELC PC/EC.* 

- Система может не прислать корректный ответ или сбросить соединение

*По CAP - CP, по PACELC - PC/EC*

А согласно PACELC-теореме, как бы вы классифицировали данные реализации?

## Задача 3

Могут ли в одной системе сочетаться принципы BASE и ACID? Почему?

*Нет. Эти принципы  противопоставляются друг другу. ACID во главу угла ставит согласованность, а BASE доступность.*

## Задача 4

Вам дали задачу написать системное решение, основой которого бы послужили:

- фиксация некоторых значений с временем жизни
- реакция на истечение таймаута

Вы слышали о key-value хранилище, которое имеет механизм [Pub/Sub](https://habr.com/ru/post/278237/). 
Что это за система? Какие минусы выбора данной системы?

*Redis реализует классический pub/sub и является key-value data storage.*

*Минусы: не поддерживается SQL, набор данных должен удобно помещаться в память.*

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---