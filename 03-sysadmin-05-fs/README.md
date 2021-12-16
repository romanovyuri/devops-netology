# Домашнее задание к занятию "3.5. Файловые системы"

1. Узнайте о [sparse](https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%B7%D1%80%D0%B5%D0%B6%D1%91%D0%BD%D0%BD%D1%8B%D0%B9_%D1%84%D0%B0%D0%B9%D0%BB) (разряженных) файлах.

    *Выполнено*

1. Могут ли файлы, являющиеся жесткой ссылкой на один объект, иметь разные права доступа и владельца? Почему?

    *Нет, не могут. Потому что жеская ссылка не является объектом файловой системы с собственными атрибутами доступа. Файл может иметь несколько жеских ссылок.*

1. Сделайте `vagrant destroy` на имеющийся инстанс Ubuntu.

    *Выполнено*

1. Используя `fdisk`, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство.

    `root@vagrant:~# fdisk /dev/sdb`

    `root@vagrant:~# fdisk /dev/sdc`

    *Выполнено*

1. Используя `sfdisk`, перенесите данную таблицу разделов на второй диск.

    `root@vagrant:~# sfdisk -d /dev/sdb | sudo sfdisk -f /dev/sdc`

1. Соберите `mdadm` RAID1 на паре разделов 2 Гб.

    `root@vagrant:~# mdadm --create /dev/md0 --level=1 --raid-devices=2 /dev/sdb1 /dev/sdc1`

1. Соберите `mdadm` RAID0 на второй паре маленьких разделов.

    `root@vagrant:~# mdadm --create /dev/md1 --level=0 --raid-devices=2 /dev/sdb5 /dev/sdc5`

1. Создайте 2 независимых PV на получившихся md-устройствах.

    `# pvcreate /dev/md0`

    `# pvcreate /dev/md1`

1. Создайте общую volume-group на этих двух PV.

    `# vgcreate vg0 /dev/md0 /dev/md1`

1. Создайте LV размером 100 Мб, указав его расположение на PV с RAID0.

    `# lvcreate --size 100m vg0 /dev/md1`

1. Создайте `mkfs.ext4` ФС на получившемся LV.

    `# mkfs.ext4 /dev/vg0/lvol0`

1. Смонтируйте этот раздел в любую директорию, например, `/tmp/new`.

    `root@vagrant:~# mkdir /tmp/new`

    `root@vagrant:~# mount /dev/vg0/lvol0 /tmp/new/`

1. Поместите туда тестовый файл, например `wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`.

    *Выполнено*

1. Прикрепите вывод `lsblk`.

    ```
    root@vagrant:~# lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
sda                    8:0    0   64G  0 disk  
├─sda1                 8:1    0  512M  0 part  /boot/efi
├─sda2                 8:2    0    1K  0 part  
└─sda5                 8:5    0 63.5G  0 part  
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
sdb                    8:16   0  2.5G  0 disk  
├─sdb1                 8:17   0    2G  0 part  
│ └─md0                9:0    0    2G  0 raid1 
└─sdb2                 8:18   0  511M  0 part  
  └─md1                9:1    0  510M  0 raid1 
    └─vg0-lvol0      253:2    0  100M  0 lvm   /tmp/new
sdc                    8:32   0  2.5G  0 disk  
├─sdc1                 8:33   0    2G  0 part  
│ └─md0                9:0    0    2G  0 raid1 
└─sdc2                 8:34   0  511M  0 part  
  └─md1                9:1    0  510M  0 raid1 
    └─vg0-lvol0      253:2    0  100M  0 lvm   /tmp/new

    ```

1. Протестируйте целостность файла:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

    

1. Используя pvmove, переместите содержимое PV с RAID0 на RAID1.

    `# pvmove /dev/md1 /dev/md0`

1. Сделайте `--fail` на устройство в вашем RAID1 md.

    `# mdadm --manage --fail /dev/md0 /dev/sdb1`

1. Подтвердите выводом `dmesg`, что RAID1 работает в деградированном состоянии.

    `# dmesg`

    ```
    [ 3970.980693] md/raid1:md0: Disk failure on sdb1, disabling device.
               md/raid1:md0: Operation continuing on 1 devices.

    ```

1. Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

    

