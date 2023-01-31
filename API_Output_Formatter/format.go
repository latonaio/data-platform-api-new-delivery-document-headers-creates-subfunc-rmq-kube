package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	var err error

	header := &Header{}
	inputHeader := sdc.Header

	// 入力ファイル
	header, err = jsonTypeConversion(header, inputHeader)
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-1
	header, err = jsonTypeConversion(header, psdc.OrdersHeader[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	// 1-2
	header, err = jsonTypeConversion(header, psdc.OrdersItem[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	header.DeliveryDocument = psdc.CalculateDeliveryDocument.DeliveryDocument
	header.DocumentDate = psdc.DocumentDate.DocumentDate
	header.InvoiceDocumentDate = psdc.InvoiceDocumentDate.InvoiceDocumentDate
	header.HeaderCompleteDeliveryIsDefined = getBoolPtr(false)
	header.HeaderDeliveryStatus = getStringPtr("NP")
	header.CreationDate = psdc.CreationDateHeader.CreationDate
	header.CreationTime = psdc.CreationTimeHeader.CreationTime
	header.LastChangeDate = psdc.LastChangeDateHeader.LastChangeDate
	header.LastChangeTime = psdc.LastChangeTimeHeader.LastChangeTime
	header.HeaderBillingStatus = getStringPtr("NP")
	header.HeaderBillingConfStatus = getStringPtr("NP")
	header.HeaderGrossWeight = psdc.HeaderGrossWeight.HeaderGrossWeight
	header.HeaderNetWeight = psdc.HeaderNetWeight.HeaderNetWeight
	header.HeaderDeliveryBlockStatus = getBoolPtr(false)
	header.HeaderIssuingBlockStatus = getBoolPtr(false)
	header.HeaderReceivingBlockStatus = getBoolPtr(false)
	header.HeaderBillingBlockStatus = getBoolPtr(false)

	return header, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
