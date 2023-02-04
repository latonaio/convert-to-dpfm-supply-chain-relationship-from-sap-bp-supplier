package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier/DPFM_API_Processing_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ConvertComplementer struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewConvertComplementer(ctx context.Context, db *database.Mysql, l *logger.Logger) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (c *ConvertComplementer) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Transaction
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-0. データ連携基盤Supply Chain Relationship TransactionとSAP Business Partner Supplierとの項目マッピング変換
		psdc.MappingTransaction, e = c.ComplementMappingTransaction(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-1. データ連携基盤Supply Chain Relationship Delivery RelationとSAP Business Partner Supplierとの項目マッピング変換
		psdc.MappingDeliveryRelation, e = c.ComplementMappingDeliveryRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <1-2. 変換元のSupplier(DeliverFromParty)のセット>
		psdc.SetSupplierDeliveryRelation = c.SetSupplierDeliveryRelation(sdc, psdc)

		// <1-1. 番号変換＞
		psdc.CodeConversionDeliveryRelation, e = c.CodeConversionDeliveryRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// BillingRelation
		// <5-3. 変換元のSupplier(BillFromParty)のセット>
		psdc.SetSupplierBillingRelation = c.SetSupplierBillingRelation(sdc, psdc)
		// <5-1. 番号変換＞
		psdc.CodeConversionBillingRelation, e = c.CodeConversionBillingRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1-3．Sellerの保持 (Supply Chain Relationship Transaction),
		psdc.ConversionData = c.ConversionData(sdc, psdc)
	}(&wg)

	// PaymentRelation
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 8-1. データ連携基盤Supply Chain Relationship Payment RelationとSAP Business Partner Supplierとの項目マッピング変換
		psdc.MappingPaymentRelation, e = c.ComplementMappingPaymentRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	c.l.Info(psdc)
	osdc, err = c.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
