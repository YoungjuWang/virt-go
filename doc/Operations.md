### VM을 중지합니다.

시작 중인 VM을 종료합니다.

VM의 실행 여부는 VM Name의 색으로 구분할 수 있습니다.

- 녹색 > 실행 중
- 적색 > 중지

```
[root@virt-go-server ~]# virt-go stop -n 20
Stop virt-go-u20-20
```

### VM을 시작합니다.

```
[root@virt-go-server ~]# virt-go start -n 20
Start virt-go-u20-20
```

### VM을 재시작합니다.

```
[root@virt-go-server ~]# virt-go restart -n 20
Stop virt-go-u20-20
Start virt-go-u20-20
```

### VM Size를 변경합니다. (Size 늘리기)

VM Size를 변경하기 위해서는 먼저 VM을 중지해야합니다.

```
[root@virt-go-server ~]# virt-go stop -n 20
Stop virt-go-u20-20
```

현재 `20` 번 VM의 Size를 확인합니다.

```
[root@virt-go-server ~]# virt-go list 
-----------------------------
 RESOURCE     STATE          
-----------------------------
 Data-Dir     /data/virt-go  
 virt-go-net  10.62.62.xxx   
 Images       u20  u20cust   
-----------------------------


---------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION 
---------------------------------------------------------
 20      virt-go-u20-20  10.62.62.20  20 GB  resize      
---------------------------------------------------------
```

`20`번 VM의 Size를 `20GB` 에서 `50GB` 로 변경하겠습니다.

```
[root@virt-go-server ~]# virt-go resize -n 20 -s 50
resize /data/virt-go/volumes/virt-go-u20-20-root.qcow2
Image resized.
/data/virt-go/volumes/virt-go-u20-20-root.qcow2 size is 50G
```

변경된 Size를 확인합니다.

```
[root@virt-go-server ~]# virt-go list 
-----------------------------
 RESOURCE     STATE          
-----------------------------
 Data-Dir     /data/virt-go  
 virt-go-net  10.62.62.xxx   
 Images       u20  u20cust   
-----------------------------


---------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION 
---------------------------------------------------------
 20      virt-go-u20-20  10.62.62.20  50 GB  resize      
---------------------------------------------------------
```

VM을 실행하여 적용 여부를 확인합니다.

```
[root@virt-go-server ~]# virt-go start -n 20
Start virt-go-u20-20
```


`50GB` 로 용량이 증가했음을 볼 수 있습니다.

```
[root@virt-go-server ~]# virt-go ssh -n 20

root@virt-go-u20-20:~# df -h /
Filesystem      Size  Used Avail Use% Mounted on
/dev/sda1        49G  1.3G   47G   3% /
```

### VM Size를 변경합니다. (Size 줄이기)

Size 를 확장하는 것과 과정은 동일합니다.

VM 중지 > Size 변경 > VM 시작

> Size 를 축소할 시 Data가 유실될 가능성이 있으므로 신중하게 작업하시기 바랍니다.
> VM이 booting되지 않을 수 있습니다.

`--shrink` option을 추가하여 Size를 `50GB` 에서 `30GB` 로 변경합니다.

```
[root@virt-go-server ~]# virt-go resize -n 20 -s 30 --shrink
resize /data/virt-go/volumes/virt-go-u20-20-root.qcow2
Image resized.
/data/virt-go/volumes/virt-go-u20-20-root.qcow2 size is 30G
```
