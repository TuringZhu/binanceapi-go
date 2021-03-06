/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API generated with [binance/binance-api-swagger/blob/master/spot_api.yaml](https://github.com/binance/binance-api-swagger/blob/master/spot_api.yaml) with commit [v1.2.0 release](https://github.com/binance/binance-api-swagger/commit/60d14be031c031600c853d5cdab86db5ab73603e)  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
Contact: qishiwenjun@163.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binanceapi

import (
	"encoding/json"
)

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	OrderListId       int64                            `json:"orderListId"`
	ContingencyType   string                           `json:"contingencyType"`
	ListStatusType    string                           `json:"listStatusType"`
	ListOrderStatus   string                           `json:"listOrderStatus"`
	ListClientOrderId string                           `json:"listClientOrderId"`
	TransactionTime   int64                            `json:"transactionTime"`
	Symbol            string                           `json:"symbol"`
	Orders            []InlineResponse2004Orders       `json:"orders"`
	OrderReports      []InlineResponse2004OrderReports `json:"orderReports"`
}

// NewInlineResponse2004 instantiates a new InlineResponse2004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004(orderListId int64, contingencyType string, listStatusType string, listOrderStatus string, listClientOrderId string, transactionTime int64, symbol string, orders []InlineResponse2004Orders, orderReports []InlineResponse2004OrderReports) *InlineResponse2004 {
	this := InlineResponse2004{}
	this.OrderListId = orderListId
	this.ContingencyType = contingencyType
	this.ListStatusType = listStatusType
	this.ListOrderStatus = listOrderStatus
	this.ListClientOrderId = listClientOrderId
	this.TransactionTime = transactionTime
	this.Symbol = symbol
	this.Orders = orders
	this.OrderReports = orderReports
	return &this
}

// NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004WithDefaults() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// GetOrderListId returns the OrderListId field value
func (o *InlineResponse2004) GetOrderListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetOrderListIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderListId, true
}

// SetOrderListId sets field value
func (o *InlineResponse2004) SetOrderListId(v int64) {
	o.OrderListId = v
}

// GetContingencyType returns the ContingencyType field value
func (o *InlineResponse2004) GetContingencyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetContingencyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContingencyType, true
}

// SetContingencyType sets field value
func (o *InlineResponse2004) SetContingencyType(v string) {
	o.ContingencyType = v
}

// GetListStatusType returns the ListStatusType field value
func (o *InlineResponse2004) GetListStatusType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetListStatusTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListStatusType, true
}

// SetListStatusType sets field value
func (o *InlineResponse2004) SetListStatusType(v string) {
	o.ListStatusType = v
}

// GetListOrderStatus returns the ListOrderStatus field value
func (o *InlineResponse2004) GetListOrderStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetListOrderStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListOrderStatus, true
}

// SetListOrderStatus sets field value
func (o *InlineResponse2004) SetListOrderStatus(v string) {
	o.ListOrderStatus = v
}

// GetListClientOrderId returns the ListClientOrderId field value
func (o *InlineResponse2004) GetListClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetListClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListClientOrderId, true
}

// SetListClientOrderId sets field value
func (o *InlineResponse2004) SetListClientOrderId(v string) {
	o.ListClientOrderId = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *InlineResponse2004) GetTransactionTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetTransactionTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *InlineResponse2004) SetTransactionTime(v int64) {
	o.TransactionTime = v
}

// GetSymbol returns the Symbol field value
func (o *InlineResponse2004) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineResponse2004) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrders returns the Orders field value
func (o *InlineResponse2004) GetOrders() []InlineResponse2004Orders {
	if o == nil {
		var ret []InlineResponse2004Orders
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetOrdersOk() (*[]InlineResponse2004Orders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orders, true
}

// SetOrders sets field value
func (o *InlineResponse2004) SetOrders(v []InlineResponse2004Orders) {
	o.Orders = v
}

// GetOrderReports returns the OrderReports field value
func (o *InlineResponse2004) GetOrderReports() []InlineResponse2004OrderReports {
	if o == nil {
		var ret []InlineResponse2004OrderReports
		return ret
	}

	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetOrderReportsOk() (*[]InlineResponse2004OrderReports, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderReports, true
}

// SetOrderReports sets field value
func (o *InlineResponse2004) SetOrderReports(v []InlineResponse2004OrderReports) {
	o.OrderReports = v
}

func (o InlineResponse2004) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orderListId"] = o.OrderListId
	}
	if true {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if true {
		toSerialize["listStatusType"] = o.ListStatusType
	}
	if true {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if true {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if true {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["orders"] = o.Orders
	}
	if true {
		toSerialize["orderReports"] = o.OrderReports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004 struct {
	value *InlineResponse2004
	isSet bool
}

func (v NullableInlineResponse2004) Get() *InlineResponse2004 {
	return v.value
}

func (v *NullableInlineResponse2004) Set(val *InlineResponse2004) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004(val *InlineResponse2004) *NullableInlineResponse2004 {
	return &NullableInlineResponse2004{value: val, isSet: true}
}

func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
