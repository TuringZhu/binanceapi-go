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

// AllOrderListOrders struct for AllOrderListOrders
type AllOrderListOrders struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

// NewAllOrderListOrders instantiates a new AllOrderListOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllOrderListOrders(symbol string, orderId int64, clientOrderId string) *AllOrderListOrders {
	this := AllOrderListOrders{}
	this.Symbol = symbol
	this.OrderId = orderId
	this.ClientOrderId = clientOrderId
	return &this
}

// NewAllOrderListOrdersWithDefaults instantiates a new AllOrderListOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllOrderListOrdersWithDefaults() *AllOrderListOrders {
	this := AllOrderListOrders{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *AllOrderListOrders) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *AllOrderListOrders) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *AllOrderListOrders) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrderId returns the OrderId field value
func (o *AllOrderListOrders) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *AllOrderListOrders) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *AllOrderListOrders) SetOrderId(v int64) {
	o.OrderId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *AllOrderListOrders) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *AllOrderListOrders) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *AllOrderListOrders) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

func (o AllOrderListOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	return json.Marshal(toSerialize)
}

type NullableAllOrderListOrders struct {
	value *AllOrderListOrders
	isSet bool
}

func (v NullableAllOrderListOrders) Get() *AllOrderListOrders {
	return v.value
}

func (v *NullableAllOrderListOrders) Set(val *AllOrderListOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableAllOrderListOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableAllOrderListOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllOrderListOrders(val *AllOrderListOrders) *NullableAllOrderListOrders {
	return &NullableAllOrderListOrders{value: val, isSet: true}
}

func (v NullableAllOrderListOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllOrderListOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
