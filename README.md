[![Go Report Card](https://goreportcard.com/badge/github.com/YoungjuWang/virt-go)](https://goreportcard.com/report/github.com/YoungjuWang/virt-go)
[![GitHub license](https://img.shields.io/github/license/YoungjuWang/virt-go?style=flat-square)](https://github.com/YoungjuWang/virt-go/blob/v2/LICENSE)

[ENG](https://github.com/YoungjuWang/virt-go/blob/v2/doc/eng/README.md)

# virt-go

<p align="center">
  <img src="https://github.com/YoungjuWang/virt-go/blob/v2/img/virt-go2.png">
</p>

'virt-go' 는 libvirt 환경에서 VM을 편리하게 관리하기 위한 프로그램으로 cloud image와 cloud-init을 이용해서 빠르게 VM 관리할 수 있도록 돕습니다.

cloud image는 아래 제 블로그 글에 있는 링크를 참고하여 준비하시기 바랍니다. 

https://yjwang.tistory.com/147


## Requirements

아래 항목들이 사전에 설치 및 실행되고 있어야합니다.

- 'libvirtd' service
- 'genisoimage' command
- 'qemu-img' commnad
- 'libvirt-lib' package 

## Index

아래 가이드를 참고하시어 사용하시기 바랍니다.

- [Installation - virt-go 환경 구성 및 Build](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Installation.md)
- [Createation - VM, Image 생성](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Creation.md)
- [Deletion - VM, Image 삭제](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Deletion.md)
- [Connection - VM에 ssh로 접속](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Connection.md)
- [Operations - VM 실행, 중지, 재시작, size 변경](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Operations.md)
- [Clear - virt-go 환경 삭제](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Clear.md)
- [Customization - Custom Image 생성, cloud-init data 변경](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Customization.md)
- [Update - Update virt-go command](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Update.md)
