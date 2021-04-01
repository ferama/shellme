package utils

import "flag"

// Flags ...
type Flags struct {
	PtyCmd *string
}

var flagValues *Flags

// GetFlags ...
func GetFlags() *Flags {
	if flagValues != nil {
		return flagValues
	}

	flagValues = &Flags{
		PtyCmd: flag.String("ptycmd", "sh", "The command to attach the pty to"),
	}

	flag.Parse()
	return flagValues
}
