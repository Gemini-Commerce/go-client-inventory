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

// ProductDeleteRequest struct for ProductDeleteRequest
type ProductDeleteRequest struct {
	Product *ProductProductEntity `json:"product,omitempty"`
}

// NewProductDeleteRequest instantiates a new ProductDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDeleteRequest() *ProductDeleteRequest {
	this := ProductDeleteRequest{}
	return &this
}

// NewProductDeleteRequestWithDefaults instantiates a new ProductDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDeleteRequestWithDefaults() *ProductDeleteRequest {
	this := ProductDeleteRequest{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ProductDeleteRequest) GetProduct() ProductProductEntity {
	if o == nil || isNil(o.Product) {
		var ret ProductProductEntity
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDeleteRequest) GetProductOk() (*ProductProductEntity, bool) {
	if o == nil || isNil(o.Product) {
    return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ProductDeleteRequest) HasProduct() bool {
	if o != nil && !isNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductProductEntity and assigns it to the Product field.
func (o *ProductDeleteRequest) SetProduct(v ProductProductEntity) {
	o.Product = &v
}

func (o ProductDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return json.Marshal(toSerialize)
}

type NullableProductDeleteRequest struct {
	value *ProductDeleteRequest
	isSet bool
}

func (v NullableProductDeleteRequest) Get() *ProductDeleteRequest {
	return v.value
}

func (v *NullableProductDeleteRequest) Set(val *ProductDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDeleteRequest(val *ProductDeleteRequest) *NullableProductDeleteRequest {
	return &NullableProductDeleteRequest{value: val, isSet: true}
}

func (v NullableProductDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


