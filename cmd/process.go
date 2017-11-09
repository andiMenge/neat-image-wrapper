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

var (
	src, noiseProfile, filterPreset, neatCliBinary string
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A small wrapper for the neat-image windows command line tool",
	Long:  `It wrapps the neat-image commandline tool for easier batch processing. Neat image must be installed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("process called")
		debug()
	},
}

func init() {
	RootCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.
	// local flags
	processCmd.Flags().StringVarP(&src, "src", "s", "", "The path to the source files")
	processCmd.Flags().StringVarP(&noiseProfile, "noiseProfile", "n", "", "The path to the neat-image camera noise profile")
	processCmd.Flags().StringVarP(&filterPreset, "filterPreset", "f", "", "The path to the neat-image filter preset")
	processCmd.Flags().StringVarP(&neatCliBinary, "neatCliBinary", "b", "", "The path to the neat-image cli binary")
}

func debug() {
	fmt.Printf("src: %v\n", src)
	fmt.Printf("noiseProfile: %v\n", noiseProfile)
	fmt.Printf("filterPreset: %v\n", filterPreset)
	fmt.Printf("neatBinary: %v\n", neatCliBinary)
}
