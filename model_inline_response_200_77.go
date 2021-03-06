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

// InlineResponse20077 struct for InlineResponse20077
type InlineResponse20077 struct {
	Asset               string `json:"asset"`
	DailyQuota          string `json:"dailyQuota"`
	LeftQuota           string `json:"leftQuota"`
	MinRedemptionAmount string `json:"minRedemptionAmount"`
}

// NewInlineResponse20077 instantiates a new InlineResponse20077 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20077(asset string, dailyQuota string, leftQuota string, minRedemptionAmount string) *InlineResponse20077 {
	this := InlineResponse20077{}
	this.Asset = asset
	this.DailyQuota = dailyQuota
	this.LeftQuota = leftQuota
	this.MinRedemptionAmount = minRedemptionAmount
	return &this
}

// NewInlineResponse20077WithDefaults instantiates a new InlineResponse20077 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20077WithDefaults() *InlineResponse20077 {
	this := InlineResponse20077{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20077) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20077) SetAsset(v string) {
	o.Asset = v
}

// GetDailyQuota returns the DailyQuota field value
func (o *InlineResponse20077) GetDailyQuota() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DailyQuota
}

// GetDailyQuotaOk returns a tuple with the DailyQuota field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetDailyQuotaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyQuota, true
}

// SetDailyQuota sets field value
func (o *InlineResponse20077) SetDailyQuota(v string) {
	o.DailyQuota = v
}

// GetLeftQuota returns the LeftQuota field value
func (o *InlineResponse20077) GetLeftQuota() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LeftQuota
}

// GetLeftQuotaOk returns a tuple with the LeftQuota field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetLeftQuotaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeftQuota, true
}

// SetLeftQuota sets field value
func (o *InlineResponse20077) SetLeftQuota(v string) {
	o.LeftQuota = v
}

// GetMinRedemptionAmount returns the MinRedemptionAmount field value
func (o *InlineResponse20077) GetMinRedemptionAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinRedemptionAmount
}

// GetMinRedemptionAmountOk returns a tuple with the MinRedemptionAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetMinRedemptionAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinRedemptionAmount, true
}

// SetMinRedemptionAmount sets field value
func (o *InlineResponse20077) SetMinRedemptionAmount(v string) {
	o.MinRedemptionAmount = v
}

func (o InlineResponse20077) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["dailyQuota"] = o.DailyQuota
	}
	if true {
		toSerialize["leftQuota"] = o.LeftQuota
	}
	if true {
		toSerialize["minRedemptionAmount"] = o.MinRedemptionAmount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20077 struct {
	value *InlineResponse20077
	isSet bool
}

func (v NullableInlineResponse20077) Get() *InlineResponse20077 {
	return v.value
}

func (v *NullableInlineResponse20077) Set(val *InlineResponse20077) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20077) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20077) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20077(val *InlineResponse20077) *NullableInlineResponse20077 {
	return &NullableInlineResponse20077{value: val, isSet: true}
}

func (v NullableInlineResponse20077) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20077) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
