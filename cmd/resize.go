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
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"libvirt.org/libvirt-go"
)

var size int

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize VM root volum. If VM is started, It will be shutdown automatically",
	Run: func(cmd *cobra.Command, args []string) {

		if size < 20 || Num < 2 {
			fmt.Println("'Size' should be bigger than 20 and 'Num' should be bigger than 2 also")
			fmt.Println("Shrink is not supported")
			os.Exit(71)
		}

		Datadir, NetAddr = GetCFG()
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			fmt.Println(err)
		}

		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
		if err != nil {
			fmt.Println(err)
		}

		for _, dom := range doms {
			domName, _ := dom.GetName()
			domStat, _ := dom.IsActive()
			splitName := strings.Split(domName, "-")
			tail := splitName[len(splitName)-1]

			if tail == strconv.Itoa(Num) && domStat == false {
				fmt.Printf("'%s' is already shutdown. resize started \n", domName)
				cmd := exec.Command("qemu-img", "resize", Datadir+"/volumes/"+domName+"root", strconv.Itoa(size)+"G")

				if err = cmd.Run(); err != nil {
					fmt.Println(err)
					os.Exit(90)
				}

				err = dom.Create()
				if err != nil {
					panic(err)
				}

			} else if tail == strconv.Itoa(Num) && domStat == true {
				var agree string
				fmt.Printf("'%s' is Active \n", domName)
				fmt.Printf("'virt-go' attempt to shutdown '%s' if you agree, enter 'yes' [yes/no] : ", domName)
				fmt.Scanf("%s", &agree)

				if agree == "yes" {
					err = dom.Destroy()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("'%s' has been shutdown and resize root volume to '%d'GB \n", domName, size)
					cmd := exec.Command("qemu-img", "resize", Datadir+"/volumes/"+domName+"root", strconv.Itoa(size)+"G")

					if err = cmd.Run(); err != nil {
						fmt.Println(err)
						os.Exit(90)
					}

					err = dom.Create()
					if err != nil {
						panic(err)
					}

				} else if agree == "no" {
					os.Exit(88)
				}
			} else {
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().IntVarP(&size, "size", "s", 20, "Number of VM Root Volume Size (GB) (required)")
	resizeCmd.MarkFlagRequired("size")
	resizeCmd.Flags().IntVarP(&Num, "number", "n", 0, "Number of VM for identification")
	resizeCmd.MarkFlagRequired("number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
