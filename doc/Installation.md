### git에서 virt-go project를 내려받습니다.

```
[root@cloud-test-5 ~]# git clone https://github.com/YoungjuWang/virt-go.git
Cloning into 'virt-go'...
remote: Enumerating objects: 403, done.
remote: Counting objects: 100% (120/120), done.
remote: Compressing objects: 100% (81/81), done.
remote: Total 403 (delta 51), reused 95 (delta 30), pack-reused 283
Receiving objects: 100% (403/403), 69.55 MiB | 11.17 MiB/s, done.
Resolving deltas: 100% (224/224), done.
```

### 서버에 go를 설치합니다.

install golang >= 1.17
https://golang.org/doc/install

### 내려 받은 virt-go project directory로 이동합니다.

```
[root@cloud-test-5 ~]# cd virt-go/
```

### virt-go에 필요한 module 들을 내려받습니다.

```
[root@cloud-test-5 virt-go]# go mod tidy
go: downloading github.com/inconshreveable/mousetrap v1.0.0
go: downloading github.com/stretchr/testify v1.7.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
...
```

### virt-go 가 설치될 준비가 됐는지 확인합니다.

아래 메시지를 확인합니다. 동일한 메시지가 나오면 설치될 준비가 끝났음을 알 수 있으며 에러 메시지는 무시하시기 바랍니다.

```
[root@cloud-test-5 virt-go]# go run main.go 
2021/10/14 10:08:31 open /etc/virt-go/virt-go.cfg: no such file or directory
exit status 1
```

### virt-go 에 필요한 환경을 구성합니다.

`install.sh` 파일을 실행하여 `virt-go` 에 필요한 폴더 및 파일과 `virt-go.cfg` 파일을 생성합니다.

```
[root@cloud-test-5 virt-go]# bash sh/install.sh 
Input virt-go data directory
ex) /etc/virt-go : /etc/virt-go

Input virt-go network address
ex) 10.62.62.0 : 10.62.62.0

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   545  100   545    0     0   1730      0 --:--:-- --:--:-- --:--:--  1730
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
Please change value below in 'user-data' file for login to virt-go server without password.
ssh_authorized_keys:
	>>	  - <pub-key>
```

### 배포된 환경을 확인합니다.

```
[root@cloud-test-5 virt-go]# ls -l /etc/virt-go
total 4
drwxr-xr-x. 2 root root 40 Oct 14 10:14 cloud-init
drwxr-xr-x. 2 root root  6 Oct 14 10:14 images
-rw-r--r--. 1 root root 40 Oct 14 10:17 virt-go.cfg
drwxr-xr-x. 2 root root  6 Oct 14 10:14 volumes

[root@cloud-test-5 virt-go]# cat /etc/virt-go/virt-go.cfg 
dataDir=/etc/virt-go
netAddr=10.62.62.0
```

### public-key를 배포합니다.

key 인증을 통해 virt-go vm에 password 없이 접근하기 위하여 ssh public key를 준비합니다.

```
[root@cloud-test-5 virt-go]# cat ~/.ssh/id_rsa.pub 
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCoMfjixSZyW5g3Z6EomG3jAsoaJlJfGBYSCC5z96YZZqVTcv2SggJnnLCSqVM00/jpLHs5cHbR74jkJaLDBT3TDKVN/ovqXk4V0eoewaDiQ0p1cmCLsmeVGt8lg2kR1PCjqMLtYFPjU9j+DfD1vIYxqzo1uQJIuOZ/1g2IdkRA63lDIZKhXr4Pr1oIhIgnvTM2Ep4imYCYKZ+kxpfF+inCoG8kjUGDg/+kUWEFXvgOF1/IAi/kvIMICLeM5wYnU68AjUi0SMgtaQN1tjIluo/S3/1OqTpWmY0jnVZ+shFTZIhgVmT9fHdJHnaCOAdTwl9SQrrALlYG8DoT+ZVVFIEeAUJH6(생략)
```

### cloud-init user-data 내용을 수정합니다. 

```
[root@cloud-test-5 virt-go]# sed 's,- <pub-key>,- '"$(cat ~/.ssh/id_rsa.pub)"',g' /etc/virt-go/cloud-init/user-data 
```


### virt-go command를 생성합니다.

```
[root@cloud-test-5 virt-go]# go build -o /usr/local/bin/virt-go
```

### virt-go command를 실행합니다.

아래 help를 진행했을 때 에러 없이 실행 돼야 정상입니다.

```
[root@cloud-test-5 virt-go]# virt-go --help

'virt-go' help user using libvirt to manage virtual resources easy and fast.
Futher informations about 'virt-go' are in https://github.com/YoungjuWang/virt-go.
Have a good day. :)

Usage:
  virt-go [command]

Available Commands:
  clear       Delete virt-go
  completion  generate the autocompletion script for the specified shell
  create      Create VM or Image
  delete      Delete VM or Image
  help        Help about any command
  init        Create virt-go network
  list        Lists of virt-go resources
  restart     Restart virt-go VM
  ssh         Connect to virt-go VM via ssh
  start       Start virt-go VM
  stop        Stop virt-go VM

Flags:
  -h, --help   help for virt-go

Use "virt-go [command] --help" for more information about a command.
```

### virt-go network를 생성합니다.

```
[root@cloud-test-5 virt-go]# virt-go init
■  Create network XML file.
■  Set 'virt-go-net' to autostart.
■  Start 'virt-go-net'.
```

### 생성된 Network를 확인합니다.

`virt-go-net` 의 STATE IP주소가 녹색으로 표시돼야 정상입니다.

```
[root@cloud-test-5 virt-go]# virt-go list
---------------------------
 RESOURCE     STATE        
---------------------------
 Data-Dir     /etc/virt-go 
 virt-go-net  10.62.62.xxx 
 Images                    
---------------------------


-------------------------------------
 NUMBER  NAME  IP  SIZE  DESCRIPTION 
-------------------------------------
-------------------------------------
```
