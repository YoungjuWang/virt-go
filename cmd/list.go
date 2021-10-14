package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

type domInfo struct {
	num  uint64
	name string
	addr string
	size string
	desc string
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists of virt-go resources",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// Get slice struct
		domInfoList := getLists(conn)

		// Sort slice
		sort.Slice(domInfoList, func(i int, j int) bool {
			return domInfoList[i].num < domInfoList[j].num
		})

		printOtherTable(conn)
		fmt.Println("\n")
		printDomTable(domInfoList)
	},
}

func getLists(conn *libvirt.Connect) []domInfo {
	domInfoList := []domInfo{}

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
			domStat, err := dom.IsActive()
			if err != nil {
				log.Fatal(err)
			}
			splitName := strings.Split(domName, "-")
			tail := splitName[len(splitName)-1]
			domAddr := g.netAddr + "." + tail
			blkInfo, err := dom.GetBlockInfo("sda", 0)
			if err != nil {
				log.Fatal(err)
			}
			blkSize := strconv.FormatUint(blkInfo.Capacity/1024/1024/1024, 10) + " GB"
			descriptionB, err := ioutil.ReadFile(g.dataDir + "/volumes/" + domName)
			if err != nil {
				log.Print(err)
			}
			description := strings.TrimSuffix(string(descriptionB), "\n")
			tailU, err := strconv.ParseUint(tail, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			if domStat {
				domName = colorGreen + domName + colorReset
			} else {
				domName = colorRed + domName + colorReset
			}
			domInfos := domInfo{tailU, domName, domAddr, blkSize, description}
			domInfoList = append(domInfoList, domInfos)
		}
	}
	return domInfoList
}

func printDomTable(domInfoList []domInfo) {
	var data [][]string
	for i := 0; i < len(domInfoList); i++ {
		data = append(data, []string{strconv.Itoa(int(domInfoList[i].num)), domInfoList[i].name, domInfoList[i].addr, domInfoList[i].size, domInfoList[i].desc})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Number", "Name", "IP", "Size", "Description"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	//table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	//table.SetHeaderLine(false)
	table.AppendBulk(data)
	table.Render()
}

func printOtherTable(conn *libvirt.Connect) {
	net, err := conn.LookupNetworkByName("virt-go-net")
	if err != nil {
		log.Fatal(err)
	}
	netName, err := net.GetName()
	if err != nil {
		log.Fatal(err)
	}
	netStat, err := net.IsActive()
	if err != nil {
		log.Fatal(err)
	}
	netAddr := g.netAddr + ".xxx"
	if netStat {
		netAddr = colorGreen + netAddr + colorReset
	} else {
		netAddr = colorRed + netAddr + colorReset
	}
	imagesB, err := os.ReadDir(g.dataDir + "/images")
	if err != nil {
		log.Fatal(err)
	}
	var imagesS string
	for _, image := range imagesB {
		imagesS += image.Name() + "  "
	}

	data := [][]string{
		{"Data-Dir", g.dataDir},
		{netName, netAddr},
		{"Images", imagesS},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Resource", "State"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	//table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	//table.SetHeaderLine(false)
	table.AppendBulk(data)
	table.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
