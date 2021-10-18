### Change Description

The description of each VM is stored as a file with the name of the VM in the volumes of the data dir.

If you change the contents of the file, you can change the Description.

check data dir path

```
[root@virt-go-server ~]# virt-go list
----------------------------
 RESOURCE STATE
----------------------------
 Data-Dir /data/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
----------------------------


------------------------------------------------------------ -------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ -------
 20 virt-go-u20-20 10.62.62.20 20 GB test server
------------------------------------------------------------ -------
```

I checked the `Data-Dir` path above.

Now, let's change the Description of VM #20.
If you change the text of that file, it will change.

```
[root@virt-go-server ~]# cat /data/virt-go/volumes/virt-go-u20-20
test server 2222
```

Confirm that it has been changed.

```
[root@virt-go-server ~]# virt-go list
----------------------------
 RESOURCE STATE
----------------------------
 Data-Dir /data/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
----------------------------


------------------------------------------------------------ ------------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ ------------
 20 virt-go-u20-20 10.62.62.20 20 GB test server 2222
------------------------------------------------------------ ------------
```


### change cloud-init data

The `user-data` file is located in the cloud-init path under the data dir.

```
[root@virt-go-server ~]# ls -l /data/virt-go/cloud-init/user-data
-rw-r--r--. 1 root root 1084 Oct 14 14:32 /data/virt-go/cloud-init/user-data
```

By changing the contents of the file, you can change the contents of cloud-init config used when creating a VM.


### Create Custom Image

Create the root volume of the existing server as an Image.

Create a file on the server `20` and run the web server.

```
root@virt-go-u20-20:~# echo "hello custom server" > /root/evidence

root@virt-go-u20-20:~# apt -y install apache2

root@virt-go-u20-20:~# echo "custom server" > /var/www/html/index.html

root@virt-go-u20-20:~# systemctl enable apache2
```

After that, when the server is converted to Image, so that cloud-init can be run again.
Initialize cloud-init.

```
root@virt-go-u20-20:~# cloud-init clean

root@virt-go-u20-20:~# cloud-init status
status: not run
```

Stop the VM to convert it to an Image.

```
[root@virt-go-server ~]# virt-go stop -n 20
```

Create the root volume of the VM as `u20cust` Image.

```
[root@virt-go-server ~]# ls -l /data/virt-go/volumes/virt-go-u20-20-root.qcow2

[root@virt-go-server ~]# virt-go create -i u20cust
Create image only.
!! u20cust doesn't exist. Create image first.

Please input base-image file full path.
ex) /base/image/path/file : /data/virt-go/volumes/virt-go-u20-20-root.qcow2

■ Create image.
⠇ Create u20cust (1.1 GB, 1341.759 MB/s)
```

Create VM ‘30’ using the image.

```
[root@virt-go-server ~]# virt-go create -i u20cust -n 30 -d "custom server"
/data/virt-go/images/u20cust already exists. Skip the image creation.

■ Create root volume.
⠇ Create virt-go-u20cust-30-root.qcow2 (1.1 GB, 1390.286 MB/s)
■ Set root volume size to 20G
/data/virt-go/volumes/virt-go-u20cust-30-root.qcow2 size is 20G
■ Create description file.
■ Create 'virt-go-u20cust-30' XML.
■ Create 'virt-go-u20cust-30' VM.
■ Start 'virt-go-u20cust-30' VM.
successfully finished
```

Connect to the server ‘30’ and check whether the changes included in the Image are applied properly.

```
root@virt-go-u20cust-30:~# cat evidence
hello custom server

root@virt-go-u20cust-30:~# curl localhost
custom server
```