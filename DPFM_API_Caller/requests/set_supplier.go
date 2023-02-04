package requests

type SetSupplier struct {
	PartnerFunction string  `json:"PartnerFunction"`
	Supplier        *string `json:"Supplier"`
}
