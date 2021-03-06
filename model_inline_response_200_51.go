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

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	Balances []InlineResponse20051Balances `json:"balances"`
}

// NewInlineResponse20051 instantiates a new InlineResponse20051 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051(balances []InlineResponse20051Balances) *InlineResponse20051 {
	this := InlineResponse20051{}
	this.Balances = balances
	return &this
}

// NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051WithDefaults() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// GetBalances returns the Balances field value
func (o *InlineResponse20051) GetBalances() []InlineResponse20051Balances {
	if o == nil {
		var ret []InlineResponse20051Balances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetBalancesOk() (*[]InlineResponse20051Balances, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *InlineResponse20051) SetBalances(v []InlineResponse20051Balances) {
	o.Balances = v
}

func (o InlineResponse20051) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balances"] = o.Balances
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051 struct {
	value *InlineResponse20051
	isSet bool
}

func (v NullableInlineResponse20051) Get() *InlineResponse20051 {
	return v.value
}

func (v *NullableInlineResponse20051) Set(val *InlineResponse20051) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051(val *InlineResponse20051) *NullableInlineResponse20051 {
	return &NullableInlineResponse20051{value: val, isSet: true}
}

func (v NullableInlineResponse20051) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
