### 실행 중인 VM에 접속합니다.

Running VMs can be identified by the color of the VM Name in the list.

- Green > Running
- Red > Stop

check list

You can't see it here, but if you look at the console, VM '20' is running.
the VM '20's name is green at the console.

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

SSH connection

If you connect as soon as the server is created, it may still be booting, so please try connecting several times or after a while.

(If a problem occurs, check with the virsh console command.)

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
