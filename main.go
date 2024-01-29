package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	getopt "rsc.io/getopt"
)

var (
	helpFlag    = flag.Bool("help", false, "Display help information")
	versionFlag = flag.Bool("version", false, "Display version information")
	clipFlag    = flag.Bool("clipboard", false, "Copy output to clipboard")
)

const version = "1.0.0"

func main() {
	getopt.Alias("h", "help")
	getopt.Alias("v", "version")
	getopt.Alias("c", "clipboard")
	getopt.Parse()

	if *helpFlag {
		printHelp()
	}

	if *versionFlag {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		fmt.Println("Error: Please specify a file to encode.")
		printHelp()
		os.Exit(1)
	}

	fileName := flag.Arg(0)
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	encodedContent := base64.StdEncoding.EncodeToString(fileContent)
	command := fmt.Sprintf(`echo "%s" | openssl base64 -d -A -out %s`, encodedContent, filepath.Base(fileName))

	if *clipFlag {
		cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo '%s' | xclip -selection clipboard`, command))
		err := cmd.Run()
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
		}
	} else {
		fmt.Println(command)
	}
}

func printHelp() {
	fmt.Println("Usage: file64copy [OPTIONS] FILE")
	fmt.Println("Encode a file to Base64 and generate a command to decode it.")
	fmt.Println("\nOptions:")
	getopt.PrintDefaults()
	os.Exit(0)
}
