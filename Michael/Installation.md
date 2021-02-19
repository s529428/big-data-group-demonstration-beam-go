# Installing Golang

## Windows
You can use [`chocolaty`](https://chocolatey.org/install) to install golang on to your Windows installation.

After you install chocolaty, run `choco install golang`. This will automate the installation process for you.
This may take a while. 

After, run `refreshenv` and close your Command Prompt/Powershell. 

## Linux (Ubuntu/WSL)
Using `apt` or `apt-get` can install golang for you.
```zsh
$ sudo apt update
$ sudo apt upgrade
$ sudo apt install golang-go
```

## Visual Studio Code Extensions
Visual Studio Code has amazing extension for all kinds of languages, and Go is no exception. Here's [Google's official extension for Go](https://marketplace.visualstudio.com/items?itemName=golang.go).

# Running Code

Here's a simple "Hello, World" program in Go (borrowed from [here](https://www.geeksforgeeks.org/hello-world-in-golang/)):
```go
package main 

//'fmt' contains functions like C's printf and scanf. 
import "fmt"
  
// Main function 
func main() { 
  
    fmt.Println("!... Hello World ...!") 
} 
```
Save this as "`helloworld.go`".

## Windows
Open a Powershell Terminal in the directory with your "`helloworld.go`" file.

```cmd
PS> go build helloworld.go
PS> .\helloworld.exe
```

If everything went right, you've gotten `"!... Hello World ...!"` in your window. 

## Linux

This can be done the same as Windows. 

Navigate to the directory with your "`helloworld.go`" file, and execute:

```bash
$ go build helloworld.go
$ ./helloworld
```
If everything is correct, you'll see `"!... Hello World ...!"` print out on your terminal.

# Reading the Documentation
If you're without internet, you can find the documentation for the built-in packages:
```bash
go docs <package>
```
Or if you have internet, you can find the online documentation [here](https://golang.org/pkg/)

