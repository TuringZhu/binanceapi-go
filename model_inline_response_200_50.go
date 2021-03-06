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

// InlineResponse20050 struct for InlineResponse20050
type InlineResponse20050 struct {
	Success bool   `json:"success"`
	TxnId   string `json:"txnId"`
}

// NewInlineResponse20050 instantiates a new InlineResponse20050 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20050(success bool, txnId string) *InlineResponse20050 {
	this := InlineResponse20050{}
	this.Success = success
	this.TxnId = txnId
	return &this
}

// NewInlineResponse20050WithDefaults instantiates a new InlineResponse20050 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20050WithDefaults() *InlineResponse20050 {
	this := InlineResponse20050{}
	return &this
}

// GetSuccess returns the Success field value
func (o *InlineResponse20050) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20050) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *InlineResponse20050) SetSuccess(v bool) {
	o.Success = v
}

// GetTxnId returns the TxnId field value
func (o *InlineResponse20050) GetTxnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20050) GetTxnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnId, true
}

// SetTxnId sets field value
func (o *InlineResponse20050) SetTxnId(v string) {
	o.TxnId = v
}

func (o InlineResponse20050) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["txnId"] = o.TxnId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20050 struct {
	value *InlineResponse20050
	isSet bool
}

func (v NullableInlineResponse20050) Get() *InlineResponse20050 {
	return v.value
}

func (v *NullableInlineResponse20050) Set(val *InlineResponse20050) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20050) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20050) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20050(val *InlineResponse20050) *NullableInlineResponse20050 {
	return &NullableInlineResponse20050{value: val, isSet: true}
}

func (v NullableInlineResponse20050) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20050) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
