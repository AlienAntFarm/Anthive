package main

//go:generate go run include.go sql/*

import (
	"github.com/alienantfarm/anthive/assets"
	"github.com/alienantfarm/anthive/ext/db"
	"github.com/alienantfarm/anthive/utils"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var verbosity int

func runAsset(assetName string) {
	asset := assets.Get(assetName)
	glog.Infof("\n%s", asset)
	_, err := db.Client().Query(asset)
	if err != nil {
		glog.Fatalf("%s", err)
	}
}

var rootCmd = &cobra.Command{
	Use:              "anthivectl",
	Short:            "Simple cli to deal with various part of anthive",
	PersistentPreRun: utils.PreRun,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init tables and types for anthive database",
	Run: func(cmd *cobra.Command, args []string) {
		runAsset("sql/init.sql")
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean every tables from anthive database",
	Run: func(cmd *cobra.Command, args []string) {
		runAsset("sql/clean.sql")
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Recreate every tables from anthive database",
	Run: func(cmd *cobra.Command, args []string) {
		runAsset("sql/remove.sql")
		runAsset("sql/init.sql")
	},
}

func main() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(resetCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("%s", err)
	}
}
