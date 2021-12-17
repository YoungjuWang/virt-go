### virt-go Disk 추가

add disk command를 이용해 별도의 Disk를 생성 및 추가할 수 있습니다.

```
[root@virt-go-server ~]# virt-go add disk -n 200 --disks "vda=10 vdb=30 vdc=20"
■  Create '/data/virt-go/volumes/virt-go-u20-200-vda.img' disk.
■  Attach '/data/virt-go/volumes/virt-go-u20-200-vda.img' disk.
■  Create '/data/virt-go/volumes/virt-go-u20-200-vdb.img' disk.
■  Attach '/data/virt-go/volumes/virt-go-u20-200-vdb.img' disk.
■  Create '/data/virt-go/volumes/virt-go-u20-200-vdc.img' disk.
■  Attach '/data/virt-go/volumes/virt-go-u20-200-vdc.img' disk.
successfully finished
```

### Virt-go Interface 추가

add net command를 이용해 별도의 Network Interface를 추가할 수 있습니다.

추가하는 Interface가 bridge인 경우 `bridge=<bridge-name>` 형식으로 입력하고 virtual network인 경우 `network=<network-name>` 형식으로 입력합니다.

```
[root@virt-go-server ~]# virt-go add net -n 200 --nets "bridge=public-br network=virt-go-net"
■  Attach 'public-br' network interface.
■  Attach 'virt-go-net' network interface.
successfully finished
```