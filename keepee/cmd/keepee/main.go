// Package main is the entry point for the keepee agent CLI.
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/milvasic/syskeeper/keepee/internal/config"
)

// Build-time variables injected via ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var cfgFile string

	root := &cobra.Command{
		Use:   "keepee",
		Short: "keepee — syskeeper monitoring agent",
		Long:  "keepee is the agent component of syskeeper. It collects system telemetry and pushes it to a keeper server.",
	}

	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to config file (default: /etc/keepee/config.yml)")

	root.AddCommand(
		newVersionCmd(),
		newRegisterCmd(&cfgFile),
		newRunCmd(&cfgFile),
	)

	return root
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("keepee version %s (commit: %s, built: %s)\n", version, commit, date)
		},
	}
}

func newRegisterCmd(cfgFile *string) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Register this agent with a keeper server",
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := loadConfig(*cfgFile)
			if err != nil {
				return err
			}

			slog.Info("registering agent", "server_url", cfg.ServerURL)
			// TODO: implement registration logic
			return nil
		},
	}
}

func newRunCmd(cfgFile *string) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Start the keepee agent and begin pushing telemetry",
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := loadConfig(*cfgFile)
			if err != nil {
				return err
			}

			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			slog.Info("starting keepee agent",
				"server_url", cfg.ServerURL,
				"push_interval", cfg.PushInterval,
				"ping_interval", cfg.PingInterval,
			)
			// TODO: implement main run loop
			return nil
		},
	}
}

func loadConfig(path string) (config.Config, error) {
	if path == "" {
		path = "/etc/keepee/config.yml"
	}

	cfg, err := config.LoadFile(path)
	if err != nil {
		// If the default path doesn't exist, use defaults silently.
		if os.IsNotExist(err) {
			return config.DefaultConfig(), nil
		}

		return config.Config{}, fmt.Errorf("loading config: %w", err)
	}

	return cfg, nil
}
