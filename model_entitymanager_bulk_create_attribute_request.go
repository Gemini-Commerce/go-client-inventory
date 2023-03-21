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

// EntitymanagerBulkCreateAttributeRequest struct for EntitymanagerBulkCreateAttributeRequest
type EntitymanagerBulkCreateAttributeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	EntityData *EntitymanagerEntityIdentifier `json:"entityData,omitempty"`
	EntityId *string `json:"entityId,omitempty"`
	Attributes []EntitymanagerAttribute `json:"attributes,omitempty"`
}

// NewEntitymanagerBulkCreateAttributeRequest instantiates a new EntitymanagerBulkCreateAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerBulkCreateAttributeRequest() *EntitymanagerBulkCreateAttributeRequest {
	this := EntitymanagerBulkCreateAttributeRequest{}
	return &this
}

// NewEntitymanagerBulkCreateAttributeRequestWithDefaults instantiates a new EntitymanagerBulkCreateAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerBulkCreateAttributeRequestWithDefaults() *EntitymanagerBulkCreateAttributeRequest {
	this := EntitymanagerBulkCreateAttributeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerBulkCreateAttributeRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerBulkCreateAttributeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEntityData returns the EntityData field value if set, zero value otherwise.
func (o *EntitymanagerBulkCreateAttributeRequest) GetEntityData() EntitymanagerEntityIdentifier {
	if o == nil || isNil(o.EntityData) {
		var ret EntitymanagerEntityIdentifier
		return ret
	}
	return *o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) GetEntityDataOk() (*EntitymanagerEntityIdentifier, bool) {
	if o == nil || isNil(o.EntityData) {
    return nil, false
	}
	return o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) HasEntityData() bool {
	if o != nil && !isNil(o.EntityData) {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given EntitymanagerEntityIdentifier and assigns it to the EntityData field.
func (o *EntitymanagerBulkCreateAttributeRequest) SetEntityData(v EntitymanagerEntityIdentifier) {
	o.EntityData = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntitymanagerBulkCreateAttributeRequest) GetEntityId() string {
	if o == nil || isNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityId) {
    return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) HasEntityId() bool {
	if o != nil && !isNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntitymanagerBulkCreateAttributeRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntitymanagerBulkCreateAttributeRequest) GetAttributes() []EntitymanagerAttribute {
	if o == nil || isNil(o.Attributes) {
		var ret []EntitymanagerAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) GetAttributesOk() ([]EntitymanagerAttribute, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntitymanagerBulkCreateAttributeRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []EntitymanagerAttribute and assigns it to the Attributes field.
func (o *EntitymanagerBulkCreateAttributeRequest) SetAttributes(v []EntitymanagerAttribute) {
	o.Attributes = v
}

func (o EntitymanagerBulkCreateAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.EntityData) {
		toSerialize["entityData"] = o.EntityData
	}
	if !isNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerBulkCreateAttributeRequest struct {
	value *EntitymanagerBulkCreateAttributeRequest
	isSet bool
}

func (v NullableEntitymanagerBulkCreateAttributeRequest) Get() *EntitymanagerBulkCreateAttributeRequest {
	return v.value
}

func (v *NullableEntitymanagerBulkCreateAttributeRequest) Set(val *EntitymanagerBulkCreateAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerBulkCreateAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerBulkCreateAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerBulkCreateAttributeRequest(val *EntitymanagerBulkCreateAttributeRequest) *NullableEntitymanagerBulkCreateAttributeRequest {
	return &NullableEntitymanagerBulkCreateAttributeRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerBulkCreateAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerBulkCreateAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

