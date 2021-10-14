### 실행 중인 VM에 접속합니다.

VM의 실행 중인 것은 list에서 VM Name의 색깔로 알 수 있습니다.

- 녹색 > 실행 중
- 적색 > 중지

list 확인

여기선 보이지 않으나 console에서 보면 `20`번 VM은 실행중입니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20          
---------------------------


------------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION    
------------------------------------------------------------
 20      virt-go-u20-20  10.62.62.20  20 GB  operation test 
------------------------------------------------------------
```

SSH 접속

서버를 생성하자마자 접속하는 경우 아직 Booting 중일 수 있기에 여러 번 혹은 조금 시간이 지난 뒤 접속해보시기 바랍니다.

(문제가 발생한 경우 virsh console 명령어로 확인하시기 바랍니다.)

```
[root@virt-go-server ~]# virt-go ssh -n 20
Warning: Permanently added '10.62.62.20' (ECDSA) to the list of known hosts.
Welcome to Ubuntu 20.04.3 LTS (GNU/Linux 5.4.0-88-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Thu Oct 14 05:18:35 UTC 2021

  System load:  0.78              Processes:             131
  Usage of /:   6.7% of 19.21GB   Users logged in:       0
  Memory usage: 4%                IPv4 address for ens2: 10.62.62.20
  Swap usage:   0%


1 update can be applied immediately.
To see these additional updates run: apt list --upgradable


Last login: Thu Oct 14 05:18:15 2021 from 10.62.62.1
root@virt-go-u20-20:~# 
```
