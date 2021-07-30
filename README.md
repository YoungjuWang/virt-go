
# Table of Contents

---

- [Intro](#intro)
- [Installation](#installation)
    + [1. Requirements](#1-requirements)
    + [2. Download 'virt-go' command](#2-download--virt-go--command)
    + [3. Install 'virt-go' command](#3-install--virt-go--command)
    + [[Option] Build 'virt-go' command from source](#-option--build--virt-go--command-from-source)
    + [4. Initialize 'virt-go' environment](#4-initialize--virt-go--environment)
- ['virt-go' usage](#-virt-go--usage)
    + [Help Command](#help-command)
    + [Create Image for VM](#create-image-for-vm)
    + [Create VM](#create-vm)
      - [1) Create VM using Image already exists](#1--create-vm-using-image-already-exists)
      - [2) Create VM using image doesn't exsit](#2--create-vm-using-image-doesn-t-exsit)
      - [3) Create VM select number of cores / memory size](#3--create-vm-select-number-of-cores---memory-size)
      - [4) Create VM with description](#4--create-vm-with-description)
      - [5) Create VM with custom cloud-init file](#5--create-vm-with-custom-cloud-init-file)
    + [Check 'virt-go' resources status](#check--virt-go--resources-status)
    + [Delete VM](#delete-vm)
    + [Delete VM Image](#delete-vm-image)
    + [VM Operations](#vm-operations)
      - [Connect to VM (ssh)](#connect-to-vm--ssh-)
      - [Connect to VM console](#connect-to-vm-console)
      - [Update 'virt-go' version](#update--virt-go--version)
      - [Resize Root Volume](#resize-root-volume)
      - [Start/Stop VM](#start-stop-vm)
- [Clear Virt-go](#clear-virt-go)

</br>

# Intro

'virt-go' 는 libvirt 환경에서 VM을 편리하게 관리하기 위한 프로그램 입니다.

cloud image와 cloud-init을 이용해서 빠르게 VM을 생성하는 것이 목적입니다.

cloud image는 아래 제 블로그 글에 있는 링크를 참고하여 준비하시기 바랍니다.
https://yjwang.tistory.com/147

</br>

# Installation

### 1. Requirements

---

```bash
- 'libvirtd' service
- 'genisoimage' command
- 'qemu-img' commnad
- 'libvirt-lib' package
```

### 2. Download 'virt-go' command

---

```bash
# wget https://github.com/YoungjuWang/virt-go/raw/master/virt-go/virt-go
```

### 3. Install 'virt-go' command

---

```bash
# chmod +x virt-go
# mv -f virt-go /usr/local/bin/
```

### [Option] Build 'virt-go' command from source

---

`golang 16.x`  이상 버전에서 진행하시기 바랍니다. (`libvirt-dev` 패키지가 추가로 필요합니다.)

```bash
# cd "$(go env GOROOT)/src"
# git clone https://github.com/YoungjuWang/virt-go.git
# cd virt-go
# go build -o /usr/local/bin/virt-go
```

### 4. Initialize 'virt-go' environment

---

help 페이지 확인

```bash
virt-go init -h
Init 'virt-go' Virtual Network and Data Directory.

Usage:
  virt-go init [flags]

Flags:
  -c, --cidr string      Network Address (required) (default "10.62.62.0")
  -d, --datadir string   Which is VM's volume, cloudinit and images stored in (required) (default "/etc/virt-go")
  -h, --help             help for init
```

VM에서 사용할 Network Address와 VM Data들이 저장될 Directory로 지정합니다.

`root` 권한이 있는 계정으로 실행해야 `directory`를 자동으로 생성합니다.

```bash
# virt-go init -c 192.168.123.0 -d /data/virt-go
```

위 command가 종료되면 아래 경로에 `virt-go` 설정 파일이 생성됩니다.

```bash
# cat /etc/virt-go/virt-go.cfg
Datadir=/data/virt-go
NetAddr=192.168.123
```

또한 지정한 Directory에 아래와 같은 구조로 구성됩니다.

```bash
# tree /data/virt-go/
/data/virt-go/
├── cloudinit
│   ├── meta-data
│   └── user-data
├── images
└── volumes
```

`<datadir>/cloudinit/user-data` file을 열어 `<pub_key>` 부분을 별도 " ' 문자 없이 update합니다.

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

</br>

# 'virt-go' usage

### Help Command

---

모든 Sub Command에서 help 페이지를 제공합니다. 필요시 확인 후 사용하시기 바랍니다.

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

### Create Image for VM

---

Image 다운로드 (Ubuntu 20.04 이미지를 사용하겠습니다.)

```bash
# wget https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img
```

Image 생성

`u20` 이름으로 Image를 생성합니다.

```bash
# virt-go create -i u20
Don't Create 'VM' Only Create Image
'u20' does not exist. 'virt-go' attempd to create image via 'base' image file.
 Enter base image full path : /home/yjwang/Downloads/focal-server-cloudimg-amd64.img
- Generate Image 'u20' from '/home/yjwang/Downloads/focal-server-cloudimg-amd64.img'
535.19 MiB / 535.19 MiB [-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 2.38 GiB p/s 400ms
```

생성된 Image 확인

```bash
# virt-go list
!!! This list only contain about 'virt-go'

Network 		 Active
=================================
virt-go-net 		 true

Images : u20 /

+-----+-----------------+----------+--------------+-----------+------------------------+
| NUM |   DOMAINNAME    |  STATUS  |  IP-ADDRESS  | ROOT-SIZE |      DESCRIPTION       |
+-----+-----------------+----------+--------------+-----------+------------------------+
+-----+-----------------+----------+--------------+-----------+------------------------+
```

### Create VM

---

#### 1) Create VM using Image already exists

```bash
# virt-go create -i u20 -n 200
- Generate Domain Root Image '/etc/virt-go/volumes/virt-go-u20-200root' from 'u20'
535.19 MiB / 535.19 MiB [-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 5.44 GiB p/s 300ms

"virt-go-u20-200" is created!
```

#### 2) Create VM using image doesn't exsit

```bash
# virt-go create -i u2004 -n 204
'u2004' is not exist. 'virt-go' attempd to create image via 'base' image file.
 Enter base image full path : /home/yjwang/Downloads/focal-server-cloudimg-amd64.img
- Generate Image 'u2004' from '/home/yjwang/Downloads/focal-server-cloudimg-amd64.img'
535.19 MiB / 535.19 MiB [------------------------------] 100.00% 4.43 GiB p/s 300ms

- Generate Domain Root Image '/etc/virt-go/volumes/virt-go-u2004-204root' from 'u2004'
535.19 MiB / 535.19 MiB [------------------------------] 100.00% 5.55 GiB p/s 300ms

"virt-go-u2004-204" is created!

```

#### 3) Create VM select number of cores / memory size

지정하지 않으면 2core 4GB mem을 기본으로 생성됩니다.

4core 8GB mem을 가지는 VM을 생성합니다.

```bash
# virt-go create -i u20 -n 210 -c 4 -m 8
- Generate Domain Root Image '/etc/virt-go/volumes/virt-go-u20-210root' from 'u20'
535.19 MiB / 535.19 MiB [------------------------------] 100.00% 3.90 GiB p/s 300ms

"virt-go-u20-210" is created!
```

#### 4) Create VM with description

`test ubuntu`라는 설명을 붙여 VM을 생성합니다.

```bash
# virt-go create -i u20 -n 60 -d "test ubuntu"
- Generate Domain Root Image '/etc/virt-go/volumes/virt-go-u20-60root' from 'u20'
535.19 MiB / 535.19 MiB [-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 5.43 GiB p/s 300ms

"virt-go-u20-60" is created!
```

`list` command에서 확인할 수 있습니다.

```bash
# virt-go list
!!! This list only contain about 'virt-go'

Network 		 Active
=================================
virt-go-net 		 true

Images : u20 / u2004 /

+-----+-------------------+----------+--------------+-----------+------------------------+
| NUM |    DOMAINNAME     |  STATUS  |  IP-ADDRESS  | ROOT-SIZE |      DESCRIPTION       |
+-----+-------------------+----------+--------------+-----------+------------------------+
|  60 | virt-go-u20-60    | Active   | 10.62.62.60  | 20 GB     | test ubuntu            |
+-----+-------------------+----------+--------------+-----------+------------------------+
```

#### 5) Create VM with custom cloud-init file

`u` option과 `t` option을 사용하여 user-data와 meta-data를 지정합니다.

지정하지 않을 시 `init` 에서 지정한 Directory에 있는 cloud-init을 사용합니다.

```bash
# virt-go create -i u20 -n 30 -u "/etc/virt-go/cloudinit/user-data" -t "/etc/virt-go/cloudinit/meta-data"
- Generate Domain Root Image '/etc/virt-go/volumes/virt-go-u20-30root' from 'u20'
535.19 MiB / 535.19 MiB [-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 5.29 GiB p/s 300ms

"virt-go-u20-30" is created!
```

</br>

### Check 'virt-go' resources status

---

vm-network / image / vm 정보 확인

```
# virt-go list
!!! This list only contain about 'virt-go'

Network          Active
=================================
virt-go-net          true

Images : u20 / u2004 /

+-----+-------------------+----------+--------------+-----------+------------------------+
| NUM |    DOMAINNAME     |  STATUS  |  IP-ADDRESS  | ROOT-SIZE |      DESCRIPTION       |
+-----+-------------------+----------+--------------+-----------+------------------------+
|  60 | virt-go-u20-60    | Active   | 10.62.62.60  | 20 GB     | test ubuntu            |
+-----+-------------------+----------+--------------+-----------+------------------------+
```

</br>

### Delete VM

---

```bash
# virt-go delete -n 200
virt-go-u20-200 shutdown!
System Already Shutdown
virt-go-u20-200 will be deleted!
delete Finished
```

</br>

### Delete VM Image

---

```
# virt-go delete -i u20
```

</br>

### VM Operations

---

#### Connect to VM (ssh)

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

####  Connect to VM console

virsh 명령을 사용하여 VM이름으로 접속합니다.

```
# virsh console virt-go-u20-63
도메인 virt-go-u20-63에 연결되었습니다
```

####  Update 'virt-go' version

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

</br>

# Clear Virt-go

VM은 사전에 삭제 해두시길 바랍니다.

삭제를 진행하지 않는 경우 Data directory는 삭제되지만 VM은 수동으로 undefine 해주어야 합니다.

```bash
# virt-go clear -h
Delete virt-network and dir about virt-go. before run this command please delete VM first.

Usage:
  virt-go clear [flags]

Flags:
  -h, --help   help for clear
```
