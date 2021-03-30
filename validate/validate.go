package validate

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Synopsis: \n \t%s command [options]\n", os.Args[0])
		fmt.Printf("Description: \n \t Choose one of these commands:\n")
		fmt.Print("\t\t help - print help message\n")
		fmt.Print("\t\t version - print info about program version\n")
		fmt.Print("\t\t run - starts a HTTP web server and serve indicated file\n")
		fmt.Print("\t\t\t --file <string> - string indicating html file\n")
	}

	log.SetPrefix("pcCLI: ")
	log.SetFlags(0)
}

func checkFilename() (string, error) {
	if len(os.Args) != 4 {
		return "", errors.New("wrong --file argument, check help message")
	}

	filename := string(os.Args[3])

	if filepath.Ext(filename) != ".html" {
		return "", errors.New("wrong file extension, required .html")
	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return "", errors.New("given file doesn't exist")
	}

	return filename, nil

}

func CheckCommand() (string, error) {

	flag.Parse()

	if len(os.Args) < 2 {
		return "", errors.New("wrong list of arguments")
	}

	arg := os.Args[1]

	switch arg {
	case "run":
		validFilename, err := checkFilename()
		if err != nil {
			return "", err
		}
		return validFilename, nil
	case "help":
		return "help", nil
	case "version":
		return "version", nil
	default:
		return "", errors.New("wrong command, check help message")
	}

}
