package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var lumosLogo string = fmt.Sprintf(`%s
⠀⠀⣶⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣿⣿⣿⣷⣤⡀⠀⠀⠀⠀⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢀⣀⣸⣿⣿⣿⣿⣿⣶⣤⣀⠀⠙⠿⣿⣿⣿⣶⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣀⠈⠻⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠈⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣀⠙⢿⣿⣿⣦⡀⠀⠀⠀⠀⠀
⠀⠀⢠⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠙⣿⣿⣿⣦⡀⠀⠀⠀
⠀⠀⠀⠙⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣬⣿⣿⣿⣷⡀⠀⠀
⠀⠀⠀⠀⠀⢀⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡛⠛⠿⣿⣷⡄⠀
⠀⠀⠀⠀⠀⠀⠙⠛⠛⠻⠿⠿⠿⠿⢿⣿⡿⠿⣿⣿⣿⣇⠀⠀⠈⠻⣿⣆
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⠉⠻⠋⠀⠀⠀⠀⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⣿⣄⠀⠀⠀⠀⠀⠀⢀⣼⣿⠃
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣶⣤⣤⣴⣾⣿⠟⠁⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠀⠀⠀⠀
%s`, Yellow, Reset)

var lumosDescription string = fmt.Sprintf(`%s
%sLumos is the incantation to a charm that can be used to produce a flash of bright white light from the tip of the wand.%s

This light can be thrown far off of the wand, illuminating the surrounding area for several minutes
after its initial casting before darkening once more.

%sLumos shows the path in the wizarding world, and similarly, the 'ls' command shows the path in the Linux world.%s %s'ls' displays the files and folders in the current working directory.%s

Usage : %s$ wand lumos%s

%sCheers!%s
`, Cyan, Yellow, Reset, Blue, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var lumosCmd = &cobra.Command{
	Use:  "lumos",
	Long: fmt.Sprintf("%s %s", lumosLogo, lumosDescription),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("ls")
	},
}

func init() {
	rootCmd.AddCommand(lumosCmd)
}
