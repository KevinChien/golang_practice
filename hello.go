package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")
}

// 建置: go build <go file>
// output: <go file name>.exe

// execute via: .\<file name>

//tpye $env:PATH += ";$(Get-Location)" under powershell can execute <go file> without symbol ".\"
