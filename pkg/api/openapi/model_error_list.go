/*
maestro Service API

maestro Service API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ErrorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorList{}

// ErrorList struct for ErrorList
type ErrorList struct {
	Kind  string  `json:"kind"`
	Page  int32   `json:"page"`
	Size  int32   `json:"size"`
	Total int32   `json:"total"`
	Items []Error `json:"items"`
}

// NewErrorList instantiates a new ErrorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorList(kind string, page int32, size int32, total int32, items []Error) *ErrorList {
	this := ErrorList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewErrorListWithDefaults instantiates a new ErrorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorListWithDefaults() *ErrorList {
	this := ErrorList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ErrorList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ErrorList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *ErrorList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ErrorList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *ErrorList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ErrorList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *ErrorList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ErrorList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ErrorList) GetItems() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetItemsOk() ([]Error, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ErrorList) SetItems(v []Error) {
	o.Items = v
}

func (o ErrorList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableErrorList struct {
	value *ErrorList
	isSet bool
}

func (v NullableErrorList) Get() *ErrorList {
	return v.value
}

func (v *NullableErrorList) Set(val *ErrorList) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorList) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorList(val *ErrorList) *NullableErrorList {
	return &NullableErrorList{value: val, isSet: true}
}

func (v NullableErrorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
