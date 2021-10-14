package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete virt-go",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		net, err := conn.LookupNetworkByName("virt-go-net")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("■  Shutdown 'virt-go-net'.")
		err = net.Destroy()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("■  Undefine 'virt-go-net'.")
		err = net.Undefine()
		if err != nil {
			log.Fatal(err)
		}

		// Delete ConfigFile
		if err := os.RemoveAll("/etc/virt-go/"); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	clearCmd.Flags().BoolP("confirm", "y", false, "Confirm to delete")
	clearCmd.MarkFlagRequired("confirm")
}
