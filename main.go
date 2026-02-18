package main

import (
	"fmt"
	"os"
	"github.com/KiranSatyaRaj/TeenyTinyGoCompiler/pkg/lexer"
)

func main() {
		src := "IF+-123foo*THEN/"
		lexer.Init(src)
		
		token, err := lexer.Trex.GetToken()
		if err != nil {
			fmt.Println(err)
			return
		}
		for token.GetKind() != lexer.EOF {
			fmt.Println(token.GetKind())
			token, err = lexer.Trex.GetToken()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
}
