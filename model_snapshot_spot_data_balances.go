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

// SnapshotSpotDataBalances struct for SnapshotSpotDataBalances
type SnapshotSpotDataBalances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// NewSnapshotSpotDataBalances instantiates a new SnapshotSpotDataBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSpotDataBalances(asset string, free string, locked string) *SnapshotSpotDataBalances {
	this := SnapshotSpotDataBalances{}
	this.Asset = asset
	this.Free = free
	this.Locked = locked
	return &this
}

// NewSnapshotSpotDataBalancesWithDefaults instantiates a new SnapshotSpotDataBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotSpotDataBalancesWithDefaults() *SnapshotSpotDataBalances {
	this := SnapshotSpotDataBalances{}
	return &this
}

// GetAsset returns the Asset field value
func (o *SnapshotSpotDataBalances) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *SnapshotSpotDataBalances) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *SnapshotSpotDataBalances) SetAsset(v string) {
	o.Asset = v
}

// GetFree returns the Free field value
func (o *SnapshotSpotDataBalances) GetFree() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *SnapshotSpotDataBalances) GetFreeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *SnapshotSpotDataBalances) SetFree(v string) {
	o.Free = v
}

// GetLocked returns the Locked field value
func (o *SnapshotSpotDataBalances) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *SnapshotSpotDataBalances) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *SnapshotSpotDataBalances) SetLocked(v string) {
	o.Locked = v
}

func (o SnapshotSpotDataBalances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["free"] = o.Free
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotSpotDataBalances struct {
	value *SnapshotSpotDataBalances
	isSet bool
}

func (v NullableSnapshotSpotDataBalances) Get() *SnapshotSpotDataBalances {
	return v.value
}

func (v *NullableSnapshotSpotDataBalances) Set(val *SnapshotSpotDataBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSpotDataBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSpotDataBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSpotDataBalances(val *SnapshotSpotDataBalances) *NullableSnapshotSpotDataBalances {
	return &NullableSnapshotSpotDataBalances{value: val, isSet: true}
}

func (v NullableSnapshotSpotDataBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSpotDataBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
