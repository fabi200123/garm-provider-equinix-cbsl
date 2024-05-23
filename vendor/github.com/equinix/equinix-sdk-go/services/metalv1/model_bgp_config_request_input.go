/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the BgpConfigRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpConfigRequestInput{}

// BgpConfigRequestInput struct for BgpConfigRequestInput
type BgpConfigRequestInput struct {
	// Autonomous System Number for local BGP deployment.
	Asn            int64                               `json:"asn"`
	DeploymentType BgpConfigRequestInputDeploymentType `json:"deployment_type"`
	// The plaintext password to share between BGP neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character
	Md5 *string `json:"md5,omitempty"`
	// A use case explanation (necessary for global BGP request review).
	UseCase              *string `json:"use_case,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpConfigRequestInput BgpConfigRequestInput

// NewBgpConfigRequestInput instantiates a new BgpConfigRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfigRequestInput(asn int64, deploymentType BgpConfigRequestInputDeploymentType) *BgpConfigRequestInput {
	this := BgpConfigRequestInput{}
	this.Asn = asn
	this.DeploymentType = deploymentType
	return &this
}

// NewBgpConfigRequestInputWithDefaults instantiates a new BgpConfigRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigRequestInputWithDefaults() *BgpConfigRequestInput {
	this := BgpConfigRequestInput{}
	return &this
}

// GetAsn returns the Asn field value
func (o *BgpConfigRequestInput) GetAsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *BgpConfigRequestInput) SetAsn(v int64) {
	o.Asn = v
}

// GetDeploymentType returns the DeploymentType field value
func (o *BgpConfigRequestInput) GetDeploymentType() BgpConfigRequestInputDeploymentType {
	if o == nil {
		var ret BgpConfigRequestInputDeploymentType
		return ret
	}

	return o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetDeploymentTypeOk() (*BgpConfigRequestInputDeploymentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentType, true
}

// SetDeploymentType sets field value
func (o *BgpConfigRequestInput) SetDeploymentType(v BgpConfigRequestInputDeploymentType) {
	o.DeploymentType = v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *BgpConfigRequestInput) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *BgpConfigRequestInput) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *BgpConfigRequestInput) SetMd5(v string) {
	o.Md5 = &v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *BgpConfigRequestInput) GetUseCase() string {
	if o == nil || IsNil(o.UseCase) {
		var ret string
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigRequestInput) GetUseCaseOk() (*string, bool) {
	if o == nil || IsNil(o.UseCase) {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *BgpConfigRequestInput) HasUseCase() bool {
	if o != nil && !IsNil(o.UseCase) {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given string and assigns it to the UseCase field.
func (o *BgpConfigRequestInput) SetUseCase(v string) {
	o.UseCase = &v
}

func (o BgpConfigRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpConfigRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asn"] = o.Asn
	toSerialize["deployment_type"] = o.DeploymentType
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.UseCase) {
		toSerialize["use_case"] = o.UseCase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpConfigRequestInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asn",
		"deployment_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBgpConfigRequestInput := _BgpConfigRequestInput{}

	err = json.Unmarshal(data, &varBgpConfigRequestInput)

	if err != nil {
		return err
	}

	*o = BgpConfigRequestInput(varBgpConfigRequestInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asn")
		delete(additionalProperties, "deployment_type")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "use_case")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpConfigRequestInput struct {
	value *BgpConfigRequestInput
	isSet bool
}

func (v NullableBgpConfigRequestInput) Get() *BgpConfigRequestInput {
	return v.value
}

func (v *NullableBgpConfigRequestInput) Set(val *BgpConfigRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigRequestInput(val *BgpConfigRequestInput) *NullableBgpConfigRequestInput {
	return &NullableBgpConfigRequestInput{value: val, isSet: true}
}

func (v NullableBgpConfigRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
