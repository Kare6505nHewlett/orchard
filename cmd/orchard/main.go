package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := newRootCmd().ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orchard",
		Short: "Orchard — a CI orchestration tool for Apple Silicon",
		Long: `Orchard is a CI orchestration tool that manages virtual machines
on Apple Silicon hardware using the Virtualization.framework.

See https://github.com/cirruslabs/orchard for upstream documentation.
Personal fork: https://github.com/nicholasgasior/orchard`,
		SilenceUsage:  true,
		SilenceErrors: true, // print errors ourselves for consistent formatting
	}

	cmd.AddCommand(
		newVersionCmd(),
	)

	return cmd
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "orchard version %s (commit: %s, built: %s)\n", version, commit, date)
		},
	}
}
