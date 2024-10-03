package wand

import (
	"fmt"

	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var reducioLogo string = fmt.Sprintf(`%s
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣷⣶⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣯⣤⣤⣬⣉⠻⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣷⣶⣦⣭⣽⣿⣶⣤⡈⠻⣷⡄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣶⣤⣄⡀⠀⠀⠀⠈⠙⢿⣷⡀⠹⣿⣆⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣧⡈⠙⢿⣿⣿⣶⣤⣄⡀⠀⠀⢻⣷⡄⢹⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣷⣶⣤⣽⣿⣿⣦⣙⠻⣷⣄⠀⢻⣷⡈⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣏⣩⣿⣿⣻⣽⣿⣷⣮⣻⣷⡈⣿⣷⣿⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣯⣿⣿⣿⣿⣿⣿⣿⣟⣿⣿⣧⣿⣿⣿⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠏⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠉⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Red, Reset)

var reducioDescription string = fmt.Sprintf(`%s
%sReducio is a shrinking charm.%s

%sSeen originally in "Harry Potter and the Goblet of Fire", Reducio basically means 'to reduce/shrink'. It results in the reduction of target's shape as soon as it is cast.%s

%sIt makes things smaller, as does the 'tar' command in linux, since it compresses or shrinks files%s

Usage: $ %sreducio /path/to/source1 /path/to/source2 ... /path/to/source-N /path/to/destination%s

%sCheers!%s
`, Cyan, Yellow, Reset, Blue, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var reducioCmd = &cobra.Command{
	Use:  "reducio",
	Long: fmt.Sprintf("%s %s", reducioLogo, reducioDescription),
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		destination := args[len(args)-1]
		sources := args[:len(args)-1]

		tarArgs := append([]string{"-czvf", destination}, sources...)

		wand.Runcommand("tar", tarArgs...)
	},
}


func init() {
	rootCmd.AddCommand(reducioCmd)
}
