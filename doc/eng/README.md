[![Go Report Card](https://goreportcard.com/badge/github.com/YoungjuWang/virt-go)](https://goreportcard.com/report/github.com/YoungjuWang/virt-go)
[![GitHub license](https://img.shields.io/github/license/YoungjuWang/virt-go?style=flat-square)](https://github.com/YoungjuWang/virt-go/blob/v2/LICENSE)

[KOR](https://github.com/YoungjuWang/virt-go/blob/v2/README.md)

# virt-go

<p align="center">
  <img src="https://github.com/YoungjuWang/virt-go/blob/v2/img/virt-go2.png">
</p>

'virt-go' is a command-line tool to conveniently manage VMs in the libvirt environment, and helps you quickly manage VMs using cloud image and cloud-init.

If you have never used cloud image, please prepare it by referring to the link in my blog post below.

https://yjwang.tistory.com/147


## Requirements

The items below must be installed and running beforehand.

- 'libvirtd' service
- 'genisoimage' command
- 'qemu-img' commnad
- 'libvirt-lib' package 

## Index

Please refer to the guide below for use.

- [Installation - Configure and Build the virt-go Environment](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Installation.md)
- [Createation - Create VM, Image](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Creation.md)
- [Deletion - Delete VM, Image](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Deletion.md)
- [Connection - Connect to VM ssh](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Connection.md)
- [Operations - Start, Stop, Restart, Resize](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Operations.md)
- [Clear - Clear virt-go environment](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Clear.md)
- [Customization - Create Custom Image, Edit cloud-init configurations, Edit VM description](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Customization.md)
- [Update - Update virt-go command](https://github.com/YoungjuWang/virt-go/blob/v2/doc/Update.md)
