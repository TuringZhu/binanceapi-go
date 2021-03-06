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

// InlineResponse20038 struct for InlineResponse20038
type InlineResponse20038 struct {
	TotalServiceCharge string                              `json:"totalServiceCharge"`
	TotalTransfered    string                              `json:"totalTransfered"`
	TransferResult     []InlineResponse20038TransferResult `json:"transferResult"`
}

// NewInlineResponse20038 instantiates a new InlineResponse20038 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038(totalServiceCharge string, totalTransfered string, transferResult []InlineResponse20038TransferResult) *InlineResponse20038 {
	this := InlineResponse20038{}
	this.TotalServiceCharge = totalServiceCharge
	this.TotalTransfered = totalTransfered
	this.TransferResult = transferResult
	return &this
}

// NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038WithDefaults() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// GetTotalServiceCharge returns the TotalServiceCharge field value
func (o *InlineResponse20038) GetTotalServiceCharge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalServiceCharge
}

// GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetTotalServiceChargeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalServiceCharge, true
}

// SetTotalServiceCharge sets field value
func (o *InlineResponse20038) SetTotalServiceCharge(v string) {
	o.TotalServiceCharge = v
}

// GetTotalTransfered returns the TotalTransfered field value
func (o *InlineResponse20038) GetTotalTransfered() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalTransfered
}

// GetTotalTransferedOk returns a tuple with the TotalTransfered field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetTotalTransferedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTransfered, true
}

// SetTotalTransfered sets field value
func (o *InlineResponse20038) SetTotalTransfered(v string) {
	o.TotalTransfered = v
}

// GetTransferResult returns the TransferResult field value
func (o *InlineResponse20038) GetTransferResult() []InlineResponse20038TransferResult {
	if o == nil {
		var ret []InlineResponse20038TransferResult
		return ret
	}

	return o.TransferResult
}

// GetTransferResultOk returns a tuple with the TransferResult field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetTransferResultOk() (*[]InlineResponse20038TransferResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferResult, true
}

// SetTransferResult sets field value
func (o *InlineResponse20038) SetTransferResult(v []InlineResponse20038TransferResult) {
	o.TransferResult = v
}

func (o InlineResponse20038) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalServiceCharge"] = o.TotalServiceCharge
	}
	if true {
		toSerialize["totalTransfered"] = o.TotalTransfered
	}
	if true {
		toSerialize["transferResult"] = o.TransferResult
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038 struct {
	value *InlineResponse20038
	isSet bool
}

func (v NullableInlineResponse20038) Get() *InlineResponse20038 {
	return v.value
}

func (v *NullableInlineResponse20038) Set(val *InlineResponse20038) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038(val *InlineResponse20038) *NullableInlineResponse20038 {
	return &NullableInlineResponse20038{value: val, isSet: true}
}

func (v NullableInlineResponse20038) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
