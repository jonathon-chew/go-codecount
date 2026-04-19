package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/jonathon-chew/go-codecount/internal/cli"
	"github.com/jonathon-chew/go-codecount/internal/utils"

	Aphrodite "github.com/jonathon-chew/Aphrodite"
)

type LanguageStats struct {
	Files         int
	Lines         int
	NonEmptyLines int
	Words         int
}

type FileResult struct {
	Language      string
	Lines         int
	NonEmptyLines int
	Words         int
}

type LanguageStatsMap map[string]*LanguageStats

const (
	LangGoLang     = "Golang"
	LangLua        = "Lua"
	LangHaskell    = "Haskell"
	LangPerl       = "Perl"
	LangDart       = "Dart"
	LangObjectiveC = "Objective-C"
	LangCCS        = "CCS"
	LangJava       = "Java"
	LangPython     = "Python"
	LangShell      = "Shell"
	LangCSharp     = "C#"
	LangSQL        = "SQL"
	LangScala      = "Scala"
	LangTypeScript = "TypeScript"
	LangPowershell = "Powershell"
	LangJulia      = "Julia"
	LangPHP        = "PHP"
	LangSwift      = "Swift"
	LangQML        = "QML"
	LangHTML       = "HTML"
	LangC          = "C"
	LangKotlin     = "Kotlin"
	LangJavaScript = "JavaScript"
	LangCPlus      = "C++"
	LangRuby       = "Ruby"
	LangMarkdown   = "Markdown"
	LangR          = "R"
	LangRust       = "Rust"
	LangZShell     = "Z Shell"
	LangJson       = "Json"
)

var extToLang = map[string]string{
	"py":    LangPython,
	"js":    LangJavaScript,
	"java":  LangJava,
	"go":    LangGoLang,
	"rs":    LangRust,
	"cpp":   LangCPlus,
	"cc":    LangCPlus,
	"cxx":   LangCPlus,
	"c":     LangC,
	"h":     LangC,
	"cs":    LangCSharp,
	"php":   LangPHP,
	"rb":    LangRuby,
	"ts":    LangTypeScript,
	"tsx":   LangTypeScript,
	"swift": LangSwift,
	"kt":    LangKotlin,
	"scala": LangScala,
	"r":     LangR,
	"dart":  LangDart,
	"hs":    LangHaskell,
	"m":     LangObjectiveC,
	"qml":   LangQML,
	"jl":    LangJulia,
	"sh":    LangShell,
	"pl":    LangPerl,
	"lua":   LangLua,
	"sql":   LangSQL,
	"mod":   LangGoLang,
	"sum":   LangGoLang,
	"html":  LangHTML,
	"ccs":   LangHTML,
	"ps1":   LangPowershell,
	"psm1":  LangPowershell,
	"psd1":  LangPowershell,
	"md":    LangMarkdown,
	"Md":    LangMarkdown,
	"zsh":   LangZShell,
	"json":  LangJson,
}

/*
Convert a int into a string, but make it human readbale by working backwards and applying commas in the right place to split up the number
*/
func HumanReadableInt(initalInt int) string {
	convertedNumber := strconv.Itoa(initalInt)
	var humanReadbleNumber string
	count := 0

	if len(convertedNumber) <= 3 {
		return convertedNumber
	}

	for i := len(convertedNumber) - 1; i >= 0; i-- {
		humanReadbleNumber = string(convertedNumber[i]) + humanReadbleNumber
		count++
		if count%3 == 0 && i != 0 {
			humanReadbleNumber = "," + humanReadbleNumber
		}
	}

	return humanReadbleNumber
}

/*
Loop through Language stats and if it exists add to it, else add it on
*/
func addToList(stats LanguageStatsMap, lines, nonEmptyLines, words int, fileExtension string) {

	if s, ok := stats[fileExtension]; ok {
		s.Files++
		s.Lines += lines
		s.NonEmptyLines += nonEmptyLines
		s.Words += words
		return
	}

	stats[fileExtension] = &LanguageStats{
		Files:         1,
		Lines:         lines,
		NonEmptyLines: nonEmptyLines,
		Words:         words,
	}

}

