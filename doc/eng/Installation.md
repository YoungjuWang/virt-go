### Download the virt-go project from git.

```
[root@virt-go-server ~]# git clone https://github.com/YoungjuWang/virt-go.git
Cloning into 'virt-go'...
remote: Enumerating objects: 403, done.
remote: Counting objects: 100% (120/120), done.
remote: Compressing objects: 100% (81/81), done.
remote: Total 403 (delta 51), reused 95 (delta 30), pack-reused 283
Receiving objects: 100% (403/403), 69.55 MiB | 11.17 MiB/s, done.
Resolving deltas: 100% (224/224), done.
```

### Install go on the server.

install golang >= 1.17
https://golang.org/doc/install

Check the installed go version.

```
[root@virt-go-server ~]# go version
go version go1.17.2 linux/amd64
```

### Move to the downloaded virt-go project directory.

```
[root@virt-go-server ~]# cd virt-go/
```

### Download the modules required for virt-go.

```
[root@virt-go-server virt-go]# go mod tidy
go: downloading github.com/inconshreveable/mousetrap v1.0.0
go: downloading github.com/stretchr/testify v1.7.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
...
```

### Make sure virt-go is ready to be installed.

Check the message below. If you see the same message, you know that the installation is ready, please ignore the error message.

```
[root@virt-go-server virt-go]# go run main.go
2021/10/14 10:08:31 open /etc/virt-go/virt-go.cfg: no such file or directory
exit status 1
```

### Configure the required environment for virt-go .

Execute the `install.sh` file to create the necessary folders and files for `virt-go` and the `virt-go.cfg` file.

```
[root@virt-go-server virt-go]# bash sh/install.sh
Input virt-go data directory
ex) /data/virt-go : /data/virt-go

Input virt-go network address
ex) 10.62.62.0 : 10.62.62.0

  % Total % Received % Xferd Average Speed ​​Time Time Time Current
                                 Dload Upload Total Spent Left Speed
100 545 100 545 0 0 1730 0 --:--:-- --:--:-- --:--:-- 1730
  % Total % Received % Xferd Average Speed ​​Time Time Time Current
                                 Dload Upload Total Spent Left Speed
  0 0 0 0 0 0 0 0 --:--:-- --:--:-- --:--:-- 0
Please change value below in 'user-data' file for login to virt-go server without password.
ssh_authorized_keys:
>> - <pub-key>
```

### Check the deployed environment.

```
[root@virt-go-server ~]# ls -l /etc/virt-go/
total 4
-rw-r--r--. 1 root root 41 Oct 14 14:29 virt-go.cfg

[root@virt-go-server ~]# ls -l /data/virt-go/
total 0
drwxr-xr-x. 2 root root 40 Oct 14 14:29 cloud-init
drwxr-xr-x. 2 root root 6 Oct 14 14:29 images
drwxr-xr-x. 2 root root 6 Oct 14 14:29 volumes

[root@virt-go-server virt-go]# cat /etc/virt-go/virt-go.cfg
dataDir=/etc/virt-go
netAddr=10.62.62.0
```

### Deploy the public-key.

Prepare ssh public key to access virt-go vm without password through key authentication.

```
[root@virt-go-server virt-go]# cat ~/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCoMfjixSZyW5g3Z6EomG3jAsoaJlJfGBYSCC5z96YZZqVTcv2SggJnnLCSqVM00 / jpLHs5cHbR74jkJaLDBT3TDKVN / ovqXk4V0eoewaDiQ0p1cmCLsmeVGt8lg2kR1PCjqMLtYFPjU9j + DfD1vIYxqzo1uQJIuOZ / 1g2IdkRA63lDIZKhXr4Pr1oIhIgnvTM2Ep4imYCYKZ + kxpfF + inCoG8kjUGDg / + kUWEFXvgOF1 / IAi / kvIMICLeM5wYnU68AjUi0SMgtaQN1tjIluo / S3 / 1OqTpWmY0jnVZ + shFTZIhgVmT9fHdJHnaCOAdTwl9SQrrALlYG8DoT + ZVVFIEeAUJH6 (omitted)
```

### Edit cloud-init user-data contents.

```
[root@virt-go-server virt-go]# sed -i 's,- <pub-key>,- '"$(cat ~/.ssh/id_rsa.pub)"',g' /data/virt-go/cloud-init/user-data
```


### Generate a virt-go command.

```
[root@virt-go-server virt-go]# go build -o /usr/local/bin/virt-go
```

### Execute the virt-go command.

It is normal when the help below is executed without errors.

```
[root@virt-go-server virt-go]# virt-go --help

'virt-go' help user using libvirt to manage virtual resources easy and fast.
Futher informations about 'virt-go' are in https://github.com/YoungjuWang/virt-go.
Have a good day. :)

Usage:
  virt-go [command]

Available Commands:
  clear Delete virt-go
  completion generate the autocompletion script for the specified shell
  create Create VM or Image
  delete Delete VM or Image
  help Help about any command
  init Create virt-go network
  list Lists of virt-go resources
  restart Restart virt-go VM
  ssh Connect to virt-go VM via ssh
  start Start virt-go VM
  stop Stop virt-go VM

Flags:
  -h, --help help for virt-go

Use "virt-go [command] --help" for more information about a command.
```

### Create a virt-go network.

```
[root@virt-go-server virt-go]# virt-go init
■ Create network XML file.
■ Set 'virt-go-net' to autostart.
■ Start 'virt-go-net'.
```

### Check the created Network.

It is normal when the STATE IP address of `virt-go-net` is displayed in green.

```
[root@virt-go-server virt-go]# virt-go list
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
-------------------
```

### Set autocompletion of virt-go command.

```
[root@virt-go-server ~]# virt-go completion bash > /etc/bash_completion.d/virt-go
[root@virt-go-server ~]# . /etc/bash_completion.d/virt-go 
```