/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"

	lockbox "github.com/orcatools/lockbox"
	"github.com/spf13/cobra"
)

// var (
// 	password  string
// 	namespace string
// 	path      string
// 	value     string
// )

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a secret value from the lockbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lb, err := lockbox.GetLockbox(args[0], os.Getenv("LOCKBOX_MASTER_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		err = lb.Unlock(namespace, code)
		if err != nil {
			log.Fatal(err)
		}

		data, err := lb.GetValue(namespace, path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringVar(&code, "code", "", "time based code to unlock the lockbox")
	getCmd.Flags().StringVar(&namespace, "namespace", "main", "namespace to put the item in")
	getCmd.Flags().StringVar(&path, "path", "/", "default path to write the item to")
	// getCmd.Flags().StringVar(&value, "value", "", "value to write to the item")
}
