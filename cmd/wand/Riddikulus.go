package wand

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var riddikulusDescription = fmt.Sprintf(`%s


⢰⣶⣶⠀⠀⠀⢀⣶⣶⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡆⠀⠀⠀⠀
⢸⣿⣿⡆⠀⠀⣾⣿⣿⠀⠀⠀⠀⠀⢀⣶⠶⣶⠀⢀⡶⠶⣶⠀⢸⡇⠀⢀⣄⡀
⢴⣿⡟⣿⡀⣼⠏⣿⣷⠆⣷⠀⢻⡇⣾⡇⠀⣿⠀⢾⡇⠀⣿⠀⢸⡇⢰⣿⣘⣷
⢸⣿⡇⠘⣿⠏⠀⢿⣿⠀⣿⡀⢸⡇⢸⣇⡠⣿⠀⢸⣇⡠⣿⠀⢸⡇⢸⣿⠀⡀
⢸⣿⠀⠀⠈⠀⠀⢸⣿⠀⠻⡧⢻⡇⠀⠉⠀⣿⠀⠀⠉⠀⣿⠀⢸⡇⠀⠙⠞⠁
⠸⡟⠀⠀⠀⠀⠀⠈⡿⠀⠀⠀⠀⠀⢴⡆⣠⠟⠀⢴⡇⣠⠟⠀⠀⠉⠀⠀⠀⠀
⠀⠃⠀⠀⠀⠀⠀⠀⠃⠀⠀⠀⠀⠀⠈⠋⠀⠀⠀⠈⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀


Riddikulus turns system stats into something humorous. This command will display system stats in a fun way, using tools like cowsay, figlet, and fortune.

You can choose different options for output with the --fun option.

Usage:
$ wand riddikulus [--fun figlet|cowsay|fortune]

`, Yellow)

var funOption string

var riddikulusCmd = &cobra.Command{
	Use:   "riddikulus",
	Short: "Display system stats in a humorous way",
	Long:  riddikulusDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if funOption == "" {
			funOption = "cowsay"
		}

		switch funOption {
		case "figlet":
			fmt.Println("Displaying system stats with figlet...")
			err = exec.Command("figlet", "System Stats").Run()
		case "cowsay":
			fmt.Println("Displaying system stats with cowsay...")
			err = exec.Command("fortune", "|", "cowsay").Run()
		case "fortune":
			fmt.Println("Displaying system stats with fortune...")
			err = exec.Command("fortune").Run()
		default:
			fmt.Println("Invalid fun option. Use 'figlet', 'cowsay', or 'fortune'.")
		}

		if err != nil {
			fmt.Println("Error displaying humorous output:", err)
		} else {
			fmt.Println("Humorous system stats displayed successfully.")
		}
	},
}

func init() {
	riddikulusCmd.Flags().StringVarP(&funOption, "fun", "f", "", "Choose your fun output (figlet|cowsay|fortune)")

	rootCmd.AddCommand(riddikulusCmd)
}
