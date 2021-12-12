# Домашнее задание к занятию "3.3. Операционные системы, лекция 1"

1. Какой системный вызов делает команда `cd`? В прошлом ДЗ мы выяснили, что `cd` не является самостоятельной  программой, это `shell builtin`, поэтому запустить `strace` непосредственно на `cd` не получится. Тем не менее, вы можете запустить `strace` на `/bin/bash -c 'cd /tmp'`. В этом случае вы увидите полный список системных вызовов, которые делает сам `bash` при старте. Вам нужно найти тот единственный, который относится именно к `cd`. Обратите внимание, что `strace` выдаёт результат своей работы в поток stderr, а не в stdout. 
  
    `strace /bin/bash -c 'cd /tmp' 2>&1 | grep "/tmp"`
     >chdir("/tmp") = 0
   
2. Попробуйте использовать команду `file` на объекты разных типов на файловой системе. Например:
    ```bash
    vagrant@netology1:~$ file /dev/tty
    /dev/tty: character special (5/0)
    vagrant@netology1:~$ file /dev/sda
    /dev/sda: block special (8/0)
    vagrant@netology1:~$ file /bin/bash
    /bin/bash: ELF 64-bit LSB shared object, x86-64
    ```
    Используя `strace` выясните, где находится база данных `file` на основании которой она делает свои догадки.

     `openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3` 

      */usr/share/misc/magic.mgc* 


4. Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).
 
 *Найти c помощью команды `lsof` PID процесса и номер файлового дескриптора FDNUM удаленного файла и обнулить файл:*

  `echo -n > /proc/$PID/fd/$FDNUM`





5. Занимают ли зомби-процессы какие-то ресурсы в ОС (CPU, RAM, IO)?

  *Имеют только PID.  Ресурсы не занимают*

7. В iovisor BCC есть утилита `opensnoop`:
    ```bash
    root@vagrant:~# dpkg -L bpfcc-tools | grep sbin/opensnoop
    /usr/sbin/opensnoop-bpfcc
    ```
На какие файлы вы увидели вызовы группы `open` за первую секунду работы утилиты? Воспользуйтесь пакетом `bpfcc-tools`   

   `$ sudo opensnoop-bpfcc -T -U`

   ```bash
           TIME(s)       UID   PID    COMM               FD ERR PATH
           0.000000000   0     1      systemd            18   0 /proc/379/cgroup
           0.000191000   0     1      systemd            18   0 /proc/765/cgroup
           1.788336000   0     766    thermald            9   0 /sys/class/thermal/thermal_zone2/temp
           1.934418000   1000  19890  CrBrowserMain      58   0 /proc/meminfo
           2.512688000   1000  2392   ThreadPoolForeg    33   0 /home/yuri/.config/google-chrome/Default/Cookies-journal
           2.553690000   1000  2392   ThreadPoolForeg    38   0 /home/yuri/.config/google-chrome/Default
           3.615735000   1000  19890  Timer-0            58   0 /home/yuri/.java/.userPrefs/.user.lock.yuri
           5.792520000   0     766    thermald            9   0 /sys/class/thermal/thermal_zone2/temp
           6.483105000   0     744    irqbalance          6   0 /proc/interrupts
           6.483276000   0     744    irqbalance          6   0 /proc/stat
           6.483372000   0     744    irqbalance          6   0 /proc/irq/24/smp_affinity
           6.483404000   0     744    irqbalance          6   0 /proc/irq/25/smp_affinity
           6.483437000   0     744    irqbalance          6   0 /proc/irq/27/smp_affinity
   ```
    
8. Какой системный вызов использует `uname -a`? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в `/proc`, где можно узнать версию ядра и релиз ОС.

 `strace uname -a` 

  `uname({sysname="Linux", nodename="MyUbuntuHost", ...}) = 0` 

  `uname({sysname="Linux", nodename="MyUbuntuHost", ...}) = 0` 


9. Чем отличается последовательность команд через `;` и через `&&` в bash? 
   Есть ли смысл использовать в bash `&&`, если применить `set -e`? 

      *Последовательность команд через `;` будет выполнена в любом случае.*  
      *В последовательности команд через `$$` каждая последующая будет выполнена только тогда, когда предидыщая закончила работу с кодом возврата "0"".

    `set -e` *завершит работу как только одна из команд завершит работу с кодом возврата "0". Поэтому использование  `&&` не имеет смысла.* 


10. Из каких опций состоит режим bash `set -euxo pipefail` и почему его хорошо было бы использовать в сценариях?

 `$ set --help`

    >-e Exit immediately if a command exits with a non-zero status.  
    >-u Treat unset variables as an error when substituting.  
    >-x Print commands and their arguments as they are executed.  
    >-o pipefail the return value of a pipeline is the status of the last command to exit with a non-zero status, or zero if no command exited with a non-zero status

    - *Остановка, в случае появления кода возврата не равного 0*
    - *Если есть неустановленные переменные вызывать ошибку*  
    - *Выводить комманды и их аргументы, как они буду выполнены*
    - *Включает возникновение ошибки в pipe, если комманда справа вернула ненулевой код выхода*

11. Используя `-o stat` для `ps`, определите, какой наиболее часто встречающийся статус у процессов в системе. В `man ps` ознакомьтесь (`/PROCESS STATE CODES`) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).

     *Наиболее часто встречающиеся статусы*
      `I<` *Неактивный поток ядра, высокий приоритет*  
      `S` *Процесс в ждущем режиме*  
      `Ss` *Процесс в ждущем режиме, лидер сессии*  
      `Ssl` *Процесс в ждущем режиме, лидер сессии, с несколькими ветками.*