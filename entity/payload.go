package entity

type Entities interface {
	User | SelfEmployedReceipt | SelfEmployedItem | InvoiceReturnItem |
		ProducerReceipt | ProducerItem | Invoice | InvoiceItem
}

type Response[Entity Entities] struct {
	Entity   Entity `json:"data"`
	Metadata struct {
		Optime string `json:"optime"`
	} `json:"metadata"`
}

type Exception struct {
	Error    string `json:"error"`
	Messages []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"messages"`
}
