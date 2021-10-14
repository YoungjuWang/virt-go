package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"

	libvirt "libvirt.org/libvirt-go"
)

type vmConfig struct {
	name     string
	num      uint8
	image    string
	desc     string
	cpu      uint
	mem      uint
	size     uint
	userData string
	metaData string
}

func (v *vmConfig) checkNumber() bool {
	if v.num == 0 {
		return true
	} else if v.num < 2 || v.num > 254 {
		log.Fatal("Invalid number. 1 < num < 255")
		return false
	}
	return false
}

func (v *vmConfig) checkImageName() {
	if strings.Contains(v.image, "-") {
		log.Fatal("Invalid image name. It can not contain '-'")
	}
}

func (v *vmConfig) getRunningVMName(conn *libvirt.Connect) string {
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
	if err != nil {
		log.Fatal(err)
	}
	for _, dom := range doms {
		domName, err := dom.GetName()
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(domName, "virt-go") {
			if strings.HasSuffix(domName, strconv.Itoa(int(v.num))) {
				return domName
			}
		}
	}
	log.Fatal("There is no server has " + strconv.Itoa(int(v.num)))
	return ""
}

type globalConfig struct {
	dataDir string
	netAddr string
}

func (g *globalConfig) getCfg() {
	cfgTextB, err := os.ReadFile("/etc/virt-go/virt-go.cfg")
	if err != nil {
		log.Fatal(err)
	}
	splitCfgText := strings.Split(string(cfgTextB), "\n")
	datadir := strings.TrimLeft(splitCfgText[0], "dataDir=")
	netaddr := strings.TrimLeft(splitCfgText[1], "netAddr=")

	g.dataDir = strings.TrimRight(datadir, "/")
	netaddrS := strings.Split(netaddr, ".")
	g.netAddr = strings.Join(netaddrS[:3], ".")
}

var v vmConfig
var g globalConfig
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorReset string = "\033[0m"

func init() {
	g.getCfg()
}
