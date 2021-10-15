package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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

func resizeVM(conn *libvirt.Connect) {
	v.name = v.getRunningVMName(conn)
	volFile := g.dataDir + "/volumes/" + v.name + "-root.qcow2"

	fmt.Println("resize", volFile)

	var cmd *exec.Cmd
	if shrink {
		cmd = exec.Command("qemu-img", "resize", "--shrink", volFile, strconv.Itoa(int(v.size))+"G")
	} else {
		cmd = exec.Command("qemu-img", "resize", volFile, strconv.Itoa(int(v.size))+"G")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(colorGreen + volFile + " size is " + strconv.Itoa(int(v.size)) + "G" + colorReset)
	time.Sleep(time.Second * 2)
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(resizeCmd)
	startCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be started")
	startCmd.MarkFlagRequired("number")
	stopCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be stopped")
	stopCmd.MarkFlagRequired("number")
	restartCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be restarted")
	restartCmd.MarkFlagRequired("number")
	resizeCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be restarted")
	resizeCmd.MarkFlagRequired("number")
	resizeCmd.Flags().UintVarP(&v.size, "size", "s", 20, "VM's Root Volume Size (GB)")
	resizeCmd.MarkFlagRequired("number")
	resizeCmd.Flags().BoolVar(&shrink, "shrink", false, "Shrink volume")
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
var shrink bool
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize virt-go VM",
	Run: func(cmd *cobra.Command, args []string) {
		v.checkNumber()

		conn, err := libvirt.NewConnect("qemu:///system")
		if err != nil {
			log.Fatal(err)
		}

		resizeVM(conn)
	},
}
