package cmd

import (
	"log"
	"os"
	"time"

	"github.com/latiif/awaitrmq/pkg/amqplookup"
	"github.com/latiif/awaitrmq/pkg/verbose"
	"github.com/spf13/cobra"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

var (
	// flags
	timeoutFlag  string
	intervalFlag string
	verboseFlag  bool
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
	rootCmd.PersistentFlags().StringVarP(&timeoutFlag, "timeout", "t", "0", "Timeout to stop waiting in milliseconds. Pass 0 to timeout in ~ 290 years.")
	rootCmd.PersistentFlags().StringVarP(&intervalFlag, "interval", "i", "2s", "Interval between attempts to check")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "When true, sets output to verbose.")
}

func doAwait(target string) {
	v := verbose.NewVerbose(verboseFlag)
	v.Publish("target", target)
	v.Publish("verbose flag", verboseFlag)
	v.Publish("interval flag", intervalFlag)
	v.Publish("timeout flag", timeoutFlag)

	intervalDuration, err := time.ParseDuration(intervalFlag)
	if err != nil {
		log.Fatalf("value of interval (%s) is invalid.\n", intervalFlag)
	}
	if intervalDuration.Seconds() < 1 {
		log.Printf("value of interval (%s) is too low. Consider increasing it above 1000ms", intervalFlag)
	}
	intervalTicker := time.NewTicker(intervalDuration)
	v.Publish("interval duration", intervalDuration)

	var timeoutTimer *time.Timer
	if timeoutFlag == "0" {
		log.Printf("timeout equals 0. It will time out in ~290 years.")
		// Maximum duration roughly 290 years
		timeoutTimer = time.NewTimer(time.Duration(1<<63 - 1))
	} else {
		timeoutDuration, err := time.ParseDuration(timeoutFlag)
		if err != nil {
			log.Fatalf("value of timeout (%s) is invalid.\n", timeoutFlag)
		}
		timeoutTimer = time.NewTimer(timeoutDuration)
		v.Publish("timeout duration", timeoutDuration)
	}
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-timeoutTimer.C:
				log.Printf("Timed out after %s.\n", timeoutFlag)
				v.Publish("Event", "timed out")
				done <- false
				return
			case <-intervalTicker.C:
				log.Printf("Attempting to find %s...", target)
				found := amqplookup.AMQPLookup(target, intervalDuration)
				log.Printf("Attempt %s.", ifThenElse(found, "succeeded", "failed"))
				if found {
					v.Publish("Event", "succeeded")
					done <- true
					return
				}
			}
		}
	}()
	successful := <-done
	v.Publish("Success", successful)
	if successful {
		v.Publish("Exit code", successExitCode)
		os.Exit(successExitCode)
	} else {
		v.Publish("Exit code", failureExitCode)
		os.Exit(failureExitCode)
	}
}

func ifThenElse(condition bool, then string, otherwise string) string {
	if condition {
		return then
	}
	return otherwise
}
