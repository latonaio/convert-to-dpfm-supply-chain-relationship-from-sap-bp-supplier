package requests

type SAPBusinessPartnerSupplierPartnerFunction struct {
	Supplier               string  `json:"Supplier"`
	PurchasingOrganization string  `json:"PurchasingOrganization"`
	PartnerCounter         string  `json:"PartnerCounter"`
	PartnerFunction        *string `json:"PartnerFunction"`
	Plant                  *string `json:"Plant"`
	DefaultPartner         *bool   `json:"DefaultPartner"`
	CreationDate           *string `json:"CreationDate"`
	ReferenceSupplier      *string `json:"ReferenceSupplier"`
	AuthorizationGroup     *string `json:"AuthorizationGroup"`
}
