package cmd

import (
	"log"
	"os"
	"time"

	"github.com/latiif/awaitrmq/pkg/dnslookup"
	"github.com/spf13/cobra"
)

var (
	// flags
	timeoutFlag   int64
	intervalFlag  int64
	dnslookupFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "awaitrmq",
	Short: "Awaits for a RabbitMQ Service",
	Long:  "Smartly awaits for a RabbitMQ Service to actually be running",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("%v is not valid. Exactly 1 argument is expected.\n", args)
		}
		doAwait(args[0])
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Int64VarP(&timeoutFlag, "timeout", "t", 0, "Timeout to stop waiting in milliseconds. Pass 0 to timeout in ~ 290 years.")
	rootCmd.PersistentFlags().Int64VarP(&intervalFlag, "interval", "i", 2000, "Interval between attempts to check in milliseconds")
	rootCmd.PersistentFlags().BoolVarP(&dnslookupFlag, "dnslookup", "l", true, "When true, awaits a succesful dnslookup before proceeding.")
}

func doAwait(target string) {
	if timeoutFlag < 0 {
		log.Fatalf("value of timeout (%d) is negative.\n", timeoutFlag)
	}
	if intervalFlag < 0 {
		log.Fatalf("value of interval (%d) is negative.\n", intervalFlag)
	}

	if intervalFlag < 1000 {
		log.Printf("value of interval (%d) is too low. Consider increasing it above 1000ms", intervalFlag)
	}

	intervalTicker := time.NewTicker(time.Duration(intervalFlag) * time.Millisecond)
	var timeoutTimer *time.Timer
	if timeoutFlag == 0 {
		log.Printf("timeout equals 0. It will time out in ~290 years.")
		// Maximum duration roughly 290 years
		timeoutTimer = time.NewTimer(time.Duration(1<<63 - 1))
	} else {
		timeoutTimer = time.NewTimer(time.Duration(timeoutFlag) * time.Millisecond)
	}
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-timeoutTimer.C:
				log.Fatalf("Timed out after %d ms.\n", timeoutFlag)
				done <- false
				return
			case <-intervalTicker.C:
				log.Printf("Attempting to find %s...", target)
				found := dnslookup.DNSLookup(target)
				if found {
					done <- true
					return
				}
			}
		}
	}()
	succesful := <-done
	if succesful {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
