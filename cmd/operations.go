package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

func startVM(conn *libvirt.Connect) {
	v.name = v.getRunningVMName(conn)
	dom, err := conn.LookupDomainByName(v.name)
	if err != nil {
		log.Fatal(err)
	}

	err = dom.Create()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colorGreen + "Start " + v.name + colorReset)
}

func stopVM(conn *libvirt.Connect) {
	v.name = v.getRunningVMName(conn)
	dom, err := conn.LookupDomainByName(v.name)
	if err != nil {
		log.Fatal(err)
	}

	err = dom.Destroy()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colorRed + "Stop " + v.name + colorReset)
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)
	startCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be started")
	startCmd.MarkFlagRequired("number")
	stopCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be stopped")
	stopCmd.MarkFlagRequired("number")
	restartCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be restarted")
	restartCmd.MarkFlagRequired("number")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start virt-go VM",
	Run: func(cmd *cobra.Command, args []string) {
		v.checkNumber()

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}

		startVM(conn)
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop virt-go VM",
	Run: func(cmd *cobra.Command, args []string) {
		v.checkNumber()

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}

		stopVM(conn)
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart virt-go VM",
	Run: func(cmd *cobra.Command, args []string) {
		v.checkNumber()

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}

		stopVM(conn)
		time.Sleep(time.Second * 2)
		startVM(conn)
	},
}
