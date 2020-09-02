package cmd

import (
	"fmt"
	"log"
	"os"

	lockbox "github.com/orcatools/lockbox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a secret value in the lockbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(fmt.Errorf("lockbox name argument is required"))
			os.Exit(1)
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			fmt.Println("invalid lockbox name")
			os.Exit(1)
		}

		// ORDER OF PRIORITY:
		// - flags
		// - environment variables
		// - config file in $HOME/.lockbox.yaml

		// check viper, which checks config and environment
		ns := viper.GetString("namespace")
		u := viper.GetString("username")
		p := viper.GetString("password")

		if namespace != "" {
			ns = namespace
		}

		if username != "" {
			u = username
		}

		if password != "" {
			p = password
		}

		// if we still don't have a namespace, username, password set, we need to error.
		if ns == "" {
			ns = "main" // this is the "default" namespace
		}

		if u == "" {
			fmt.Println(fmt.Errorf("a username is required"))
			os.Exit(1)
		}

		if p == "" {
			fmt.Println(fmt.Errorf("a password is required"))
			os.Exit(1)
		}

		lb, err := lockbox.GetLockbox(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = lb.Unlock(ns, u, p, code)
		if err != nil {
			log.Fatal(err)
		}

		err = lb.SetValue([]byte(path), []byte(value))
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVar(&code, "code", "", "time based code to unlock the lockbox")
	setCmd.Flags().StringVar(&namespace, "namespace", "", "namespace to put the item in")
	setCmd.Flags().StringVar(&path, "path", "", "default path to write the item to")
	setCmd.Flags().StringVar(&value, "value", "", "value to write to the item")
	setCmd.Flags().StringVar(&username, "username", "", "user's username")
	setCmd.Flags().StringVar(&password, "password", "", "user's password")

	setCmd.MarkFlagRequired("code")
	setCmd.MarkFlagRequired("path")
	setCmd.MarkFlagRequired("value")
}
