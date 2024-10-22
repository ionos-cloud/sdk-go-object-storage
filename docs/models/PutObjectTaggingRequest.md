# PutObjectTaggingRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TagSet** | Pointer to [**[]Tag**](Tag.md) | Contains the tag set. | [optional] |

## Methods

### NewPutObjectTaggingRequest

`func NewPutObjectTaggingRequest() *PutObjectTaggingRequest`

NewPutObjectTaggingRequest instantiates a new PutObjectTaggingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectTaggingRequestWithDefaults

`func NewPutObjectTaggingRequestWithDefaults() *PutObjectTaggingRequest`

NewPutObjectTaggingRequestWithDefaults instantiates a new PutObjectTaggingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagSet

`func (o *PutObjectTaggingRequest) GetTagSet() []Tag`

GetTagSet returns the TagSet field if non-nil, zero value otherwise.

### GetTagSetOk

`func (o *PutObjectTaggingRequest) GetTagSetOk() (*[]Tag, bool)`

GetTagSetOk returns a tuple with the TagSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSet

`func (o *PutObjectTaggingRequest) SetTagSet(v []Tag)`

SetTagSet sets TagSet field to given value.

### HasTagSet

`func (o *PutObjectTaggingRequest) HasTagSet() bool`

HasTagSet returns a boolean if a field has been set.

