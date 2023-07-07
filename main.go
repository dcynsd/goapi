package main

import (
	"fmt"
	"os"

	"goapi/bootstrap"
	"goapi/cmd"
	"goapi/cmd/make"
	"goapi/pkg/config"
	"goapi/pkg/console"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "goapi",
		Short: "A basic project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {

			config.SetupConfig(command.CommandPath())

			bootstrap.SetupLogger()

			bootstrap.SetupDB()

			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
