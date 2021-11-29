### VM List 확인

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

### 새로운 Image 생성

Image를 생성하기 전에 cloud-img를 다운받습니다. 예시에서는 ubuntu-20.04 Image를 사용하겠습니다.

```
[root@virt-go-server ~]# wget https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img
```

다운 받은 Image를 사용하여 `virt-go` Image를 생성합니다.

```
[root@virt-go-server ~]# virt-go create -i u20
Create image only.
!! u20 doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /root/focal-server-cloudimg-amd64.img

■  Create image.
⠸ Create u20 (523 MB, 1339.060 MB/s) 
```

생성된 Image가 list에 표기되는지 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20          
---------------------------


-------------------------------------
 NUMBER  NAME  IP  SIZE  DESCRIPTION 
-------------------------------------
-------------------------------------
```

### (Option) 새로운 Image 생성 [None Image]

경우에 따라 base-image 가 없는 즉 OS가 설치되지 않은 Image가 필요한 경우가 있습니다.

우선 빈 이미지 qcow2를 생성합니다.

```
[root@virt-go-server ~]# qemu-img create -f qcow2 none-base.qcow2 5G

[root@virt-go-server ~]# qemu-img info none-base.qcow2
image: none-base.qcow2
file format: qcow2
virtual size: 5.0G (5368709120 bytes)
disk size: 196K
cluster_size: 65536
Format specific information:
    compat: 1.1
    lazy refcounts: false
    refcount bits: 16
    corrupt: false
```

이후 해당 qcow2를 활용하여 virt-go image를 생성합니다.

```
[root@virt-go-server ~]# virt-go create -i none
Create image only.
!! none doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : none-base.qcow2

■  Create image.
```

생성된 none image를 확인합니다.

```
[root@virt-go-server ~] virt-go list
---------------------------------------------
 RESOURCE     STATE                          
---------------------------------------------
 Data-Dir     /data/virt-go                  
 virt-go-net  10.62.62.xxx                   
 Images       c83  none  r79  u20  u20cust   
 ```

### Image로 VM 생성

위에서 생성한 `u20` Image를 활용하여 VM을 생성합니다.

virt-go는 IP주소의 가장 마지막 자리 D-class Number를 이용하여 VM을 구분합니다. 즉 30번 VM을 생성하면 해당 VM의 IP주소 가장 마지막 자리 숫자는 30이 됩니다.

이러한 이유로 사용 및 생성 가능한 VM은 1~254 까지입니다.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 30 -d "test server"
/etc/virt-go/images/u20 already exists. Skip the image creation.

■  Create root volume.
⠸ Create virt-go-u20-30-root.qcow2 (491 MB, 1257.487 MB/s) 
■  Set root volume size to 20G
/etc/virt-go/volumes/virt-go-u20-30-root.qcow2 size is 20G
■  Create description file.
■  Create 'virt-go-u20-30' XML.
■  Create 'virt-go-u20-30' VM.
■  Start 'virt-go-u20-30' VM.
successfully finished
```

생성된 VM을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20          
---------------------------


---------------------------------------------------------
 NUMBER  NAME            IP           SIZE   DESCRIPTION 
---------------------------------------------------------
 30      virt-go-u20-30  10.62.62.30  20 GB  test server 
---------------------------------------------------------
```

### Image와 VM 동시에 생성

새로운 Image와 VM을 동시에 생성할 수도 있습니다.

이전에 다운받은 ubuntu-20.04 Image를 사용하여 `u2004` Image를 생성하고 해당 Image로 `40`번 VM을 생성해보겠습니다.

```
[root@virt-go-server ~]# virt-go create -i u2004 -n 40 -d "test 40 server"
!! u2004 doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /root/focal-server-cloudimg-amd64.img

■  Create image.
⠸ Create u2004 (504 MB, 1289.958 MB/s) 
■  Create root volume.
⠸ Create virt-go-u2004-40-root.qcow2 (522 MB, 1337.285 MB/s) 
■  Set root volume size to 20G
/etc/virt-go/volumes/virt-go-u2004-40-root.qcow2 size is 20G
■  Create description file.
■  Create 'virt-go-u2004-40' XML.
■  Create 'virt-go-u2004-40' VM.
■  Start 'virt-go-u2004-40' VM.
successfully finished
```

생성된 VM을 확인합니다.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images       u20  u2004   
---------------------------


--------------------------------------------------------------
 NUMBER  NAME              IP           SIZE   DESCRIPTION    
--------------------------------------------------------------
 30      virt-go-u20-30    10.62.62.30  20 GB  test server    
 40      virt-go-u2004-40  10.62.62.40  20 GB  test 40 server 
--------------------------------------------------------------
```


### OS 영역 외 Data Disk가 있는 VM 생성

아래와 같이 command를 실행하면 10G 3개의 별도 disk가 서버에 추가됩니다.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 50 --disks "10 10 10"
```

아래와 같이 command를 실행하면 10G 1개, 20G 2개의 별도 disk가 서버에 추가됩니다.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 50 --disks "10 20 20"
```


### 추가 Option

VM 생성 시 Spec을 수정할 수 있도록 다양한 Option을 개발해놓았으니 `help` command를 확인해보시기 바랍니다.

```
[root@virt-go-server ~]# virt-go create --help
Create VM or Image

Usage:
  virt-go create [flags]

Flags:
  -n, --number uint8       Number, VM will use
  -i, --image string       Image, VM will use (required)
  -d, --desc string        Description
  -c, --cpu uint           number of core (default 2)
  -m, --mem uint           size of memory (GB) (default 4)
  -s, --size uint          VM's Root Volume Size (GB) (default 20)
      --user-data string   cloud-init user-data (default "/etc/virt-go/cloud-init/user-data")
      --meta-data string   cloud-init meta-data (default "/etc/virt-go/cloud-init/meta-data")
      --disks string       additional disk list (default "none")
  -h, --help               help for create
```
