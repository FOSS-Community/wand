package wand

import (
	"fmt"

	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var expelliarmusLogo string = fmt.Sprintf(`%s 
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⠖⠉⠁⠀⠀⠀⠈⠉⠲⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣀⣀⡴⠡⠤⠤⣝⡒⠒⣚⡥⠤⠤⠀⠹⠤⠤⠤⢤⣀⡀⠀⠀⠀⠀
⠀⡠⢚⣩⠤⡄⠀⣠⢺⣓⢦⠈⠉⠀⡴⣺⡒⣄⠀⠀⡔⠋⠑⠲⢭⡲⣄⠀⠀
⡼⢱⠋⠀⠀⠱⠂⢣⡻⠼⡸⠀⠀⠀⢇⠧⢝⡼⠀⢠⠃⠀⠀⢀⡤⠥⢬⣣⡀
⠉⠙⠒⠲⢤⡀⠀⠀⢉⣍⡤⠄⠀⢀⣌⣉⠉⠀⠀⠞⠀⠀⣠⠎⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠑⠤⡄⠛⠈⢳⡀⠀⡴⠚⠈⠙⡆⢠⠤⠴⠊⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⣄⢀⠀⠧⠜⠀⠀⠀⠈⣡⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠘⢦⡑⠦⠤⠤⠒⢁⠜⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠒⠒⠒⠚⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Cyan, Reset)

var expelliarmusDescription string = fmt.Sprintf(`%s
%sThe Expelliarmus spell disarms an opponent, making them drop whatever they are holding.%s

%sIn the "wand" library, this command will correspond to killing a process in Linux. :)%s

Usage : %s$ wand expelliarmus <process_id>%s

%sCheers!%s
`, Cyan, Magenta, Reset, Yellow, Reset, Green, Reset, Red, Reset)

var expelliarmusCmd = &cobra.Command{
	Use:  "expelliarmus",
	Long: fmt.Sprintf("%s %s", expelliarmusLogo, expelliarmusDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		// List of critical processes that should not be killed
		criticalProcesses := []string{
			"init",
			"systemd",
			"cron",    
			"sshd",
		}
		for _, criticalProcess := range criticalProcesses {
			if arg == criticalProcess {
				fmt.Printf("Error: Attempting to kill critical system process: %s\n", criticalProcess)
				return
			}
		}
		wand.Runcommand("kill", arg)
	},
}

func init() {
	rootCmd.AddCommand(expelliarmusCmd)
}
