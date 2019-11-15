Goto [Russian description](#как-найти-ip-камеру-после-прошивки)

## How to find IP camera after firmware update
After updating camera firmware, DHCP is usually turned on and the camera reboots with a new IP address. This utility allows to quickly find out new device location.

Before flashing the camera, you need to run the utility with `-c` parameter to make a quick scan of the local network.

```bash
./camdiff -c
Own IP: 192.168.10.245
Scanning network...
Found 36 hosts, saved to hosts file
```

After flashing, you need to run the utility again (without parameters). New hosts are displayed with a `+` sign, disappeared ones with a `-` sign.

```bash
./camdiff
Own IP: 192.168.10.245
Scanning network...
- 192.168.10.3
+ 192.168.10.197
```

## Supported platforms
Tested on Linux (Ubuntu 18.04), Windows 7 (x64).

Binary releases for Linux, Windows [are available](https://github.com/krabiswabbie/camdiff/releases/tag/v0.1).


## Quick scan
Scan of the local subnet (255.255.255.0) is performed with a timeout of 1s for the following ports only:

```bash
21 - ftp
22 - ssh
23 - telnet
80 - http
90 - dnsix
135 - msrpc
139 - netbios-ssn
443 - https
445 - microsoft-ds
554 - rtsp
8000 - http-alt
8022 - oa-system
8080 - http-proxy
```

This scan is designed for a quick search of local IP cameras and may not detect all the (other) devices.


## Как найти IP-камеру после прошивки
При обновлении прошивки камеры, обычно включается DHCP и камера перезагружается с новым IP-адресом. Данная утилита позволяет быстро установить новый адрес устройства. 

Перед прошивкой камеры необходимо запустить утилиту с ключом `-c` для быстрого сканирования локальной сети и сохранения результатов. 

```bash
./camdiff -c
Own IP: 192.168.10.245
Scanning network...
Found 36 hosts, saved to hosts file
```

После прошивки нужно снова запустить утилиту (без параметров). Новые хосты выводятся со знаком `+`, исчезнувшие со знаком `-`.

```bash
./camdiff
Own IP: 192.168.10.245
Scanning network...
- 192.168.10.3
+ 192.168.10.197
```

## Поддерживаемые платформы
Тестировался под Ubuntu 18.04, и Windows 7 (x64).

Исполняемые файлы для Linux, Windows можно скачать в разделе [Release](https://github.com/krabiswabbie/camdiff/releases/tag/v0.1).

## Быстрое сканирование
Сканирование локальной подсети (255.255.255.0) выполняется с таймаутом 1с по отдельным портам:

```bash
21 - ftp
22 - ssh
23 - telnet
80 - http
90 - dnsix
135 - msrpc
139 - netbios-ssn
443 - https
445 - microsoft-ds
554 - rtsp
8000 - http-alt
8022 - oa-system
8080 - http-proxy
```

Сканирование предназначено для быстрого поиска IP-камер в локальной сети и может не обнаруживать отдельные устройства.