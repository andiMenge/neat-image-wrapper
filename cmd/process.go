// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A small wrapper for the neat-image windows command line tool",
	Long:  `It wrapps the neat-image commandline tool for easier batch processing. Neat image must be installed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("process called")
	},
}

func init() {
	RootCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags
	processCmd.Flags().String("src", "", "The path to the source files")
	processCmd.Flags().String("noiseProfile", "", "The path to the neat-image camera noise profile")
	processCmd.Flags().String("filterPreset", "", "The path to the neat-image filter preset")
}
