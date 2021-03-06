package cmd

import (
	"fmt"
	"io/ioutil"
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
			log.Fatalln(fmt.Errorf("lockbox name argument is required"))
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatalln("invalid lockbox name")
		}

		// ORDER OF PRIORITY:
		// - flags
		// - environment variables
		// - config file in $HOME/.lockbox.yaml

		// check viper, which checks config and environment
		ns := viper.GetString("namespace")
		u := viper.GetString("username")
		p := viper.GetString("password")
		var v []byte

		if namespace != "" {
			ns = namespace
		}

		if username != "" {
			u = username
		}

		if password != "" {
			p = password
		}

		if value != "" {
			v = []byte(value)
		}

		if value == "" && filepath != "" {
			filev, err := ioutil.ReadFile(filepath)
			if err != nil {
				log.Fatalln(err)
			}
			v = filev
		}

		if value == "" && filepath == "" {
			log.Fatalln("a value or a file must be provided")
		}

		// if we still don't have a namespace, username, password set, we need to error.
		if ns == "" {
			ns = "main" // this is the "default" namespace
		}

		if u == "" {
			log.Fatalln(fmt.Errorf("a username is required"))
		}

		if p == "" {
			log.Fatalln(fmt.Errorf("a password is required"))
		}

		lb, err := lockbox.GetLockbox(args[0])
		if err != nil {
			log.Fatalln(err)
		}
		mfa, err := lb.CheckMFA(ns)
		if err != nil {
			log.Fatal(err)
		}
		if mfa {
			err = lb.UnlockWithMFA(ns, u, p, code)
		} else {
			err = lb.Unlock(ns, u, p)
		}

		if err != nil {
			log.Fatalln(err)
		}

		err = lb.SetValue([]byte(path), v)
		if err != nil {
			log.Fatalln(err)
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
	setCmd.Flags().StringVar(&filepath, "file", "", "path to file")

	setCmd.MarkFlagRequired("path")
}
