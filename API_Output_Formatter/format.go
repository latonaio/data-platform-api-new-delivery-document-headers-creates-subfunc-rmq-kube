package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"
	"reflect"

	"golang.org/x/xerrors"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Header, error) {
	var err error
	ordersHeaderMap := StructArrayToMap(psdc.OrdersHeader, "OrderID")

	headers := make([]*Header, 0)
	for i, deliveryDocument := range psdc.CalculateDeliveryDocument {
		header := &Header{}
		inputHeader := sdc.Header

		// 入力ファイル
		header, err = jsonTypeConversion(header, inputHeader)
		if err != nil {
			return nil, xerrors.Errorf("request create error: %w", err)
		}

		orderID := deliveryDocument.OrderID
		orderItem := deliveryDocument.OrderItem

		// 1-1
		header, err = jsonTypeConversion(header, ordersHeaderMap[orderID])
		if err != nil {
			return nil, xerrors.Errorf("request create error: %w", err)
		}

		// 1-2
		for _, ordersItem := range psdc.OrdersItem {
			if ordersItem.OrderID == orderID && ordersItem.OrderItem == orderItem {
				header, err = jsonTypeConversion(header, ordersItem)
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}
				break
			}
		}

		header.DeliveryDocument = psdc.CalculateDeliveryDocument[i].DeliveryDocument
		header.OrderID = nil
		header.OrderItem = nil
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
		header.IsCancelled = getBoolPtr(false)
		header.IsMarkedForDeletion = getBoolPtr(false)

		headers = append(headers, header)

	}

	return headers, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func StructArrayToMap[T any](data []T, key string) map[any]T {
	res := make(map[any]T, len(data))

	for _, value := range data {
		m := StructToMap[T](&value, key)
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func StructToMap[T any](data interface{}, key string) map[any]T {
	res := make(map[any]T)
	elem := reflect.Indirect(reflect.ValueOf(data).Elem())
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		if field == key {
			rv := reflect.ValueOf(elem.Field(i).Interface())
			if rv.Kind() == reflect.Ptr {
				if rv.IsNil() {
					return nil
				}
			}
			value := reflect.Indirect(elem.Field(i)).Interface()
			var dist T
			res[value], _ = jsonTypeConversion(dist, elem.Interface())
			break
		}
	}

	return res
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
