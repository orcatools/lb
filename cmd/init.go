package cmd

import (
	"fmt"
	"log"

	lockbox "github.com/orcatools/lockbox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new lockbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: make this error message better by checking args length
		if len(args) != 1 {
			log.Fatalln(fmt.Errorf("lockbox name argument is required"))
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
		otp, err := lb.Init(ns, u, p)
		if err != nil {
			log.Fatalln(err)
		}
		err = lb.Close()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(fmt.Sprintf("OTP SECRET: %v", otp.Secret()))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&namespace, "namespace", "", "the namespace to use")
	initCmd.Flags().StringVar(&username, "username", "", "user's username")
	initCmd.Flags().StringVar(&password, "password", "", "user's password")
}
