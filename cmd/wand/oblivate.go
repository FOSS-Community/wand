package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var obliviateLogo string = fmt.Sprintf(`%s
⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⡴⠶⠚⠛⠛⠛⠛⠓⠶⢦⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣠⡴⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⢦⣄⠀⠀⠀⠀⠀
⠀⠀⠀⣠⡾⠋⠀⢀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⣶⡍⠙⣷⣦⠀⠙⢷⣄⠀⠀⠀
⠀⠀⣴⠋⠀⢠⣾⣿⡟⠉⠙⣿⣷⣆⠀⠀⠀⠀⠙⠁⠀⣼⡿⠀⠀⠀⠹⣦⠀⠀
⠀⣼⠃⠀⢠⣿⣿⣿⠀⠀⠀⢸⣿⣿⣇⠀⠀⠀⠀⠈⠛⣧⣄⠀⠀⠀⠀⠘⣧⠀
⢰⡏⠀⠀⢸⣿⣿⣿⠀⠀⠀⢸⣿⣿⣿⡀⠀⢠⣶⡆⠀⣸⣿⠇⠀⠀⠀⠀⢹⡆
⣾⠁⠀⠀⠸⣿⣿⣿⡀⠀⠀⣸⣿⣿⣿⡇⠀⠀⠛⠦⠶⠛⠋⠀⠀⠀⠀⠀⠀⣷
⣿⠀⠀⠀⠀⠙⢿⣿⣷⣄⣠⢿⣿⣿⣿⡇⠀⠀⠉⠉⢉⣉⠉⠁⠀⠀⠀⠀⠀⣿
⢿⠀⠀⠀⠀⠀⠀⠈⠉⠉⠀⢸⣿⣿⣿⠁⠀⠀⠀⢠⣿⣿⠀⠀⠀⠀⠀⠀⠀⡿
⠸⣇⠀⠀⠀⣴⣾⣶⣆⠀⠀⣸⣿⣿⡏⠀⠀⠀⣠⠏⣿⣿⠀⠀⠀⠀⠀⠀⣸⠇
⠀⢻⡄⠀⠀⢿⣿⡿⠋⠀⢀⣿⣿⠏⠀⠀⠀⢰⣏⣀⣿⣿⣀⠀⠀⠀⠀⢠⡟⠀
⠀⠀⠻⣆⠀⠈⠙⠛⠶⠞⠛⠋⠁⠀⠀⠀⠀⠈⠉⠉⣿⣿⠉⠀⠀⠀⣰⠟⠀⠀
⠀⠀⠀⠙⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠛⠉⠀⣠⡾⠋⠀⠀⠀
⠀⠀⠀⠀⠀⠙⠳⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⠞⠋⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠳⠶⢤⣤⣤⣤⣤⡤⠶⠞⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Yellow, Reset)

var obliviateDescription string = fmt.Sprintf(`%s
%sOriginating from Harry Potter and the Chamber of Secrets, Obliviate is a term meaning 'forget'.%s

Obliviate is a memory charm, resulting in the erasure of the recipient's memory. One's memory vanishes as soon as this charm is cast.

%sIt makes one forget things or memories, and similarly, the 'clear' command in the Linux terminal 'obliviates' the screen.%s %s'clear' erases the terminal's output.%s

Usage : %s$ wand obliviate%s

%sCheers!%s
`, Cyan, Yellow, Reset, Blue, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var obliviateCmd = &cobra.Command{
	Use:  "obliviate",
	Long: fmt.Sprintf("%s %s", obliviateLogo, obliviateDescription),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("clear")
	},
}

func init() {
	rootCmd.AddCommand(obliviateCmd)
}
