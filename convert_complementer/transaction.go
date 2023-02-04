package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Transactionの処理
func (c *ConvertComplementer) ComplementMappingTransaction(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingTransaction, error) {
	res, err := psdc.ConvertToMappingTransaction(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionTransaction(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionTransaction, error) {
	var data []dpfm_api_processing_formatter.CodeConversionTransaction

	for _, supplier := range sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier {
		var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
		var args []interface{}

		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Supplier", "Seller", supplier.Supplier))

		repeat := strings.Repeat("?,", len(dataKey)-1) + "?,"
		for _, v := range dataKey {
			args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
		}

		rows, err := c.db.Query(
			`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
			CodeConvertFrom, CodeConvertTo, BusinessPartner
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
			WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
		if err != nil {
			return nil, err
		}

		datum, err := psdc.ConvertToCodeConversionTransaction(dataQueryGets)
		if err != nil {
			return nil, err
		}

		data = append(data, *datum)
	}

	return &data, nil
}
