package cmd

import (
	"strconv"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	user string
	host string
)

// https://stackoverflow.com/questions/61293342/execute-ssh-in-golang
// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Connect to virt-go VM via ssh",
	Run: func(cmd *cobra.Command, args []string) {
		Datadir, NetAddr = GetCFG()
		hostNetAddr := NetAddr + "." + strconv.Itoa(Num)

		scmd := exec.Command("ssh", user+"@"+hostNetAddr, "-p", "22", "-o", "StrictHostKeyChecking=no")
    scmd.Stdin = os.Stdin
    scmd.Stdout = os.Stdout
    scmd.Stderr = os.Stderr

		err := scmd.Run()
		if err != nil {
		panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVarP(&user, "user", "u", "root", "ssh user")
	sshCmd.Flags().IntVarP(&Num, "number", "n", 0, "Number of VM for identification")
	sshCmd.MarkFlagRequired("number")
}
