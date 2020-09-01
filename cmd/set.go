package cmd

import (
	"log"
	"os"

	lockbox "github.com/orcatools/lockbox"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a secret value in the lockbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lb, err := lockbox.GetLockbox(args[0], os.Getenv("LOCKBOX_MASTER_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		err = lb.Unlock(namespace, code, salt)
		if err != nil {
			log.Fatal(err)
		}

		err = lb.SetValue([]byte(value), namespace, path, salt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setCmd.Flags().StringVar(&code, "code", "", "time based code to unlock the lockbox")
	setCmd.Flags().StringVar(&namespace, "namespace", "main", "namespace to put the item in")
	setCmd.Flags().StringVar(&path, "path", "/", "default path to write the item to")
	setCmd.Flags().StringVar(&value, "value", "", "value to write to the item")
	setCmd.Flags().StringVar(&salt, "salt", "", "salt to add extra layer of security")
}
