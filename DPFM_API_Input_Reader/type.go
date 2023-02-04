package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey             string                    `json:"connection_key"`
	Result                    bool                      `json:"result"`
	RedisKey                  string                    `json:"redis_key"`
	Filepath                  string                    `json:"filepath"`
	APIStatusCode             int                       `json:"api_status_code"`
	RuntimeSessionID          string                    `json:"runtime_session_id"`
	BusinessPartnerID         *int                      `json:"business_partner"`
	ServiceLabel              string                    `json:"service_label"`
	SAPBusinessPartnerGeneral SAPBusinessPartnerGeneral `json:"BusinessPartner"`
	APISchema                 string                    `json:"api_schema"`
	Accepter                  []string                  `json:"accepter"`
	Deleted                   bool                      `json:"deleted"`
}

type SAPBusinessPartnerGeneral struct {
	BusinessPartner               string                       `json:"BusinessPartner"`
	Customer                      *string                      `json:"Customer"`
	Supplier                      *string                      `json:"Supplier"`
	AcademicTitle                 *string                      `json:"AcademicTitle"`
	AuthorizationGroup            *string                      `json:"AuthorizationGroup"`
	BusinessPartnerCategory       *string                      `json:"BusinessPartnerCategory"`
	BusinessPartnerFullName       *string                      `json:"BusinessPartnerFullName"`
	BusinessPartnerGrouping       *string                      `json:"BusinessPartnerGrouping"`
	BusinessPartnerName           *string                      `json:"BusinessPartnerName"`
	CorrespondenceLanguage        *string                      `json:"CorrespondenceLanguage"`
	CreationDate                  *string                      `json:"CreationDate"`
	CreationTime                  *string                      `json:"CreationTime"`
	FirstName                     *string                      `json:"FirstName"`
	Industry                      *string                      `json:"Industry"`
	IsFemale                      *bool                        `json:"IsFemale"`
	IsMale                        *bool                        `json:"IsMale"`
	IsNaturalPerson               *string                      `json:"IsNaturalPerson"`
	IsSexUnknown                  *bool                        `json:"IsSexUnknown"`
	GenderCodeName                *string                      `json:"GenderCodeName"`
	Language                      *string                      `json:"Language"`
	LastChangeDate                *string                      `json:"LastChangeDate"`
	LastChangeTime                *string                      `json:"LastChangeTime"`
	LastName                      *string                      `json:"LastName"`
	OrganizationBPName1           *string                      `json:"OrganizationBPName1"`
	OrganizationBPName2           *string                      `json:"OrganizationBPName2"`
	OrganizationBPName3           *string                      `json:"OrganizationBPName3"`
	OrganizationBPName4           *string                      `json:"OrganizationBPName4"`
	OrganizationFoundationDate    *string                      `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string                      `json:"OrganizationLiquidationDate"`
	SearchTerm1                   *string                      `json:"SearchTerm1"`
	SearchTerm2                   *string                      `json:"SearchTerm2"`
	AdditionalLastName            *string                      `json:"AdditionalLastName"`
	BirthDate                     *string                      `json:"BirthDate"`
	BusinessPartnerBirthplaceName *string                      `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string                      `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      *bool                        `json:"BusinessPartnerIsBlocked"`
	BusinessPartnerType           *string                      `json:"BusinessPartnerType"`
	GroupBusinessPartnerName1     *string                      `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string                      `json:"GroupBusinessPartnerName2"`
	IndependentAddressID          *string                      `json:"IndependentAddressID"`
	MiddleName                    *string                      `json:"MiddleName"`
	NameCountry                   *string                      `json:"NameCountry"`
	PersonFullName                *string                      `json:"PersonFullName"`
	PersonNumber                  *string                      `json:"PersonNumber"`
	IsMarkedForArchiving          *bool                        `json:"IsMarkedForArchiving"`
	BusinessPartnerIDByExtSystem  *string                      `json:"BusinessPartnerIDByExtSystem"`
	TradingPartner                *string                      `json:"TradingPartner"`
	SAPBusinessPartnerSupplier    []SAPBusinessPartnerSupplier `json:"BusinessPartnerSupplier"`
}

type SAPBusinessPartnerSupplier struct {
	Supplier                                         string                                             `json:"Supplier"`
	AuthorizationGroup                               *string                                            `json:"AuthorizationGroup"`
	CreationDate                                     *string                                            `json:"CreationDate"`
	Customer                                         *string                                            `json:"Customer"`
	PaymentIsBlockedForSupplier                      *bool                                              `json:"PaymentIsBlockedForSupplier"`
	PostingIsBlocked                                 *bool                                              `json:"PostingIsBlocked"`
	PurchasingIsBlocked                              *bool                                              `json:"PurchasingIsBlocked"`
	SupplierAccountGroup                             *string                                            `json:"SupplierAccountGroup"`
	SupplierFullName                                 *string                                            `json:"SupplierFullName"`
	SupplierName                                     *string                                            `json:"SupplierName"`
	BirthDate                                        *string                                            `json:"BirthDate"`
	DeletionIndicator                                *bool                                              `json:"DeletionIndicator"`
	Industry                                         *string                                            `json:"Industry"`
	IsNaturalPerson                                  *string                                            `json:"IsNaturalPerson"`
	SupplierCorporateGroup                           *string                                            `json:"SupplierCorporateGroup"`
	SupplierProcurementBlock                         *string                                            `json:"SupplierProcurementBlock"`
	SAPBusinessPartnerSupplierPartnerFunction        []SAPBusinessPartnerSupplierPartnerFunction        `json:"SAPBusinessPartnerSupplierPartnerFunction"`
	SAPBusinessPartnerSupplierPurchasingOrganization []SAPBusinessPartnerSupplierPurchasingOrganization `json:"SAPBusinessPartnerSupplierPurchasingOrganization"`
	SAPBusinessPartnerSupplierCompany                []SAPBusinessPartnerSupplierCompany                `json:"SAPBusinessPartnerSupplierCompany"`
}

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
