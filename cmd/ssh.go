package cmd

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var user string

// https://stackoverflow.com/questions/61293342/execute-ssh-in-golang
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Connect to virt-go VM via ssh",
	Run: func(cmd *cobra.Command, args []string) {
		vmNetAddr := g.netAddr + "." + strconv.Itoa(int(v.num))

		scmd := exec.Command("ssh", user+"@"+vmNetAddr, "-p", "22", "-o", "StrictHostKeyChecking=no")
		scmd.Stdin = os.Stdin
		scmd.Stdout = os.Stdout
		scmd.Stderr = os.Stderr

		err := scmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVarP(&user, "user", "u", "root", "ssh user")
	sshCmd.Flags().Uint8VarP(&v.num, "number", "n", 0, "VM having this number will be connected")
	sshCmd.MarkFlagRequired("number")
}
