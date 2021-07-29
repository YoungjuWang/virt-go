package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	libvirtxml "libvirt.org/libvirt-go-xml"
)

// GenISO generate iso file with cloud-init
func GenISO(Num int, image string, userData string, metaData string) string {
	Datadir, NetAddr = GetCFG()
	isoTmpName := Datadir + "/cloudinit/virt-go-" + image + "-" + strconv.Itoa(Num) + "cloud-init.iso"

	if _, err := os.Stat(isoTmpName); os.IsExist(err) {
		if err := os.Remove(isoTmpName); err != nil {
			fmt.Println("Cloudinit iso file remove error")
		}
	}

	if _, err := os.Stat("/usr/bin/genisoimage"); os.IsNotExist(err) {
		fmt.Println("'genisoimage' command is not exists. please install it before")
		panic(err)
	}

	// genisoimage -output seed.iso -volid cidata -joliet -rock user-data meta-data
	cmd := exec.Command(
		"genisoimage",
		"-output", isoTmpName,
		"-volid", "cidata",
		"-joliet",
		"-rock", userData, metaData,
	)

	if err := cmd.Run(); err != nil {
		fmt.Println(err, "??")
	}

	return isoTmpName
}

// GenISOXML generate ISO XML file for Domain
func GenISOXML(isoTmpName string) (xmldoc string, err error) {
	var drive uint
	drive = 0
	domcfg := &libvirtxml.DomainDisk{
		Device: "cdrom",
		Driver: &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "raw"},
		Source: &libvirtxml.DomainDiskSource{
			File: &libvirtxml.DomainDiskSourceFile{
				File: isoTmpName}},
		BackingStore: &libvirtxml.DomainDiskBackingStore{},
		Target:       &libvirtxml.DomainDiskTarget{Dev: "hda", Bus: "ide"},
		ReadOnly:     &libvirtxml.DomainDiskReadOnly{},
		Address: &libvirtxml.DomainAddress{
			Drive: &libvirtxml.DomainAddressDrive{
				Controller: &drive, Bus: &drive, Target: &drive, Unit: &drive}},
	}
	return domcfg.Marshal()
}

func GenDomXML(image string, Num int, domImage string, cpu int, mem int, macAddr string) (xmldoc string) {
	var drive uint
	drive = 0

	vmName := "virt-go-" + image + "-" + strconv.Itoa(Num)
	if mem < 0 || cpu < 0 {
		fmt.Println("'mem' or 'cpu' values is less than 0. Change value to positive number automatically.")
		mem = -mem
		cpu = -cpu
	}

	domcfg := &libvirtxml.Domain{
		Type:        "kvm",
		Name:        vmName,
		Description: "testMachine",
		Memory:      &libvirtxml.DomainMemory{Value: uint(mem), Unit: "GiB", DumpCore: "on"},
		VCPU:        &libvirtxml.DomainVCPU{Value: uint(cpu)},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{Arch: "x86_64", Type: "hvm"},
		},
		CPU:      &libvirtxml.DomainCPU{Mode: "host-model"},
		OnReboot: "restart",
		Devices: &libvirtxml.DomainDeviceList{
			// Disk List
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{Type: "qcow2"},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: domImage}},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda", Bus: "sata"},
					Boot: &libvirtxml.DomainDeviceBoot{Order: 1},
					Address: &libvirtxml.DomainAddress{
						Drive: &libvirtxml.DomainAddressDrive{
							Controller: &drive, Bus: &drive, Target: &drive, Unit: &drive}},
				},
				{
					Device:   "cdrom",
					Driver:   &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "raw"},
					Target:   &libvirtxml.DomainDiskTarget{Dev: "hda", Bus: "ide"},
					ReadOnly: &libvirtxml.DomainDiskReadOnly{},
					Address: &libvirtxml.DomainAddress{
						Drive: &libvirtxml.DomainAddressDrive{
							Controller: &drive, Bus: &drive, Target: &drive, Unit: &drive}},
				},
			},
			// Network Interface List
			Interfaces: []libvirtxml.DomainInterface{
				{
					MAC:    &libvirtxml.DomainInterfaceMAC{Address: macAddr},
					Source: &libvirtxml.DomainInterfaceSource{Network: &libvirtxml.DomainInterfaceSourceNetwork{Network: "virt-go-net"}},
					Model:  &libvirtxml.DomainInterfaceModel{Type: "virtio"},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				{
					Target: &libvirtxml.DomainConsoleTarget{Type: "serial"},
				},
			},
		},
	}

	xmldoc, err := domcfg.Marshal()
	if err != nil {
		fmt.Println(err)
	}
	return xmldoc
}

func GetMAC(Num int) (macAddr string) {
	switch {
	case Num < 2:
		fmt.Println("Number should be larger or equal than 2")
		os.Exit(20)
	case Num >= 2 && Num < 10:
		macAddr = "02:00:AA:AA:AA:0" + strconv.Itoa(Num)
	case Num >= 10 && Num < 100:
		macAddr = "02:00:AA:AA:AA:" + strconv.Itoa(Num)
	case Num >= 100 && Num < 200:
		Num = Num - 100
		macAddr = "02:00:AA:AA:AB:" + strconv.Itoa(Num)
	case Num >= 200 && Num < 255:
		Num = Num - 200
		macAddr = "02:00:AA:AA:AC:" + strconv.Itoa(Num)
	}
	return macAddr
}
