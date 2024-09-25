package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var legilimensLogo string = fmt.Sprintf(`%s 
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⠙⠛⣠⣍⡻⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣃⣒⣾⣿⣿⣿⣶⡿⠿⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⡿⠿⠿⢿⣿⣯⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣷⣶⣶⣶⣽⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣾⣿⣿⣟⠿⠿⠯⠽⠗⣶⣹⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣷⣽⣛⣛⣶⣶⠾⣛⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⡿⢽⣛⡛⠋⠉⣀⣼⣿⣿⣟⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣷⣦⣷⣶⣾⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⡿⢿⠿⠿⡛⠉⠙⣿⣿⣿⣿⠛⠫⣿⣼⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣶⣶⣶⣶⣿⣿⣿⣿⣷⣮⣭⣉⠉⠉⣛⣻⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⡙⠻⠿⠿⠿⠿⠿⠿⠿⣿⣿⣿⣿⣿⣷⡿⠿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣀⣀⣀⣀⣀⣀⡀⠀⣰⣿⣿⡿⣿⣦⣄⣀⡀⠀⠐⠒⠑⠒⠲⠉⢛⣚⣉⣥⠶⣿⣷⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢀⣴⣶⠛⠉⠉⠁⢀⡈⠻⣯⡉⠙⣿⣿⣿⣿⣤⣉⠉⠉⠉⢙⠛⠛⠻⠯⠭⠽⣻⠉⢃⣐⣤⠾⣿⣧⣀⣀⣀⣤⣤⣤⣤⣄⣀⠀⠀⠀
⠈⢿⣿⣧⡄⠀⠀⠀⠈⠓⠬⠿⠆⠳⠌⠙⠛⠿⢿⠛⢶⣶⣤⣤⣤⣤⣴⣶⣛⣉⣤⣤⡴⠶⠟⠻⠯⠽⢿⣿⠟⠛⠛⠛⠿⢿⣿⣦⣄
⠀⠀⠙⠿⣧⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠂⢀⠄⠀⠀⠉⠀⠀⠀⠀⠉⠉⠉⣁⡤⣀⡀⠀⠀⠀⢀⡴⠟⠁⡤⣬⣁⡀⠒⢲⣻⡿⠋
⠀⠀⠀⠀⠈⠙⠻⣶⣠⣀⠀⠀⠀⣠⣤⣶⠞⠯⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠁⠁⠈⢳⡀⠀⠚⠉⠀⠀⠀⢀⣽⣿⣯⡝⠟⠋⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠾⢿⢿⣿⡟⢀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠆⠀⠀⣠⣤⡤⠶⠛⠋⠉⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠛⠒⠶⠶⠶⠶⣶⣦⣴⣶⠦⠤⠶⠶⠒⠒⠚⠋⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Magenta, Reset)

var legilimensDescription string = fmt.Sprintf(`%s
%sLegilemency or Legilimens is the act of magically navigating through many layers of a person's mind and correctly interpreting one's findings.%s

Muggles often call this 'mind-reading'.

%sIt reads what is beneath the flesh of the person in the wizarding world, and similarly 'cat' does the same in the Linux world.%s %s'cat' reads the contents hidden beneath the flesh of a file. :P%s

Usage : %s$ wand legilimens filename%s

%sCheers!%s
`, Cyan, Magenta, Reset, Yellow, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var legilimensCmd = &cobra.Command{
	Use:  "legilimens",
	Long: fmt.Sprintf("%s %s", legilimensLogo, legilimensDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		wand.Runcommand("cat", arg)
	},
}

func init() {
	rootCmd.AddCommand(legilimensCmd)
}
