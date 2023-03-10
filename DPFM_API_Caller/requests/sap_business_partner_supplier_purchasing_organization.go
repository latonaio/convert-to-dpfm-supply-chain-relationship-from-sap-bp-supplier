package requests

type SAPBusinessPartnerSupplierPurchasingOrganization struct {
	Supplier                       string  `json:"Supplier"`
	PurchasingOrganization         string  `json:"PurchasingOrganization"`
	IncotermsClassification        *string `json:"IncotermsClassification"`
	InvoiceIsGoodsReceiptBased     *bool   `json:"InvoiceIsGoodsReceiptBased"`
	PaymentTerms                   *string `json:"PaymentTerms"`
	PurOrdAutoGenerationIsAllowed  *bool   `json:"PurOrdAutoGenerationIsAllowed"`
	PurchaseOrderCurrency          *string `json:"PurchaseOrderCurrency"`
	PurchasingGroup                *string `json:"PurchasingGroup"`
	ShippingCondition              *string `json:"ShippingCondition"`
	SupplierPhoneNumber            *string `json:"SupplierPhoneNumber"`
	SupplierRespSalesPersonName    *string `json:"SupplierRespSalesPersonName"`
	PurchasingIsBlockedForSupplier *bool   `json:"PurchasingIsBlockedForSupplier"`
	DeletionIndicator              *bool   `json:"DeletionIndicator"`
}
