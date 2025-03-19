package main

import (
	"fmt"
	"reflect"
)

// RescueOptions represents various rescue options.
type RescueOptions struct {
	CDM           bool `comment:"Forces the target IOM to grab the profile script from the CDM."`
	DebugProfile  bool `comment:"Sends a whack command to start the profile script in debug."`
	DryRun        bool `comment:"Runs the rescue but prints out the whack commands instead of sending them to the target node."`
	NoReboot      bool `comment:"Stops the target node from unmounting the file system and rebooting after the rescue."`
	LeaveCdmFiles bool `comment:"Stops the rescuer from removing symlinks to files that are copied to the CDM."`
	NoBootNet     bool `comment:"Prints the last whack command to the console so that it can be copied into the whack command terminal."`
	RmRsync       bool `comment:"Remote rsync forces the target node to connect to the network master simulating a cluster expansion or remote chassis rescue."`
	SkipFw        bool `comment:"To skip the initial firmware check and upgrade."`
	Initramfs     bool `comment:"Creates a ramfs in the profile script."`
}

// StringWithComments returns a string representation of RescueOptions,
// including comments for each boolean field.
func (r RescueOptions) StringWithComments() string {
	var result string

	// Use reflection to iterate over struct fields
	rt := reflect.TypeOf(r)
	rv := reflect.ValueOf(r)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i).Bool()

		comment, ok := field.Tag.Lookup("comment")
		if !ok {
			comment = ""
		}

		result += fmt.Sprintf("%s: %t  // %s\n", field.Name, value, comment)
	}

	return result
}

func main() {
	// Example usage:
	options := RescueOptions{
		CDM:          true,
		DebugProfile: false,
		DryRun:       true,
	}

	// Print the struct with comments
	fmt.Println(options.StringWithComments())
}
