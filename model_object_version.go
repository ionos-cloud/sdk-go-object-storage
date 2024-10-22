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
	"time"
)

import "encoding/xml"

// ObjectVersion The version of an object.
type ObjectVersion struct {
	XMLName xml.Name `xml:"Version"`
	// Entity tag that identifies the object's data. Objects with different object data will have different entity tags. The entity tag is an opaque string. The entity tag may or may not be an MD5 digest of the object data. If the entity tag is not an MD5 digest of the object data, it will contain one or more nonhexadecimal characters and/or will consist of less than 32 or more than 32 hexadecimal digits.
	ETag *string `json:"ETag,omitempty" xml:"ETag"`
	// Size in bytes of the object
	Size         *int32                     `json:"Size,omitempty" xml:"Size"`
	StorageClass *ObjectVersionStorageClass `json:"StorageClass,omitempty" xml:"StorageClass"`
	// The object key.
	Key *string `json:"Key,omitempty" xml:"Key"`
	// Version ID of an object.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId"`
	// Specifies whether the object is (true) or is not (false) the latest version of an object.
	IsLatest *bool `json:"IsLatest,omitempty" xml:"IsLatest"`
	// Creation date of the object.
	LastModified *IonosTime `json:"LastModified,omitempty" xml:"LastModified"`
	Owner        *Owner     `json:"Owner,omitempty" xml:"Owner"`
}

// NewObjectVersion instantiates a new ObjectVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectVersion() *ObjectVersion {
	this := ObjectVersion{}

	return &this
}

// NewObjectVersionWithDefaults instantiates a new ObjectVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectVersionWithDefaults() *ObjectVersion {
	this := ObjectVersion{}
	return &this
}

// GetETag returns the ETag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectVersion) GetETag() *string {
	if o == nil {
		return nil
	}

	return o.ETag

}

// GetETagOk returns a tuple with the ETag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetETagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ETag, true
}

// SetETag sets field value
func (o *ObjectVersion) SetETag(v string) {

	o.ETag = &v

}

// HasETag returns a boolean if a field has been set.
func (o *ObjectVersion) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ObjectVersion) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *ObjectVersion) SetSize(v int32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *ObjectVersion) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for ObjectVersionStorageClass will be returned
func (o *ObjectVersion) GetStorageClass() *ObjectVersionStorageClass {
	if o == nil {
		return nil
	}

	return o.StorageClass

}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetStorageClassOk() (*ObjectVersionStorageClass, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageClass, true
}

// SetStorageClass sets field value
func (o *ObjectVersion) SetStorageClass(v ObjectVersionStorageClass) {

	o.StorageClass = &v

}

// HasStorageClass returns a boolean if a field has been set.
func (o *ObjectVersion) HasStorageClass() bool {
	if o != nil && o.StorageClass != nil {
		return true
	}

	return false
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectVersion) GetKey() *string {
	if o == nil {
		return nil
	}

	return o.Key

}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Key, true
}

// SetKey sets field value
func (o *ObjectVersion) SetKey(v string) {

	o.Key = &v

}

// HasKey returns a boolean if a field has been set.
func (o *ObjectVersion) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// GetVersionId returns the VersionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObjectVersion) GetVersionId() *string {
	if o == nil {
		return nil
	}

	return o.VersionId

}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.VersionId, true
}

// SetVersionId sets field value
func (o *ObjectVersion) SetVersionId(v string) {

	o.VersionId = &v

}

// HasVersionId returns a boolean if a field has been set.
func (o *ObjectVersion) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// GetIsLatest returns the IsLatest field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ObjectVersion) GetIsLatest() *bool {
	if o == nil {
		return nil
	}

	return o.IsLatest

}

// GetIsLatestOk returns a tuple with the IsLatest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetIsLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.IsLatest, true
}

// SetIsLatest sets field value
func (o *ObjectVersion) SetIsLatest(v bool) {

	o.IsLatest = &v

}

// HasIsLatest returns a boolean if a field has been set.
func (o *ObjectVersion) HasIsLatest() bool {
	if o != nil && o.IsLatest != nil {
		return true
	}

	return false
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ObjectVersion) GetLastModified() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModified == nil {
		return nil
	}
	return &o.LastModified.Time

}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModified == nil {
		return nil, false
	}
	return &o.LastModified.Time, true

}

// SetLastModified sets field value
func (o *ObjectVersion) SetLastModified(v time.Time) {

	o.LastModified = &IonosTime{v}

}

// HasLastModified returns a boolean if a field has been set.
func (o *ObjectVersion) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for Owner will be returned
func (o *ObjectVersion) GetOwner() *Owner {
	if o == nil {
		return nil
	}

	return o.Owner

}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectVersion) GetOwnerOk() (*Owner, bool) {
	if o == nil {
		return nil, false
	}

	return o.Owner, true
}

// SetOwner sets field value
func (o *ObjectVersion) SetOwner(v Owner) {

	o.Owner = &v

}

// HasOwner returns a boolean if a field has been set.
func (o *ObjectVersion) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

func (o ObjectVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ETag != nil {
		toSerialize["ETag"] = o.ETag
	}

	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}

	if o.StorageClass != nil {
		toSerialize["StorageClass"] = o.StorageClass
	}

	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}

	if o.VersionId != nil {
		toSerialize["VersionId"] = o.VersionId
	}

	if o.IsLatest != nil {
		toSerialize["IsLatest"] = o.IsLatest
	}

	if o.LastModified != nil {
		toSerialize["LastModified"] = o.LastModified
	}

	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	return json.Marshal(toSerialize)
}

type NullableObjectVersion struct {
	value *ObjectVersion
	isSet bool
}

func (v NullableObjectVersion) Get() *ObjectVersion {
	return v.value
}

func (v *NullableObjectVersion) Set(val *ObjectVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectVersion(val *ObjectVersion) *NullableObjectVersion {
	return &NullableObjectVersion{value: val, isSet: true}
}

func (v NullableObjectVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
