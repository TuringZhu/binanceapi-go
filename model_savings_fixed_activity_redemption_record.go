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

// SavingsFixedActivityRedemptionRecord struct for SavingsFixedActivityRedemptionRecord
type SavingsFixedActivityRedemptionRecord struct {
	Items []map[string]interface{}
}

// NewSavingsFixedActivityRedemptionRecord instantiates a new SavingsFixedActivityRedemptionRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavingsFixedActivityRedemptionRecord() *SavingsFixedActivityRedemptionRecord {
	this := SavingsFixedActivityRedemptionRecord{}
	return &this
}

// NewSavingsFixedActivityRedemptionRecordWithDefaults instantiates a new SavingsFixedActivityRedemptionRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavingsFixedActivityRedemptionRecordWithDefaults() *SavingsFixedActivityRedemptionRecord {
	this := SavingsFixedActivityRedemptionRecord{}
	return &this
}

func (o SavingsFixedActivityRedemptionRecord) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *SavingsFixedActivityRedemptionRecord) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableSavingsFixedActivityRedemptionRecord struct {
	value *SavingsFixedActivityRedemptionRecord
	isSet bool
}

func (v NullableSavingsFixedActivityRedemptionRecord) Get() *SavingsFixedActivityRedemptionRecord {
	return v.value
}

func (v *NullableSavingsFixedActivityRedemptionRecord) Set(val *SavingsFixedActivityRedemptionRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSavingsFixedActivityRedemptionRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSavingsFixedActivityRedemptionRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavingsFixedActivityRedemptionRecord(val *SavingsFixedActivityRedemptionRecord) *NullableSavingsFixedActivityRedemptionRecord {
	return &NullableSavingsFixedActivityRedemptionRecord{value: val, isSet: true}
}

func (v NullableSavingsFixedActivityRedemptionRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavingsFixedActivityRedemptionRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
