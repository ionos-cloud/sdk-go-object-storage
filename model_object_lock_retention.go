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

// ObjectLockRetention A Retention configuration for an object.
type ObjectLockRetention struct {
	XMLName xml.Name `xml:"Retention"`
	// Indicates the Retention mode for the specified object.
	Mode *string `json:"Mode,omitempty" xml:"Mode"`
	// The date on which this Object Lock Retention will expire.
	RetainUntilDate *string `json:"RetainUntilDate,omitempty" xml:"RetainUntilDate"`
}

// NewObjectLockRetention instantiates a new ObjectLockRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLockRetention() *ObjectLockRetention {
	this := ObjectLockRetention{}

	return &this
}

// NewObjectLockRetentionWithDefaults instantiates a new ObjectLockRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLockRetentionWithDefaults() *ObjectLockRetention {
	this := ObjectLockRetention{}
	return &this
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectLockRetention) GetMode() *string {
	if o == nil {
		return nil
	}

	return o.Mode

}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectLockRetention) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Mode, true
}

// SetMode sets field value
func (o *ObjectLockRetention) SetMode(v string) {

	o.Mode = &v

}

// HasMode returns a boolean if a field has been set.
func (o *ObjectLockRetention) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// GetRetainUntilDate returns the RetainUntilDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectLockRetention) GetRetainUntilDate() *string {
	if o == nil {
		return nil
	}

	return o.RetainUntilDate

}

// GetRetainUntilDateOk returns a tuple with the RetainUntilDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectLockRetention) GetRetainUntilDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.RetainUntilDate, true
}

// SetRetainUntilDate sets field value
func (o *ObjectLockRetention) SetRetainUntilDate(v string) {

	o.RetainUntilDate = &v

}

// HasRetainUntilDate returns a boolean if a field has been set.
func (o *ObjectLockRetention) HasRetainUntilDate() bool {
	if o != nil && o.RetainUntilDate != nil {
		return true
	}

	return false
}

func (o ObjectLockRetention) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	if o.RetainUntilDate != nil {
		toSerialize["RetainUntilDate"] = o.RetainUntilDate
	}

	return json.Marshal(toSerialize)
}

type NullableObjectLockRetention struct {
	value *ObjectLockRetention
	isSet bool
}

func (v NullableObjectLockRetention) Get() *ObjectLockRetention {
	return v.value
}

func (v *NullableObjectLockRetention) Set(val *ObjectLockRetention) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLockRetention) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLockRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLockRetention(val *ObjectLockRetention) *NullableObjectLockRetention {
	return &NullableObjectLockRetention{value: val, isSet: true}
}

func (v NullableObjectLockRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLockRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
