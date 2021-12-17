package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get virt-go version",
	Run: func(cmd *cobra.Command, args []string) {
		version := "2.0.04"
		gitVersion := func() string {
			// https://raw.githubusercontent.com/YoungjuWang/virt-go/v2/version.txt
			resp, err := http.Get("https://raw.githubusercontent.com/YoungjuWang/virt-go/v2/version.txt")
			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			return string(data)
		}

		fmt.Println("latest version : ", gitVersion())
		fmt.Println("installed version : ", version, "\n")

		if version != gitVersion() {
			fmt.Println(colorRed, "You need to update virt-go", colorReset)
		} else {
			fmt.Println(colorGreen, "You have latest virt-go", colorReset)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
