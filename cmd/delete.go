package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	libvirt "libvirt.org/libvirt-go"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete VM or Image",
	Run: func(cmd *cobra.Command, args []string) {
		if v.num == 0 && v.image == "" {
			cmd.Help()
		}
		checked := v.checkNumber()
		if checked {
			deleteImage()
		} else {
			if v.image == "" {
				deleteVM()
			} else {
				deleteVM()
				deleteImage()
			}
		}

		fmt.Println(colorGreen + "successfully finished" + colorReset)
	},
}

func deleteVM() {
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

	domName, err := dom.GetName()
	if err != nil {
		log.Fatal(err)
	}

	domStat, err := dom.IsActive()
	if err != nil {
		log.Fatal(err)
	}

	if domStat {
		// Destroy.
		fmt.Printf("■  Shutdown %s.\n", domName)
		err = dom.Destroy()
		if err != nil {
			log.Fatal("err")
		}
	} else {
		fmt.Println("System Already Shutdown")
	}

	// Undefine.
	fmt.Printf("■  Undefine %s.\n", domName)
	err = dom.Undefine()
	if err != nil {
		log.Fatal(err)
	}

	// Delete root volume.
	fmt.Printf("■  Delete %s root volume.\n", domName)
	err = os.Remove(g.dataDir + "/volumes/" + domName + "-root.qcow2")
	if err != nil {
		log.Fatal(err)
	}

	// Delete cloud-init ISO file.
	fmt.Printf("■  Delete %s cloud-init iso file.\n", domName)
	err = os.Remove(g.dataDir + "/volumes/" + domName + "-init.iso")
	if err != nil {
		log.Fatal(err)
	}

	// Delete description file.
	fmt.Printf("■  Delete %s description file.\n", domName)
	err = os.Remove(g.dataDir + "/volumes/" + domName)
	if err != nil {
		log.Fatal(err)
	}

	// Delete volume file.
	files, err := filepath.Glob(g.dataDir + "/volumes/" + domName + "*.img")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Printf("■  Delete %s.\n", f)
		err := os.Remove(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func deleteImage() {
	// Delete Image.
	fmt.Printf("■  Delete %s.\n", v.image)
	err := os.Remove(g.dataDir + "/images/" + v.image)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having the number will be deleted")
	deleteCmd.Flags().StringVarP(&v.image, "image", "i", "", "Image, will be deleted")
}
