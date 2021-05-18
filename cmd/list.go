/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
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

		Datadir, NetAddr = GetCFG()

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		fmt.Printf("!!! This list only contain about 'virt-go' \n\n")

		// net list
		nets, err := conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_PERSISTENT)
		if err != nil {
			fmt.Println(err)
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

		// image list
		images, err := os.ReadDir(Datadir + "/images")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", "Images : ")
		//fmt.Printf("=================\n")

		for _, image := range images {
			fmt.Print(image.Name(), " / ")
		}
		fmt.Printf("\n\n")

		// dom list
		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
		if err != nil {
			fmt.Println(err)
		}

		data := [][]string{}
		for _, dom := range doms {
			domName, _ := dom.GetName()
			if !(strings.Contains(domName, "virt-go")) {
				continue
			}
			domStat, _ := dom.IsActive()
			splitName := strings.Split(domName, "-")
			tail := splitName[len(splitName)-1]

			data = append(data, []string{domName, strconv.FormatBool(domStat), NetAddr + "." + tail})
			//fmt.Printf("%s\t%t\t\t%s\n", domName, domStat, NetAddr+"."+tail)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "IsActive", "IP"})
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
