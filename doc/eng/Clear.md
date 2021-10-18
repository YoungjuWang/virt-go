### Delete virt-go 

Deleting virt-go will delete all files in `virt-go-net` and `/etc/virt-go/`.

If the data directory of the VM is set to `/etc/virt-go`, an error may occur because there is no VM disk, so please proceed with caution.

```
[root@virt-go-server ~]# virt-go clear -y
■  Shutdown 'virt-go-net'.
■  Undefine 'virt-go-net'.
```

Delete `virt-go` command.

```
[root@virt-go-server ~]# rm /usr/local/bin/virt-go 
rm: remove regular file '/usr/local/bin/virt-go'? y
```

If you wish to restore, proceed with [Installation](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Installation.md) again.

**However, if the data directory is set to `/etc/virt-go`, the VM will not be restored.**
