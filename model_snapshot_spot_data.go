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

// SnapshotSpotData struct for SnapshotSpotData
type SnapshotSpotData struct {
	Balances        []SnapshotSpotDataBalances `json:"balances"`
	TotalAssetOfBtc string                     `json:"totalAssetOfBtc"`
}

// NewSnapshotSpotData instantiates a new SnapshotSpotData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSpotData(balances []SnapshotSpotDataBalances, totalAssetOfBtc string) *SnapshotSpotData {
	this := SnapshotSpotData{}
	this.Balances = balances
	this.TotalAssetOfBtc = totalAssetOfBtc
	return &this
}

// NewSnapshotSpotDataWithDefaults instantiates a new SnapshotSpotData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotSpotDataWithDefaults() *SnapshotSpotData {
	this := SnapshotSpotData{}
	return &this
}

// GetBalances returns the Balances field value
func (o *SnapshotSpotData) GetBalances() []SnapshotSpotDataBalances {
	if o == nil {
		var ret []SnapshotSpotDataBalances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *SnapshotSpotData) GetBalancesOk() (*[]SnapshotSpotDataBalances, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *SnapshotSpotData) SetBalances(v []SnapshotSpotDataBalances) {
	o.Balances = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value
func (o *SnapshotSpotData) GetTotalAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *SnapshotSpotData) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAssetOfBtc, true
}

// SetTotalAssetOfBtc sets field value
func (o *SnapshotSpotData) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = v
}

func (o SnapshotSpotData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotSpotData struct {
	value *SnapshotSpotData
	isSet bool
}

func (v NullableSnapshotSpotData) Get() *SnapshotSpotData {
	return v.value
}

func (v *NullableSnapshotSpotData) Set(val *SnapshotSpotData) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSpotData) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSpotData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSpotData(val *SnapshotSpotData) *NullableSnapshotSpotData {
	return &NullableSnapshotSpotData{value: val, isSet: true}
}

func (v NullableSnapshotSpotData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSpotData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
