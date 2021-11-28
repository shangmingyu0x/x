# Dockerfile补充说明
## 远端官方库
[官方镜像地址](https://hub.docker.com/repository/docker/shangmingyu/httpserver)
```
docker pull shangmingyu/httpserver:v0.5
```
## 通过 nsenter 进入容器查看 IP 配置
```
root@ubuntu:/home/ubuntu# docker container ls | grep shangmingyu/httpserver:v0.5
21233b4d1ee9   shangmingyu/httpserver:v0.5                         "/app/httpserver"        34 seconds ago   Up 33 seconds   8888/tcp   elated_benz

root@ubuntu:/home/ubuntu# docker container top 21233b4d1ee9
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                208753              208719              0                   17:14               ?                   00:00:00            /app/httpserver

root@ubuntu:/home/ubuntu# docker container top 21233b4d1ee9
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                208753              208719              0                   17:14               ?                   00:00:00            /app/httpserver

root@ubuntu:/home/ubuntu# nsenter -t 208753 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
19: eth0@if20: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
``` 