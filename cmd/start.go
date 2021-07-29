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
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"libvirt.org/libvirt-go"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start virt-go VM",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}

		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
		if err != nil {
			log.Fatal(err)
		}

		count := 0

		for _, dom := range doms {
			domName, _ := dom.GetName()
			if strings.Contains(domName, strconv.Itoa(Num)) {
				err = dom.Create()
				if err != nil {
					log.Fatal(err)
				}
				count++
				fmt.Println(domName, " is Started !")
			} else {
				continue
			}
		}

		if count == 0 {
			fmt.Println("There is no server contains ", Num, " in Name")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVarP(&Num, "number", "n", 0, "Number of VM for identification (required)")
	startCmd.MarkFlagRequired("number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
