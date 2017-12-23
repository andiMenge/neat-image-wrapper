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
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type neatConfig struct {
	binary       string
	noiseProfile string
	filterPreset string
	args         string
}

type neatError struct {
	files []string
	count int
}

var (
	neatImage    = neatConfig{}
	processError = neatError{}
	jpgs         = make([]string, 0) //initialize array
	src          string
)

const JpgRegEx = "(?i)\\.(jpg|jpeg)" //(?i)=case insensitive

// processCmd represents the process command. Main logic is in here
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "A small wrapper for the neat-image windows command line tool",
	Long:  `It wrapps the neat-image commandline tool for easier batch processing. Neat image must be installed`,
	Run: func(cmd *cobra.Command, args []string) {
		// Welcome message
		fmt.Println("## NEAT-IMAGE-WRAPPER ##")

		// walk the src dir and look for JPEG images
		err := walkSrcDir()
		if err != nil {
			fmt.Errorf("Error: %s ", err)
			os.Exit(1)
		}

		//call neat image on the found JPEG images
		processJpgs(jpgs)

		// print debug stats
		//debug()
		printStats()
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(processCmd)
	// Here you will define your flags and configuration settings.
	// local flags
	processCmd.Flags().StringVarP(&src, "src", "s", "", "The path to the source files")
	processCmd.Flags().StringVarP(&neatImage.noiseProfile, "noiseProfile", "n", "", "The path to the neat-image camera noise profile")
	processCmd.Flags().StringVarP(&neatImage.filterPreset, "filterPreset", "f", "", "The path to the neat-image filter preset")
	processCmd.Flags().StringVarP(&neatImage.binary, "neatCliBinary", "b", "", "The path to the neat-image cli binary")
	processCmd.Flags().StringVarP(&neatImage.args, "output args", "o", "-oi", "NeatImage output arguments")

}

func debug() {
	fmt.Printf("src: %v\n", src)
	fmt.Printf("noiseProfile: %v\n", neatImage.noiseProfile)
	fmt.Printf("filterPreset: %v\n", neatImage.filterPreset)
	fmt.Printf("neatBinary: %v\n", neatImage.binary)
	for _, i := range jpgs {
		fmt.Printf("%v\n", i)
	}
}

func walkSrcDir() error {
	err := filepath.Walk(src, findJpgs) // parses dir recursive and execs findJpgs for every element in dir
	if err != nil {
		return fmt.Errorf("Error walking directory: %s ", err)
	}
	return nil
}

// check if a path contains a jpeg signature
func isJpg(path string) bool {
	r, err := regexp.Compile(JpgRegEx)
	if err != nil {
		fmt.Errorf("Error compiling RegEx: %s ", err)
	}

	if r.MatchString(path) {
		return true
	} else {
		return false
	}
}

func findJpgs(path string, f os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("foo: %s ", err)
	}

	if isJpg(path) {
		jpgs = append(jpgs, path)
	}
	return nil
}

func processJpgs(jpgs []string) {
	// set error counter to zero
	processError.count = 0

	// loop over jpgs and call neat on them
	for _, jpg := range jpgs {
		out, err := exec.Command(neatImage.binary, jpg, neatImage.noiseProfile, neatImage.filterPreset, neatImage.args).Output()
		if err != nil {
			processError.count++
			processError.files = append(processError.files, jpg)
			fmt.Printf("%s\n", out)
		}
		fmt.Printf("%s\n", out)
	}
}

func printStats() {
	fmt.Printf("\n\n## PROCESS STATS ##\nErrorCount: %v\nErrorFiles: %v\n", processError.count, processError.files)
}
