package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/integrii/flaggy"
	"github.com/saintfish/chardet"
)

var appExec = "NFO-to-UTF8"
var appVersion string
var fileName string
var spaces = false
var verbose = false
var linebreaks = false

func init() {
	appExec, _ = os.Executable()
	flaggy.SetName(filepath.Base(appExec))
	flaggy.SetDescription("A command line tool to convert NFO files from CP437 to UTF-8 encoding")
	flaggy.AddPositionalValue(&fileName, "NFO", 1, true, "Path to the NFO file to be converted")
	flaggy.Bool(&spaces, "s", "spaces", "Convert spaces to non-breaking spaces")
	flaggy.Bool(&linebreaks, "l", "linebreaks", "Convert line breaks to correct characters for the system (LF for Linux/Mac and CRLF for Windows)")
	flaggy.Bool(&verbose, "v", "verbose", "Show verbose output")
	if appVersion != "" {
		flaggy.SetVersion(appVersion)
	}
	flaggy.Parse()
}

func main() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		exit(err)
	}
	if encoding := detectEncoding(file); encoding == "CP437" {
		utf8File := cp437toUTF8(file, spaces)
		err = ioutil.WriteFile(fileName, []byte(utf8File), 0666)
		if err != nil {
			exit(err)
		} else {
			if verbose {
				fmt.Printf("File %s succesfully converted to UTF-8\n", fileName)
			}
		}
	} else {
		if verbose {
			fmt.Printf("File %s is not CP437 encoded, exiting...\n", fileName)
		}
		exit(nil)
	}
}

func detectEncoding(data []byte) string {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(data)
	if err != nil {
		if err.Error() == "Charset not detected." {
			if verbose {
				fmt.Println("Charset not detected")
				fmt.Println("Assuming it is CP437")
			}
			return "CP437"
		}
		exit(fmt.Errorf("failed decoding data from file: %s", err))
	}
	if verbose {
		if verbose {
			fmt.Printf("Detected encoding: %s\n", result.Charset)
		}
	}
	if strings.Contains(result.Charset, "ISO-8859") || strings.Contains(result.Charset, "windows") || strings.Contains(result.Charset, "KOI8") || strings.Contains(result.Charset, "IBM") || result.Charset == "Shift_JIS" {
		if verbose {
			fmt.Println("Assuming it is CP437")
		}
		return "CP437"
	}
	return result.Charset
}

func cp437toUTF8(b []byte, convertSpaces bool) string {
	space := " "
	if convertSpaces {
		if verbose {
			fmt.Println("Replacing spaces with non-breaking spaces")
		}
		space = "\u00A0"
	}
	var cp437 = []rune("\u0000☺☻♥♦♣♠•◘○\u000A♂♀\u000D♬☼►◄↕‼¶§▬↨↑↓→←∟↔▲▼" + space + "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~⌂ÇüéâäàåçêëèïîìÄÅÉæÆôöòûùÿÖÜ¢£¥₧ƒáíóúñÑªº¿⌐¬½¼¡«»\u2591\u2592\u2593│┤╡╢╖╕╣║╗╝╜╛┐└┴┬├─┼╞╟╚╔╩╦╠═╬╧╨╤╥╙╘╒╓╫╪┘┌█▄▌▐▀αßΓπΣσµτΦΘΩδ∞φε∩≡±≥≤⌠⌡÷≈°∙·√ⁿ²■\u00A0")
	runes := make([]rune, len(b))
	for i := range runes {
		runes[i] = cp437[b[i]]
	}
	utf8 := string(runes)
	if linebreaks {
		if verbose {
			fmt.Println("Replacing line break characters")
		}
		utf8 = strings.ReplaceAll(utf8, "\u000D\u000A", "\u000A")
		if runtime.GOOS == "windows" {
			utf8 = strings.ReplaceAll(utf8, "\u000A", "\u000D\u000A")
		}
	}
	return utf8
}

func exit(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
