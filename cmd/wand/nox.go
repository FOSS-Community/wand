package wand

import (
	"fmt"

	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var noxLogo string = fmt.Sprintf(`%s
 ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⢶⣶⣤⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⡿⠛⠋⠙⠛⠿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⡇⠀⠀⠀⠀⠀⠀⠈⢿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣴⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣾⣿⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢠⣿⣿⠏⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣷⣄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣾⣿⠏⠀⠀⠀⠀⠀⢠⣴⣾⣿⠿⢿⣿⣷⣶⣤⡀⠀⠀⠀⠀⠘⣿⣿⣦⡀⠀⠀⠀⠀
⠀⠀⢰⣿⡏⠀⠀⠀⠀⠀⣴⣿⡿⠋⠀⠀⠀⠙⢿⣿⣿⣿⣶⣤⣄⠀⣀⣽⣿⣿⣿⣄⠀⠀⠀
⠀⠀⢸⣿⡇⠀⠀⠀⢀⣾⣿⠏⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀
⠀⠀⠘⣿⣇⠀⠀⠀⢸⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣿⣿⣿⠛⠉⠛⢿⣿⣿⣿
⠀⠀⠀⠘⣿⣇⠀⠀⢸⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⡿⠛⠻⢿⣧⣤⣤⣼⣿⣿⡇
⠀⠀⠀⠀⠘⣿⣧⠀⠘⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣧⠀⠀⠀⠉⠉⠙⢿⣿⣿⡇
⠀⠀⠀⠀⠀⠘⣿⣧⠀⠹⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣷⡀⠀⠀⠀⠀⢸⣿⣿⡇
⠀⠀⠀⠀⠀⠀⠘⣿⣧⠀⠘⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⣷⡀⠀⠀⠀⢸⣿⣿⠇
⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⡀⠀⠙⠳⢤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣦⡀⠀⣸⣿⡟⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣄⠀⠀⠀⠙⠛⠶⣤⣄⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣶⣿⣿⠃⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣷⡀⠀⠀⠀⠀⠀⠉⠛⠷⣤⡀⠀⠀⠀⠈⠻⣿⣿⡿⠋⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣶⣤⣤⣶⣿⡟⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠋⠉⠉⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⣿⣿⣿⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⣿⣿⣿⣿⣿⣶⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠿⠿⠿⠿⠿⠿⠿⠿⠛⠋⠀⠀⠀⠀⠀
%s`, White, Reset)

var noxDescription string = fmt.Sprintf(`%s
%sNox is a The Wand-Extinguishing Charm.%s

It is a spell which causes the light at the end of the caster's wand to be extinguished.

%snox ends your path in the wizarding world%s %sTyping "nox" in your terminal will make the computer shutdown.%s

Usage : %s$ wand nox%s

%sCheers!%s
`, Cyan, Yellow, Reset, Blue, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var noxCmd = &cobra.Command{
	Use:  "nox",
	Long: fmt.Sprintf("%s %s", noxLogo, noxDescription),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("shutdown", "-h", "now")
	},
}

func init() {
	rootCmd.AddCommand(noxCmd)
}
