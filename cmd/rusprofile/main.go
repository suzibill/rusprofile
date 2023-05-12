package main

import (
	"github.com/suzibill/rusprofile/internal/server"
)

func main() {
	//companyInfo, err := parser.GetCompanyInfo("5609026406")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%#+v", companyInfo)
	server.StartServer()
}
