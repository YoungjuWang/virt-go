### Stop the VM.

Shut down any VMs that are starting.

Whether a VM is running can be distinguished by the color of the VM Name.

- Green > Running
- Red > Stop

```
[root@virt-go-server ~]# virt-go stop -n 20
Stop virt-go-u20-20
```

### Start the VM.

```
[root@virt-go-server ~]# virt-go start -n 20
Start virt-go-u20-20
```

### Restart the VM.

```
[root@virt-go-server ~]# virt-go restart -n 20
Stop virt-go-u20-20
Start virt-go-u20-20
```

### Change the VM Size. (Increase Size)

To change the VM Size, you must first stop the VM.

```
[root@virt-go-server ~]# virt-go stop -n 20
Stop virt-go-u20-20
```

Check the size of the current '20' VM.

```
[root@virt-go-server ~]# virt-go list
-----------------------------
 RESOURCE STATE
-----------------------------
 Data-Dir /data/virt-go
 virt-go-net 10.62.62.xxx
 Images u20 u20cust
-----------------------------


------------------------------------------------------------ -------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ -------
 20 virt-go-u20-20 10.62.62.20 20 GB resize
------------------------------------------------------------ -------
```

We will change the size of VM ‘20’ from ‘20GB’ to ‘50GB’.

```
[root@virt-go-server ~]# virt-go resize -n 20 -s 50
resize /data/virt-go/volumes/virt-go-u20-20-root.qcow2
Image resized.
/data/virt-go/volumes/virt-go-u20-20-root.qcow2 size is 50G
```

Check the changed size.

```
[root@virt-go-server ~]# virt-go list
-----------------------------
 RESOURCE STATE
-----------------------------
 Data-Dir /data/virt-go
 virt-go-net 10.62.62.xxx
 Images u20 u20cust
-----------------------------


------------------------------------------------------------ -------
 NUMBER NAME IP SIZE DESCRIPTION
------------------------------------------------------------ -------
 20 virt-go-u20-20 10.62.62.20 50 GB resize
------------------------------------------------------------ -------
```

Run the VM to see if it applies.

```
[root@virt-go-server ~]# virt-go start -n 20
Start virt-go-u20-20
```


You can see that the capacity has increased to '50GB'.

```
[root@virt-go-server ~]# virt-go ssh -n 20

root@virt-go-u20-20:~# df -h /
Filesystem Size Used Avail Use% Mounted on
/dev/sda1 49G 1.3G 47G 3% /
```

### Change the VM Size. (Reduce Size)

The process is the same as expanding Size .

Stop VM > Change Size > Start VM

> When reducing the size, there is a possibility of data loss, so please work carefully.
> The VM may not boot.

Add `--shrink` option to change Size from `50GB` to `30GB`.

```
[root@virt-go-server ~]# virt-go resize -n 20 -s 30 --shrink
resize /data/virt-go/volumes/virt-go-u20-20-root.qcow2
Image resized.
/data/virt-go/volumes/virt-go-u20-20-root.qcow2 size is 30G
```