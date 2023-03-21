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

// EntitymanagerCreateEntityResponse struct for EntitymanagerCreateEntityResponse
type EntitymanagerCreateEntityResponse struct {
	AttributeWriteErrors *EntitymanagerAttributeWriteErrors `json:"attributeWriteErrors,omitempty"`
	Entity *EntitymanagerEntity `json:"entity,omitempty"`
}

// NewEntitymanagerCreateEntityResponse instantiates a new EntitymanagerCreateEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerCreateEntityResponse() *EntitymanagerCreateEntityResponse {
	this := EntitymanagerCreateEntityResponse{}
	return &this
}

// NewEntitymanagerCreateEntityResponseWithDefaults instantiates a new EntitymanagerCreateEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerCreateEntityResponseWithDefaults() *EntitymanagerCreateEntityResponse {
	this := EntitymanagerCreateEntityResponse{}
	return &this
}

// GetAttributeWriteErrors returns the AttributeWriteErrors field value if set, zero value otherwise.
func (o *EntitymanagerCreateEntityResponse) GetAttributeWriteErrors() EntitymanagerAttributeWriteErrors {
	if o == nil || isNil(o.AttributeWriteErrors) {
		var ret EntitymanagerAttributeWriteErrors
		return ret
	}
	return *o.AttributeWriteErrors
}

// GetAttributeWriteErrorsOk returns a tuple with the AttributeWriteErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateEntityResponse) GetAttributeWriteErrorsOk() (*EntitymanagerAttributeWriteErrors, bool) {
	if o == nil || isNil(o.AttributeWriteErrors) {
    return nil, false
	}
	return o.AttributeWriteErrors, true
}

// HasAttributeWriteErrors returns a boolean if a field has been set.
func (o *EntitymanagerCreateEntityResponse) HasAttributeWriteErrors() bool {
	if o != nil && !isNil(o.AttributeWriteErrors) {
		return true
	}

	return false
}

// SetAttributeWriteErrors gets a reference to the given EntitymanagerAttributeWriteErrors and assigns it to the AttributeWriteErrors field.
func (o *EntitymanagerCreateEntityResponse) SetAttributeWriteErrors(v EntitymanagerAttributeWriteErrors) {
	o.AttributeWriteErrors = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntitymanagerCreateEntityResponse) GetEntity() EntitymanagerEntity {
	if o == nil || isNil(o.Entity) {
		var ret EntitymanagerEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerCreateEntityResponse) GetEntityOk() (*EntitymanagerEntity, bool) {
	if o == nil || isNil(o.Entity) {
    return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntitymanagerCreateEntityResponse) HasEntity() bool {
	if o != nil && !isNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntitymanagerEntity and assigns it to the Entity field.
func (o *EntitymanagerCreateEntityResponse) SetEntity(v EntitymanagerEntity) {
	o.Entity = &v
}

func (o EntitymanagerCreateEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AttributeWriteErrors) {
		toSerialize["attributeWriteErrors"] = o.AttributeWriteErrors
	}
	if !isNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerCreateEntityResponse struct {
	value *EntitymanagerCreateEntityResponse
	isSet bool
}

func (v NullableEntitymanagerCreateEntityResponse) Get() *EntitymanagerCreateEntityResponse {
	return v.value
}

func (v *NullableEntitymanagerCreateEntityResponse) Set(val *EntitymanagerCreateEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerCreateEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerCreateEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerCreateEntityResponse(val *EntitymanagerCreateEntityResponse) *NullableEntitymanagerCreateEntityResponse {
	return &NullableEntitymanagerCreateEntityResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerCreateEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerCreateEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

