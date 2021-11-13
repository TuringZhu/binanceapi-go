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

// CapitalConfigGetallNetworkList struct for CapitalConfigGetallNetworkList
type CapitalConfigGetallNetworkList struct {
	AddressRegex string `json:"addressRegex"`
	Coin         string `json:"coin"`
	// shown only when \"depositEnable\" is false.
	DepositDesc   string `json:"depositDesc"`
	DepositEnable bool   `json:"depositEnable"`
	IsDefault     bool   `json:"isDefault"`
	MemoRegex     string `json:"memoRegex"`
	// min number for balance confirmation.
	MinConfirm         int64  `json:"minConfirm"`
	Name               string `json:"name"`
	ResetAddressStatus bool   `json:"resetAddressStatus"`
	SpecialTips        string `json:"specialTips"`
	// confirmation number for balance unlock.
	UnLockConfirm int64 `json:"unLockConfirm"`
	// shown only when \"withdrawEnable\" is false
	WithdrawDesc            string `json:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"`
}

// NewCapitalConfigGetallNetworkList instantiates a new CapitalConfigGetallNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalConfigGetallNetworkList(addressRegex string, coin string, depositDesc string, depositEnable bool, isDefault bool, memoRegex string, minConfirm int64, name string, resetAddressStatus bool, specialTips string, unLockConfirm int64, withdrawDesc string, withdrawEnable bool, withdrawFee string, withdrawIntegerMultiple string, withdrawMax string, withdrawMin string, sameAddress bool) *CapitalConfigGetallNetworkList {
	this := CapitalConfigGetallNetworkList{}
	this.AddressRegex = addressRegex
	this.Coin = coin
	this.DepositDesc = depositDesc
	this.DepositEnable = depositEnable
	this.IsDefault = isDefault
	this.MemoRegex = memoRegex
	this.MinConfirm = minConfirm
	this.Name = name
	this.ResetAddressStatus = resetAddressStatus
	this.SpecialTips = specialTips
	this.UnLockConfirm = unLockConfirm
	this.WithdrawDesc = withdrawDesc
	this.WithdrawEnable = withdrawEnable
	this.WithdrawFee = withdrawFee
	this.WithdrawIntegerMultiple = withdrawIntegerMultiple
	this.WithdrawMax = withdrawMax
	this.WithdrawMin = withdrawMin
	this.SameAddress = sameAddress
	return &this
}

// NewCapitalConfigGetallNetworkListWithDefaults instantiates a new CapitalConfigGetallNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalConfigGetallNetworkListWithDefaults() *CapitalConfigGetallNetworkList {
	this := CapitalConfigGetallNetworkList{}
	return &this
}

// GetAddressRegex returns the AddressRegex field value
func (o *CapitalConfigGetallNetworkList) GetAddressRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressRegex
}

// GetAddressRegexOk returns a tuple with the AddressRegex field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetAddressRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressRegex, true
}

// SetAddressRegex sets field value
func (o *CapitalConfigGetallNetworkList) SetAddressRegex(v string) {
	o.AddressRegex = v
}

// GetCoin returns the Coin field value
func (o *CapitalConfigGetallNetworkList) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetCoinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *CapitalConfigGetallNetworkList) SetCoin(v string) {
	o.Coin = v
}

// GetDepositDesc returns the DepositDesc field value
func (o *CapitalConfigGetallNetworkList) GetDepositDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositDesc
}

// GetDepositDescOk returns a tuple with the DepositDesc field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetDepositDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositDesc, true
}

// SetDepositDesc sets field value
func (o *CapitalConfigGetallNetworkList) SetDepositDesc(v string) {
	o.DepositDesc = v
}

// GetDepositEnable returns the DepositEnable field value
func (o *CapitalConfigGetallNetworkList) GetDepositEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DepositEnable
}

// GetDepositEnableOk returns a tuple with the DepositEnable field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetDepositEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositEnable, true
}

// SetDepositEnable sets field value
func (o *CapitalConfigGetallNetworkList) SetDepositEnable(v bool) {
	o.DepositEnable = v
}

// GetIsDefault returns the IsDefault field value
func (o *CapitalConfigGetallNetworkList) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *CapitalConfigGetallNetworkList) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetMemoRegex returns the MemoRegex field value
func (o *CapitalConfigGetallNetworkList) GetMemoRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemoRegex
}

// GetMemoRegexOk returns a tuple with the MemoRegex field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetMemoRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoRegex, true
}

// SetMemoRegex sets field value
func (o *CapitalConfigGetallNetworkList) SetMemoRegex(v string) {
	o.MemoRegex = v
}

// GetMinConfirm returns the MinConfirm field value
func (o *CapitalConfigGetallNetworkList) GetMinConfirm() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinConfirm
}

// GetMinConfirmOk returns a tuple with the MinConfirm field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetMinConfirmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinConfirm, true
}

// SetMinConfirm sets field value
func (o *CapitalConfigGetallNetworkList) SetMinConfirm(v int64) {
	o.MinConfirm = v
}

// GetName returns the Name field value
func (o *CapitalConfigGetallNetworkList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CapitalConfigGetallNetworkList) SetName(v string) {
	o.Name = v
}

// GetResetAddressStatus returns the ResetAddressStatus field value
func (o *CapitalConfigGetallNetworkList) GetResetAddressStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResetAddressStatus
}

// GetResetAddressStatusOk returns a tuple with the ResetAddressStatus field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetResetAddressStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetAddressStatus, true
}

