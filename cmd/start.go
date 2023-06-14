package cmd

import (
	"github.com/fzxiehui/net_to_uart/log"
	"github.com/spf13/cobra"
)

var configFile string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start net_to_uart",
	Long:  `start net_to_uart`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start net_to_uart")
		log.Debug(args)
		log.Debug(configFile)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// config.yml
	startCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file")
}