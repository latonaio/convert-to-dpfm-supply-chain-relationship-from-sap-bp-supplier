package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Processing_Formatter"
)

func (c *ConvertComplementer) SetValue(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var transaction *[]dpfm_api_output_formatter.Transaction
	var deliveryRelation *[]dpfm_api_output_formatter.DeliveryRelation
	var billingRelation *[]dpfm_api_output_formatter.BillingRelation
	var paymentRelation *[]dpfm_api_output_formatter.PaymentRelation
	var err error

	transaction, err = dpfm_api_output_formatter.ConvertToTransaction(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	deliveryRelation, err = dpfm_api_output_formatter.ConvertToDeliveryRelation(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	billingRelation, err = dpfm_api_output_formatter.ConvertToBillingRelation(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	paymentRelation, err = dpfm_api_output_formatter.ConvertToPaymentRelation(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Transaction:      *transaction,
		DeliveryRelation: *deliveryRelation,
		BillingRelation:  *billingRelation,
		PaymentRelation:  *paymentRelation,
	}

	return osdc, nil
}