// SetResetAddressStatus sets field value
func (o *CapitalConfigGetallNetworkList) SetResetAddressStatus(v bool) {
	o.ResetAddressStatus = v
}

// GetSpecialTips returns the SpecialTips field value
func (o *CapitalConfigGetallNetworkList) GetSpecialTips() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecialTips
}

// GetSpecialTipsOk returns a tuple with the SpecialTips field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetSpecialTipsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecialTips, true
}

// SetSpecialTips sets field value
func (o *CapitalConfigGetallNetworkList) SetSpecialTips(v string) {
	o.SpecialTips = v
}

// GetUnLockConfirm returns the UnLockConfirm field value
func (o *CapitalConfigGetallNetworkList) GetUnLockConfirm() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnLockConfirm
}

// GetUnLockConfirmOk returns a tuple with the UnLockConfirm field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetUnLockConfirmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnLockConfirm, true
}

// SetUnLockConfirm sets field value
func (o *CapitalConfigGetallNetworkList) SetUnLockConfirm(v int64) {
	o.UnLockConfirm = v
}

// GetWithdrawDesc returns the WithdrawDesc field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawDesc
}

// GetWithdrawDescOk returns a tuple with the WithdrawDesc field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawDesc, true
}

// SetWithdrawDesc sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawDesc(v string) {
	o.WithdrawDesc = v
}

// GetWithdrawEnable returns the WithdrawEnable field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WithdrawEnable
}

// GetWithdrawEnableOk returns a tuple with the WithdrawEnable field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawEnable, true
}

// SetWithdrawEnable sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawEnable(v bool) {
	o.WithdrawEnable = v
}

// GetWithdrawFee returns the WithdrawFee field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawFee, true
}

// SetWithdrawFee sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawFee(v string) {
	o.WithdrawFee = v
}

// GetWithdrawIntegerMultiple returns the WithdrawIntegerMultiple field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawIntegerMultiple() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawIntegerMultiple
}

// GetWithdrawIntegerMultipleOk returns a tuple with the WithdrawIntegerMultiple field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawIntegerMultipleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawIntegerMultiple, true
}

// SetWithdrawIntegerMultiple sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawIntegerMultiple(v string) {
	o.WithdrawIntegerMultiple = v
}

// GetWithdrawMax returns the WithdrawMax field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawMax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawMax
}

// GetWithdrawMaxOk returns a tuple with the WithdrawMax field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawMax, true
}

// SetWithdrawMax sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawMax(v string) {
	o.WithdrawMax = v
}

// GetWithdrawMin returns the WithdrawMin field value
func (o *CapitalConfigGetallNetworkList) GetWithdrawMin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawMin
}

// GetWithdrawMinOk returns a tuple with the WithdrawMin field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetWithdrawMinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawMin, true
}

// SetWithdrawMin sets field value
func (o *CapitalConfigGetallNetworkList) SetWithdrawMin(v string) {
	o.WithdrawMin = v
}

// GetSameAddress returns the SameAddress field value
func (o *CapitalConfigGetallNetworkList) GetSameAddress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SameAddress
}

// GetSameAddressOk returns a tuple with the SameAddress field value
// and a boolean to check if the value has been set.
func (o *CapitalConfigGetallNetworkList) GetSameAddressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SameAddress, true
}

// SetSameAddress sets field value
func (o *CapitalConfigGetallNetworkList) SetSameAddress(v bool) {
	o.SameAddress = v
}

func (o CapitalConfigGetallNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addressRegex"] = o.AddressRegex
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["depositDesc"] = o.DepositDesc
	}
	if true {
		toSerialize["depositEnable"] = o.DepositEnable
	}
	if true {
		toSerialize["isDefault"] = o.IsDefault
	}
	if true {
		toSerialize["memoRegex"] = o.MemoRegex
	}
	if true {
		toSerialize["minConfirm"] = o.MinConfirm
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["resetAddressStatus"] = o.ResetAddressStatus
	}
	if true {
		toSerialize["specialTips"] = o.SpecialTips
	}
	if true {
		toSerialize["unLockConfirm"] = o.UnLockConfirm
	}
	if true {
		toSerialize["withdrawDesc"] = o.WithdrawDesc
	}
	if true {
		toSerialize["withdrawEnable"] = o.WithdrawEnable
	}
	if true {
		toSerialize["withdrawFee"] = o.WithdrawFee
	}
	if true {
		toSerialize["withdrawIntegerMultiple"] = o.WithdrawIntegerMultiple
	}
	if true {
		toSerialize["withdrawMax"] = o.WithdrawMax
	}
	if true {
		toSerialize["withdrawMin"] = o.WithdrawMin
	}
	if true {
		toSerialize["sameAddress"] = o.SameAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCapitalConfigGetallNetworkList struct {
	value *CapitalConfigGetallNetworkList
	isSet bool
}

func (v NullableCapitalConfigGetallNetworkList) Get() *CapitalConfigGetallNetworkList {
	return v.value
}

func (v *NullableCapitalConfigGetallNetworkList) Set(val *CapitalConfigGetallNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalConfigGetallNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalConfigGetallNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalConfigGetallNetworkList(val *CapitalConfigGetallNetworkList) *NullableCapitalConfigGetallNetworkList {
	return &NullableCapitalConfigGetallNetworkList{value: val, isSet: true}
}

func (v NullableCapitalConfigGetallNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalConfigGetallNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}