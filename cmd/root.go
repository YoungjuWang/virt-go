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
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "virt-go",
	Long: `
'virt-go' help user using libvirt to manage virtual resources easy and fast.
Futher informations about 'virt-go' are in https://github.com/YoungjuWang/virt-go.
Have a good day. :)
`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
