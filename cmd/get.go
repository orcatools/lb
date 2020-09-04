package cmd

import (
	"fmt"
	"log"
	"os"

	lockbox "github.com/orcatools/lockbox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a secret value from the lockbox",
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
			log.Fatalln(fmt.Errorf("a username is required"))
		}

		if p == "" {
			log.Fatalln(fmt.Errorf("a password is required"))
		}

		lb, err := lockbox.GetLockbox(args[0])
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		data, err := lb.GetValue([]byte(path))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(data))
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVar(&code, "code", "", "time based code to unlock the lockbox")
	getCmd.Flags().StringVar(&path, "path", "", "path to write the item to")
	getCmd.Flags().StringVar(&namespace, "namespace", "", "namespace to put the item in")
	getCmd.Flags().StringVar(&username, "username", "", "user's username")
	getCmd.Flags().StringVar(&password, "password", "", "user's password")
	getCmd.MarkFlagRequired("path")
}
