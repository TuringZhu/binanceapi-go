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

// InlineResponse20012 struct for InlineResponse20012
type InlineResponse20012 struct {
	AssetFullName  string `json:"assetFullName"`
	AssetName      string `json:"assetName"`
	IsBorrowable   bool   `json:"isBorrowable"`
	IsMortgageable bool   `json:"isMortgageable"`
	UserMinBorrow  string `json:"userMinBorrow"`
	UserMinRepay   string `json:"userMinRepay"`
}

// NewInlineResponse20012 instantiates a new InlineResponse20012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012(assetFullName string, assetName string, isBorrowable bool, isMortgageable bool, userMinBorrow string, userMinRepay string) *InlineResponse20012 {
	this := InlineResponse20012{}
	this.AssetFullName = assetFullName
	this.AssetName = assetName
	this.IsBorrowable = isBorrowable
	this.IsMortgageable = isMortgageable
	this.UserMinBorrow = userMinBorrow
	this.UserMinRepay = userMinRepay
	return &this
}

// NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012WithDefaults() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// GetAssetFullName returns the AssetFullName field value
func (o *InlineResponse20012) GetAssetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetFullName
}

// GetAssetFullNameOk returns a tuple with the AssetFullName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetAssetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetFullName, true
}

// SetAssetFullName sets field value
func (o *InlineResponse20012) SetAssetFullName(v string) {
	o.AssetFullName = v
}

// GetAssetName returns the AssetName field value
func (o *InlineResponse20012) GetAssetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetAssetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetName, true
}

// SetAssetName sets field value
func (o *InlineResponse20012) SetAssetName(v string) {
	o.AssetName = v
}

// GetIsBorrowable returns the IsBorrowable field value
func (o *InlineResponse20012) GetIsBorrowable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBorrowable
}

// GetIsBorrowableOk returns a tuple with the IsBorrowable field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetIsBorrowableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBorrowable, true
}

// SetIsBorrowable sets field value
func (o *InlineResponse20012) SetIsBorrowable(v bool) {
	o.IsBorrowable = v
}

// GetIsMortgageable returns the IsMortgageable field value
func (o *InlineResponse20012) GetIsMortgageable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMortgageable
}

// GetIsMortgageableOk returns a tuple with the IsMortgageable field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetIsMortgageableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMortgageable, true
}

// SetIsMortgageable sets field value
func (o *InlineResponse20012) SetIsMortgageable(v bool) {
	o.IsMortgageable = v
}

// GetUserMinBorrow returns the UserMinBorrow field value
func (o *InlineResponse20012) GetUserMinBorrow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserMinBorrow
}

// GetUserMinBorrowOk returns a tuple with the UserMinBorrow field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetUserMinBorrowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMinBorrow, true
}

// SetUserMinBorrow sets field value
func (o *InlineResponse20012) SetUserMinBorrow(v string) {
	o.UserMinBorrow = v
}

// GetUserMinRepay returns the UserMinRepay field value
func (o *InlineResponse20012) GetUserMinRepay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserMinRepay
}

// GetUserMinRepayOk returns a tuple with the UserMinRepay field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetUserMinRepayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMinRepay, true
}

// SetUserMinRepay sets field value
func (o *InlineResponse20012) SetUserMinRepay(v string) {
	o.UserMinRepay = v
}

func (o InlineResponse20012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assetFullName"] = o.AssetFullName
	}
	if true {
		toSerialize["assetName"] = o.AssetName
	}
	if true {
		toSerialize["isBorrowable"] = o.IsBorrowable
	}
	if true {
		toSerialize["isMortgageable"] = o.IsMortgageable
	}
	if true {
		toSerialize["userMinBorrow"] = o.UserMinBorrow
	}
	if true {
		toSerialize["userMinRepay"] = o.UserMinRepay
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012 struct {
	value *InlineResponse20012
	isSet bool
}

func (v NullableInlineResponse20012) Get() *InlineResponse20012 {
	return v.value
}

func (v *NullableInlineResponse20012) Set(val *InlineResponse20012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012(val *InlineResponse20012) *NullableInlineResponse20012 {
	return &NullableInlineResponse20012{value: val, isSet: true}
}

func (v NullableInlineResponse20012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
