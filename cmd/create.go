package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	progressbar "github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
	libvirtxml "libvirt.org/libvirt-go-xml"
)

func createVM() {
	v.checkImageName()
	checked := v.checkNumber()
	if checked {
		fmt.Println("Create image only.")
		imageFile := g.dataDir + "/images/" + v.image
		if _, err := os.Stat(imageFile); os.IsNotExist(err) {
			createImage(imageFile)
		} else {
			log.Fatal(imageFile + " already exists.")
		}
		os.Exit(0)
	}

	checkVMExists()
	v.name = "virt-go-" + v.image + "-" + strconv.Itoa(int(v.num))
	volFile := createVolume() // =createImage() also.
	cloudInitIso := createCloudInitISO()
	macAddr := func() string {
		var addr string
		addrTail := v.num
		switch {
		case addrTail >= 2 && addrTail < 10:
			addr = "02:00:AA:AA:AA:0" + strconv.Itoa(int(addrTail))
		case addrTail >= 10 && addrTail < 100:
			addr = "02:00:AA:AA:AA:" + strconv.Itoa(int(addrTail))
		case addrTail >= 100 && addrTail < 200:
			addrTail = addrTail - 100
			addr = "02:00:AA:AA:AB:" + strconv.Itoa(int(addrTail))
		case addrTail >= 200 && addrTail < 255:
			addrTail = addrTail - 200
			addr = "02:00:AA:AA:AC:" + strconv.Itoa(int(addrTail))
		}
		return addr
	}

	// Gen Domain XML.
	var diskControllerAddr uint = 0
	domCfg := &libvirtxml.Domain{
		Type:        "kvm",
		Name:        v.name,
		Description: "This Virtual Machine is created by 'virt-go'",
		Memory:      &libvirtxml.DomainMemory{Value: uint(v.mem), Unit: "GiB", DumpCore: "on"},
		VCPU:        &libvirtxml.DomainVCPU{Value: uint(v.cpu)},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{Arch: "x86_64", Type: "hvm"},
		},
		// For Hot-Plug Feature.
		Features: &libvirtxml.DomainFeatureList{
			ACPI:   &libvirtxml.DomainFeature{},
			APIC:   &libvirtxml.DomainFeatureAPIC{},
			VMPort: &libvirtxml.DomainFeatureState{State: "off"},
		},
		CPU:      &libvirtxml.DomainCPU{Mode: "host-model"},
		OnReboot: "restart",
		Devices: &libvirtxml.DomainDeviceList{
			// Disks.
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{Type: "qcow2"},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: volFile}},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda", Bus: "sata"},
					Boot: &libvirtxml.DomainDeviceBoot{Order: 1},
					Address: &libvirtxml.DomainAddress{
						Drive: &libvirtxml.DomainAddressDrive{
							Controller: &diskControllerAddr, Bus: &diskControllerAddr, Target: &diskControllerAddr, Unit: &diskControllerAddr}},
				},
				{
					Device: "cdrom",
					Driver: &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "raw"},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: cloudInitIso}},
					Target:   &libvirtxml.DomainDiskTarget{Dev: "hda", Bus: "ide"},
					ReadOnly: &libvirtxml.DomainDiskReadOnly{},
					Address: &libvirtxml.DomainAddress{
						Drive: &libvirtxml.DomainAddressDrive{
							Controller: &diskControllerAddr, Bus: &diskControllerAddr, Target: &diskControllerAddr, Unit: &diskControllerAddr}},
				},
			},
			// Network Interfaces.
			Interfaces: []libvirtxml.DomainInterface{
				{
					MAC:    &libvirtxml.DomainInterfaceMAC{Address: macAddr()},
					Source: &libvirtxml.DomainInterfaceSource{Network: &libvirtxml.DomainInterfaceSourceNetwork{Network: "virt-go-net"}},
					Model:  &libvirtxml.DomainInterfaceModel{Type: "virtio"},
				},
			},
			// Serial Console Devices.
			Consoles: []libvirtxml.DomainConsole{
				{
					Target: &libvirtxml.DomainConsoleTarget{Type: "serial"},
				},
			},
		},
	}
	fmt.Printf("■  Create '%s' XML.\n", v.name)
	domXML, err := domCfg.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	// Define Domain via XML created before.
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("■  Create '%s' VM.\n", v.name)
	dom, err := conn.DomainDefineXML(domXML)
	if err != nil {
		log.Fatal(err)
	}

	// Start Domain.
	fmt.Printf("■  Start '%s' VM.\n", v.name)
	err = dom.Create()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(colorGreen + "successfully finished" + colorReset)
}

