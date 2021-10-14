[root@cloud-test-5 ~]# git clone https://github.com/YoungjuWang/virt-go.git
Cloning into 'virt-go'...
remote: Enumerating objects: 403, done.
remote: Counting objects: 100% (120/120), done.
remote: Compressing objects: 100% (81/81), done.
remote: Total 403 (delta 51), reused 95 (delta 30), pack-reused 283
Receiving objects: 100% (403/403), 69.55 MiB | 11.17 MiB/s, done.
Resolving deltas: 100% (224/224), done.

install golang >= 1.17
https://golang.org/doc/install

[root@cloud-test-5 ~]# cd virt-go/

[root@cloud-test-5 virt-go]# go mod tidy
go: downloading github.com/inconshreveable/mousetrap v1.0.0
go: downloading github.com/stretchr/testify v1.7.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
...

[root@cloud-test-5 virt-go]# go run main.go 
2021/10/14 10:08:31 open /etc/virt-go/virt-go.cfg: no such file or directory
exit status 1


