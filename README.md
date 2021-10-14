# virt-go

![](https://github.com/YoungjuWang/virt-go/blob/v2/img/virt-go.png)

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
- [Operations - VM 실행, 중지, 재시작](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Operations.md)
- [Clear - virt-go 환경 삭제](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Clear.md)

**!! Clear시 VM에서 사용한 `Volume` 과 `VM` 은 수동으로 삭제해야합니다.**

## Release note

```
- 소스를 객체지향 형식으로 정리했습니다.
- virt-go logo가 추가됐습니다.
- list command의 output을 변경했습니다.
- disk 및 interface를 hot-plug 지원되도록 수정했습니다.
- restart command를 추가했습니다.
- cloud-init 초기 data파일을 init이 아닌 git에서 내려받도록 수정했습니다.
- build된 실행파일을 다운 받는 것이 아닌 사용자가 직접 build하도록 가이드를 수정했습니다.
```
