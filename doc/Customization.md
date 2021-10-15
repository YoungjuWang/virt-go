### Description 변경

각 VM의 description은 data dir의 volumes 내에 VM이름 명의 파일로 저장돼있습니다.

해당 파일의 내용을 변경하면 Description 변경이 가능합니다.

data dir 경로 확인

```
[root@virt-go-server ~]# virt-go list
----------------------------
 RESOURCE     STATE         
----------------------------
 Data-Dir     /data/virt-go 
 virt-go-net  10.62.62.xxx  
 Images       u20           
----------------------------


---------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION 
---------------------------------------------------------
 20      virt-go-u20-20  10.62.62.20  20 GB  test server 
---------------------------------------------------------
```

위에서 `Data-Dir` 경로를 확인했습니다.

이제 `20` 번 VM의 Description을 변경해보겠습니다.
해당 파일의 text를 변경하면 변경됩니다.

```
[root@virt-go-server ~]# cat /data/virt-go/volumes/virt-go-u20-20
test server 2222
```

변경됐음을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
----------------------------
 RESOURCE     STATE         
----------------------------
 Data-Dir     /data/virt-go 
 virt-go-net  10.62.62.xxx  
 Images       u20           
----------------------------


--------------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION      
--------------------------------------------------------------
 20      virt-go-u20-20  10.62.62.20  20 GB  test server 2222 
--------------------------------------------------------------
```


### cloud-init data 변경

`user-data` 파일은 data dir 하위 cloud-init 경로에 위치해있습니다.

```
[root@virt-go-server ~]# ls -l /data/virt-go/cloud-init/user-data 
-rw-r--r--. 1 root root 1084 Oct 14 14:32 /data/virt-go/cloud-init/user-data
```

해당 파일의 내용을 변경하면 VM 생성 시 사용되는 cloud-init config의 내용을 변경할 수 있습니다.


### Custom Image 만들기

기존 서버의 root volume을 Image로 생성합니다.

`20` 번 서버에 파일 생성 및 web 서버를 실행시킵니다.

```
root@virt-go-u20-20:~# echo "hello custom server" > /root/evidence

root@virt-go-u20-20:~# apt -y install apache2

root@virt-go-u20-20:~# echo "custom server" > /var/www/html/index.html 

root@virt-go-u20-20:~# systemctl enable apache2
```

이후 해당 서버가 Image로 변환될 때 cloud-init이 다시 실행될 수 있도록
cloud-init을 초기화합니다.

```
root@virt-go-u20-20:~# cloud-init clean

root@virt-go-u20-20:~# cloud-init status
status: not run
```

VM을 Image로 변환하기 위해 중지합니다.

```
[root@virt-go-server ~]# virt-go stop -n 20
```

해당 VM의 root volume을 `u20cust` Image로 생성합니다.

```
[root@virt-go-server ~]# ls -l /data/virt-go/volumes/virt-go-u20-20-root.qcow2 

[root@virt-go-server ~]# virt-go create -i u20cust
Create image only.
!! u20cust doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /data/virt-go/volumes/virt-go-u20-20-root.qcow2

■  Create image.
⠇ Create u20cust (1.1 GB, 1341.759 MB/s) 
```

해당 Image를 이용하여 `30` 번 VM을 생성합니다.

```
[root@virt-go-server ~]# virt-go create -i u20cust -n 30 -d "custom server"
/data/virt-go/images/u20cust already exists. Skip the image creation.

■  Create root volume.
⠇ Create virt-go-u20cust-30-root.qcow2 (1.1 GB, 1390.286 MB/s) 
■  Set root volume size to 20G
/data/virt-go/volumes/virt-go-u20cust-30-root.qcow2 size is 20G
■  Create description file.
■  Create 'virt-go-u20cust-30' XML.
■  Create 'virt-go-u20cust-30' VM.
■  Start 'virt-go-u20cust-30' VM.
successfully finished
```

`30` 번 서버에 접속하여 Image에 포함된 변경사항들이 잘 적용됐는지 확인합니다.

```
root@virt-go-u20cust-30:~# cat evidence 
hello custom server

root@virt-go-u20cust-30:~# curl localhost
custom server
```
