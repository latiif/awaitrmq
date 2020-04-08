package cmd

import "github.com/spf13/cobra"

import "fmt"

var (
	timeoutFlag   float64
	intervalFlag  float64
	dnslookupFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "awaitrmq",
	Short: "Awaits for a RabbitMQ Service",
	Long:  "Smartly awaits for a RabbitMQ Service to actually be running",
	Run: func(cmd *cobra.Command, args []string) {
		doAwait()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Float64VarP(&timeoutFlag, "timeout", "t", 0, "Timeout to stop waiting in milliseconds. Pass 0 to never timeout.")
	rootCmd.PersistentFlags().Float64VarP(&intervalFlag, "interval", "i", 2000, "Interval between attempts to check in milliseconds")
	rootCmd.PersistentFlags().BoolVarP(&dnslookupFlag, "dnslookup", "l", true, "When true, awaits a succesful dnslookup before proceeding.")
}

func doAwait() {
	//TODO Implement logic
	fmt.Println("TODO")
}
