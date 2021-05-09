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
		fmt.Println(andotp.FormatError("No input file provided\nSee -h for available options."))
		os.Exit(0)
	}

	if *encryptPtr && *decryptPtr {
		fmt.Println(andotp.FormatError("Please provide any one of encrypt (-e) or decrypt (-d) flag"))
		os.Exit(0)
	}

	if *passwordPtr == "" {
		*passwordPtr = getPassword()
		if *passwordPtr == "" {
			fmt.Println(andotp.FormatError("No password provided."))
			os.Exit(0)
		}
	}

	if *encryptPtr {

		plaintext, err := andotp.ReadFile(*inputFilePtr)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}

		filecontent, err := andotp.Encrypt(plaintext, *passwordPtr)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}

		processfile(filecontent, *outputFilePtr)

	} else if *decryptPtr {

		encryptedtext, err := andotp.ReadFile(*inputFilePtr)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}

		filecontent, err := andotp.Decrypt(encryptedtext, *passwordPtr)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(0)
		}

		processfile(filecontent, *outputFilePtr)

	} else {
		fmt.Println(andotp.FormatError("Please provide encrypt (-e) or decrypt (-d) flag"))
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
		fmt.Print(andotp.FormatError(err.Error()))
	}
	fmt.Print("\n")
	return string(pass)
}
