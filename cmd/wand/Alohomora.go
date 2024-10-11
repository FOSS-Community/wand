package wand

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var alohomoraDescription = fmt.Sprintf(`%s

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣤⣤⣶⣶⣶⣿⣿⣿⣿⣿⣿⣷⣶⣶⣦⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣶⣿⠿⠛⠛⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠛⠻⠿⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⡿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⡿⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⠖⠛⠛⢿⣷⣶⣄⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⢀⣀⣤⣤⣤⣤⣤⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣷⡄⠀⠀⢻⣿⣿⣷⠀⠀⠀⠀⠙⢿⣷⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣴⣿⠟⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⠟⠛⠛⢿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⠿⠟⠁⠀⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠈⠻⣿⣆⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣾⣿⠃⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⢻⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⠘⣿⣧⡀⠀⠀⠀
⠀⠀⢀⣾⡿⠁⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠶⠶⣿⣏⣁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣷⡀⠀⠀
⠀⠀⣾⣿⠃⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣧⠀⠀
⠀⣸⣿⠇⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⢀⣤⣤⣄⠀⠀⠀⢸⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡆⠀
⢀⣿⡟⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⣾⣿⣿⣿⠀⠀⠀⣸⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⠀
⢸⣿⡇⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣷⠀⠀⠀⠀⠘⢿⣿⣁⣀⣀⣴⣿⣿⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡇
⣾⣿⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣧
⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠰⠶⠶⠶⠶⠶⠶⠶⠶⠶⠶⠶⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿
⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣿⣿⣿⣷⣶⣶⡿⠿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿
⢿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠛⠛⠉⠁⠀⠀⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿
⢸⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡇
⠈⣿⣧⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣶⣶⣦⣄⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⢠⣾⡟⢹⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⠁
⠀⢹⣿⡄⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⣠⣿⡟⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡏⠀
⠀⠀⢿⣿⡀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⣾⣿⣿⣿⣿⠟⠀⠀⠀⠀⠀⠀⠀⣴⣿⠏⠀⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⡿⠀⠀
⠀⠀⠈⢿⣷⡀⠀⠀⠀⠀⢻⣿⣿⣿⣿⠋⠀⠀⠀⠀⢀⣼⣿⣿⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⢰⣿⣏⣀⣀⣀⣼⣿⣿⣿⣤⣤⡀⠀⠀⠀⠀⠀⠀⠀⢀⣾⡿⠁⠀⠀
⠀⠀⠀⠈⢿⣿⡄⠀⠀⠀⠀⠙⠻⢿⣿⣷⣶⣶⣶⣾⣿⣿⡿⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⢻⣿⣿⣿⠋⠉⠁⠀⠀⠀⠀⠀⠀⢠⣿⡿⠁⠀⠀⠀
⠀⠀⠀⠀⠀⠻⣿⣦⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⢀⣴⣿⠟⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠛⠿⠿⠿⠿⠿⠿⠿⠃⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣿⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣿⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢿⣿⣶⣦⣤⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣤⣴⣶⣿⠿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠻⠿⠿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠟⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀


Alohomora is used to unlock file or directory permissions, making them writable.

Usage:
$ wand alohomora <file_or_directory>

This command will prompt for confirmation before changing permissions.
`, Yellow)

var alohomoraCmd = &cobra.Command{
	Use:   "alohomora <file_or_directory>",
	Short: "Change file or directory permissions to make them writable",
	Long:  alohomoraDescription,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		// Confirm with the user before changing permissions
		fmt.Printf("Are you sure you want to unlock permissions for '%s'? (y/n): ", target)
		reader := bufio.NewReader(os.Stdin)
		confirmation, _ := reader.ReadString('\n')
		confirmation = strings.TrimSpace(confirmation)

		if strings.ToLower(confirmation) == "y" {
			// Change file/directory permissions to writable using chmod
			fmt.Printf("Unlocking permissions for '%s'...\n", target)
			err := exec.Command("chmod", "+w", target).Run()

			if err != nil {
				fmt.Println("Error changing permissions:", err)
			} else {
				fmt.Println("Permissions changed successfully. The file/directory is now writable.")
			}
		} else {
			fmt.Println("Operation canceled.")
		}
	},
}

func init() {
	// Add the "Alohomora" command to the root command
	rootCmd.AddCommand(alohomoraCmd)
}
