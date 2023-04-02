# go-andotp
CLI program to encrypt/decrypt [andOTP](https://github.com/andOTP/andOTP) files.

## Installation

<details>
<summary><b>Linux</b></summary>

Download:
* [x86_64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-linux-x86_64) Intel or AMD 64-Bit CPU
  ```shell
  curl -L "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-linux-x86_64" \
       -o "go-andotp" && \
  chmod +x "go-andotp"
  ```
* [arm64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-linux-arm64) Arm-based 64-Bit CPU (i.e. in Raspberry Pi)
  ```shell
  curl -L "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-linux-arm64" \
       -o "go-andotp" && \
  chmod +x "go-andotp"
  ```

To determine your OS version, run `getconf LONG_BIT` or `uname -m` at the command line.
</details>

<details>
<summary><b>macOS</b></summary>

Download:
* [x86_64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-macos-x86_64) Intel 64-bit
  ```shell
  curl -L "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-macos-x86_64" \
       -o "go-andotp" && \
  chmod +x "go-andotp"
  ```
* [arm64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-macos-arm64) Apple silicon 64-bit
  ```shell
  curl -L "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-macos-arm64" \
       -o "go-andotp" && \
  chmod +x "go-andotp"
  ```

To determine your OS version, run `uname -m` at the command line.
</details>

<details>
<summary><b>Windows</b></summary>

Download:
* [x86_64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-windows-x86_64.exe) Intel or AMD 64-Bit CPU
   ```powershell
   Invoke-WebRequest -Uri "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-windows-x86_64.exe" -OutFile "go-andotp.exe"
   ```
* [arm64](https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-windows-arm64.exe) Arm-based 64-Bit CPU
   ```powershell
   Invoke-WebRequest -Uri "https://github.com/RijulGulati/go-andotp/releases/latest/download/go-andotp-windows-arm64.exe" -OutFile "go-andotp.exe"
   ```
To determine your OS version, run `echo %PROCESSOR_ARCHITECTURE%` at the command line.
</details>

<details>
<summary><b>Go</b></summary>

```shell
go install github.com/grijul/go-andotp
```
</details>

## Usage
```text
Usage: go-andotp -i <INPUT_FILE> {-e|-d} [-o <OUT_FILE>] [-p PASSWORD]

  -d    Decrypt file
  -e    Encrypt file.
  -i string
        Input File
  -o string
        Output File. If no file is provided, output is printed to STDOUT
  -p string
        Encryption Password. This option can be skipped to get password prompt.
```

## Examples
- Encrypt JSON file (Password is asked after hitting ```Enter```. Password is not echoed)
```shell
go-andotp -e -i file.json -o file.json.aes
```
- Encrypt JSON file (Password is entered through CLI)
```shell
go-andotp -e -i file.json -o file.json.aes -p testpass
```
- Decrypt JSON file
```shell
go-andotp -d -i file.aes.json -o file.json
```
- Decrypt JSON file and print json to console
```shell
go-andotp -d -i file.aes.json
```

## Using go-andotp as library
go-andotp can be used as library as well. It implements ```Encrypt()``` and ```Decrypt()``` functions to encrypt/decrypt text (respectively).
It's documentation is available at: https://pkg.go.dev/github.com/grijul/go-andotp/andotp

Example usage:
```go
import "github.com/grijul/go-andotp/andotp"

func main() {
    andotp.Encrypt(...)
    andotp.Decrypt(...)
}
```

## Build
Compile `go-andotp` on your computer:

```shell
go build -o go-andotp main.go
```

To compile `go-andotp` for another platform please set the `GOARCH` and `GOOS` environmental variables.
Example:
```shell
GOOS=windows GOARCH=amd64 go build -o go-andotp.exe main.go
```

To compile `go-andotp` for Windows, macOS and Linux you can use the script `build.sh`:
```shell
bash build.sh
```

More help: <https://go.dev/doc/install/source#environment>

# License
[MIT](https://github.com/grijul/go-andotp/blob/main/LICENSE)
