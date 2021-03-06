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

// MarginTrade struct for MarginTrade
type MarginTrade struct {
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              int64  `json:"id"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	OrderId         int64  `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Symbol          string `json:"symbol"`
	IsIsolated      bool   `json:"isIsolated"`
	Time            int64  `json:"time"`
}

// NewMarginTrade instantiates a new MarginTrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginTrade(commission string, commissionAsset string, id int64, isBestMatch bool, isBuyer bool, isMaker bool, orderId int64, price string, qty string, symbol string, isIsolated bool, time int64) *MarginTrade {
	this := MarginTrade{}
	this.Commission = commission
	this.CommissionAsset = commissionAsset
	this.Id = id
	this.IsBestMatch = isBestMatch
	this.IsBuyer = isBuyer
	this.IsMaker = isMaker
	this.OrderId = orderId
	this.Price = price
	this.Qty = qty
	this.Symbol = symbol
	this.IsIsolated = isIsolated
	this.Time = time
	return &this
}

// NewMarginTradeWithDefaults instantiates a new MarginTrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginTradeWithDefaults() *MarginTrade {
	this := MarginTrade{}
	return &this
}

// GetCommission returns the Commission field value
func (o *MarginTrade) GetCommission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetCommissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *MarginTrade) SetCommission(v string) {
	o.Commission = v
}

// GetCommissionAsset returns the CommissionAsset field value
func (o *MarginTrade) GetCommissionAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetCommissionAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommissionAsset, true
}

// SetCommissionAsset sets field value
func (o *MarginTrade) SetCommissionAsset(v string) {
	o.CommissionAsset = v
}

// GetId returns the Id field value
func (o *MarginTrade) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MarginTrade) SetId(v int64) {
	o.Id = v
}

// GetIsBestMatch returns the IsBestMatch field value
func (o *MarginTrade) GetIsBestMatch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetIsBestMatchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBestMatch, true
}

// SetIsBestMatch sets field value
func (o *MarginTrade) SetIsBestMatch(v bool) {
	o.IsBestMatch = v
}

// GetIsBuyer returns the IsBuyer field value
func (o *MarginTrade) GetIsBuyer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetIsBuyerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBuyer, true
}

// SetIsBuyer sets field value
func (o *MarginTrade) SetIsBuyer(v bool) {
	o.IsBuyer = v
}

// GetIsMaker returns the IsMaker field value
func (o *MarginTrade) GetIsMaker() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetIsMakerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMaker, true
}

// SetIsMaker sets field value
func (o *MarginTrade) SetIsMaker(v bool) {
	o.IsMaker = v
}

// GetOrderId returns the OrderId field value
func (o *MarginTrade) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *MarginTrade) SetOrderId(v int64) {
	o.OrderId = v
}

// GetPrice returns the Price field value
func (o *MarginTrade) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *MarginTrade) SetPrice(v string) {
	o.Price = v
}

// GetQty returns the Qty field value
func (o *MarginTrade) GetQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *MarginTrade) SetQty(v string) {
	o.Qty = v
}

// GetSymbol returns the Symbol field value
func (o *MarginTrade) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarginTrade) SetSymbol(v string) {
	o.Symbol = v
}

// GetIsIsolated returns the IsIsolated field value
func (o *MarginTrade) GetIsIsolated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetIsIsolatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsIsolated, true
}

// SetIsIsolated sets field value
func (o *MarginTrade) SetIsIsolated(v bool) {
	o.IsIsolated = v
}

// GetTime returns the Time field value
func (o *MarginTrade) GetTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *MarginTrade) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *MarginTrade) SetTime(v int64) {
	o.Time = v
}

func (o MarginTrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commission"] = o.Commission
	}
	if true {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}
	if true {
		toSerialize["isBuyer"] = o.IsBuyer
	}
	if true {
		toSerialize["isMaker"] = o.IsMaker
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["qty"] = o.Qty
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableMarginTrade struct {
	value *MarginTrade
	isSet bool
}

func (v NullableMarginTrade) Get() *MarginTrade {
	return v.value
}

func (v *NullableMarginTrade) Set(val *MarginTrade) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginTrade(val *MarginTrade) *NullableMarginTrade {
	return &NullableMarginTrade{value: val, isSet: true}
}

func (v NullableMarginTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
