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

// InlineResponse20083 struct for InlineResponse20083
type InlineResponse20083 struct {
	Asset       string `json:"asset"`
	Interest    string `json:"interest"`
	LendingType string `json:"lendingType"`
	ProductName string `json:"productName"`
	Time        int64  `json:"time"`
}

// NewInlineResponse20083 instantiates a new InlineResponse20083 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20083(asset string, interest string, lendingType string, productName string, time int64) *InlineResponse20083 {
	this := InlineResponse20083{}
	this.Asset = asset
	this.Interest = interest
	this.LendingType = lendingType
	this.ProductName = productName
	this.Time = time
	return &this
}

// NewInlineResponse20083WithDefaults instantiates a new InlineResponse20083 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20083WithDefaults() *InlineResponse20083 {
	this := InlineResponse20083{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20083) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20083) SetAsset(v string) {
	o.Asset = v
}

// GetInterest returns the Interest field value
func (o *InlineResponse20083) GetInterest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interest
}

// GetInterestOk returns a tuple with the Interest field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetInterestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interest, true
}

// SetInterest sets field value
func (o *InlineResponse20083) SetInterest(v string) {
	o.Interest = v
}

// GetLendingType returns the LendingType field value
func (o *InlineResponse20083) GetLendingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LendingType
}

// GetLendingTypeOk returns a tuple with the LendingType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetLendingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LendingType, true
}

// SetLendingType sets field value
func (o *InlineResponse20083) SetLendingType(v string) {
	o.LendingType = v
}

// GetProductName returns the ProductName field value
func (o *InlineResponse20083) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *InlineResponse20083) SetProductName(v string) {
	o.ProductName = v
}

// GetTime returns the Time field value
func (o *InlineResponse20083) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20083) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *InlineResponse20083) SetTime(v int64) {
	o.Time = v
}

func (o InlineResponse20083) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["interest"] = o.Interest
	}
	if true {
		toSerialize["lendingType"] = o.LendingType
	}
	if true {
		toSerialize["productName"] = o.ProductName
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20083 struct {
	value *InlineResponse20083
	isSet bool
}

func (v NullableInlineResponse20083) Get() *InlineResponse20083 {
	return v.value
}

func (v *NullableInlineResponse20083) Set(val *InlineResponse20083) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20083) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20083) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20083(val *InlineResponse20083) *NullableInlineResponse20083 {
	return &NullableInlineResponse20083{value: val, isSet: true}
}

func (v NullableInlineResponse20083) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20083) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
