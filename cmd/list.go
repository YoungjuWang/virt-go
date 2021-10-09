package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

// vmCmd represents the vm command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// Network.
		nets, err := conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_PERSISTENT)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s \t\t %s\n", "Network", "Active")
		fmt.Printf("=================================\n")
		for _, net := range nets {
			netName, _ := net.GetName()
			if !(strings.Contains(netName, "virt-go")) {
				continue
			}
			netStat, _ := net.IsActive()
			fmt.Printf("%s \t\t %t\n", netName, netStat)
		}
		fmt.Printf("\n")

		// Images.
		images, err := os.ReadDir(Datadir + "/images")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", "Images : ")

		for _, image := range images {
			fmt.Print(image.Name(), " / ")
		}
		fmt.Printf("\n\n")

		// dom list
		dil := []domInfo{}
		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
		if err != nil {
			log.Fatal(err)
		}

		// Set Domain Information Slice
		for _, dom := range doms {
			domName, _ := dom.GetName()
			if !(strings.Contains(domName, "virt-go")) {
				continue
			}
			domStat, _ := dom.IsActive()
			splitName := strings.Split(domName, "-")
			tail := splitName[len(splitName)-1]
			domAddr := NetAddr + "." + tail
			domAddrToIP := net.ParseIP(domAddr)
			blkInfo, _ := dom.GetBlockInfo("sda", 0)
			blkSize := strconv.FormatUint(blkInfo.Capacity/1024/1024/1024, 10) + " GB"
			descriptionB, _ := ioutil.ReadFile(Datadir + "/volumes/" + domName + "desc")
			description := string(descriptionB)

			var colorStat string
			if domStat {
				domS := strconv.FormatBool(domStat)
				domS = "Active"
				colorStat = green(domS)
			} else {
				domS := strconv.FormatBool(domStat)
				domS = "Inactive"
				colorStat = red(domS)
			}
			di := domInfo{tail, domName, colorStat, domAddrToIP, blkSize, description}
			dil = append(dil, di)
		}

		// Sort via addr
		sort.Slice(dil, func(i int, j int) bool {
			return bytes.Compare(dil[i].addr, dil[j].addr) < 0
		})

		// Formatting table
		data := [][]string{}
		for _, d := range dil {
			sda := d.addr.String()
			data = append(data, []string{d.num, d.name, d.state, sda, d.size, d.desc})
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Num", "DomainName", "Status", "IP-Address", "Root-Size", "Description"})
		table.AppendBulk(data)
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
