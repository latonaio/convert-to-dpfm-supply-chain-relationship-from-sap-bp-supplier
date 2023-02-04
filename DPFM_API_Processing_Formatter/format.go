package dpfm_api_processing_formatter

import (
	"convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// データ連携基盤とSAP Business Partner Customerとの項目マッピング
// Transaction
func (psdc *SDC) ConvertToMappingTransaction(sdc *dpfm_api_input_reader.SDC) (*[]MappingTransaction, error) {
	var res []MappingTransaction
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier

	systemDate := GetSystemDatePtr()

	for _, dataSupplier := range dataSupplier {
		dataPurchasingOrganization := dataSupplier.SAPBusinessPartnerSupplierPurchasingOrganization
		for _, dataPurchasingOrganization := range dataPurchasingOrganization {

			res = append(res, MappingTransaction{
				Buyer:                  *sdc.BusinessPartnerID,
				TransactionCurrency:    dataPurchasingOrganization.PurchaseOrderCurrency,
				PaymentTerms:           dataPurchasingOrganization.PaymentTerms,
				Incoterms:              dataPurchasingOrganization.IncotermsClassification,
				AccountAssignmentGroup: dataSupplier.SupplierAccountGroup,
				CreationDate:           systemDate,
				LastChangeDate:         systemDate,
			})
		}
	}

	return &res, nil
}

// DeliveryRelation
func (psdc *SDC) ConvertToMappingDeliveryRelation(sdc *dpfm_api_input_reader.SDC) (*[]MappingDeliveryRelation, error) {
	var res []MappingDeliveryRelation
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier

	systemDate := GetSystemDatePtr()

	for _, dataSupplier := range dataSupplier {

		res = append(res, MappingDeliveryRelation{
			Supplier:       dataSupplier.Supplier,
			Buyer:          *sdc.BusinessPartnerID,
			DeliverToParty: *sdc.BusinessPartnerID,
			CreationDate:   systemDate,
			LastChangeDate: systemDate,
		})
	}

	return &res, nil
}

// BillingRelation
func (psdc *SDC) ConvertToMappingBillingRelation(sdc *dpfm_api_input_reader.SDC) (*[]MappingBillingRelation, error) {
	var res []MappingBillingRelation
	data := sdc.SAPBusinessPartnerGeneral
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier

	systemDate := GetSystemDatePtr()

	for _, dataSupplier := range dataSupplier {
		res = append(res, MappingBillingRelation{
			Supplier:       dataSupplier.Supplier,
			Buyer:          *sdc.BusinessPartnerID,
			BillToParty:    *sdc.BusinessPartnerID,
			BillToCountry:  data.NameCountry,
			CreationDate:   systemDate,
			LastChangeDate: systemDate,
		})
	}

	return &res, nil
}

// PaymentRelation
func (psdc *SDC) ConvertToMappingPaymentRelation(sdc *dpfm_api_input_reader.SDC) (*[]MappingPaymentRelation, error) {
	var res []MappingPaymentRelation
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier

	systemDate := GetSystemDatePtr()

	for _, dataSupplier := range dataSupplier {
		dataSupplierCompany := dataSupplier.SAPBusinessPartnerSupplierCompany
		for _, dataSupplierCompany := range dataSupplierCompany {
			res = append(res, MappingPaymentRelation{
				Supplier:                 dataSupplier.Supplier,
				Buyer:                    *sdc.BusinessPartnerID,
				SupplierForBillFromParty: dataSupplier.Supplier,
				BillToParty:              *sdc.BusinessPartnerID,
				Payer:                    *sdc.BusinessPartnerID,
				PayerHouseBank:           dataSupplierCompany.HouseBank,
				CreationDate:             systemDate,
				LastChangeDate:           systemDate,
			})
		}
	}

	return &res, nil
}

// 番号処理
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	defer rows.Close()
	var res []CodeConversionQueryGets

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CodeConversionQueryGets{}

		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionTransaction(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionTransaction, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionTransaction{}

	pm.Supplier = dataMap["Seller"].CodeConvertFrom
	pm.Seller, err = ParseInt(dataMap["Seller"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionTransaction{
		Supplier: data.Supplier,
		Seller:   data.Seller,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionDeliveryRelation(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionDeliveryRelation, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionDeliveryRelation{}

	pm.DeliverFromParty, err = ParseInt(dataMap["DeliverFromParty"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionDeliveryRelation{
		DeliverFromParty: data.DeliverFromParty,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionBillingRelation(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionBillingRelation, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionBillingRelation{}

	pm.BillFromParty, err = ParseInt(dataMap["BillFromParty"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionBillingRelation{
		BillFromParty: data.BillFromParty,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConversionData() *[]ConversionData {
	var res []ConversionData

	for _, v := range *psdc.CodeConversionTransaction {
		for _, w := range *psdc.CodeConversionBillingRelation {
			pm := &requests.ConversionData{}

			pm.Supplier = v.Supplier
			pm.Seller = v.Seller
			pm.SupplierForBillFromParty = w.SupplierForBillFromParty
			pm.BillFromParty = w.BillFromParty

			data := pm
			res = append(res, ConversionData{
				Supplier:                 data.Supplier,
				Seller:                   data.Seller,
				SupplierForBillFromParty: data.SupplierForBillFromParty,
				BillFromParty:            data.BillFromParty,
			})
		}
	}

	return &res
}

// 個別処理
func (psdc *SDC) ConvertToSetSupplier(supplierPartner dpfm_api_input_reader.SAPBusinessPartnerSupplierPartnerFunction) *SetSupplier {
	pm := &requests.SetSupplier{}

	pm.PartnerFunction = *supplierPartner.PartnerFunction
	pm.Supplier = &supplierPartner.Supplier

	data := pm
	res := SetSupplier{
		PartnerFunction: data.PartnerFunction,
		Supplier:        data.Supplier,
	}

	return &res
}

func (psdc *SDC) ConvertToSetSupplierDeliveryRelation(supplier []SetSupplier) *SetSupplierDeliveryRelation {
	pm := &requests.SetSupplierDeliveryRelation{}

	supplierMap := make(map[string]SetSupplier, len(supplier))
	for _, v := range supplier {
		supplierMap[v.PartnerFunction] = v
	}

	pm.SupplierForDeliverFromParty = supplierMap["OA"].Supplier

	data := pm
	res := SetSupplierDeliveryRelation{
		SupplierForDeliverFromParty: data.SupplierForDeliverFromParty,
	}

	return &res
}

func (psdc *SDC) ConvertToSetSupplierBillingRelation(supplier []SetSupplier) *SetSupplierBillingRelation {
	pm := &requests.SetSupplierBillingRelation{}

	supplierMap := make(map[string]SetSupplier, len(supplier))
	for _, v := range supplier {
		supplierMap[v.PartnerFunction] = v
	}

	pm.SupplierForBillFromParty = supplierMap["IP"].Supplier

	data := pm
	res := SetSupplierBillingRelation{
		SupplierForBillFromParty: data.SupplierForBillFromParty,
	}

	return &res
}
