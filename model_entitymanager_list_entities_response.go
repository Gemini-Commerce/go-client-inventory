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

// EntitymanagerListEntitiesResponse struct for EntitymanagerListEntitiesResponse
type EntitymanagerListEntitiesResponse struct {
	Entities []EntitymanagerEntity `json:"entities,omitempty"`
	NextPage *int32 `json:"nextPage,omitempty"`
}

// NewEntitymanagerListEntitiesResponse instantiates a new EntitymanagerListEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListEntitiesResponse() *EntitymanagerListEntitiesResponse {
	this := EntitymanagerListEntitiesResponse{}
	return &this
}

// NewEntitymanagerListEntitiesResponseWithDefaults instantiates a new EntitymanagerListEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListEntitiesResponseWithDefaults() *EntitymanagerListEntitiesResponse {
	this := EntitymanagerListEntitiesResponse{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesResponse) GetEntities() []EntitymanagerEntity {
	if o == nil || isNil(o.Entities) {
		var ret []EntitymanagerEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesResponse) GetEntitiesOk() ([]EntitymanagerEntity, bool) {
	if o == nil || isNil(o.Entities) {
    return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesResponse) HasEntities() bool {
	if o != nil && !isNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntitymanagerEntity and assigns it to the Entities field.
func (o *EntitymanagerListEntitiesResponse) SetEntities(v []EntitymanagerEntity) {
	o.Entities = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *EntitymanagerListEntitiesResponse) GetNextPage() int32 {
	if o == nil || isNil(o.NextPage) {
		var ret int32
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListEntitiesResponse) GetNextPageOk() (*int32, bool) {
	if o == nil || isNil(o.NextPage) {
    return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *EntitymanagerListEntitiesResponse) HasNextPage() bool {
	if o != nil && !isNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given int32 and assigns it to the NextPage field.
func (o *EntitymanagerListEntitiesResponse) SetNextPage(v int32) {
	o.NextPage = &v
}

func (o EntitymanagerListEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !isNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerListEntitiesResponse struct {
	value *EntitymanagerListEntitiesResponse
	isSet bool
}

func (v NullableEntitymanagerListEntitiesResponse) Get() *EntitymanagerListEntitiesResponse {
	return v.value
}

func (v *NullableEntitymanagerListEntitiesResponse) Set(val *EntitymanagerListEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListEntitiesResponse(val *EntitymanagerListEntitiesResponse) *NullableEntitymanagerListEntitiesResponse {
	return &NullableEntitymanagerListEntitiesResponse{value: val, isSet: true}
}

func (v NullableEntitymanagerListEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

