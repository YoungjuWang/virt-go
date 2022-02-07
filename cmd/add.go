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
	Use:     "disk",
	Short:   "Add disks to VM",
	Example: "virt-go add disk -n 10 --disks \"vda=10 vdb=20\"",
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

		// function in create.go
		createAdditionalDisks(v.disks, v.name, dom)
		fmt.Println(colorGreen + "successfully finished" + colorReset)
	},
}

var netCmd = &cobra.Command{
	Use:     "net",
	Short:   "Add network to VM",
	Example: "virt-go add net -n 10 --nets \"bridge=public-br network=virt-go-net\"",
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

		// function in create.go
		attachAdditionalNetworks(v.nets, v.name, dom)
		fmt.Println(colorGreen + "successfully finished" + colorReset)
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
	diskCmd.MarkFlagRequired("disks")

	// netCmd
	addCmd.AddCommand(netCmd)
	netCmd.Flags().SortFlags = false
	netCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "Number, VM will use")
	netCmd.MarkFlagRequired("number")
	netCmd.Flags().StringVar(&v.nets, "nets", "none", "additional network list")
	netCmd.MarkFlagRequired("nets")
}
