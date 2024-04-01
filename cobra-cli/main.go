package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			times, _ := cmd.Flags().GetInt("times")
			if times > 100 {
				return errors.New("times is too high")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			times, _ := cmd.Flags().GetInt("times")
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdEcho.Flags().IntP("times", "t", 1, "times to echo the input")
	// cmdEcho.MarkFlagRequired("times")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	rootCmd.Execute()
}
