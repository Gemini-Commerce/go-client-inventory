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

// EntitymanagerListAttributeOptionsResponse struct for EntitymanagerListAttributeOptionsResponse
type EntitymanagerListAttributeOptionsResponse struct {
	Options []EntitymanagerAttributeOption `json:"options,omitempty"`
}

// NewEntitymanagerListAttributeOptionsResponse instantiates a new EntitymanagerListAttributeOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListAttributeOptionsResponse() *EntitymanagerListAttributeOptionsResponse {
	this := EntitymanagerListAttributeOptionsResponse{}
	return &this
}

// NewEntitymanagerListAttributeOptionsResponseWithDefaults instantiates a new EntitymanagerListAttributeOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListAttributeOptionsResponseWithDefaults() *EntitymanagerListAttributeOptionsResponse {
	this := EntitymanagerListAttributeOptionsResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerListAttributeOptionsResponse) GetOptions() []EntitymanagerAttributeOption {
	if o == nil || isNil(o.Options) {
		var ret []EntitymanagerAttributeOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListAttributeOptionsResponse) GetOptionsOk() ([]EntitymanagerAttributeOption, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerListAttributeOptionsResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerAttributeOption and assigns it to the Options field.
func (o *EntitymanagerListAttributeOptionsResponse) SetOptions(v []EntitymanagerAttributeOption) {
	o.Options = v
}

func (o EntitymanagerListAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerListAttributeOptionsResponse struct {
	value *EntitymanagerListAttributeOptionsResponse
	isSet bool
}

func (v NullableEntitymanagerListAttributeOptionsResponse) Get() *EntitymanagerListAttributeOptionsResponse {
	return v.value
}

func (v *NullableEntitymanagerListAttributeOptionsResponse) Set(val *EntitymanagerListAttributeOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListAttributeOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListAttributeOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListAttributeOptionsResponse(val *EntitymanagerListAttributeOptionsResponse) *NullableEntitymanagerListAttributeOptionsResponse {
	return &NullableEntitymanagerListAttributeOptionsResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerListAttributeOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListAttributeOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


