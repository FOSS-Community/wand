package wand

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var expectoPatronumDescription = fmt.Sprintf(`%s
⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣤⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠰
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⠀⣠⡿⠋⠁⡀⠀⠙⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠂⠀⠀⠀⠀⠀⢀⣴⠟⠀⠀⠀⣧⠀⠀⠀⠙⢷⣄⡀⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⣤⡶⠟⠁⣾⡏⢣⡀⢻⣦⣀⠀⠀⠀⠉⠛⠶⣄⠀⠀⣤⡘⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡿⠋⠉⠁⠀⠀⠀⠘⣧⡀⠙⢬⡻⣿⣷⡀⢐⣶⣦⡀⠸⣇⠀⠌⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠙⠲⢦⣄⣸⡟⠉⠉⠉⠙⢻⣆⢻⣮⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣼⠿⡄⠈⠀⠀⠀⣀⣀⣀⣀⣀⠀⠀⡀⢻⡆⠀⠀⠤⠀⠀⠀⠙⢷⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⢀⣴⡿⠋⠁⠈⠀⠀⠀⢀⣾⣫⣥⡤⢤⣤⣄⠀⠹⣼⣧⠀⠀⠀⠀⠀⠀⠀⠠⠉⠙⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣤⣼⣯⠉⠀⠀⠀⠀⠀⢀⣼⠟⠋⢀⣠⣆⣀⠉⢷⣆⣾⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⢏⠉⠉⠻⣟⠒⠀⠀⣠⣴⠾⠛⠁⣠⣠⣿⣿⣿⣿⣿⣆⠹⡎⢿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⠀⠀⠀⠀⠀⠉⠀⠀⠠⢸⣏⣾⣷⣄⠀⠈⠑⢠⡾⠋⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⡏⡦⠙⠈⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣻⣿⣮⡀⠀⣾⡇⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⢇⡀⠀⢸⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡼⣷⣿⣿⣷⣄⣿⢧⣤⠿⣿⣿⣿⣿⣛⣛⣯⣿⣿⢟⡋⠁⠀⠘⢷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠐⠒⠪⠽⠉⠉⠈⠁⠀⠈⠉⠛⠮⠭⠉⠉⠀⠀⠀⡅⠀⠆⠀⠈⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣷⣶⠾⠁⠀⠀⠀⣰⣄⣠⣠⡤⣀⡀⠈⠠⠀⠀⠒⣒⠒⠦⢤⣄⣄⣻⡇⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⢃⣴⣶⣾⣷⣾⣿⣿⣿⣿⣿⣿⣿⣾⣷⣶⣾⣿⠿⠿⠛⠛⠷⣬⡙⢻⣿⣄⠀⠀⠀⠀⠀⠈⠀⠀⠀⠄⠀⠀⠀⠀⠀⠀⠐⠀
⠀⠀⠀⠀⠀⢸⡿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⡎⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠛⠋⢉⣁⡀⠀⠀⠀⠀⠒⠛⢋⣫⠽⠟⠛⢛⣛⣛⣳⣶⢾⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⢷⡘⣧⣀⣀⡀⣄⠀⠀⠀⣺⡾⠋⠀⠉⢋⣙⡻⡿⠟⠛⣛⣉⠥⠂⢀⣠⡾⠿⠶⠀⠀⠀⠀⣀⠤⢚⣩⣤⠴⠾⠛⠛⠋⠉⠙⠛⠻⠿⣝⣷⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠈⠳⣌⠻⣯⣻⡟⢷⡶⣿⠋⢳⡤⠶⠞⠿⣍⡛⠛⠲⠶⠤⣤⣶⠞⠋⠀⠀⠀⢀⣀⠴⢀⣨⡴⠞⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⢿⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣟⠳⣬⡻⣿⣆⠹⠄⠀⠀⠆⠀⠀⠀⠀⠉⠑⠀⠔⠚⡏⠀⠀⠀⠀⠤⠒⣉⣤⠶⠿⠛⠚⠓⠒⠶⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⡿⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠐⠂⠀⠀⣤⡌⠻⣮⡙⠿⣦⣄⣀⡀⠀⠀⠀⢰⠀⠀⠀⠀⠀⠸⠀⠀⢀⣤⠶⠋⠁⠀⠀⠀⠀⢀⢀⣀⣠⣤⣤⣤⣤⣄⣀⣀⣤⣤⠶⠞⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣄⠀⠀⠀⠈⠻⢦⣀⡉⠉⠻⠙⠁⠤⠬⠀⠒⠀⣀⣀⣤⡤⢾⠛⠳⠖⠒⢒⠶⠒⠒⠚⠛⠋⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠈⠀⠤⠀⠉⠛⠛⠲⠶⠾⠶⠶⠾⠛⠛⠉⠁⠀⠀⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
Expecto Patronum summons a Patronus for protection by creating a backup of your important files.
You can specify a destination for the backup, or use the default location ~/backup.

Usage:
$ wand expecto-patronum /path/to/important/files --destination /path/to/backup/location
`, Yellow)

// defaultBackupLocation is the default directory for storing backups if none is provided
var defaultBackupLocation = filepath.Join(os.Getenv("HOME"), "backup")

// expectoPatronumCmd represents the "Expecto Patronum" command
var expectoPatronumCmd = &cobra.Command{
	Use:   "expecto-patronum",
	Short: "Summon a backup of important files",
	Long:  expectoPatronumDescription,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Source file to be backed up
		source := args[0]

		// Get the backup destination from the flag or use the default location
		destination, _ := cmd.Flags().GetString("destination")
		if destination == "" {
			destination = defaultBackupLocation
		}

		// Ensure the destination directory exists
		if _, err := os.Stat(destination); os.IsNotExist(err) {
			fmt.Printf("Creating backup directory at %s\n", destination)
			err := os.MkdirAll(destination, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating backup directory:", err)
				return
			}
		}

		// Perform the file copy operation
		backupFile(source, destination)
	},
}

// backupFile copies the source file to the destination directory using the "cp" command
func backupFile(source, destination string) {
	// Get the filename from the source path
	fileName := filepath.Base(source)

	// Full destination path for the file
	destPath := filepath.Join(destination, fileName)

	// Run the "cp" command to copy the file
	fmt.Printf("Backing up %s to %s\n", source, destination)
	wand.Runcommand("cp", source, destPath)

	// Optional: Add additional checks to ensure the copy was successful
	_, err := os.Stat(destPath)
	if os.IsNotExist(err) {
		fmt.Println("Backup failed: file not copied.")
	} else {
		fmt.Println("Backup successful!")
	}
}

func init() {
	// Add the "Expecto Patronum" command to the root command
	rootCmd.AddCommand(expectoPatronumCmd)

	// Define a flag for specifying the backup destination
	expectoPatronumCmd.Flags().StringP("destination", "d", "", "Specify the backup destination")
}
