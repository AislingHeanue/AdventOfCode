// Code generated by 'go run ./gen'; DO NOT EDIT
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day1"
)
	

func addDays(root *cobra.Command) {
	day1.AddCommandsTo(root)
}


func CreateCommand() *cobra.Command {
	var startTime time.Time

	root := &cobra.Command{
		Use:     "2023",
		Short:   "Advent of Code 2023",
		Long:    "Go implementation of my solutions for the 2023 Advent of Code\nGeneration template provided by github.com/nlowe",
		Example: "go run main.go 1a",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			startTime = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Time: %v\n", time.Since(startTime))
		},
	}

	addDays(root)

	return root
}
