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
	"strings"

	"github.com/spf13/cobra"
	"libvirt.org/libvirt-go"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete virt-network and dir about virt-go. before run this command please delete VM first.",
	Run: func(cmd *cobra.Command, args []string) {

		// Get Datadir and NetAddr
		Datadir, NetAddr = GetCFG()

		// Create libvirt Connection
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		// Destroy virt-go-net and Undefine
		nets, err := conn.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_PERSISTENT)
		if err != nil {
			fmt.Println(err)
		}
		for _, net := range nets {
			netName, _ := net.GetName()
			if (strings.Contains(netName, "virt-go")) {
				fmt.Println("Destroy virt-go-net")
				if err := net.Destroy(); err != nil {panic(err)}

				fmt.Println("Undefine virt-go-net")
				if err := net.Undefine(); err != nil {panic(err)}
			}
			continue
		}

		// Delete Datadir
		if err := os.RemoveAll(Datadir); err !=nil {fmt.Println(err)}

		// Delete ConfigFile
		if err := os.RemoveAll("/etc/virt-go/"); err !=nil {fmt.Println(err)}

	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
