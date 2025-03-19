package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	node := flag.String("n", "", "Node ID of the target node to be rescued.")
	cage := flag.String("c", "", "Cage ID of the target IOM to rescue.")
	bay := flag.String("b", "", "Bay of the target IOM to rescue.")
	chassis := flag.String("chassis", "", "Chassis assembly serial number of the target IOM to rescue.")
	helpFlag := flag.Bool("h", false, "Print help text.")
	clusterExp := flag.Bool("ex", false, "Cluster expansion expands the nodes in the cluster.")
	// Define optional flags
	user := flag.String("User", "", "Specify the user that the task will be tracked with.")
	task := flag.Bool("task", false, "Track the progress of process in the task structure.")
	initramfs := flag.Bool("initramfs", false, "Creates a ramfs in the profile script.")
	// Define debug flags
	cdm := flag.Bool("cdm", false, "Forces the profile script to come from the CDM.")
	debugProfile := flag.Bool("debug_profile", false, "Starts the profile script in debug.")
	dryrun := flag.Bool("dryrun", false, "Runs the rescue but prints out the whack commands instead of sending them.")
	leaveCdmFiles := flag.Bool("leave_cdm_files", false, "Stops the rescuer from removing symlinks to files that are copied to the CDM.")
	noBootNet := flag.Bool("no_boot_net", false, "Prints the last whack command to the console so that it can be copied into the whack command terminal.")
	noreboot := flag.Bool("noreboot", false, "Stops the target node from unmounting the file system and rebooting after the rescue.")
	rmRsync := flag.Bool("rm_rsync", false, "Remote rsync forces the target node to connect to the network master simulating a cluster expansion or remote chassis rescue.")
	skipFw := flag.Bool("skip_fw", false, "To skip the initial firmware check and upgrade.")

	// Parse command-line flags
	flag.Parse()

	// Handle help flag
	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	// Here you can perform additional validation or processing if needed
	// For example, checking if required flags are present

	// Example validation (you can customize as per your needs)
	if *node == "" && *cage == "" && *chassis == "" {
		fmt.Println("Error: Please specify at least one of -n, -c, or -chassis.")
		os.Exit(1)
	}

	// Print out the parsed arguments (for testing purposes)
	fmt.Println("Parsed arguments:")
	fmt.Println("-n:", *node)
	fmt.Println("-c:", *cage)
	fmt.Println("-b:", *bay)
	fmt.Println("-chassis:", *chassis)
	fmt.Println("-ex:", *clusterExp)
	fmt.Println("-User:", *user)
	fmt.Println("-task:", *task)
	fmt.Println("-initramfs:", *initramfs)
	fmt.Println("-cdm:", *cdm)
	fmt.Println("-debug_profile:", *debugProfile)
	fmt.Println("-dryrun:", *dryrun)
	fmt.Println("-leave_cdm_files:", *leaveCdmFiles)
	fmt.Println("-no_boot_net:", *noBootNet)
	fmt.Println("-noreboot:", *noreboot)
	fmt.Println("-rm_rsync:", *rmRsync)
	fmt.Println("-skip_fw:", *skipFw)
}
