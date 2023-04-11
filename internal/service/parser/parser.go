package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type CompanyInfo struct {
	INN         string
	KPP         string
	Name        string
	CEOFullname string
}

func GetCompanyInfo(INN string) (CompanyInfo, error) {
	if !IsValidINN(INN) {
		return CompanyInfo{}, fmt.Errorf("invalid INN: %s", INN)
	}
	companyURL := fmt.Sprintf("https://www.rusprofile.ru/search?query=%s", INN)
	resp, err := http.Get(companyURL)
	if err != nil {
		return CompanyInfo{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return CompanyInfo{}, err
	}

	title := doc.Find("title").Text()
	if title == INN+" - результаты поиска на Rusprofile.ru" {
		return CompanyInfo{}, fmt.Errorf("company with INN %s not found", INN)
	}

	var companyInfo CompanyInfo
	companyInfo.INN = INN
	companyInfo.KPP = doc.Find("#clip_kpp").Text()
	companyInfo.Name = doc.Find("h1[itemprop='name']").Text()
	companyInfo.Name = strings.ReplaceAll(strings.TrimSpace(companyInfo.Name), "\"", "")
	companyInfo.CEOFullname = doc.Find(".company-info__text a span").Text()

	return companyInfo, nil
}

func IsValidINN(INN string) bool {
	if len(INN) != 10 {
		return false
	}
	for _, c := range INN {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
