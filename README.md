
### Table of Contents

---

- [소개](#소개)
- [설치](#설치)
- [사용법](#사용법)
- [운영](#운영)

### 소개

---

'virt-go' 는 libvirt 환경에서 VM을 편리하게 관리하기 위한 프로그램 입니다.



### 설치

---


Prerequisite

```bash
- 'libvirtd' service
- 'genisoimage' command
- 'qemu-img' commnad
- 'libvirt-lib' package
```


Download Manager Command

```bash
# wget https://github.com/YoungjuWang/virt-go/raw/master/virt-go/virt-go
```


Install Manager Command

```bash
# chmod +x virt-go
# mv -f virt-go /usr/local/bin/
```


(Option) Build virt-go from source


golang 16.x 이상 version에서 진행하시기 바랍니다.
```
# cd "$(go env GOROOT)/src"
# git clone https://github.com/YoungjuWang/virt-go.git
# cd virt-go
# go build -o /usr/local/bin/virt-go
```


### 사용법

---



#### 0. 모든 command에는 Help Page를 제공합니다.

```bash
# virt-go --help
virt-go is inspired by 'fast-vm' by Ondrej
Use it for your test machine in libvirt!

Usage:
  virt-go [command]

Available Commands:
  clear       Delete virt-network and dir about virt-go. before run this command please delete VM first.
  completion  Generate completion script
  create      Create VM (Virtual Machine)
  delete      Delete Selected VM
  help        Help about any command
  init        Init 'virt-go' environment
  list        list
  resize      Resize VM root volum. If VM is started, It will be shutdown automatically
  start       Start virt-go VM
  stop        Stop 'virt-go' VM

Flags:
  -h, --help   help for virt-go

Use "virt-go [command] --help" for more information about a command.
```



#### 1. init command를 사용하여 시스템을 초기화합니다.


VM에서 사용할 Network Address와 VM Data들이 저장될 Directory로 지정합니다.

```bash
# virt-go init -c 192.168.123.0 -d /data/virt-go
```


위 command가 종료되면 아래 경로에 `virt-go` 설정 파일이 생성되며

```bash
# cat /etc/virt-go/virt-go.cfg
Datadir=/data/virt-go
NetAddr=192.168.123
```


지정한 Directory에 아래와 같이 구조가 생성됩니다.

```bash
# tree /data/virt-go/
/data/virt-go/
├── cloudinit
│   ├── meta-data
│   └── user-data
├── images
└── volumes
```


이후 `user-data`에서 key-file을 update합니다.

```bash
#cloud-config
users:
  - name: root
    ssh_authorized_keys:
      - <pub_key>

user: cloud-user
chpasswd:
  list: |
    root:testtest
    cloud-user:testtest
  expire: False
ssh_pwauth: True

growpart:
  mode: auto
  devices: ["/"]
  ignore_growroot_disabled: false

runcmd:
  - sed -i '/PermitRootLogin prohibit-password/a\PermitRootLogin yes' /etc/ssh/sshd_config
  - sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
  - reboot
```



#### 2. create command를 이용하여 VM을 생성합니다.


별도로 지정하지 않는이상`4GB Mem` `2 CPU` `/data/virt-go/cloudinit/user-data` `/data/virt-go/cloudinit/meta-data` 파일을 사용하여 VM이 생성됩니다.

VM 이름은 반드시 숫자 `2 ~ 254` 범위 내에서 지정해야 하며 해당 번호를 가진 `MAC`과 `IP`를 가지게 됩니다.


Create Command는 많은 Option이 있으므로 help를 사전에 살펴보신 후 사용하시길 권장드립니다.


2-1. 사전에 없는 Image로 VM을 생성하는 경우.
VM을 생성할 때 없는 Image 이름을 지정하면 자동으로 Image를 생성하고자 합니다.
Image의 `base` 가 될 파일을 선택하면 Image 및 VM을 동시에 생성할 수 있습니다.

```
# virt-go create -i u20 -n 62
'u20' is not exist. 'virt-go' attempd to create image via 'base' image file. 
 Enter base image full path : /usr/vm-template/focal-server-cloudimg-amd64.qcow2
- Generate Image 'u20' from '/usr/vm-template/focal-server-cloudimg-amd64.qcow2' 
1.30 GiB / 1.30 GiB [-------------------------------------------------------------------------------------------] 100.00% 922.03 MiB p/s 1.6s

- Generate Domain Root Image '/data/virt-go/volumes/virt-go-u20-62root' from 'u20' 
1.30 GiB / 1.30 GiB [-------------------------------------------------------------------------------------------] 100.00% 944.25 MiB p/s 1.6s

"virt-go-u20-62" is created! 
```



2-2 사전에 있는 Image로 VM을 생성하는 경우

```
# virt-go create -i u20 -n 63
- Generate Domain Root Image '/data/virt-go/volumes/virt-go-u20-63root' from 'u20' 
1.30 GiB / 1.30 GiB [-------------------------------------------------------------------------------------------] 100.00% 624.37 MiB p/s 2.3s

"virt-go-u20-63" is created! 
```



#### 3.생성한 Resource 정보 확인.


```
# virt-go list
!!! This list only contain about 'virt-go' 

Network 		 Active
=================================
virt-go-net 		 true

Images : u18 / u20 / 

+-----+-----------------+--------+--------------+-----------+-------------+
| NUM |   DOMAINNAME    | STATUS |  IP-ADDRESS  | ROOT-SIZE | DESCRIPTION |
+-----+-----------------+--------+--------------+-----------+-------------+
| 100 | virt-go-u20-100 | Active | 10.62.62.100 | 20 GB     |             |
| 101 | virt-go-u18-101 | Active | 10.62.62.101 | 20 GB     |             |
| 150 | virt-go-u20-150 | Active | 10.62.62.150 | 20 GB     | test server |
+-----+-----------------+--------+--------------+-----------+-------------+
```



#### 4. Resource 삭제.


Image 삭제

```
# virt-go delete -i u20
```


VM 삭제

```
# virt-go delete -n 62
virt-go-u20-62 shutdown!
virt-go-u20-62 will be deleted!
delete Finished
```


### 운영

---



#### 서버 접속

- Default Account ID / PW : cloud-user / testtest
- Default Admin ID / PW : root / testtest

`init` 과정에서 ssh-key를 수정했다면 `virt-go list` 로 확인되는 서버 IP로 ssh 접속하면 바로 접속이 돼야합니다.
만약 수정하지 않았다면 재 수정 후 VM을 재 생성해보시기 바랍니다.

```
# ssh 192.168.123.63
The authenticity of host '192.168.123.63 (192.168.123.63)' can't be established.
ECDSA key fingerprint is SHA256:Z/sptVgUPGNaJXWgp4I4sGtChwg3FM4DAQLRCBXDb0Y.
Are you sure you want to continue connecting (yes/no/[fingerprint])?  yes
```



####  console 접속


virsh 명령을 빌려 VM이름으로 접속하면 됩니다.

```
# virsh console virt-go-u20-63
도메인 virt-go-u20-63에 연결되었습니다
```



####  virt-go 업데이트


별도의 Migration 및 중지 없이 Update 가능합니다.

```
# rm virt-go
# wget https://github.com/YoungjuWang/virt-go/raw/master/virt-go/virt-go
# chmod +x virt-go
# mv -f virt-go /usr/local/bin/
```



#### Resize Root Volume


만약 VM이 실행 중이라면 동의를 구하고 자동으로 Shutdown 후 Start 됩니다.
Shrink는 지원하지 않고 확장만 가능합니다. Shrink는 `qemu-img` command를 사용 하시기 바랍니다.

```
# virt-go resize -n 30 -s 50
```



####  Start/Stop VM


Stop VM

```
# virt-go stop -n 90
virt-go-u20-90  is Stopped !
```


Start VM

```
# virt-go start -n 90
virt-go-u20-90  is Started !
```



#### Virt-go 삭제

virt-go 삭제


> VM은 사전에 삭제 해두시길 바랍니다. 삭제를 진행하지 않는 경우 Datadir은 삭제되지만 VM은 수동으로 undefine 해주어야 합니다.

```
# virt-go clear -h
Delete virt-network and dir about virt-go. before run this command please delete VM first.

Usage:
  virt-go clear [flags]

Flags:
  -h, --help   help for clear
```
