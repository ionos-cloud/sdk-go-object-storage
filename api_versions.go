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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VersionsApiService VersionsApi service
type VersionsApiService service

type ApiListObjectVersionsRequest struct {
	ctx              context.Context
	ApiService       *VersionsApiService
	bucket           string
	delimiter        *string
	encodingType     *string
	keyMarker        *string
	maxKeys          *int32
	prefix           *string
	versionIdMarker  *string
	maxKeys2         *string
	keyMarker2       *string
	versionIdMarker2 *string
}

// A delimiter is a character that you specify to group keys. All keys that contain the same string between the &#x60;prefix&#x60; and the first occurrence of the delimiter are grouped under a single result element in CommonPrefixes. These groups are counted as one result against the max-keys limitation. These keys are not returned elsewhere in the response.
func (r ApiListObjectVersionsRequest) Delimiter(delimiter string) ApiListObjectVersionsRequest {
	r.delimiter = &delimiter
	return r
}

func (r ApiListObjectVersionsRequest) EncodingType(encodingType string) ApiListObjectVersionsRequest {
	r.encodingType = &encodingType
	return r
}

// Specifies the key to start with when listing objects in a bucket.
func (r ApiListObjectVersionsRequest) KeyMarker(keyMarker string) ApiListObjectVersionsRequest {
	r.keyMarker = &keyMarker
	return r
}

// Sets the maximum number of keys returned in the response. By default the operation returns up to 1,000 key names. The response might contain fewer keys but will never contain more. If additional keys satisfy the search criteria, but were not returned because max-keys was exceeded, the response contains &amp;lt;isTruncated&amp;gt;true&amp;lt;/isTruncated&amp;gt;. To return the additional keys, see key-marker and version-id-marker.
func (r ApiListObjectVersionsRequest) MaxKeys(maxKeys int32) ApiListObjectVersionsRequest {
	r.maxKeys = &maxKeys
	return r
}

// Use this parameter to select only those keys that begin with the specified prefix. You can use prefixes to separate a bucket into different groupings of keys. (You can think of using prefix to make groups in the same way you&#39;d use a folder in a file system.) You can use prefix with delimiter to roll up numerous objects into a single result under CommonPrefixes.
func (r ApiListObjectVersionsRequest) Prefix(prefix string) ApiListObjectVersionsRequest {
	r.prefix = &prefix
	return r
}

// Specifies the object version you want to start listing from.
func (r ApiListObjectVersionsRequest) VersionIdMarker(versionIdMarker string) ApiListObjectVersionsRequest {
	r.versionIdMarker = &versionIdMarker
	return r
}

// Pagination limit
func (r ApiListObjectVersionsRequest) MaxKeys2(maxKeys2 string) ApiListObjectVersionsRequest {
	r.maxKeys2 = &maxKeys2
	return r
}

// Pagination token
func (r ApiListObjectVersionsRequest) KeyMarker2(keyMarker2 string) ApiListObjectVersionsRequest {
	r.keyMarker2 = &keyMarker2
	return r
}

// Pagination token
func (r ApiListObjectVersionsRequest) VersionIdMarker2(versionIdMarker2 string) ApiListObjectVersionsRequest {
	r.versionIdMarker2 = &versionIdMarker2
	return r
}

func (r ApiListObjectVersionsRequest) Execute() (*ListObjectVersionsOutput, *APIResponse, error) {
	return r.ApiService.ListObjectVersionsExecute(r)
}

/*
ListObjectVersions ListObjectVersions

<p>Returns metadata about all versions of the objects in a bucket. You can also use request parameters as selection criteria to return metadata about a subset of all the object versions.</p> <important> <p> To use this operation, you must have permissions to perform the `ListBucketVersions` operation. Be aware of the name difference. </p> </important> <note> <p> A 200 OK response can contain valid or invalid XML. Make sure to design your application to parse the contents of the response and handle it appropriately.</p> </note> <p>To use this operation, you must have READ access to the bucket.</p>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bucket
	@return ApiListObjectVersionsRequest
*/
func (a *VersionsApiService) ListObjectVersions(ctx context.Context, bucket string) ApiListObjectVersionsRequest {
	return ApiListObjectVersionsRequest{
		ApiService: a,
		ctx:        ctx,
		bucket:     bucket,
	}
}

// Execute executes the request
//
//	@return ListObjectVersionsOutput
func (a *VersionsApiService) ListObjectVersionsExecute(r ApiListObjectVersionsRequest) (*ListObjectVersionsOutput, *APIResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListObjectVersionsOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VersionsApiService.ListObjectVersions")
	if err != nil {
		gerr := GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/{Bucket}?versions"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", parameterValueToString(r.bucket, "bucket"), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if Strlen(r.bucket) < 3 {
		return localVarReturnValue, nil, reportError("bucket must have at least 3 elements")
	}
	if Strlen(r.bucket) > 63 {
		return localVarReturnValue, nil, reportError("bucket must have less than 63 elements")
	}

	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "")
	}
	if r.encodingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "encoding-type", r.encodingType, "")
	}
	if r.keyMarker != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key-marker", r.keyMarker, "")
	}
	if r.maxKeys != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-keys", r.maxKeys, "")
	}
	if r.prefix != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prefix", r.prefix, "")
	}
	if r.versionIdMarker != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version-id-marker", r.versionIdMarker, "")
	}
	if r.maxKeys2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxKeys", r.maxKeys2, "")
	}
	if r.keyMarker2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "KeyMarker", r.keyMarker2, "")
	}
	if r.versionIdMarker2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "VersionIdMarker", r.versionIdMarker2, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)
	localVarAPIResponse := &APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "ListObjectVersions",
	}
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
