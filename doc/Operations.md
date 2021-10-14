### VM을 중지합니다.

시작 중인 VM을 종료합니다.

VM의 실행 여부는 VM Name의 색으로 구분할 수 있습니다.

- 녹색 > 실행 중
- 적색 > 중지

```
[root@virt-go-server ~]# virt-go stop -n 20
Stop virt-go-u20-20
```

### VM을 시작합니다.

```
[root@virt-go-server ~]# virt-go start -n 20
Start virt-go-u20-20
```

### VM을 재시작합니다.

```
[root@virt-go-server ~]# virt-go restart -n 20
Stop virt-go-u20-20
Start virt-go-u20-20
```
