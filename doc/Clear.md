### virt-go 삭제

virt-go 삭제 시 `virt-go-net` 과 `/etc/virt-go/` 에 있는 모든 파일이 삭제됩니다.

만약 VM의 data direcctory를 `/etc/virt-go` 로 설정한 경우 VM Disk가 없어 Error가 발생할 수 있으니 유의하여 진행하시길 바랍니다.

```
[root@virt-go-server ~]# virt-go clear -y
■  Shutdown 'virt-go-net'.
■  Undefine 'virt-go-net'.
```

`virt-go` command 삭제

```
[root@virt-go-server ~]# rm /usr/local/bin/virt-go 
rm: remove regular file '/usr/local/bin/virt-go'? y
```

복구를 희망하시는 경우 [Installation](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Installation.md) 을 다시 진행하시면 됩니다.

**단 data directory를 `/etc/virt-go` 로 설정하신 경우 VM은 복구되지 않습니다.**
