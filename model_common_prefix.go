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

// CommonPrefix Container for all (if there are any) keys between Prefix and the next occurrence of the string specified by a delimiter. CommonPrefixes lists keys that act like subdirectories in the directory specified by Prefix. For example, if the prefix is `notes/` and the delimiter is a slash (`/“) as in `notes/summer/july“, the common prefix is `notes/summer/“.
type CommonPrefix struct {
	XMLName xml.Name `xml:"CommonPrefixes"`
	// Object key prefix that identifies one or more objects to which this rule applies. Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests.
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix"`
}

// NewCommonPrefix instantiates a new CommonPrefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonPrefix() *CommonPrefix {
	this := CommonPrefix{}

	return &this
}

// NewCommonPrefixWithDefaults instantiates a new CommonPrefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonPrefixWithDefaults() *CommonPrefix {
	this := CommonPrefix{}
	return &this
}

// GetPrefix returns the Prefix field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonPrefix) GetPrefix() *string {
	if o == nil {
		return nil
	}

	return o.Prefix

}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonPrefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Prefix, true
}

// SetPrefix sets field value
func (o *CommonPrefix) SetPrefix(v string) {

	o.Prefix = &v

}

// HasPrefix returns a boolean if a field has been set.
func (o *CommonPrefix) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

func (o CommonPrefix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}

	return json.Marshal(toSerialize)
}

type NullableCommonPrefix struct {
	value *CommonPrefix
	isSet bool
}

func (v NullableCommonPrefix) Get() *CommonPrefix {
	return v.value
}

func (v *NullableCommonPrefix) Set(val *CommonPrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonPrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonPrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonPrefix(val *CommonPrefix) *NullableCommonPrefix {
	return &NullableCommonPrefix{value: val, isSet: true}
}

func (v NullableCommonPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonPrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
