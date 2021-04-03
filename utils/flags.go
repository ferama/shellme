package utils

import "flag"

// Flags ...
type Flags struct {
	PtyCmd *string
	Debug  *bool
}

var flagValues *Flags

// GetFlags ...
func GetFlags() *Flags {
	if flagValues != nil {
		return flagValues
	}

	flagValues = &Flags{
		PtyCmd: flag.String("ptycmd", "sh", "The command to attach the pty to"),
		Debug:  flag.Bool("debug", false, "Enable debug mode"),
	}

	flag.Parse()
	return flagValues
}
