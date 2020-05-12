package verbose

import "log"

// Verbose logs events in the program
type Verbose struct {
	active bool
}

// NewVerbose instantiates a new Verbose
func NewVerbose(shouldPublish bool) *Verbose {
	return &Verbose{active: shouldPublish}
}

// Publish logs the event if Verbose is active
func (v *Verbose) Publish(entry string, value interface{}) {
	if v.active {
		log.Printf("[INFO] %s: %v\n", entry, value)
	}
}