func createImage(imageFile string) {
	matched, err := regexp.MatchString("^[0-9]", v.image)
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		log.Fatal("Number cannot be at the beginning of the Image Name.")
	}

	fmt.Println("!! " + v.image + " doesn't exist. Create image first.\n")
	// Get basFile
	var baseFile string
	fmt.Print("Please input base-image file full path.\nex) /base/image/path/file : ")
	fmt.Scan(&baseFile)
	if _, err := os.Stat(baseFile); os.IsNotExist(err) {
		log.Fatal(baseFile + " doesn't exist.")
	}

	// Copy baseFile to imageFile
	in, err := os.Open(baseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Println("\n■  Create image.")
	bar := progressbar.DefaultBytes(-1, "Create "+v.image)
	_, err = io.Copy(io.MultiWriter(out, bar), in)
	if err != nil {
		log.Fatal(err)
	}
}

func createVolume() string {
	// Copy image to root volume.
	volFile := g.dataDir + "/volumes/" + v.name + "-root.qcow2"
	if _, err := os.Stat(volFile); err == nil {
		log.Fatal(volFile + " already exists.")
	}

	imageFile := g.dataDir + "/images/" + v.image
	if _, err := os.Stat(imageFile); os.IsNotExist(err) {
		createImage(imageFile)
	} else {
		fmt.Println(imageFile + " already exists. Skip the image creation.")
	}
	in, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(volFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Println("\n■  Create root volume.")
	bar := progressbar.DefaultBytes(-1, "Create "+v.name+"-root.qcow2")
	_, err = io.Copy(io.MultiWriter(out, bar), in)
	if err != nil {
		log.Fatal(err)
	}

	// Resize root volume.
	fmt.Println("\n■  Set root volume size to " + strconv.Itoa(int(v.size)) + "G")
	cmd := exec.Command("qemu-img", "resize", volFile, strconv.Itoa(int(v.size))+"G")
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(volFile + " size is " + strconv.Itoa(int(v.size)) + "G")

	// Create Description file.
	fmt.Println("■  Create description file.")
	descFile := g.dataDir + "/volumes/" + v.name
	descOut, err := os.Create(descFile)
	if err != nil {
		log.Fatal(err)
	}
	defer descOut.Close()
	descOut.WriteString(v.desc)

	return volFile
}

func createCloudInitISO() string {
	cloudInitIso := g.dataDir + "/volumes/" + v.name + "-init.iso"

	if _, err := os.Stat("/usr/bin/genisoimage"); os.IsNotExist(err) {
		log.Fatal("'genisoimage' command doesn't exist.Please install the command before.")
	}

	// genisoimage -output cloudInitIso.iso -volid cidata -joliet -rock user-data meta-data
	cmd := exec.Command("genisoimage", "-output", cloudInitIso, "-volid", "cidata", "-joliet", "-rock", v.userData, v.metaData)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return cloudInitIso
}

func checkVMExists() {
	files, err := ioutil.ReadDir(g.dataDir + "/volumes")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		isExist := strings.Contains(fileName, "-"+strconv.Itoa(int(v.num))+"-")
		if isExist {
			log.Fatal(strconv.Itoa(int(v.num)) + " is already used. Please use another number.")
		}
	}
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Disable sorted flags
	createCmd.Flags().SortFlags = false

	// Flags
	createCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "Number, VM will use.")
	createCmd.Flags().StringVarP(&v.image, "image", "i", "", "Image, VM will use. (required)")
	createCmd.MarkFlagRequired("image")
	createCmd.Flags().StringVarP(&v.desc, "desc", "d", "", "Description")
	createCmd.Flags().UintVarP(&v.cpu, "cpu", "c", 2, "number of core")
	createCmd.Flags().UintVarP(&v.mem, "mem", "m", 4, "size of memory (GB)")
	createCmd.Flags().UintVarP(&v.size, "size", "s", 20, "VM's Root Volume Size (GB).")
	createCmd.Flags().StringVar(&v.bridge, "bridge", "", "Bridge Interface")
	createCmd.Flags().StringVar(&v.userData, "user-data", g.dataDir+"/cloudinit/user-data", "cloud-init user-data")
	createCmd.Flags().StringVar(&v.metaData, "meta-data", g.dataDir+"/cloudinit/meta-data", "cloud-init meta-data")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create VM or Image",
	Run: func(cmd *cobra.Command, args []string) {
		createVM()
	},
}
