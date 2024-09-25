package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var accioLogo string = fmt.Sprintf(`%s   
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⡴⠶⠶⠶⠶⠶⠶⠶⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡾⠋⠀⠀⠀⠐⢦⣄⣀⣀⣤⣴⠞⠀⠀⠀⠀⠀⠙⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣀⣤⣤⠶⠶⠟⠁⠘⠛⠛⠛⠛⠶⣤⣀⣀⣤⡶⠟⠛⠉⡙⠛⠂⠀⠀⠛⠛⠉⢉⡛⠛⠳⠶⣤⣄⡀⠀⠀⠀⠀⠀
⠀⢀⣠⠾⠛⣉⣤⣤⣄⠀⠀⠀⣤⠾⣿⡛⠷⣄⠀⠀⠀⠀⠀⣠⡶⢻⡛⠷⣄⠀⠀⠀⠀⣴⠟⠋⠙⠛⠳⢦⣌⡛⢶⣄⠀⠀⠀
⢠⡞⢁⣴⠟⠉⠁⠀⢿⣄⠀⢸⡏⢰⡟⢻⡆⢹⡆⠀⠀⠀⠀⡟⢰⡟⠛⣧⢹⡇⠀⠀⢀⡏⠀⠀⠀⠀⠀⠀⠈⠻⠆⠙⢷⡄⠀
⡿⠀⢾⠁⠀⠀⠀⠀⠈⠋⠀⠈⢷⣌⡛⠛⣡⡾⠀⠀⠀⠀⠈⢿⣌⠻⠟⣋⣼⠃⠀⠀⣸⠇⠀⠀⠀⠀⠀⣠⡾⠛⠳⠶⣦⣿⣆
⠛⠛⠳⠶⠶⣤⣄⡀⠀⠀⠀⠀⠀⠉⢙⡋⣉⣀⡀⠀⠀⠀⠀⠀⠉⠛⠛⠋⠁⠀⠀⢠⡟⠀⠀⠀⠀⢀⡼⠋⠀⠀⠀⠀⠀⠈⠁
⠀⠀⠀⠀⠀⠀⠈⠻⣦⡀⠀⠀⠀⣤⠞⠓⣿⠉⠁⠀⠀⠀⠛⢻⡆⠶⣤⣄⠀⠀⠀⠀⠀⠀⠀⣀⡴⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠘⠻⢦⣤⡀⠛⠁⠀⠙⢷⡄⠀⠀⣠⠶⠛⠁⠀⠀⠉⣷⠀⢀⣤⠴⠶⠛⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢷⡀⠀⠀⠀⠘⣧⠀⣰⡏⠀⠀⠀⠀⠀⠰⠏⣠⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢷⡄⠀⣄⠀⠙⠓⠋⠀⠀⠀⠀⡀⠀⠀⣴⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣦⡉⠻⠦⣤⣀⣤⣤⠶⠛⠁⣠⡾⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⢦⣄⣀⠀⠀⣀⣠⡴⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Cyan, Reset)

var accioDescription string = fmt.Sprintf(`%s
%sThe Summoning Charm (Accio) was a charm that caused a target at a distance from the caster to levitate or fly over to them.%s

This spell needs thought behind it; the object must be clear in the caster's mind before trying to summon.

%sThis spell is one of the oldest spells known to wizarding society.%s

%sAccio 'gets' the object and so does 'wget' in the Linux world.%s

Usage : %s$ wand accio LINK%s

%sCheers!%s
`, Cyan, Blue, Reset, Yellow, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var accioCmd = &cobra.Command{
	Use:  "accio",
	Long: fmt.Sprintf("%s %s", accioLogo, accioDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		wand.Runcommand("wget", arg)
	},
}

func init() {
	rootCmd.AddCommand(accioCmd)
}
