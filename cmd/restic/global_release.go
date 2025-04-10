//go:build false && !debug && !profile
// +build false,!debug,!profile

package main

import "github.com/spf13/cobra"

func registerProfiling(_ *cobra.Command) {
	// No profiling in release mode
}
