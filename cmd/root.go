package cmd

import (
	"fmt"
	"os"

	"github.com/AdrianThePirate/WSC---Website-Checker/internal/util"
	"github.com/spf13/cobra"
)

var (
	httpMode bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "wsc <url/ip>",
	Example: "wcs example.com",
	Short:   "Checks url/ip site for issues",
	Long: `WebSite Checker (wsc) is a CLI tool to check website health.
This application is a crawler that checks a given website for:
	* Dead links
	* Long load times
	* Multi-redirects`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain, err := util.ValidateDomain(args[0], httpMode)
		if err != nil {
			return err
		}

		err = util.ValidateConnection(domain)
		if err != nil {
			return err
		}
		fmt.Println(domain)
		return nil
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
	rootCmd.Flags().BoolVarP(&httpMode, "http", "s", false, "assume http instead of https")
}
