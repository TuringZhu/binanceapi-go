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

// InlineResponse20019UserAssets struct for InlineResponse20019UserAssets
type InlineResponse20019UserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// NewInlineResponse20019UserAssets instantiates a new InlineResponse20019UserAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20019UserAssets(asset string, borrowed string, free string, interest string, locked string, netAsset string) *InlineResponse20019UserAssets {
	this := InlineResponse20019UserAssets{}
	this.Asset = asset
	this.Borrowed = borrowed
	this.Free = free
	this.Interest = interest
	this.Locked = locked
	this.NetAsset = netAsset
	return &this
}

// NewInlineResponse20019UserAssetsWithDefaults instantiates a new InlineResponse20019UserAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20019UserAssetsWithDefaults() *InlineResponse20019UserAssets {
	this := InlineResponse20019UserAssets{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20019UserAssets) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20019UserAssets) SetAsset(v string) {
	o.Asset = v
}

// GetBorrowed returns the Borrowed field value
func (o *InlineResponse20019UserAssets) GetBorrowed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetBorrowedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Borrowed, true
}

// SetBorrowed sets field value
func (o *InlineResponse20019UserAssets) SetBorrowed(v string) {
	o.Borrowed = v
}

// GetFree returns the Free field value
func (o *InlineResponse20019UserAssets) GetFree() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetFreeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *InlineResponse20019UserAssets) SetFree(v string) {
	o.Free = v
}

// GetInterest returns the Interest field value
func (o *InlineResponse20019UserAssets) GetInterest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interest
}

// GetInterestOk returns a tuple with the Interest field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetInterestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interest, true
}

// SetInterest sets field value
func (o *InlineResponse20019UserAssets) SetInterest(v string) {
	o.Interest = v
}

// GetLocked returns the Locked field value
func (o *InlineResponse20019UserAssets) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *InlineResponse20019UserAssets) SetLocked(v string) {
	o.Locked = v
}

// GetNetAsset returns the NetAsset field value
func (o *InlineResponse20019UserAssets) GetNetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20019UserAssets) GetNetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetAsset, true
}

// SetNetAsset sets field value
func (o *InlineResponse20019UserAssets) SetNetAsset(v string) {
	o.NetAsset = v
}

func (o InlineResponse20019UserAssets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["borrowed"] = o.Borrowed
	}
	if true {
		toSerialize["free"] = o.Free
	}
	if true {
		toSerialize["interest"] = o.Interest
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["netAsset"] = o.NetAsset
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20019UserAssets struct {
	value *InlineResponse20019UserAssets
	isSet bool
}

func (v NullableInlineResponse20019UserAssets) Get() *InlineResponse20019UserAssets {
	return v.value
}

func (v *NullableInlineResponse20019UserAssets) Set(val *InlineResponse20019UserAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20019UserAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20019UserAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20019UserAssets(val *InlineResponse20019UserAssets) *NullableInlineResponse20019UserAssets {
	return &NullableInlineResponse20019UserAssets{value: val, isSet: true}
}

func (v NullableInlineResponse20019UserAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20019UserAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
