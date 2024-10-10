package wand

import (
	"fmt"

	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var expectoPatronumLogo string = fmt.Sprintf(`%s
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡿⠛⠛⠻⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⠟⠀⠀⠀⠀⠈⢿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⡇⠀⠀⠀⠀⠀⠀⠸⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⢻⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣷⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣿⣆⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⠀⠀⠀⠀⠀⠀⠀`, Blue)

var expectoPatronumDescription string = fmt.Sprintf(`%s
%sThe Expecto Patronum is a spell used to a spell in the Harry Potter series that summons a guardian to protect the caster from dark creatures.%s


%sIt protects objects, creating a backup of them, and so does the 'tar' command in the Linux terminal.%s %s'tar' creates an archive of files and directories in the Linux world.%s

Usage : %s$ wand expecto-patronum /path/to/source /path/to/archive.tar%s

%sCheers!%s
`, Cyan, Blue, Reset, Yellow, Reset, Green, Reset, Red, Reset, Magenta, Reset)

var expectoCmd = &cobra.Command{
	Use:  "expecto-patronum",
	Long: fmt.Sprintf("%s %s", expectoPatronumLogo, expectoPatronumDescription),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputDirectory := args[0]
		archiveFile := args[1]
		// Run "tar cvf archiveFile -C inputDirectory ."
		wand.Runcommand("tar", "cvf", archiveFile, "-C", inputDirectory, ".")
	},
}

func init() {
	rootCmd.AddCommand(expectoCmd)
}
