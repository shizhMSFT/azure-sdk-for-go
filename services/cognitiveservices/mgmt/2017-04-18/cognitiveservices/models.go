package cognitiveservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// KeyName enumerates the values for key name.
type KeyName string

const (
	// Key1 ...
	Key1 KeyName = "Key1"
	// Key2 ...
	Key2 KeyName = "Key2"
)

// PossibleKeyNameValues returns an array of possible values for the KeyName const type.
func PossibleKeyNameValues() [2]KeyName {
	return [2]KeyName{Key1, Key2}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// BingAutosuggestv7 ...
	BingAutosuggestv7 Kind = "Bing.Autosuggest.v7"
	// BingCustomSearch ...
	BingCustomSearch Kind = "Bing.CustomSearch"
	// BingSearchv7 ...
	BingSearchv7 Kind = "Bing.Search.v7"
	// BingSpeech ...
	BingSpeech Kind = "Bing.Speech"
	// BingSpellCheckv7 ...
	BingSpellCheckv7 Kind = "Bing.SpellCheck.v7"
	// ComputerVision ...
	ComputerVision Kind = "ComputerVision"
	// ContentModerator ...
	ContentModerator Kind = "ContentModerator"
	// CustomSpeech ...
	CustomSpeech Kind = "CustomSpeech"
	// CustomVisionPrediction ...
	CustomVisionPrediction Kind = "CustomVision.Prediction"
	// CustomVisionTraining ...
	CustomVisionTraining Kind = "CustomVision.Training"
	// Emotion ...
	Emotion Kind = "Emotion"
	// Face ...
	Face Kind = "Face"
	// LUIS ...
	LUIS Kind = "LUIS"
	// QnAMaker ...
	QnAMaker Kind = "QnAMaker"
	// SpeakerRecognition ...
	SpeakerRecognition Kind = "SpeakerRecognition"
	// SpeechTranslation ...
	SpeechTranslation Kind = "SpeechTranslation"
	// TextAnalytics ...
	TextAnalytics Kind = "TextAnalytics"
	// TextTranslation ...
	TextTranslation Kind = "TextTranslation"
	// WebLM ...
	WebLM Kind = "WebLM"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() [19]Kind {
	return [19]Kind{BingAutosuggestv7, BingCustomSearch, BingSearchv7, BingSpeech, BingSpellCheckv7, ComputerVision, ContentModerator, CustomSpeech, CustomVisionPrediction, CustomVisionTraining, Emotion, Face, LUIS, QnAMaker, SpeakerRecognition, SpeechTranslation, TextAnalytics, TextTranslation, WebLM}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() [6]ProvisioningState {
	return [6]ProvisioningState{Creating, Deleting, Failed, Moving, ResolvingDNS, Succeeded}
}

// QuotaUsageStatus enumerates the values for quota usage status.
type QuotaUsageStatus string

const (
	// Blocked ...
	Blocked QuotaUsageStatus = "Blocked"
	// Included ...
	Included QuotaUsageStatus = "Included"
	// InOverage ...
	InOverage QuotaUsageStatus = "InOverage"
	// Unknown ...
	Unknown QuotaUsageStatus = "Unknown"
)

// PossibleQuotaUsageStatusValues returns an array of possible values for the QuotaUsageStatus const type.
func PossibleQuotaUsageStatusValues() [4]QuotaUsageStatus {
	return [4]QuotaUsageStatus{Blocked, Included, InOverage, Unknown}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// F0 ...
	F0 SkuName = "F0"
	// P0 ...
	P0 SkuName = "P0"
	// P1 ...
	P1 SkuName = "P1"
	// P2 ...
	P2 SkuName = "P2"
	// S0 ...
	S0 SkuName = "S0"
	// S1 ...
	S1 SkuName = "S1"
	// S2 ...
	S2 SkuName = "S2"
	// S3 ...
	S3 SkuName = "S3"
	// S4 ...
	S4 SkuName = "S4"
	// S5 ...
	S5 SkuName = "S5"
	// S6 ...
	S6 SkuName = "S6"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() [11]SkuName {
	return [11]SkuName{F0, P0, P1, P2, S0, S1, S2, S3, S4, S5, S6}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Free ...
	Free SkuTier = "Free"
	// Premium ...
	Premium SkuTier = "Premium"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() [3]SkuTier {
	return [3]SkuTier{Free, Premium, Standard}
}

// UnitType enumerates the values for unit type.
type UnitType string

const (
	// Bytes ...
	Bytes UnitType = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UnitType = "BytesPerSecond"
	// Count ...
	Count UnitType = "Count"
	// CountPerSecond ...
	CountPerSecond UnitType = "CountPerSecond"
	// Milliseconds ...
	Milliseconds UnitType = "Milliseconds"
	// Percent ...
	Percent UnitType = "Percent"
	// Seconds ...
	Seconds UnitType = "Seconds"
)

// PossibleUnitTypeValues returns an array of possible values for the UnitType const type.
func PossibleUnitTypeValues() [7]UnitType {
	return [7]UnitType{Bytes, BytesPerSecond, Count, CountPerSecond, Milliseconds, Percent, Seconds}
}

// Account cognitive Services Account is an Azure resource representing the provisioned account, its type, location
// and SKU.
type Account struct {
	autorest.Response `json:"-"`
	// Etag - Entity Tag
	Etag *string `json:"etag,omitempty"`
	// ID - The id of the created account
	ID *string `json:"id,omitempty"`
	// Kind - Type of cognitive service account.
	Kind *string `json:"kind,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Name - The name of the created account
	Name *string `json:"name,omitempty"`
	// AccountProperties - Properties of Cognitive Services account.
	*AccountProperties `json:"properties,omitempty"`
	// Sku - The SKU of Cognitive Services account.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Etag != nil {
		objectMap["etag"] = a.Etag
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Kind != nil {
		objectMap["kind"] = a.Kind
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				a.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "kind":
			if v != nil {
				var kind string
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				a.Kind = &kind
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		}
	}

	return nil
}

// AccountCreateParameters the parameters to provide for the account.
type AccountCreateParameters struct {
	// Sku - Required. Gets or sets the SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
	// Kind - Required. Gets or sets the Kind of the resource. Possible values include: 'BingAutosuggestv7', 'BingCustomSearch', 'BingSearchv7', 'BingSpeech', 'BingSpellCheckv7', 'ComputerVision', 'ContentModerator', 'CustomSpeech', 'CustomVisionPrediction', 'CustomVisionTraining', 'Emotion', 'Face', 'LUIS', 'QnAMaker', 'SpeakerRecognition', 'SpeechTranslation', 'TextAnalytics', 'TextTranslation', 'WebLM'
	Kind Kind `json:"kind,omitempty"`
	// Location - Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update the request will succeed.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Properties - Must exist in the request. Must be an empty object. Must not be null.
	Properties interface{} `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Sku != nil {
		objectMap["sku"] = acp.Sku
	}
	objectMap["kind"] = acp.Kind
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	objectMap["properties"] = acp.Properties
	return json.Marshal(objectMap)
}

// AccountEnumerateSkusResult the list of cognitive services accounts operation response.
type AccountEnumerateSkusResult struct {
	autorest.Response `json:"-"`
	// Value - Gets the list of Cognitive Services accounts and their properties.
	Value *[]ResourceAndSku `json:"value,omitempty"`
}

// AccountKeys the access keys for the cognitive services account.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Gets the value of key 1.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Gets the value of key 2.
	Key2 *string `json:"key2,omitempty"`
}

// AccountListResult the list of cognitive services accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of accounts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - Gets the list of Cognitive Services accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListResultIterator) Response() AccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListResultIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer() (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) Next() error {
	next, err := page.fn(page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListResultPage) Values() []Account {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// AccountProperties properties of Cognitive Services account.
type AccountProperties struct {
	// ProvisioningState - Gets the status of the cognitive services account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Moving', 'Deleting', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Endpoint - Endpoint of the created account.
	Endpoint *string `json:"endpoint,omitempty"`
	// InternalID - The internal identifier.
	InternalID *string `json:"internalId,omitempty"`
}

// AccountUpdateParameters the parameters to provide for the account.
type AccountUpdateParameters struct {
	// Sku - Gets or sets the SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	return json.Marshal(objectMap)
}

// CheckSkuAvailabilityParameter check SKU availability parameter.
type CheckSkuAvailabilityParameter struct {
	// Skus - The SKU of the resource.
	Skus *[]SkuName `json:"skus,omitempty"`
	// Kind - The Kind of the resource. Possible values include: 'BingAutosuggestv7', 'BingCustomSearch', 'BingSearchv7', 'BingSpeech', 'BingSpellCheckv7', 'ComputerVision', 'ContentModerator', 'CustomSpeech', 'CustomVisionPrediction', 'CustomVisionTraining', 'Emotion', 'Face', 'LUIS', 'QnAMaker', 'SpeakerRecognition', 'SpeechTranslation', 'TextAnalytics', 'TextTranslation', 'WebLM'
	Kind Kind `json:"kind,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
}

// CheckSkuAvailabilityResult check SKU availability result.
type CheckSkuAvailabilityResult struct {
	// Kind - The Kind of the resource. Possible values include: 'BingAutosuggestv7', 'BingCustomSearch', 'BingSearchv7', 'BingSpeech', 'BingSpellCheckv7', 'ComputerVision', 'ContentModerator', 'CustomSpeech', 'CustomVisionPrediction', 'CustomVisionTraining', 'Emotion', 'Face', 'LUIS', 'QnAMaker', 'SpeakerRecognition', 'SpeechTranslation', 'TextAnalytics', 'TextTranslation', 'WebLM'
	Kind Kind `json:"kind,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
	// SkuName - The SKU of Cognitive Services account. Possible values include: 'F0', 'P0', 'P1', 'P2', 'S0', 'S1', 'S2', 'S3', 'S4', 'S5', 'S6'
	SkuName SkuName `json:"skuName,omitempty"`
	// SkuAvailable - Indicates the given SKU is available or not.
	SkuAvailable *bool `json:"skuAvailable,omitempty"`
	// Reason - Reason why the SKU is not available.
	Reason *string `json:"reason,omitempty"`
	// Message - Additional error message.
	Message *string `json:"message,omitempty"`
}

// CheckSkuAvailabilityResultList check SKU availability result list.
type CheckSkuAvailabilityResultList struct {
	autorest.Response `json:"-"`
	// Value - Check SKU availability result list.
	Value *[]CheckSkuAvailabilityResult `json:"value,omitempty"`
}

// Error cognitive Services error object.
type Error struct {
	// Error - The error body.
	Error *ErrorBody `json:"error,omitempty"`
}

// ErrorBody cognitive Services error body.
type ErrorBody struct {
	// Code - error code
	Code *string `json:"code,omitempty"`
	// Message - error message
	Message *string `json:"message,omitempty"`
}

// MetricName a metric name.
type MetricName struct {
	// Value - The name of the metric.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - The friendly name of the metric.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// OperationDisplayInfo the operation supported by Cognitive Services.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Microsoft Cognitive Services.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Cognitive Services.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Cognitive Services.
	Display *OperationDisplayInfo `json:"display,omitempty"`
	// Origin - The origin of the operation.
	Origin *string `json:"origin,omitempty"`
	// Properties - Additional properties.
	Properties interface{} `json:"properties,omitempty"`
}

// OperationEntityListResult the list of cognitive services accounts operation response.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The list of operations.
	Value *[]OperationEntity `json:"value,omitempty"`
}

// OperationEntityListResultIterator provides access to a complete listing of OperationEntity values.
type OperationEntityListResultIterator struct {
	i    int
	page OperationEntityListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationEntityListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationEntityListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationEntityListResultIterator) Response() OperationEntityListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationEntityListResultIterator) Value() OperationEntity {
	if !iter.page.NotDone() {
		return OperationEntity{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (oelr OperationEntityListResult) IsEmpty() bool {
	return oelr.Value == nil || len(*oelr.Value) == 0
}

// operationEntityListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (oelr OperationEntityListResult) operationEntityListResultPreparer() (*http.Request, error) {
	if oelr.NextLink == nil || len(to.String(oelr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(oelr.NextLink)))
}

// OperationEntityListResultPage contains a page of OperationEntity values.
type OperationEntityListResultPage struct {
	fn   func(OperationEntityListResult) (OperationEntityListResult, error)
	oelr OperationEntityListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationEntityListResultPage) Next() error {
	next, err := page.fn(page.oelr)
	if err != nil {
		return err
	}
	page.oelr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationEntityListResultPage) NotDone() bool {
	return !page.oelr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationEntityListResultPage) Response() OperationEntityListResult {
	return page.oelr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationEntityListResultPage) Values() []OperationEntity {
	if page.oelr.IsEmpty() {
		return nil
	}
	return *page.oelr.Value
}

// RegenerateKeyParameters regenerate key parameters.
type RegenerateKeyParameters struct {
	// KeyName - key name to generate (Key1|Key2). Possible values include: 'Key1', 'Key2'
	KeyName KeyName `json:"keyName,omitempty"`
}

// ResourceAndSku cognitive Services resource type and SKU.
type ResourceAndSku struct {
	// ResourceType - Resource Namespace and Type
	ResourceType *string `json:"resourceType,omitempty"`
	// Sku - The SKU of Cognitive Services account.
	Sku *Sku `json:"sku,omitempty"`
}

// Sku the SKU of the cognitive services account.
type Sku struct {
	// Name - Gets or sets the sku name. Required for account creation, optional for update. Possible values include: 'F0', 'P0', 'P1', 'P2', 'S0', 'S1', 'S2', 'S3', 'S4', 'S5', 'S6'
	Name SkuName `json:"name,omitempty"`
	// Tier - Gets the sku tier. This is based on the SKU name. Possible values include: 'Free', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
}

// Usage the usage data for a usage request.
type Usage struct {
	// Unit - The unit of the metric. Possible values include: 'Count', 'Bytes', 'Seconds', 'Percent', 'CountPerSecond', 'BytesPerSecond', 'Milliseconds'
	Unit UnitType `json:"unit,omitempty"`
	// Name - The name information for the metric.
	Name *MetricName `json:"name,omitempty"`
	// QuotaPeriod - The quota period used to summarize the usage values.
	QuotaPeriod *string `json:"quotaPeriod,omitempty"`
	// Limit - Maximum value for this metric.
	Limit *int32 `json:"limit,omitempty"`
	// CurrentValue - Current value for this metric.
	CurrentValue *int32 `json:"currentValue,omitempty"`
	// NextResetTime - Next reset time for current quota.
	NextResetTime *string `json:"nextResetTime,omitempty"`
	// Status - Cognitive Services account quota usage status. Possible values include: 'Included', 'Blocked', 'InOverage', 'Unknown'
	Status QuotaUsageStatus `json:"status,omitempty"`
}

// UsagesResult the response to a list usage request.
type UsagesResult struct {
	autorest.Response `json:"-"`
	// Value - The list of usages for Cognitive Service account.
	Value *[]Usage `json:"value,omitempty"`
}
