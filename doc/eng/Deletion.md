### Delete Image

Delete the created Image.

I will delete the `u2004` Image shown in the list.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20 u2004
---------------------------
...
```

Delete in progress

```
[root@virt-go-server ~]# virt-go delete -i u2004
■ Delete u2004.
successfully finished
```

Confirm that it has been deleted from the list.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
---------------------------
...
```

### Delete VM

Let's delete VM '30'.

```
[root@virt-go-server ~]# virt-go delete -n 30
■ Shutdown virt-go-u20-30.
■ Undefine virt-go-u20-30.
■ Delete virt-go-u20-30 volume.
■ Delete virt-go-u20-30 cloud-init iso file.
■ Delete virt-go-u20-30 description file.
successfully finished
```

Confirm that it has been deleted from the list.

```
[root@virt-go-server ~]# virt-go list
---------------------------
 RESOURCE STATE
---------------------------
 Data-Dir /etc/virt-go
 virt-go-net 10.62.62.xxx
 Images u20
---------------------------


------------------------------------------------------------ ------------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ ------------
 40 virt-go-u2004-40 10.62.62.40 20 GB test 40 server
------------------------------------------------------------ ------------
```

### Delete Image and VM at the same time

I will delete the `u20` Image and the `40` VM at the same time.

```
[root@virt-go-server ~]# virt-go delete -n 40 -i u20
■ Shutdown virt-go-u2004-40.
■ Undefine virt-go-u2004-40.
■ Delete virt-go-u2004-40 volume.
■ Delete virt-go-u2004-40 cloud-init iso file.
■ Delete virt-go-u2004-40 description file.
■ Delete u20.
successfully finished
```

Confirm that it has been deleted from the list.

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