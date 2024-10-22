/*
 * IONOS Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

import "encoding/xml"

// ListAllMyBucketsResult struct for ListAllMyBucketsResult
type ListAllMyBucketsResult struct {
	XMLName xml.Name  `xml:"ListAllMyBucketsResult"`
	Owner   *Owner    `json:"Owner,omitempty" xml:"Owner"`
	Buckets *[]Bucket `json:"Buckets,omitempty" xml:"Buckets>Bucket"`
}

// NewListAllMyBucketsResult instantiates a new ListAllMyBucketsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllMyBucketsResult() *ListAllMyBucketsResult {
	this := ListAllMyBucketsResult{}

	return &this
}

// NewListAllMyBucketsResultWithDefaults instantiates a new ListAllMyBucketsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllMyBucketsResultWithDefaults() *ListAllMyBucketsResult {
	this := ListAllMyBucketsResult{}
	return &this
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for Owner will be returned
func (o *ListAllMyBucketsResult) GetOwner() *Owner {
	if o == nil {
		return nil
	}

	return o.Owner

}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAllMyBucketsResult) GetOwnerOk() (*Owner, bool) {
	if o == nil {
		return nil, false
	}

	return o.Owner, true
}

// SetOwner sets field value
func (o *ListAllMyBucketsResult) SetOwner(v Owner) {

	o.Owner = &v

}

// HasOwner returns a boolean if a field has been set.
func (o *ListAllMyBucketsResult) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// GetBuckets returns the Buckets field value
// If the value is explicit nil, the zero value for []Bucket will be returned
func (o *ListAllMyBucketsResult) GetBuckets() *[]Bucket {
	if o == nil {
		return nil
	}

	return o.Buckets

}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAllMyBucketsResult) GetBucketsOk() (*[]Bucket, bool) {
	if o == nil {
		return nil, false
	}

	return o.Buckets, true
}

// SetBuckets sets field value
func (o *ListAllMyBucketsResult) SetBuckets(v []Bucket) {

	o.Buckets = &v

}

// HasBuckets returns a boolean if a field has been set.
func (o *ListAllMyBucketsResult) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

func (o ListAllMyBucketsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	if o.Buckets != nil {
		toSerialize["Buckets"] = o.Buckets
	}

	return json.Marshal(toSerialize)
}

type NullableListAllMyBucketsResult struct {
	value *ListAllMyBucketsResult
	isSet bool
}

func (v NullableListAllMyBucketsResult) Get() *ListAllMyBucketsResult {
	return v.value
}

func (v *NullableListAllMyBucketsResult) Set(val *ListAllMyBucketsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllMyBucketsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllMyBucketsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllMyBucketsResult(val *ListAllMyBucketsResult) *NullableListAllMyBucketsResult {
	return &NullableListAllMyBucketsResult{value: val, isSet: true}
}

func (v NullableListAllMyBucketsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllMyBucketsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
