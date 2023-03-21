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

// EntitymanagerUpdateAttributeOptionsRequest struct for EntitymanagerUpdateAttributeOptionsRequest
type EntitymanagerUpdateAttributeOptionsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ListCode *string `json:"listCode,omitempty"`
	Options []EntitymanagerAttributeOption `json:"options,omitempty"`
}

// NewEntitymanagerUpdateAttributeOptionsRequest instantiates a new EntitymanagerUpdateAttributeOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerUpdateAttributeOptionsRequest() *EntitymanagerUpdateAttributeOptionsRequest {
	this := EntitymanagerUpdateAttributeOptionsRequest{}
	return &this
}

// NewEntitymanagerUpdateAttributeOptionsRequestWithDefaults instantiates a new EntitymanagerUpdateAttributeOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerUpdateAttributeOptionsRequestWithDefaults() *EntitymanagerUpdateAttributeOptionsRequest {
	this := EntitymanagerUpdateAttributeOptionsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerUpdateAttributeOptionsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetListCode returns the ListCode field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetListCode() string {
	if o == nil || isNil(o.ListCode) {
		var ret string
		return ret
	}
	return *o.ListCode
}

// GetListCodeOk returns a tuple with the ListCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetListCodeOk() (*string, bool) {
	if o == nil || isNil(o.ListCode) {
    return nil, false
	}
	return o.ListCode, true
}

// HasListCode returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) HasListCode() bool {
	if o != nil && !isNil(o.ListCode) {
		return true
	}

	return false
}

// SetListCode gets a reference to the given string and assigns it to the ListCode field.
func (o *EntitymanagerUpdateAttributeOptionsRequest) SetListCode(v string) {
	o.ListCode = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetOptions() []EntitymanagerAttributeOption {
	if o == nil || isNil(o.Options) {
		var ret []EntitymanagerAttributeOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) GetOptionsOk() ([]EntitymanagerAttributeOption, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *EntitymanagerUpdateAttributeOptionsRequest) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []EntitymanagerAttributeOption and assigns it to the Options field.
func (o *EntitymanagerUpdateAttributeOptionsRequest) SetOptions(v []EntitymanagerAttributeOption) {
	o.Options = v
}

func (o EntitymanagerUpdateAttributeOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ListCode) {
		toSerialize["listCode"] = o.ListCode
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerUpdateAttributeOptionsRequest struct {
	value *EntitymanagerUpdateAttributeOptionsRequest
	isSet bool
}

func (v NullableEntitymanagerUpdateAttributeOptionsRequest) Get() *EntitymanagerUpdateAttributeOptionsRequest {
	return v.value
}

func (v *NullableEntitymanagerUpdateAttributeOptionsRequest) Set(val *EntitymanagerUpdateAttributeOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerUpdateAttributeOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerUpdateAttributeOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerUpdateAttributeOptionsRequest(val *EntitymanagerUpdateAttributeOptionsRequest) *NullableEntitymanagerUpdateAttributeOptionsRequest {
	return &NullableEntitymanagerUpdateAttributeOptionsRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerUpdateAttributeOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerUpdateAttributeOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


