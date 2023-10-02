# argsparsego

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/masa23/argsparsego)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## 概要 

Go言語でコマンドライン引数をパースするためのライブラリです。  
exec.Command()で実行するコマンドの引数をパースすることを想定しています。  


## 使い方

```go
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/masa23/argsparsego"
)

func main() {
	args := "hoge \"fuga piyo\""

	parsed, err := argsparsego.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command("echo", parsed...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

## ライセンス
MIT