package data_service

import "fmt"

type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("send xml data")
}
