### Check VM List

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images
---------------------------


-------------------------------------
 NUMBER NAME IP SIZE DESCRIPTION
-------------------------------------
-------------------------------------
```

### Create New Image

Download cloud-img before creating an image. In the example, we will use ubuntu-20.04 Image.

```
[root@virt-go-server ~]# wget https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img
```

Create a `virt-go` Image using the downloaded Image.

```
[root@virt-go-server ~]# virt-go create -i u20
Create image only.
!! u20 doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /root/focal-server-cloudimg-amd64.img

■ Create image.
⠸ Create u20 (523 MB, 1339.060 MB/s)
```

Check that the created Image is displayed in the list.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
---------------------------


-------------------------------------
 NUMBER NAME IP SIZE DESCRIPTION
-------------------------------------
-------------------------------------
```

### (Option) Create a new Image [None Image]

In some cases, you need an image without a base-image i.e. no OS installed.

First create an empty image qcow2.

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

After that, create a virt-go image using the corresponding qcow2.

```
[root@virt-go-server ~]# virt-go create -i none
Create image only.
!! none doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : none-base.qcow2

■ Create image.
```

Check the created none image.

```
[root@virt-go-server ~] virt-go list
---------------------------------------------
 RESOURCE STATE
---------------------------------------------
 Data-Dir /data/virt-go
 virt-go-net 10.62.62.xxx
 Images c83 none r79 u20 u20cust
 ```


### Create VM with Image

Create a VM using the `u20` Image created above.

virt-go identifies VMs using the last digit D-class number of the IP address. That is, if VM 30 is created, the last digit of the VM's IP address will be 30.

For this reason, the number of VMs that can be used and created is 1 to 254.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 30 -d "test server"
/etc/virt-go/images/u20 already exists. Skip the image creation.

■ Create root volume.
⠸ Create virt-go-u20-30-root.qcow2 (491 MB, 1257.487 MB/s)
■ Set root volume size to 20G
/etc/virt-go/volumes/virt-go-u20-30-root.qcow2 size is 20G
■ Create description file.
■ Create 'virt-go-u20-30' XML.
■ Create 'virt-go-u20-30' VM.
■ Start 'virt-go-u20-30' VM.
successfully finished
```

Check the created VM.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
---------------------------


------------------------------------------------------------ -------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ -------
 30 virt-go-u20-30 10.62.62.30 20 GB test server
------------------------------------------------------------ -------
```

### Create Image and VM at the same time

You can also create a new Image and VM at the same time.

Using the previously downloaded ubuntu-20.04 Image, we will create a `u2004` Image and create a VM `40` with that image.

```
[root@virt-go-server ~]# virt-go create -i u2004 -n 40 -d "test 40 server"
!! u2004 doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /root/focal-server-cloudimg-amd64.img

■ Create image.
⠸ Create u2004 (504 MB, 1289.958 MB/s)
■ Create root volume.
⠸ Create virt-go-u2004-40-root.qcow2 (522 MB, 1337.285 MB/s)
■ Set root volume size to 20G
/etc/virt-go/volumes/virt-go-u2004-40-root.qcow2 size is 20G
■ Create description file.
■ Create 'virt-go-u2004-40' XML.
■ Create 'virt-go-u2004-40' VM.
■ Start 'virt-go-u2004-40' VM.
successfully finished
```

Check the created VM.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20 u2004
---------------------------


------------------------------------------------------------ ------------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ ------------
 30 virt-go-u20-30 10.62.62.30 20 GB test server
 40 virt-go-u2004-40 10.62.62.40 20 GB test 40 server
------------------------------------------------------------ ------------
```


### Create a VM with Data Disk outside the OS area

If you execute the command as below, 3 separate disks of 10G are added to the server.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 50 --disks "vda=10 vdb=10 vdc=10"
```

If you execute the command as below, 1 10G and 2 20G separate disks are added to the server.

```
[root@virt-go-server ~]# virt-go create -i u20 -n 50 --disks "vda=10 vdb=20 vdc=20"
```


### Additional Options

We have developed various options so that you can modify the Spec when creating a VM, so please check the `help` command.

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