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

// InlineResponse20091Data struct for InlineResponse20091Data
type InlineResponse20091Data struct {
	ConfigDetails []InlineResponse20091DataConfigDetails `json:"configDetails"`
	TotalNum      int64                                  `json:"totalNum"`
	PageSize      int64                                  `json:"pageSize"`
}

// NewInlineResponse20091Data instantiates a new InlineResponse20091Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091Data(configDetails []InlineResponse20091DataConfigDetails, totalNum int64, pageSize int64) *InlineResponse20091Data {
	this := InlineResponse20091Data{}
	this.ConfigDetails = configDetails
	this.TotalNum = totalNum
	this.PageSize = pageSize
	return &this
}

// NewInlineResponse20091DataWithDefaults instantiates a new InlineResponse20091Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091DataWithDefaults() *InlineResponse20091Data {
	this := InlineResponse20091Data{}
	return &this
}

// GetConfigDetails returns the ConfigDetails field value
func (o *InlineResponse20091Data) GetConfigDetails() []InlineResponse20091DataConfigDetails {
	if o == nil {
		var ret []InlineResponse20091DataConfigDetails
		return ret
	}

	return o.ConfigDetails
}

// GetConfigDetailsOk returns a tuple with the ConfigDetails field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20091Data) GetConfigDetailsOk() (*[]InlineResponse20091DataConfigDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigDetails, true
}

// SetConfigDetails sets field value
func (o *InlineResponse20091Data) SetConfigDetails(v []InlineResponse20091DataConfigDetails) {
	o.ConfigDetails = v
}

// GetTotalNum returns the TotalNum field value
func (o *InlineResponse20091Data) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20091Data) GetTotalNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *InlineResponse20091Data) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetPageSize returns the PageSize field value
func (o *InlineResponse20091Data) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20091Data) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *InlineResponse20091Data) SetPageSize(v int64) {
	o.PageSize = v
}

func (o InlineResponse20091Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configDetails"] = o.ConfigDetails
	}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20091Data struct {
	value *InlineResponse20091Data
	isSet bool
}

func (v NullableInlineResponse20091Data) Get() *InlineResponse20091Data {
	return v.value
}

func (v *NullableInlineResponse20091Data) Set(val *InlineResponse20091Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091Data(val *InlineResponse20091Data) *NullableInlineResponse20091Data {
	return &NullableInlineResponse20091Data{value: val, isSet: true}
}

func (v NullableInlineResponse20091Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
