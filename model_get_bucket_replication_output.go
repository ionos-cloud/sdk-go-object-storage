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

// GetBucketReplicationOutput struct for GetBucketReplicationOutput
type GetBucketReplicationOutput struct {
	XMLName                  xml.Name                  `xml:"GetBucketReplicationOutput"`
	ReplicationConfiguration *ReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration"`
}

// NewGetBucketReplicationOutput instantiates a new GetBucketReplicationOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketReplicationOutput() *GetBucketReplicationOutput {
	this := GetBucketReplicationOutput{}

	return &this
}

// NewGetBucketReplicationOutputWithDefaults instantiates a new GetBucketReplicationOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketReplicationOutputWithDefaults() *GetBucketReplicationOutput {
	this := GetBucketReplicationOutput{}
	return &this
}

// GetReplicationConfiguration returns the ReplicationConfiguration field value
// If the value is explicit nil, the zero value for ReplicationConfiguration will be returned
func (o *GetBucketReplicationOutput) GetReplicationConfiguration() *ReplicationConfiguration {
	if o == nil {
		return nil
	}

	return o.ReplicationConfiguration

}

// GetReplicationConfigurationOk returns a tuple with the ReplicationConfiguration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBucketReplicationOutput) GetReplicationConfigurationOk() (*ReplicationConfiguration, bool) {
	if o == nil {
		return nil, false
	}

	return o.ReplicationConfiguration, true
}

// SetReplicationConfiguration sets field value
func (o *GetBucketReplicationOutput) SetReplicationConfiguration(v ReplicationConfiguration) {

	o.ReplicationConfiguration = &v

}

// HasReplicationConfiguration returns a boolean if a field has been set.
func (o *GetBucketReplicationOutput) HasReplicationConfiguration() bool {
	if o != nil && o.ReplicationConfiguration != nil {
		return true
	}

	return false
}

func (o GetBucketReplicationOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReplicationConfiguration != nil {
		toSerialize["ReplicationConfiguration"] = o.ReplicationConfiguration
	}

	return json.Marshal(toSerialize)
}

type NullableGetBucketReplicationOutput struct {
	value *GetBucketReplicationOutput
	isSet bool
}

func (v NullableGetBucketReplicationOutput) Get() *GetBucketReplicationOutput {
	return v.value
}

func (v *NullableGetBucketReplicationOutput) Set(val *GetBucketReplicationOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketReplicationOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketReplicationOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketReplicationOutput(val *GetBucketReplicationOutput) *NullableGetBucketReplicationOutput {
	return &NullableGetBucketReplicationOutput{value: val, isSet: true}
}

func (v NullableGetBucketReplicationOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketReplicationOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
