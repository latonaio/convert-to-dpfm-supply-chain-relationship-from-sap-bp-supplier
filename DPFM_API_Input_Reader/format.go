package dpfm_api_input_reader

import (
	"convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.SAPBusinessPartnerGeneral {
	data := sdc.SAPBusinessPartnerGeneral
	return &requests.SAPBusinessPartnerGeneral{
		BusinessPartner:               data.BusinessPartner,
		Customer:                      data.Customer,
		Supplier:                      data.Supplier,
		AcademicTitle:                 data.AcademicTitle,
		AuthorizationGroup:            data.AuthorizationGroup,
		BusinessPartnerCategory:       data.BusinessPartnerCategory,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerGrouping:       data.BusinessPartnerGrouping,
		BusinessPartnerName:           data.BusinessPartnerName,
		CorrespondenceLanguage:        data.CorrespondenceLanguage,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		FirstName:                     data.FirstName,
		Industry:                      data.Industry,
		IsFemale:                      data.IsFemale,
		IsMale:                        data.IsMale,
		IsNaturalPerson:               data.IsNaturalPerson,
		IsSexUnknown:                  data.IsSexUnknown,
		GenderCodeName:                data.GenderCodeName,
		Language:                      data.Language,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		LastName:                      data.LastName,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		AdditionalLastName:            data.AdditionalLastName,
		BirthDate:                     data.BirthDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		BusinessPartnerType:           data.BusinessPartnerType,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		IndependentAddressID:          data.IndependentAddressID,
		MiddleName:                    data.MiddleName,
		NameCountry:                   data.NameCountry,
		PersonFullName:                data.PersonFullName,
		PersonNumber:                  data.PersonNumber,
		IsMarkedForArchiving:          data.IsMarkedForArchiving,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		TradingPartner:                data.TradingPartner,
	}
}

func (sdc *SDC) ConvertToSupplier(num int) *requests.SAPBusinessPartnerSupplier {
	data := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier[num]
	return &requests.SAPBusinessPartnerSupplier{
		Supplier:                    data.Supplier,
		AuthorizationGroup:          data.AuthorizationGroup,
		CreationDate:                data.CreationDate,
		Customer:                    data.Customer,
		PaymentIsBlockedForSupplier: data.PaymentIsBlockedForSupplier,
		PostingIsBlocked:            data.PostingIsBlocked,
		PurchasingIsBlocked:         data.PurchasingIsBlocked,
		SupplierAccountGroup:        data.SupplierAccountGroup,
		SupplierFullName:            data.SupplierFullName,
		SupplierName:                data.SupplierName,
		BirthDate:                   data.BirthDate,
		DeletionIndicator:           data.DeletionIndicator,
		Industry:                    data.Industry,
		IsNaturalPerson:             data.IsNaturalPerson,
		SupplierCorporateGroup:      data.SupplierCorporateGroup,
		SupplierProcurementBlock:    data.SupplierProcurementBlock,
	}
}

func (sdc *SDC) ConvertToSupplierPartnerFunction(num, pnum int) *requests.SAPBusinessPartnerSupplierPartnerFunction {
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier[num]
	data := dataSupplier.SAPBusinessPartnerSupplierPartnerFunction[pnum]

	return &requests.SAPBusinessPartnerSupplierPartnerFunction{
		Supplier:               dataSupplier.Supplier,
		PurchasingOrganization: data.PurchasingOrganization,
		PartnerCounter:         data.PartnerCounter,
		PartnerFunction:        data.PartnerFunction,
		Plant:                  data.Plant,
		DefaultPartner:         data.DefaultPartner,
		CreationDate:           data.CreationDate,
		ReferenceSupplier:      data.ReferenceSupplier,
		AuthorizationGroup:     data.AuthorizationGroup,
	}
}

func (sdc *SDC) ConvertToSupplierPurchasingOrganization(num, pnum int) *requests.SAPBusinessPartnerSupplierPurchasingOrganization {
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier[num]
	data := dataSupplier.SAPBusinessPartnerSupplierPurchasingOrganization[pnum]
	return &requests.SAPBusinessPartnerSupplierPurchasingOrganization{
		Supplier:                       dataSupplier.Supplier,
		PurchasingOrganization:         data.PurchasingOrganization,
		IncotermsClassification:        data.IncotermsClassification,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		PaymentTerms:                   data.PaymentTerms,
		PurOrdAutoGenerationIsAllowed:  data.PurOrdAutoGenerationIsAllowed,
		PurchaseOrderCurrency:          data.PurchaseOrderCurrency,
		PurchasingGroup:                data.PurchasingGroup,
		ShippingCondition:              data.ShippingCondition,
		SupplierPhoneNumber:            data.SupplierPhoneNumber,
		SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
		PurchasingIsBlockedForSupplier: data.PurchasingIsBlockedForSupplier,
		DeletionIndicator:              data.DeletionIndicator,
	}
}

func (sdc *SDC) ConvertToSupplierCompany(num, pnum int) *requests.SAPBusinessPartnerSupplierCompany {
	dataSupplier := sdc.SAPBusinessPartnerGeneral.SAPBusinessPartnerSupplier[num]
	data := dataSupplier.SAPBusinessPartnerSupplierCompany[num]
	return &requests.SAPBusinessPartnerSupplierCompany{
		Supplier:                    dataSupplier.Supplier,
		CompanyCode:                 data.CompanyCode,
		PaymentBlockingReason:       data.PaymentBlockingReason,
		PaymentMethodsList:          data.PaymentMethodsList,
		PaymentTerms:                data.PaymentTerms,
		ClearCustomerSupplier:       data.ClearCustomerSupplier,
		HouseBank:                   data.HouseBank,
		ReconciliationAccount:       data.ReconciliationAccount,
		SupplierIsBlockedForPosting: data.SupplierIsBlockedForPosting,
		DeletionIndicator:           data.DeletionIndicator,
	}
}
