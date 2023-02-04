package requests

type ConversionData struct {
	Supplier                 string `json:"Supplier"`
	Seller                   int    `json:"Seller"`
	SupplierForBillFromParty string `json:"SupplierForBillFromParty"`
	BillFromParty            int    `json:"BillFromParty"`
}
