/*
Copyright © 2023 Huge Works <hugeworkscompany@gmail.com>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const aboutCommandName = "about"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quact",
	Short: "Query All Cloud Things",
	Long:  `quact is a tool by Huge Works that helps you discover, inventory, and analyze your cloud resources.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		aboutIsSet := cmd.Flags().Lookup(aboutCommandName).Changed

		if aboutIsSet {
			aboutText := "Huge Works <hugeworkscompany@gmail.com>\nhttps://github.com/huge-works"

			fmt.Println(aboutText)
		} else {
			fmt.Println("Running quact.")
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.quact.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP(aboutCommandName, "a", false, "Print info about the author.")
}
