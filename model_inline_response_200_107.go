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

// InlineResponse200107 struct for InlineResponse200107
type InlineResponse200107 struct {
	QuoteAsset string  `json:"quoteAsset"`
	BaseAsset  string  `json:"baseAsset"`
	QuoteQty   float64 `json:"quoteQty"`
	BaseQty    float64 `json:"baseQty"`
	Price      float64 `json:"price"`
	Slippage   float64 `json:"slippage"`
	Fee        float64 `json:"fee"`
}

// NewInlineResponse200107 instantiates a new InlineResponse200107 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200107(quoteAsset string, baseAsset string, quoteQty float64, baseQty float64, price float64, slippage float64, fee float64) *InlineResponse200107 {
	this := InlineResponse200107{}
	this.QuoteAsset = quoteAsset
	this.BaseAsset = baseAsset
	this.QuoteQty = quoteQty
	this.BaseQty = baseQty
	this.Price = price
	this.Slippage = slippage
	this.Fee = fee
	return &this
}

// NewInlineResponse200107WithDefaults instantiates a new InlineResponse200107 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200107WithDefaults() *InlineResponse200107 {
	this := InlineResponse200107{}
	return &this
}

// GetQuoteAsset returns the QuoteAsset field value
func (o *InlineResponse200107) GetQuoteAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetQuoteAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteAsset, true
}

// SetQuoteAsset sets field value
func (o *InlineResponse200107) SetQuoteAsset(v string) {
	o.QuoteAsset = v
}

// GetBaseAsset returns the BaseAsset field value
func (o *InlineResponse200107) GetBaseAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseAsset
}

// GetBaseAssetOk returns a tuple with the BaseAsset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetBaseAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseAsset, true
}

// SetBaseAsset sets field value
func (o *InlineResponse200107) SetBaseAsset(v string) {
	o.BaseAsset = v
}

// GetQuoteQty returns the QuoteQty field value
func (o *InlineResponse200107) GetQuoteQty() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetQuoteQtyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteQty, true
}

// SetQuoteQty sets field value
func (o *InlineResponse200107) SetQuoteQty(v float64) {
	o.QuoteQty = v
}

// GetBaseQty returns the BaseQty field value
func (o *InlineResponse200107) GetBaseQty() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BaseQty
}

// GetBaseQtyOk returns a tuple with the BaseQty field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetBaseQtyOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQty, true
}

// SetBaseQty sets field value
func (o *InlineResponse200107) SetBaseQty(v float64) {
	o.BaseQty = v
}

// GetPrice returns the Price field value
func (o *InlineResponse200107) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InlineResponse200107) SetPrice(v float64) {
	o.Price = v
}

// GetSlippage returns the Slippage field value
func (o *InlineResponse200107) GetSlippage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Slippage
}

// GetSlippageOk returns a tuple with the Slippage field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetSlippageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slippage, true
}

// SetSlippage sets field value
func (o *InlineResponse200107) SetSlippage(v float64) {
	o.Slippage = v
}

// GetFee returns the Fee field value
func (o *InlineResponse200107) GetFee() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200107) GetFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *InlineResponse200107) SetFee(v float64) {
	o.Fee = v
}

func (o InlineResponse200107) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if true {
		toSerialize["baseAsset"] = o.BaseAsset
	}
	if true {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if true {
		toSerialize["baseQty"] = o.BaseQty
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["slippage"] = o.Slippage
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200107 struct {
	value *InlineResponse200107
	isSet bool
}

func (v NullableInlineResponse200107) Get() *InlineResponse200107 {
	return v.value
}

func (v *NullableInlineResponse200107) Set(val *InlineResponse200107) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200107) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200107) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200107(val *InlineResponse200107) *NullableInlineResponse200107 {
	return &NullableInlineResponse200107{value: val, isSet: true}
}

func (v NullableInlineResponse200107) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200107) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
