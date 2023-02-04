package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Processing_Formatter"
)

// Mapping Payment Relationの処理
func (c *ConvertComplementer) ComplementMappingPaymentRelation(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingPaymentRelation, error) {
	res, err := psdc.ConvertToMappingPaymentRelation(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
