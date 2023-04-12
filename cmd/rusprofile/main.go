package main

import (
	"fmt"
	"github.com/suzibill/rusprofile/internal/service/parser"
)

func main() {
	companyInfo, err := parser.GetCompanyInfo("5609026406")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#+v", companyInfo)
}
