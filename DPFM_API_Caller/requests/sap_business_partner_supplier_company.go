package requests

type SAPBusinessPartnerSupplierCompany struct {
	Supplier                    string  `json:"Supplier"`
	CompanyCode                 string  `json:"CompanyCode"`
	PaymentBlockingReason       *string `json:"PaymentBlockingReason"`
	PaymentMethodsList          *string `json:"PaymentMethodsList"`
	PaymentTerms                *string `json:"PaymentTerms"`
	ClearCustomerSupplier       *string `json:"ClearCustomerSupplier"`
	HouseBank                   *string `json:"HouseBank"`
	ReconciliationAccount       *string `json:"ReconciliationAccount"`
	SupplierIsBlockedForPosting *string `json:"SupplierIsBlockedForPosting"`
	DeletionIndicator           *string `json:"DeletionIndicator"`
}
