### Image 삭제

생성된 Image를 삭제합니다.

list에서 보이는 `u2004` Image를 삭제하겠습니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20  u2004   
---------------------------
...
```

삭제 진행

```
[root@virt-go-server ~]# virt-go delete -i u2004
■  Delete u2004.
successfully finished
```

list에서 삭제됐음을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20          
---------------------------
...
```

### VM 삭제

`30`번 VM을 삭제하겠습니다.

```
[root@virt-go-server ~]# virt-go delete -n 30
■  Shutdown virt-go-u20-30.
■  Undefine virt-go-u20-30.
■  Delete virt-go-u20-30 volume.
■  Delete virt-go-u20-30 cloud-init iso file.
■  Delete virt-go-u20-30 description file.
successfully finished
```

list에서 삭제됐음을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20          
---------------------------


--------------------------------------------------------------
 NUMBER  NAME              IP           SIZE   DESCRIPTION    
--------------------------------------------------------------
 40      virt-go-u2004-40  10.62.62.40  20 GB  test 40 server 
--------------------------------------------------------------
```

### Image와 VM을 동시에 삭제

`u20` Image와 `40`번 VM을 동시에 삭제하겠습니다.

```
[root@virt-go-server ~]# virt-go delete -n 40 -i u20
■  Shutdown virt-go-u2004-40.
■  Undefine virt-go-u2004-40.
■  Delete virt-go-u2004-40 volume.
■  Delete virt-go-u2004-40 cloud-init iso file.
■  Delete virt-go-u2004-40 description file.
■  Delete u20.
successfully finished
```

list에서 삭제됐음을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images                    
---------------------------


-------------------------------------
 NUMBER  NAME  IP  SIZE  DESCRIPTION 
-------------------------------------
-------------------------------------
```