func countWords(b []byte) int {
	inWord := false
	count := 0

	for _, c := range b {
		if (c >= 'A' && c <= 'Z') ||
			(c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') ||
			c == '_' {

			if !inWord {
				count++
				inWord = true
			}
		} else {
			inWord = false
		}
	}
	return count
}

func worker(fileChan <-chan string, resultChan chan<- FileResult, cliFlags cli.Flags) {
	for path := range fileChan {

		// Open file
		f, err := os.Open(path)
		if err != nil {
			if !cliFlags.IgnoreError {
				fmt.Println("error opening:", path)
			}
			continue
		}

		var lines, nonEmptyLines, words int
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Bytes()

			if len(line) > 0 {
				nonEmptyLines++
				words += countWords(line)
			}
			lines++
		}

		f.Close() // ⚠️ no defer in worker loop

		// Extension logic
		ext := filepath.Ext(path)
		if strings.HasPrefix(ext, ".") {
			ext = ext[1:]
		} else {
			ext = "Other"
		}

		lang, ok := extToLang[ext]
		if !ok {
			lang = ext
		}

		resultChan <- FileResult{
			Language:      lang,
			Lines:         lines,
			NonEmptyLines: nonEmptyLines,
			Words:         words,
		}
	}
}

func main() {

	var cliFlags cli.Flags
	if len(os.Args[1:]) >= 1 {
		cliFlags = cli.Cli(os.Args[1:])
	}

	root := "./"
	stats := make(LanguageStatsMap)
	var finalPrint []string
	var totalLines, totalNonEmptyLines, totalFiles, totalWords, biggestLangLength, biggestNumberOfFilesLength, biggestNumberOfNonEmptyLinesLength int
	var biggestNumberOfWordsLength int = len("No. words:")

	fileChan := make(chan string, 100)
	resultChan := make(chan FileResult, 100)

	numWorkers := runtime.NumCPU() * 2 // or runtime.NumCPU()*2

	for i := 0; i < numWorkers; i++ {
		go worker(fileChan, resultChan, cliFlags)
	}

	go func() {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err // stop on error
			}

			for _, ignoreFolder := range cliFlags.IgnoreFolders {
				if strings.Contains(path, ignoreFolder+"/") || strings.Contains(path, ignoreFolder+"\\") {
					return nil
				}
			}

			for _, ignoreFile := range cliFlags.IgnoreFiles {
				if strings.Contains(d.Name(), ignoreFile) {
					return nil
				}
			}

			var allowed bool
			if len(cliFlags.Exclusive) > 0 {
				for _, exclusiveFileType := range cliFlags.Exclusive {
					if strings.Contains(d.Name(), exclusiveFileType) {
						allowed = true
						break
					}
				}

				if !allowed {
					return nil
				}
			}

			fileChan <- path
			return nil
		})

		close(fileChan) // 🚨 important
	}()

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Go(func() {
			worker(fileChan, resultChan, cliFlags)
		})
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		addToList(stats, result.Lines, result.NonEmptyLines, result.Words, result.Language)
	}

	biggestNumberOfFilesLength = len(HumanReadableInt(biggestNumberOfFilesLength))
	biggestNumberOfWordsLength = len(HumanReadableInt(biggestNumberOfWordsLength))
	biggestNumberOfNonEmptyLinesLength = len(HumanReadableInt(biggestNumberOfNonEmptyLinesLength))

	sentence := fmt.Sprintf("%%-%ds %%-%ds %%-%ds %%-%ds %%s\n", biggestLangLength+len("Totals:"), biggestNumberOfFilesLength+len("No. files:"), biggestNumberOfWordsLength+len("No. words:"), biggestNumberOfNonEmptyLinesLength+len("No. Non Empty Lines:"))

	Aphrodite.PrintBold("Cyan", fmt.Sprintf(sentence, "Name: ", "No. Files:", "No. Words: ", "No. Non Empty Lines:", "No. Lines:"))

	for _, LanguageName := range utils.SortedKeys(stats) {
		printresult := stats[LanguageName]

		results_sentence := fmt.Sprintf(sentence, LanguageName, HumanReadableInt(printresult.Files), HumanReadableInt(printresult.Words), HumanReadableInt(printresult.NonEmptyLines), HumanReadableInt(printresult.Lines))
		finalPrint = append(finalPrint, results_sentence)

		totalFiles += printresult.Files
		totalLines += printresult.Lines
		totalNonEmptyLines += printresult.NonEmptyLines
		totalWords += printresult.Words
	}

	Aphrodite.PrintColour("Green", strings.Join(finalPrint, ""))

	total_sentence := fmt.Sprintf(sentence, "Totals:", HumanReadableInt(totalFiles), HumanReadableInt(totalWords), HumanReadableInt(totalNonEmptyLines), HumanReadableInt(totalLines))

	Aphrodite.PrintBoldHighIntensity("Yellow", total_sentence)
}
