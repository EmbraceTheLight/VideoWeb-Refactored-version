package iputil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	searchAPI    = "http://ip-api.com/json/"
	searchSuffix = "?lang=zh-CN"
)

// IPInfo contains information about an IP address.
type IPInfo struct {
	Query     string `json:"query"`     //待查询的IP
	Status    string `json:"status"`    //查询状态，如success
	Continent string `json:"continent"` //大洲	//Mobile  bool `json:"mobile"`
	Country   string `json:"country"`
	City      string `json:"city"`
	//Proxy   bool `json:"proxy"`
	//Hosting bool `json:"hosting"`
	//Lat           float64 `json:"lat"`       //纬度
	//Lon           float64 `json:"lon"`       //经度
	//ContinentCode string  `json:"continentCode"`
	//CountryCode   string  `json:"countryCode"`
	//Region        string  `json:"region"`
	//RegionName    string  `json:"regionName"`
	//District      string  `json:"district"`
	//Zip           string  `json:"zip"`
	//Timezone      string  `json:"timezone"`
	//Offset        string  `json:"offset"`
	//Currency      string  `json:"currency"`
	//ISP           string  `json:"ISP"`
	//Org           string  `json:"org"`
	//As            string  `json:"as"`
	//AsName        string  `json:"asName"`
}

// GetIPInfo returns information about an IP address.
func GetIPInfo(ip string) (*IPInfo, error) {
	searchURL := searchAPI + ip + searchSuffix
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo IPInfo
	err = json.Unmarshal(out, &ipInfo)
	if err != nil {
		fmt.Println("err in Unmarshal:", err)
		return nil, err
	}

	return &ipInfo, nil
}

// GetCountryAndCity returns the city and continent of an IP address.
func GetCountryAndCity(ip string) (country, city string, err error) {
	info, err := GetIPInfo(ip)
	if err != nil {
		return "", "", err
	}
	return info.City, info.Continent, nil
}
