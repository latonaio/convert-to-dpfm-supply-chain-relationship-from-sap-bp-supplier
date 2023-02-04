package requests

type CodeConversionBillingRelation struct {
	SupplierForBillFromParty string `json:"SupplierForBillFromParty"`
	BillFromParty            int    `json:"BillFromParty"`
}
