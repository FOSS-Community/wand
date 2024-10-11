package wand

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Color string

const (
	Reset   Color = "\033[0m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	Grey    Color = "\033[37m"
	White   Color = "\033[97m"
)

var version = "0.0.1"

var commandlist = []string{
	fmt.Sprintf("\t%s1.\tlumos%s\n", Grey, Reset),
	fmt.Sprintf("\t%s2.\toblivate%s\n", Green, Reset),
	fmt.Sprintf("\t%s3.\tgeminio%s\n", Yellow, Reset),
	fmt.Sprintf("\t%s4.\tlegilimens%s\n", Blue, Reset),
	fmt.Sprintf("\t%s5.\tmobiliarbus%s\n", Magenta, Reset),
	fmt.Sprintf("\t%s6.\tavada-kedavra%s\n", Cyan, Reset),
	fmt.Sprintf("\t%s7.\taccio%s\n", Red, Reset),
	fmt.Sprintf("\t%s8.\tnox%s\n", White, Reset),
	fmt.Sprintf("\t%s9.\reducio%s\n", Red, Reset),
}

var logowand string = `
                          _ 
                         | |
 __      ____ _ _ __   __| |
 \ \ /\ / / _  | '_ \ / _  |
  \ V  V / (_| | | | | (_| |
   \_/\_/ \__,_|_| |_|\__,_|
`

var Short string = fmt.Sprintf("%s\n%s wand - a magical tool to learn linux commands %s \n", logowand, Green, Reset)
var Long string = fmt.Sprintf("%s\n%s wand is a CLI tool that will help you to lean linux command by potterifying them %s \n", logowand, Magenta, Reset)

var rootCmd = &cobra.Command{
	Use:     "wand",
	Short:   Short,
	Long:    Long,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s\n", Blue, logowand, Reset)
		fmt.Printf("%s\tPotter-ify your Linux Experience!%s\n", Yellow, Reset)
		fmt.Printf("%s\tType 'wand SPELL_NAME --help' for help%s\n", White, Reset)
		fmt.Println()
		fmt.Printf("%sFollowing are the spells you can use: %s\n", White, Reset)
		fmt.Println()
		for _, v := range commandlist {
			fmt.Printf("%s\n", v)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s Whoops. Some error occurs '%s'\n %s", Red, err, Reset)
		os.Exit(1)
	}
}
