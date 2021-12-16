package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add disks or network interfaces to VM",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Add disks to VM",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		v.name = v.getRunningVMName(conn)
		dom, err := conn.LookupDomainByName(v.name)
		if err != nil {
			log.Fatal(err)
		}
		defer dom.Free()

		createAdditionalDisks(v.disks, v.name, dom)
		fmt.Println(colorGreen + "successfully finished" + colorReset)
	},
}

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Add network to VM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Network")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// diskCmd
	addCmd.AddCommand(diskCmd)
	diskCmd.Flags().SortFlags = false
	diskCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "Number, VM will use")
	diskCmd.MarkFlagRequired("number")
	diskCmd.Flags().StringVar(&v.disks, "disks", "none", "additional disk list")
	//diskCmd.MarkFlagRequired("disks")

	// netCmd
	addCmd.AddCommand(netCmd)
}
