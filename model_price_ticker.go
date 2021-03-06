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

// PriceTicker struct for PriceTicker
type PriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// NewPriceTicker instantiates a new PriceTicker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTicker(symbol string, price string) *PriceTicker {
	this := PriceTicker{}
	this.Symbol = symbol
	this.Price = price
	return &this
}

// NewPriceTickerWithDefaults instantiates a new PriceTicker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTickerWithDefaults() *PriceTicker {
	this := PriceTicker{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *PriceTicker) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *PriceTicker) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *PriceTicker) SetSymbol(v string) {
	o.Symbol = v
}

// GetPrice returns the Price field value
func (o *PriceTicker) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *PriceTicker) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *PriceTicker) SetPrice(v string) {
	o.Price = v
}

func (o PriceTicker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullablePriceTicker struct {
	value *PriceTicker
	isSet bool
}

func (v NullablePriceTicker) Get() *PriceTicker {
	return v.value
}

func (v *NullablePriceTicker) Set(val *PriceTicker) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTicker) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTicker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTicker(val *PriceTicker) *NullablePriceTicker {
	return &NullablePriceTicker{value: val, isSet: true}
}

func (v NullablePriceTicker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTicker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
