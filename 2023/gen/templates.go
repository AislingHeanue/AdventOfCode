package main

const (
	rootTemplate = `// Code generated by 'go run ./gen'; DO NOT EDIT
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	{{ range $day := seq 1 .N }}
	"github.com/AislingHeanue/Advent-Of-Code/2023/core/day{{ $day }}"
	{{- end }}
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
)
	

func addDays(root *cobra.Command) {
    {{- range $day := seq 1 .N }}
	day{{ $day }}.AddCommandsTo(root)
    {{- end }}
}


func CreateCommand() *cobra.Command {
	var startTime time.Time

	root := &cobra.Command{
		Use:     "2023",
		Short:   "Advent of Code 2023",
		Long:    "Go implementation of my solutions for the 2023 Advent of Code\nGeneration template provided by github.com/nlowe",
		Example: "go run . 1a",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			startTime = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Time: %v\n", time.Since(startTime))
			util.AwaitClosure()
		},
	}

	addDays(root)

	return root
}
`

	glueTemplate = `package day{{ .N }}

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	root.AddCommand(aCommand())
    root.AddCommand(bCommand())
}
`

	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func {{ .AB | toLower }}Command() *cobra.Command {
    return &cobra.Command{
        Use:   "{{ .N }}{{ .AB | toLower }}",
        Short: "Day {{ .N }}, Problem {{ .AB }}",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", part{{ .AB | toUpper }}(core.FromFile()))
        },
    }
}

func part{{ .AB | toUpper }}(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
    panic("Not implemented!")
}
`

	testTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/stretchr/testify/require"
)

func Test{{ .AB }}(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral("foobar")

	result := part{{ .AB | toUpper }}(input)

	require.Equal(t, 42, result)
	util.AwaitClosure()
}
`

	benchmarkTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
)


func BenchmarkA(b *testing.B) {
	util.ForceNoWindow = true
	for i := 0; i < b.N; i++ {
		_ = partA(core.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	util.ForceNoWindow = true
	for i := 0; i < b.N; i++ {
		_ = partB(core.FromFile())
	}
}`
)
