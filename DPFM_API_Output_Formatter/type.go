package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Transaction      []Transaction      `json:"Transaction"`
	DeliveryRelation []DeliveryRelation `json:"DeliveryRelation"`
	BillingRelation  []BillingRelation  `json:"BillingRelation"`
	PaymentRelation  []PaymentRelation  `json:"PaymentRelation"`
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
