package dpfm_api_processing_formatter

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
	MappingTransaction             *[]MappingTransaction             `json:"MappingTransaction"`
	CodeConversionTransaction      *[]CodeConversionTransaction      `json:"CodeConversionHeader"`
	MappingDeliveryRelation        *[]MappingDeliveryRelation        `json:"MappingDeliveryRelation"`
	CodeConversionDeliveryRelation *[]CodeConversionDeliveryRelation `json:"CodeConversionDeliveryRelation"`
	SetSupplierDeliveryRelation    *SetSupplierDeliveryRelation      `json:"SetSupplierDeliveryRelation"`
	MappingBillingRelation         *[]MappingBillingRelation         `json:"MappingBillingRelation"`
	CodeConversionBillingRelation  *[]CodeConversionBillingRelation  `json:"CodeConversionBillingRelation"`
	SetSupplierBillingRelation     *SetSupplierBillingRelation       `json:"SetSupplierBillingRelation"`
	MappingPaymentRelation         *[]MappingPaymentRelation         `json:"MappingPaymentRelation"`
	ConversionData                 *[]ConversionData                 `json:"ConversionData"`
}

type Transaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	Incoterms                 *string `json:"Incoterms"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	QuotationIsBlocked        *bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked            *bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked         *bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked          *bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryRelation struct {
	SupplyChainRelationshipID         int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int     `json:"Buyer"`
	Seller                            int     `json:"Seller"`
	DeliverToParty                    int     `json:"DeliverToParty"`
	DeliverFromParty                  int     `json:"DeliverFromParty"`
	DefaultRelation                   *bool   `json:"DefaultRelation"`
	CreationDate                      *string `json:"CreationDate"`
	LastChangeDate                    *string `json:"LastChangeDate"`
	IsMarkedForDeletion               *bool   `json:"IsMarkedForDeletion"`
}

type BillingRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	BillToCountry                    *string `json:"BillToCountry"`
	BillFromCountry                  *string `json:"BillFromCountry"`
	IsExportImport                   *bool   `json:"IsExportImport"`
	TransactionTaxCategory           *string `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string `json:"TransactionTaxClassification"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type PaymentRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	Payer                            int     `json:"Payer"`
	Payee                            int     `json:"Payee"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	PayerHouseBank                   *string `json:"PayerHouseBank"`
	PayerHouseBankAccount            *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank                   *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount            *string `json:"PayeeHouseBankAccount"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

// 項目マッピング変換
type MappingTransaction struct {
	Buyer                  int     `json:"Buyer"`
	TransactionCurrency    *string `json:"TransactionCurrency"`
	PaymentTerms           *string `json:"PaymentTerms"`
	Incoterms              *string `json:"Incoterms"`
	AccountAssignmentGroup *string `json:"AccountAssignmentGroup"`
	CreationDate           *string `json:"CreationDate"`
	LastChangeDate         *string `json:"LastChangeDate"`
}

type MappingDeliveryRelation struct {
	Supplier       string  `json:"Supplier"`
	Buyer          int     `json:"Buyer"`
	DeliverToParty int     `json:"DeliverToParty"`
	CreationDate   *string `json:"CreationDate"`
	LastChangeDate *string `json:"LastChangeDate"`
}

type MappingBillingRelation struct {
	Supplier       string  `json:"Supplier"`
	Buyer          int     `json:"Buyer"`
	BillToParty    int     `json:"BillToParty"`
	BillToCountry  *string `json:"BillToCountry"`
	CreationDate   *string `json:"CreationDate"`
	LastChangeDate *string `json:"LastChangeDate"`
}

type MappingPaymentRelation struct {
	Supplier                 string  `json:"Supplier"`
	Buyer                    int     `json:"Buyer"`
	SupplierForBillFromParty string  `json:"SupplierForBillFromParty"`
	BillToParty              int     `json:"BillToParty"`
	Payer                    int     `json:"Payer"`
	PayerHouseBank           *string `json:"PayerHouseBank"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
}

// 番号変換
type CodeConversionKey struct {
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionQueryGets struct {
	CodeConversionID  int    `json:"CodeConversionID"`
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	CodeConvertTo     string `json:"CodeConvertTo"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionTransaction struct {
	Supplier string `json:"Supplier"`
	Seller   int    `json:"Seller"`
}

type CodeConversionDeliveryRelation struct {
	DeliverFromParty int `json:"DeliverFromParty"`
}

type CodeConversionBillingRelation struct {
	SupplierForBillFromParty string `json:"SupplierForBillFromParty"`
	BillFromParty            int    `json:"BillFromParty"`
}

type ConversionData struct {
	Supplier                 string `json:"Supplier"`
	Seller                   int    `json:"Seller"`
	SupplierForBillFromParty string `json:"SupplierForBillFromParty"`
	BillFromParty            int    `json:"BillFromParty"`
}

// 個別処理
type SetSupplier struct {
	PartnerFunction string  `json:"PartnerFunction"`
	Supplier        *string `json:"Supplier"`
}

type SetSupplierBillingRelation struct {
	SupplierForBillFromParty *string `json:"SupplierForBillFromParty"`
}

type SetSupplierDeliveryRelation struct {
	SupplierForDeliverFromParty *string `json:"SupplierForDeliverFromParty"`
}
