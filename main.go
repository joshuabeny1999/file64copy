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

var version string = "development"

func main() {
	getopt.Alias("h", "help")
	getopt.Alias("v", "version")
	getopt.Alias("c", "clipboard")
	getopt.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		printErrorAndHelp("Please specify a file to encode.")
	}

	fileName := flag.Arg(0)
	encodedContent, err := encodeFile(fileName)
	if err != nil {
		printErrorAndExit("Error reading file:", err)
	}

	command := generateCommand(encodedContent, fileName)

	if *clipFlag {
		err := copyToClipboard(command)
		if err != nil {
			printErrorAndExit("Failed to copy to clipboard:", err)
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
}

func printVersion() {
	fmt.Println("Version:", version)
}

func printErrorAndHelp(message string) {
	fmt.Println("Error:", message)
	printHelp()
	os.Exit(1)
}

func printErrorAndExit(message string, err error) {
	fmt.Println(message, err)
	os.Exit(1)
}

func encodeFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	encodedContent := base64.StdEncoding.EncodeToString(fileContent)
	return encodedContent, nil
}

func generateCommand(encodedContent, fileName string) string {
	return fmt.Sprintf(`echo "%s" | openssl base64 -d -A -out %s`, encodedContent, filepath.Base(fileName))
}

func copyToClipboard(command string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo '%s' | xclip -selection clipboard`, command))
	return cmd.Run()
}
