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

// InlineResponse20081 struct for InlineResponse20081
type InlineResponse20081 struct {
	Asset           string `json:"asset"`
	CanTransfer     bool   `json:"canTransfer"`
	CreateTimestamp int64  `json:"createTimestamp"`
	Duration        int64  `json:"duration"`
	EndTime         int64  `json:"endTime"`
	Interest        string `json:"interest"`
	InterestRate    string `json:"interestRate"`
	Lot             int64  `json:"lot"`
	PositionId      int64  `json:"positionId"`
	Principal       string `json:"principal"`
	ProjectId       string `json:"projectId"`
	ProjectName     string `json:"projectName"`
	PurchaseTime    int64  `json:"purchaseTime"`
	RedeemDate      string `json:"redeemDate"`
	StartTime       int64  `json:"startTime"`
	Status          string `json:"status"`
	Type            string `json:"type"`
}

// NewInlineResponse20081 instantiates a new InlineResponse20081 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20081(asset string, canTransfer bool, createTimestamp int64, duration int64, endTime int64, interest string, interestRate string, lot int64, positionId int64, principal string, projectId string, projectName string, purchaseTime int64, redeemDate string, startTime int64, status string, type_ string) *InlineResponse20081 {
	this := InlineResponse20081{}
	this.Asset = asset
	this.CanTransfer = canTransfer
	this.CreateTimestamp = createTimestamp
	this.Duration = duration
	this.EndTime = endTime
	this.Interest = interest
	this.InterestRate = interestRate
	this.Lot = lot
	this.PositionId = positionId
	this.Principal = principal
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.PurchaseTime = purchaseTime
	this.RedeemDate = redeemDate
	this.StartTime = startTime
	this.Status = status
	this.Type = type_
	return &this
}

// NewInlineResponse20081WithDefaults instantiates a new InlineResponse20081 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20081WithDefaults() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20081) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20081) SetAsset(v string) {
	o.Asset = v
}

// GetCanTransfer returns the CanTransfer field value
func (o *InlineResponse20081) GetCanTransfer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanTransfer
}

// GetCanTransferOk returns a tuple with the CanTransfer field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetCanTransferOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanTransfer, true
}

// SetCanTransfer sets field value
func (o *InlineResponse20081) SetCanTransfer(v bool) {
	o.CanTransfer = v
}

// GetCreateTimestamp returns the CreateTimestamp field value
func (o *InlineResponse20081) GetCreateTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetCreateTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTimestamp, true
}

// SetCreateTimestamp sets field value
func (o *InlineResponse20081) SetCreateTimestamp(v int64) {
	o.CreateTimestamp = v
}

// GetDuration returns the Duration field value
func (o *InlineResponse20081) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *InlineResponse20081) SetDuration(v int64) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value
func (o *InlineResponse20081) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *InlineResponse20081) SetEndTime(v int64) {
	o.EndTime = v
}

// GetInterest returns the Interest field value
func (o *InlineResponse20081) GetInterest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interest
}

// GetInterestOk returns a tuple with the Interest field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetInterestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interest, true
}

// SetInterest sets field value
func (o *InlineResponse20081) SetInterest(v string) {
	o.Interest = v
}

// GetInterestRate returns the InterestRate field value
func (o *InlineResponse20081) GetInterestRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetInterestRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterestRate, true
}

// SetInterestRate sets field value
func (o *InlineResponse20081) SetInterestRate(v string) {
	o.InterestRate = v
}

// GetLot returns the Lot field value
func (o *InlineResponse20081) GetLot() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Lot
}

// GetLotOk returns a tuple with the Lot field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetLotOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lot, true
}

// SetLot sets field value
func (o *InlineResponse20081) SetLot(v int64) {
	o.Lot = v
}

// GetPositionId returns the PositionId field value
func (o *InlineResponse20081) GetPositionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetPositionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionId, true
}

// SetPositionId sets field value
func (o *InlineResponse20081) SetPositionId(v int64) {
	o.PositionId = v
}

// GetPrincipal returns the Principal field value
func (o *InlineResponse20081) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *InlineResponse20081) SetPrincipal(v string) {
	o.Principal = v
}

// GetProjectId returns the ProjectId field value
func (o *InlineResponse20081) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *InlineResponse20081) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
func (o *InlineResponse20081) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *InlineResponse20081) SetProjectName(v string) {
	o.ProjectName = v
}

// GetPurchaseTime returns the PurchaseTime field value
func (o *InlineResponse20081) GetPurchaseTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PurchaseTime
}

// GetPurchaseTimeOk returns a tuple with the PurchaseTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetPurchaseTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseTime, true
}

// SetPurchaseTime sets field value
func (o *InlineResponse20081) SetPurchaseTime(v int64) {
	o.PurchaseTime = v
}

// GetRedeemDate returns the RedeemDate field value
func (o *InlineResponse20081) GetRedeemDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedeemDate
}

// GetRedeemDateOk returns a tuple with the RedeemDate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetRedeemDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedeemDate, true
}

// SetRedeemDate sets field value
func (o *InlineResponse20081) SetRedeemDate(v string) {
	o.RedeemDate = v
}

// GetStartTime returns the StartTime field value
func (o *InlineResponse20081) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *InlineResponse20081) SetStartTime(v int64) {
	o.StartTime = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20081) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20081) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *InlineResponse20081) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20081) SetType(v string) {
	o.Type = v
}

func (o InlineResponse20081) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["canTransfer"] = o.CanTransfer
	}
	if true {
		toSerialize["createTimestamp"] = o.CreateTimestamp
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["interest"] = o.Interest
	}
	if true {
		toSerialize["interestRate"] = o.InterestRate
	}
	if true {
		toSerialize["lot"] = o.Lot
	}
	if true {
		toSerialize["positionId"] = o.PositionId
	}
	if true {
		toSerialize["principal"] = o.Principal
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["projectName"] = o.ProjectName
	}
	if true {
		toSerialize["purchaseTime"] = o.PurchaseTime
	}
	if true {
		toSerialize["redeemDate"] = o.RedeemDate
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20081 struct {
	value *InlineResponse20081
	isSet bool
}

func (v NullableInlineResponse20081) Get() *InlineResponse20081 {
	return v.value
}

func (v *NullableInlineResponse20081) Set(val *InlineResponse20081) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20081) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20081) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20081(val *InlineResponse20081) *NullableInlineResponse20081 {
	return &NullableInlineResponse20081{value: val, isSet: true}
}

func (v NullableInlineResponse20081) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20081) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
