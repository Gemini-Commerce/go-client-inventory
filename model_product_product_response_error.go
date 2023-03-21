/*
product/product.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// ProductProductResponseError struct for ProductProductResponseError
type ProductProductResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewProductProductResponseError instantiates a new ProductProductResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductProductResponseError() *ProductProductResponseError {
	this := ProductProductResponseError{}
	return &this
}

// NewProductProductResponseErrorWithDefaults instantiates a new ProductProductResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductProductResponseErrorWithDefaults() *ProductProductResponseError {
	this := ProductProductResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductProductResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductProductResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductProductResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProductProductResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProductResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProductProductResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProductProductResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ProductProductResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableProductProductResponseError struct {
	value *ProductProductResponseError
	isSet bool
}

func (v NullableProductProductResponseError) Get() *ProductProductResponseError {
	return v.value
}

func (v *NullableProductProductResponseError) Set(val *ProductProductResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableProductProductResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableProductProductResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductProductResponseError(val *ProductProductResponseError) *NullableProductProductResponseError {
	return &NullableProductProductResponseError{value: val, isSet: true}
}

func (v NullableProductProductResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductProductResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


