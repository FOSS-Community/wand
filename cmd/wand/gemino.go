package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var geminoLogo string = fmt.Sprintf(`%s 
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣼⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⡍⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠃⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣤⣾⠿⠛⠛⠛⠿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⠿⠟⠛⠛⠿⢷⣦⡀⠀⠀⠀⠀
⠀⠀⢠⣿⢟⡄⠀⠀⠀⠀⠀⠀⠹⣷⡄⠀⠀⠀⠀⠀⢀⣾⠟⡁⠀⠀⠀⠀⠀⠀⠙⢿⣆⠀⠀⠀
⠀⠀⣿⠇⣾⠁⠀⠀⠀⠀⠀⠀⠀⢹⣷⣾⠿⠛⠻⣷⣾⡟⣸⡇⠀⠀⠀⠀⠀⠀⠀⠈⣿⡄⠀⠀
⠿⠿⣿⡀⣿⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠁⠀⠀⠀⠈⢿⡇⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿⠿⠂
⠀⠀⢻⣧⠹⣧⡀⠀⠀⠀⠀⠀⢀⣾⡏⠀⠀⠀⠀⠀⠸⣿⡘⢷⣄⠀⠀⠀⠀⠀⠀⣰⡿⠀⠀⠀
⠀⠀⠀⠻⣷⣬⣛⠁⠀⠀⣀⣤⣾⠏⠀⠀⠀⠀⠀⠀⠀⠙⢿⣦⣝⡃⠀⠀⢀⣠⣾⠟⠁⠀⠀⠀
⠀⠀⠀⠀⠈⠙⠛⠿⠿⠿⠛⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⠿⠿⠛⠛⠁⠀⠀⠀⠀⠀
%s`, Blue, Reset)

var geminoDescription string = fmt.Sprintf(`%s
%sThe Geminio Curse or Doubling Charm (Geminio) is a spell used to duplicate an object.%s

It can also be used to bewitch an object into multiplying repeatedly when touched, though how one would produce the latter effect is still unknown.

%sIt duplicates objects, creating copies of them, and so does the 'cp' command in the Linux terminal.%s %s'cp' copies files and directories in the Linux world.%s

Usage : %s$ wand geminio /path/to/source /path/to/destination%s

%sCheers!%s
`, Cyan, Blue, Reset, Yellow, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var geminoCmd = &cobra.Command{
	Use:  "gemino",
	Long: fmt.Sprintf("%s %s", geminoLogo, geminoDescription),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]
		wand.Runcommand("cp", source, destination)
	},
}

func init() {
	rootCmd.AddCommand(geminoCmd)
}
