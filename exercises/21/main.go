package main

import (
	"wb1/exercises/21/data"
	data_service "wb1/exercises/21/data-service"
)

type JsonDocumentAdapter struct {
	*data.JsonDocument
}

func (doc JsonDocumentAdapter) ConvertToXml() string {
	return "<xml>...</xml>"
}

func (adapter *JsonDocumentAdapter) SendXmlData() {
	s := adapter.ConvertToXml()
	println("send json convert to xml", s)
}

func NewJsonDocumentAdapter(json *data.JsonDocument) *JsonDocumentAdapter {
	return &JsonDocumentAdapter{json}
}

func main() {
	jsonDoc := NewJsonDocumentAdapter(&data.JsonDocument{})
	data_service.AnalyticalDataService(jsonDoc).SendXmlData()
	xmlDoc := data_service.XmlDocument{}
	data_service.AnalyticalDataService(xmlDoc).SendXmlData()
}
