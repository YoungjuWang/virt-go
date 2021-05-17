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

	"github.com/spf13/cobra"
	"libvirt.org/libvirt-go"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Selected VM",
	Run: func(cmd *cobra.Command, args []string) {
		Datadir, NetAddr = GetCFG()

		if image != "" {
			// Delete Image
			err := os.Remove(Datadir + "/images/" + image)
			os.Exit(0)
			if err != nil {
				fmt.Println("image can't be deleted : ", err)
			}
		} else if image == "" && Num == 0 {
			cmd.Help()
			os.Exit(33)
		}

		// Get libvirt connection
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
		if err != nil {
			fmt.Println(err)
		}
		for _, dom := range doms {
			domName, _ := dom.GetName()
			if (strings.Contains(domName, "virt-go")) && strings.Contains(domName, strconv.Itoa(Num)) {
				fmt.Println(domName, "shutdown!")
				err = dom.Destroy()
				if err != nil {
					fmt.Println("System Already Shutdown")
				}

				fmt.Println(domName, "will be deleted!")
				err = dom.Undefine()
				if err != nil {
					panic(err)
				}

				err = os.Remove(Datadir + "/cloudinit/" + domName + "cloud-init.iso")
				if err != nil {
					panic(err)
				}
				err = os.Remove(Datadir + "/volumes/" + domName + "root")
				if err != nil {
					panic(err)
				}
			}
		}
		// destroy
		// undefine
		// delete Iso File
		// delete Data File
		fmt.Println("delete Finished")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&Num, "number", "n", 0, "Number of VM for identification (required)")
	deleteCmd.Flags().StringVarP(&image, "image", "i", "", "Image that VM will use (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
