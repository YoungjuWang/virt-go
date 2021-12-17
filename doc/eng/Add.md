### Add virt-go disk

You can use the add disk command to create and add disks.

```
[root@virt-go-server ~]# add virt-go disk -n 200 --disks "vda=10 vdb=30 vdc=20"
■ Create the disk '/data/virt-go/volumes/virt-go-u20-200-vda.img'.
■ Attach the disk '/data/virt-go/volumes/virt-go-u20-200-vda.img'.
■ Create the disk '/data/virt-go/volumes/virt-go-u20-200-vdb.img'.
■ Attach the '/data/virt-go/volumes/virt-go-u20-200-vdb.img' disk.
■ Create the disk '/data/virt-go/volumes/virt-go-u20-200-vdc.img'.
■ Attach the disk '/data/virt-go/volumes/virt-go-u20-200-vdc.img'.
completed successfully
```

### Add Virt-go interface

You can add network interfaces that you can add.

If the interface to be added is a bridge, enter it in the format `bridge=<bridge-name>`, if it is a virtual network, enter it in the format `network=<network-name>`.

```
[root@virt-go-server ~]# virt-go add net -n 200 --nets "bridge=public-br network=virt-go-net"
■ Connect the 'public-br' network interface.
■ Connect the 'virt-go-net' network interface.
completed successfully
```