### Check virt-go version

```
[root@virt-go-server virt-go]# virt-go version
latest version : 2.0.02
installed version : 2.0.01

  You need to update virt-go
  ```

  If you see the message `You need to update virt-go` in the above command result, proceed as follows.


### virt-go pull latest git file

```
[root@virt-go-server virt-go]# git pull
remote: Enumrating objects: 79, done.
remote: Counting objects: 100% (79/79), done.
remote: Compressing objects: 100% (42/42), done.
remote: Total 66 (delta 41), reused 48 (delta 23), pack-reused 0
Unpacking objects: 100% (66/66), done.
From https://github.com/YoungjuWang/virt-go
   b2cac75..255dc16 v2 -> origin/v2
Updating b2cac75..255dc16
Fast-forward
 LICENSE | 202 +++++++++++++++++++++++++++++++++++++++++++++++++++ +++++++++++++++++++++++++++++++++++++++
 README.md | 22 ++++++----
 cloud-init/user-data | 2 +-
 cmd/list.go | 6 +--
 cmd/operations.go | 46 +++++++++++++++++++++
 cmd/root.go | 2 +-
 doc/Customization.md | 143 +++++++++++++++++++++++++++++++++++++++++++++++++++ ++++++++++
 doc/Installation.md | 24 +++++++---
 img/virt-go2.png | Bin 0 -> 61440 bytes
 sh/install.sh | 2 +-
 10 files changed, 428 insertions(+), 21 deletions(-)
 create mode 100644 LICENSE
 create mode 100644 doc/Customization.md
 create mode 100644 img/virt-go2.png
```

### virt-go rebuild

You only need to rebuild the command regardless of the VM.

```
[root@virt-go-server virt-go]# go build -o /usr/local/bin/virt-go
```

### Check for Update

```
[root@virt-go-server virt-go]# virt-go version
latest version : 2.0.02
installed version : 2.0.02

  You have latest virt-go
```