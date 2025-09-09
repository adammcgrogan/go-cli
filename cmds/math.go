package cmds

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var mathCmd = &cobra.Command{
	Use:   "math",
	Short: "Perform basic math operations",
	Long:  `A parent command for various math-related subcommands like add, subtract, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a math subcommand (e.g., 'add').")
	},
}

var addCmd = &cobra.Command{
	Use:   "add [numbers to add]",
	Short: "Adds a series of numbers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: '%s' is not a valid integer.\n", arg)
				return
			}
			sum += num
		}
		fmt.Printf("%d\n", sum)
	},
}

var subtractCmd = &cobra.Command{
	Use:   "subtract [initial number] [numbers to subtract...]",
	Short: "Subtracts a series of numbers from the first number",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initialNum, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid integer.\n", args[0])
			return
		}
		if len(args) == 1 {
			fmt.Printf("The result is: %d\n", initialNum)
			return
		}
		for _, arg := range args[1:] {
			num, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: '%s' is not a valid integer.\n", arg)
				return
			}
			initialNum -= num
		}
		fmt.Printf("%d\n", initialNum)
	},
}

var multiplyCmd = &cobra.Command{
	Use:   "multiply [numbers to multiply]",
	Short: "Multiplies a series of numbers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		product := 1
		for _, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: '%s' is not a valid integer.\n", arg)
				return
			}
			product *= num
		}
		fmt.Printf("%d\n", product)
	},
}

var divideCmd = &cobra.Command{
	Use:   "divide [initial number] [numbers to divide by...]",
	Short: "Divides the first number by a series of numbers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quotient, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid number.\n", args[0])
			return
		}
		if len(args) == 1 {
			fmt.Printf("The result is: %f\n", quotient)
			return
		}
		for _, arg := range args[1:] {
			divisor, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Error: '%s' is not a valid number.\n", arg)
				return
			}
			if divisor == 0 {
				fmt.Println("Error: Cannot divide by zero.")
				return // Stop execution immediately.
			}
			quotient /= divisor
		}
		fmt.Printf("%f\n", quotient)
	},
}

func init() {
	mathCmd.AddCommand(addCmd)
	mathCmd.AddCommand(subtractCmd)
	mathCmd.AddCommand(multiplyCmd)
	mathCmd.AddCommand(divideCmd)
	rootCmd.AddCommand(mathCmd)
}
