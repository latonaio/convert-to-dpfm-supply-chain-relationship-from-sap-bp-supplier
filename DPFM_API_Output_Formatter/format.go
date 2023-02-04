package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToTransaction(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Transaction, error) {
	var transactions []Transaction
	mappingTransaction := psdc.MappingTransaction
	codeConversionTransaction := psdc.CodeConversionTransaction

	for i := range *mappingTransaction {
		transaction := Transaction{}

		data, err := json.Marshal((*mappingTransaction)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &transaction)
		if err != nil {
			return nil, err
		}

		transaction.Seller = (*codeConversionTransaction)[i].Seller

		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}

func ConvertToDeliveryRelation(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]DeliveryRelation, error) {
	var deliveryRelations []DeliveryRelation
	mappingDeliveryRelation := psdc.MappingDeliveryRelation
	codeConversionDeliveryRelation := psdc.CodeConversionDeliveryRelation
	conversionData := psdc.ConversionData

	for i := range *mappingDeliveryRelation {
		deliveryRelation := DeliveryRelation{}

		data, err := json.Marshal((*mappingDeliveryRelation)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &deliveryRelation)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.Supplier == (*mappingDeliveryRelation)[i].Supplier {
				deliveryRelation.Seller = v.Seller
				break
			}
		}
		deliveryRelation.DeliverFromParty = (*codeConversionDeliveryRelation)[i].DeliverFromParty

		deliveryRelations = append(deliveryRelations, deliveryRelation)
	}

	return &deliveryRelations, nil
}

func ConvertToBillingRelation(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]BillingRelation, error) {
	var billingRelations []BillingRelation
	mappingBillingRelation := psdc.MappingBillingRelation
	codeConversionBillingRelation := psdc.CodeConversionBillingRelation
	conversionData := psdc.ConversionData

	for i := range *mappingBillingRelation {
		billingRelation := BillingRelation{}

		data, err := json.Marshal((*mappingBillingRelation)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &billingRelation)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.Supplier == (*mappingBillingRelation)[i].Supplier {
				billingRelation.Seller = v.Seller
				break
			}
		}
		billingRelation.BillFromParty = (*codeConversionBillingRelation)[i].BillFromParty

		billingRelations = append(billingRelations, billingRelation)
	}

	return &billingRelations, nil
}

func ConvertToPaymentRelation(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]PaymentRelation, error) {
	var paymentRelations []PaymentRelation
	mappingPaymentRelation := psdc.MappingPaymentRelation
	conversionData := psdc.ConversionData

	for i := range *mappingPaymentRelation {
		paymentRelation := PaymentRelation{}

		data, err := json.Marshal((*mappingPaymentRelation)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &paymentRelation)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.Supplier == (*mappingPaymentRelation)[i].Supplier && v.SupplierForBillFromParty == (*mappingPaymentRelation)[i].SupplierForBillFromParty {
				paymentRelation.Seller = v.Seller
				paymentRelation.BillFromParty = v.BillFromParty
				break
			}
		}

		paymentRelations = append(paymentRelations, paymentRelation)
	}

	return &paymentRelations, nil
}
