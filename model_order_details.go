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

// OrderDetails struct for OrderDetails
type OrderDetails struct {
	Symbol  string `json:"symbol"`
	OrderId int64  `json:"orderId"`
	// Unless OCO, value will be -1
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                int64  `json:"time"`
	UpdateTime          int64  `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
}

// NewOrderDetails instantiates a new OrderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetails(symbol string, orderId int64, orderListId int64, clientOrderId string, price string, origQty string, executedQty string, cummulativeQuoteQty string, status string, timeInForce string, type_ string, side string, stopPrice string, icebergQty string, time int64, updateTime int64, isWorking bool, origQuoteOrderQty string) *OrderDetails {
	this := OrderDetails{}
	this.Symbol = symbol
	this.OrderId = orderId
	this.OrderListId = orderListId
	this.ClientOrderId = clientOrderId
	this.Price = price
	this.OrigQty = origQty
	this.ExecutedQty = executedQty
	this.CummulativeQuoteQty = cummulativeQuoteQty
	this.Status = status
	this.TimeInForce = timeInForce
	this.Type = type_
	this.Side = side
	this.StopPrice = stopPrice
	this.IcebergQty = icebergQty
	this.Time = time
	this.UpdateTime = updateTime
	this.IsWorking = isWorking
	this.OrigQuoteOrderQty = origQuoteOrderQty
	return &this
}

// NewOrderDetailsWithDefaults instantiates a new OrderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailsWithDefaults() *OrderDetails {
	this := OrderDetails{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *OrderDetails) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *OrderDetails) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrderId returns the OrderId field value
func (o *OrderDetails) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderDetails) SetOrderId(v int64) {
	o.OrderId = v
}

// GetOrderListId returns the OrderListId field value
func (o *OrderDetails) GetOrderListId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetOrderListIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderListId, true
}

// SetOrderListId sets field value
func (o *OrderDetails) SetOrderListId(v int64) {
	o.OrderListId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *OrderDetails) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *OrderDetails) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetPrice returns the Price field value
func (o *OrderDetails) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *OrderDetails) SetPrice(v string) {
	o.Price = v
}

// GetOrigQty returns the OrigQty field value
func (o *OrderDetails) GetOrigQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetOrigQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrigQty, true
}

// SetOrigQty sets field value
func (o *OrderDetails) SetOrigQty(v string) {
	o.OrigQty = v
}

// GetExecutedQty returns the ExecutedQty field value
func (o *OrderDetails) GetExecutedQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetExecutedQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutedQty, true
}

// SetExecutedQty sets field value
func (o *OrderDetails) SetExecutedQty(v string) {
	o.ExecutedQty = v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value
func (o *OrderDetails) GetCummulativeQuoteQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CummulativeQuoteQty, true
}

// SetCummulativeQuoteQty sets field value
func (o *OrderDetails) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = v
}

// GetStatus returns the Status field value
func (o *OrderDetails) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderDetails) SetStatus(v string) {
	o.Status = v
}

// GetTimeInForce returns the TimeInForce field value
func (o *OrderDetails) GetTimeInForce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetTimeInForceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInForce, true
}

// SetTimeInForce sets field value
func (o *OrderDetails) SetTimeInForce(v string) {
	o.TimeInForce = v
}

// GetType returns the Type field value
func (o *OrderDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderDetails) SetType(v string) {
	o.Type = v
}

// GetSide returns the Side field value
func (o *OrderDetails) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *OrderDetails) SetSide(v string) {
	o.Side = v
}

// GetStopPrice returns the StopPrice field value
func (o *OrderDetails) GetStopPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetStopPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopPrice, true
}

// SetStopPrice sets field value
func (o *OrderDetails) SetStopPrice(v string) {
	o.StopPrice = v
}

// GetIcebergQty returns the IcebergQty field value
func (o *OrderDetails) GetIcebergQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetIcebergQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IcebergQty, true
}

// SetIcebergQty sets field value
func (o *OrderDetails) SetIcebergQty(v string) {
	o.IcebergQty = v
}

// GetTime returns the Time field value
func (o *OrderDetails) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *OrderDetails) SetTime(v int64) {
	o.Time = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *OrderDetails) GetUpdateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetUpdateTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *OrderDetails) SetUpdateTime(v int64) {
	o.UpdateTime = v
}

// GetIsWorking returns the IsWorking field value
func (o *OrderDetails) GetIsWorking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWorking
}

// GetIsWorkingOk returns a tuple with the IsWorking field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetIsWorkingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWorking, true
}

// SetIsWorking sets field value
func (o *OrderDetails) SetIsWorking(v bool) {
	o.IsWorking = v
}

// GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field value
func (o *OrderDetails) GetOrigQuoteOrderQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrigQuoteOrderQty
}

// GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field value
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetOrigQuoteOrderQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrigQuoteOrderQty, true
}

// SetOrigQuoteOrderQty sets field value
func (o *OrderDetails) SetOrigQuoteOrderQty(v string) {
	o.OrigQuoteOrderQty = v
}

func (o OrderDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["orderListId"] = o.OrderListId
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["origQty"] = o.OrigQty
	}
	if true {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if true {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["side"] = o.Side
	}
	if true {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if true {
		toSerialize["icebergQty"] = o.IcebergQty
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if true {
		toSerialize["isWorking"] = o.IsWorking
	}
	if true {
		toSerialize["origQuoteOrderQty"] = o.OrigQuoteOrderQty
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDetails struct {
	value *OrderDetails
	isSet bool
}

func (v NullableOrderDetails) Get() *OrderDetails {
	return v.value
}

func (v *NullableOrderDetails) Set(val *OrderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetails(val *OrderDetails) *NullableOrderDetails {
	return &NullableOrderDetails{value: val, isSet: true}
}

func (v NullableOrderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
