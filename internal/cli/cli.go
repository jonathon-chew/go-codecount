package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	aphrodite "github.com/jonathon-chew/Aphrodite"
)

type Flags struct {
	IgnoreFolders []string
	IgnoreFiles   []string
	IncludeFiles  []string
	Exclusive     []string
	IgnoreError   bool
}

// This is the function to manage command line arguments
func Cli(Arguments []string) Flags {

	var FlagArguments Flags

	for numberOfArguments := 0; numberOfArguments < len(Arguments); numberOfArguments++ {
		flag := Arguments[numberOfArguments]

		switch flag {

		case "--version", "-v":
			versionNumber := "v0.5.1"
			fmt.Printf("Version %s", versionNumber)
			os.Exit(0)

		case "--help", "-h":

			aphrodite.PrintBold("Cyan", "Version\n")
			aphrodite.PrintInfo("--version " + "-v\n\n")

			aphrodite.PrintBold("Cyan", "Help\n")
			aphrodite.PrintInfo("--help " + "-h\n\n")

			aphrodite.PrintBold("Cyan", "Ignore a folder\n")
			aphrodite.PrintInfo("--ignore " + "-i \n")
			fmt.Print("Followed by the name to exclude - this can be a partial as long as it's the last part of the folder name\neg. .git will ignore any folders called .git, it would also work\n\n")

			aphrodite.PrintBold("Cyan", "Ignore a file\n")
			aphrodite.PrintInfo("--ignore-file " + "-if\n")
			fmt.Print("Followed by the name to exclude - this can be a partial as long as it's the last part of the file name\neg. README will ignore any file called README, READ would also work, this also works/applies with file extensions\n\n")

			aphrodite.PrintBold("Cyan", "Include a file\n")
			aphrodite.PrintInfo("--include " + "-in\n")
			fmt.Print("Followed by the name to include - this can be a partial as long as it's the last part of the file name\neg. README will ignore any file called README, READ would also work, this also works/applies with file extensions\n\n")

			aphrodite.PrintBold("Cyan", "Exclusive a file name\n")
			aphrodite.PrintInfo("--exclusive " + "-e\n")
			fmt.Print("Followed by the name to include - this can be a partial as long as it's the last part of the file name\neg. README will ignore any file called README, READ would also work, this also works/applies with file extensions.\nIf you're looking for file types include the . eg. .go for all go files\n\n")

			/* 			aphrodite.PrintBold("Cyan", "Include file\n")
			   			aphrodite.PrintInfo("--include" + "-in\n")
			   			fmt.Print("Followed by the exact name to include - this can be a partial as long as it's the last part of the file name\neg. README will include any file called README, READ would also work, this also works/applies with file extensions\n\n")
			*/
			os.Exit(0)

		case "--ignore", "-i":
			if numberOfArguments+1 >= len(Arguments) {
				log.Print("[ERROR]: no file found after -i flag")
				return FlagArguments
			}

			for i := numberOfArguments + 1; i < len(Arguments); i++ {
				if !strings.HasPrefix(Arguments[i], "-") {
					FlagArguments.IgnoreFolders = append(
						FlagArguments.IgnoreFolders,
						Arguments[i],
					)
					numberOfArguments++
				} else {
					break
				}
			}

		case "--include", "-in":
			if numberOfArguments+1 >= len(Arguments) {
				log.Print("[ERROR]: no file found after -in flag")
				return FlagArguments
			}

			for i := numberOfArguments + 1; i < len(Arguments); i++ {
				if !strings.HasPrefix(Arguments[i], "-") {
					FlagArguments.IncludeFiles = append(
						FlagArguments.IncludeFiles,
						Arguments[i],
					)
					numberOfArguments++
				} else {
					break
				}
			}

		case "--ignore-file", "-if":
			if numberOfArguments+1 >= len(Arguments) {
				log.Print("[ERROR]: no file found after -if flag")
				return FlagArguments
			}

			for i := numberOfArguments + 1; i < len(Arguments); i++ {
				if !strings.HasPrefix(Arguments[i], "-") {
					FlagArguments.IgnoreFiles = append(
						FlagArguments.IgnoreFiles,
						Arguments[i],
					)
					numberOfArguments++
				} else {
					break
				}
			}

		case "--exclusive", "-e":
			if numberOfArguments+1 >= len(Arguments) {
				log.Print("[ERROR]: no file found after -if flag")
				return FlagArguments
			}

			for i := numberOfArguments + 1; i < len(Arguments); i++ {
				if !strings.HasPrefix(Arguments[i], "-") {
					FlagArguments.Exclusive = append(
						FlagArguments.Exclusive,
						Arguments[i],
					)
					numberOfArguments++
				} else {
					break
				}
			}

		case "--ignore-error", "-ie":
			FlagArguments.IgnoreError = true

		default:
			log.Println("Unable to deal with argument:", flag)
		}
	}

	return FlagArguments
}
