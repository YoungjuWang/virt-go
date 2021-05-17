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
	"syscall"

	"github.com/spf13/cobra"
	"libvirt.org/libvirt-go"
)

var (
	cidr    string
	Datadir string
	NetAddr string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init 'virt-go' environment",
	Long:  `Init 'virt-go' Virtual Network and Data Directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 마지막에 있는 '/' 문자 삭제
		Datadir = strings.TrimRight(Datadir, "/")

		syscall.Umask(0)
		// Create Data dir
		err := os.MkdirAll(Datadir+"/cloudinit", 0777)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll(Datadir+"/volumes", 0777)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll(Datadir+"/images", 0777)
		if err != nil {
			panic(err)
		}

		// Create Config Directory
		err = os.MkdirAll("/etc/virt-go", 0777)
		if err != nil {
			panic(err)
		}

		// Create Config File
		virtGoCFG, err := os.Create("/etc/virt-go/virt-go.cfg")
		if err != nil {
			panic(err)
		}

		// Create Sample Coludinit data
		userDataFile, err := os.Create(Datadir + "/cloudinit/user-data")
		if err != nil {
			panic(err)
		}
		userData := `#cloud-config
users:
  - name: root
    ssh_authorized_keys:
      - <pub-key>
password: testtest
chpasswd:
  list: |
    root:testtest
  expire: False
ssh_pwauth: True
runcmd:
  - growpart /dev/sda 1`
		userDataFile.WriteString(userData)

		metaDataFile, err := os.Create(Datadir + "/cloudinit/meta-data")
		if err != nil {
			panic(err)
		}
		metaData := ""
		metaDataFile.WriteString(metaData)

		// Create Network
		adds := strings.Split(cidr, ".")
		NetAddr = strings.Join(adds[:3], ".")

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		net, err := conn.NetworkDefineXML(GenNetwork(NetAddr))
		if err != nil {
			fmt.Println(err)
		}

		err = net.SetAutostart(true)
		if err != nil {
			fmt.Println(err)
		}

		err = net.Create()
		if err != nil {
			fmt.Println(err)
		}

		cfgData := "Datadir=" + Datadir + "\n" + "NetAddr=" + NetAddr
		virtGoCFG.WriteString(cfgData)

		// Generate /etc/hosts
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&cidr, "cidr", "c", "10.62.62.0", "Network Address (required)")
	initCmd.MarkFlagRequired("cidr")
	initCmd.Flags().StringVarP(&Datadir, "datadir", "d", "/etc/virt-go", "Which is VM's volume, cloudinit and images stored in (required)")
	initCmd.MarkFlagRequired("datadir")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
