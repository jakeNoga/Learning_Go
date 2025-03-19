package main

import (
	"flag"
	"os"
	"reflect"
)

type RescueDebugFlags struct {
	CDM           bool `option:"-cdm" helpText:"Forces the target IOM to grab the profile script from the CDM."`
	DebugProfile  bool `option:"-debug_profile" helpText:"Sends a whack command to start the profile script in debug."`
	DryRun        bool `option:"-dry_run" helpText:"Runs the rescue but prints out the whack commands instead of sending them to the target node."`
	NoReboot      bool `option:"-no_reboot" helpText:"Stops the target node from unmounting the file system and rebooting after the rescue."`
	LeaveCdmFiles bool `option:"-leave_cdm_files" helpText:"Stops the rescuer from removing symlinks to files that are copied to the CDM."`
	NoBootNet     bool `option:"-no_boot_net" helpText:"Prints the last whack command to the console so that it can be copied into the whack command terminal."`
	RmRsync       bool `option:"-rm_rsync" helpText:"Remote rsync forces the target node to connect to the network master simulating a cluster expansion or remote chassis rescue."`
	SkipFw        bool `option:"-skip_fw" helpText:"To skip the initial firmware check and upgrade."`
	Initramfs     bool `option:"-initramfs" helpText:"Sends a whack command to creates a ramfs in the profile script."`
}

func main() {
	var rescueArgs struct {
		ThisRescueFlags RescueDebugFlags
	}

	// Using reflection to iterate through fields of RescueDebugFlags
	fs := flag.NewFlagSet("debug flags", flag.ExitOnError)
	val := reflect.ValueOf(&rescueArgs.ThisRescueFlags).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("option")
		name := tag[1:] // Removing the leading '-'

		if field.Kind() == reflect.Bool {
			fs.BoolVar(field.Addr().Interface().(*bool), name, field.Bool(), "")
		}
	}

	// Parse command-line arguments
	fs.Parse(os.Args[1:])

	// Accessing the values
	println("CDM:", rescueArgs.ThisRescueFlags.CDM)
	println("DebugProfile:", rescueArgs.ThisRescueFlags.DebugProfile)
}
