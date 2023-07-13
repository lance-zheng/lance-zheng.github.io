package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "Generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello cobra!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("fail to init cli")
	}
}
