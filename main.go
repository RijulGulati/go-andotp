package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/grijul/go-andotp/andotp"
	"golang.org/x/term"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s -i <INPUT_FILE> {-e|-d} [-o <OUT_FILE>] [-p PASSWORD]\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	encryptPtr := flag.Bool("e", false, "Encrypt file.")
	decryptPtr := flag.Bool("d", false, "Decrypt file")
	inputFilePtr := flag.String("i", "", "Input File")
	passwordPtr := flag.String("p", "", "Encryption Password. This option can be skipped to get password prompt.")
	outputFilePtr := flag.String("o", "", "Output File. If no file is provided, output is printed to STDOUT")

	flag.Parse()

	if *inputFilePtr == "" {
		fmt.Println(formatError("No input file provided\nSee -h for available options."))
		os.Exit(0)
	}

	if *encryptPtr && *decryptPtr {
		fmt.Println(formatError("Please provide any one of encrypt (-e) or decrypt (-d) flag"))
		os.Exit(0)
	}

	if *passwordPtr == "" {
		*passwordPtr = getPassword()
		if *passwordPtr == "" {
			fmt.Println(formatError("No password provided."))
			os.Exit(0)
		}
	}

	if *encryptPtr {
		plaintext := readFile(*inputFilePtr)
		filecontent, err := andotp.Encrypt(plaintext, *passwordPtr)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}

		processfile(filecontent, *outputFilePtr)

	} else if *decryptPtr {
		encryptedtext := readFile(*inputFilePtr)
		filecontent, err := andotp.Decrypt(encryptedtext, *passwordPtr)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}
		processfile(filecontent, *outputFilePtr)

	} else {
		fmt.Println(formatError("Please provide encrypt (-e) or decrypt (-d) flag"))
		os.Exit(0)
	}
}

func processfile(filecontent []byte, outputfile string) {
	if outputfile == "" {
		fmt.Printf("%s", filecontent)
	} else {
		ioutil.WriteFile(outputfile, filecontent, 0644)
	}
}

func getPassword() string {
	fmt.Print("Password: ")
	pass, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Print(formatError(err.Error()))
	}
	fmt.Print("\n")
	return string(pass)
}

func readFile(file string) []byte {
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return f
}

func formatError(e string) error {
	return fmt.Errorf("error: %s", e)
}
