package storage

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

type AccountStatus string

const (
	Available   AccountStatus = "Available"
	Unavailable AccountStatus = "Unavailable"
)

type AccountType string

const (
	PremiumLRS    AccountType = "Premium_LRS"
	StandardGRS   AccountType = "Standard_GRS"
	StandardLRS   AccountType = "Standard_LRS"
	StandardRAGRS AccountType = "Standard_RAGRS"
	StandardZRS   AccountType = "Standard_ZRS"
)

type KeyName string

const (
	Key1 KeyName = "key1"
	Key2 KeyName = "key2"
)

type ProvisioningState string

const (
	Creating     ProvisioningState = "Creating"
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	Succeeded    ProvisioningState = "Succeeded"
)

type Reason string

const (
	AccountNameInvalid Reason = "AccountNameInvalid"
	AlreadyExists      Reason = "AlreadyExists"
)

type UsageUnit string

const (
	Bytes           UsageUnit = "Bytes"
	BytesPerSecond  UsageUnit = "BytesPerSecond"
	Count           UsageUnit = "Count"
	CountsPerSecond UsageUnit = "CountsPerSecond"
	Percent         UsageUnit = "Percent"
	Seconds         UsageUnit = "Seconds"
)

// The CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	NameAvailable     bool   `json:"nameAvailable,omitempty"`
	Reason            Reason `json:"reason,omitempty"`
	Message           string `json:"message,omitempty"`
}

// The storage account.
type StorageAccount struct {
	autorest.Response `json:"-"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	Location          string            `json:"location,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Properties        struct {
		ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
		AccountType       AccountType       `json:"accountType,omitempty"`
		PrimaryEndpoints  struct {
			Blob  string `json:"blob,omitempty"`
			Queue string `json:"queue,omitempty"`
			Table string `json:"table,omitempty"`
		} `json:"primaryEndpoints,omitempty"`
		PrimaryLocation     string        `json:"primaryLocation,omitempty"`
		StatusOfPrimary     AccountStatus `json:"statusOfPrimary,omitempty"`
		LastGeoFailoverTime date.Time     `json:"lastGeoFailoverTime,omitempty"`
		SecondaryLocation   string        `json:"secondaryLocation,omitempty"`
		StatusOfSecondary   AccountStatus `json:"statusOfSecondary,omitempty"`
		CreationTime        date.Time     `json:"creationTime,omitempty"`
		CustomDomain        struct {
			Name         string `json:"name,omitempty"`
			UseSubDomain bool   `json:"useSubDomain,omitempty"`
		} `json:"customDomain,omitempty"`
		SecondaryEndpoints struct {
			Blob  string `json:"blob,omitempty"`
			Queue string `json:"queue,omitempty"`
			Table string `json:"table,omitempty"`
		} `json:"secondaryEndpoints,omitempty"`
	} `json:"properties,omitempty"`
}

type StorageAccountCheckNameAvailabilityParameters struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// The parameters to provide for the account.
type StorageAccountCreateParameters struct {
	Id         string            `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Type       string            `json:"type,omitempty"`
	Location   string            `json:"location,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
	Properties struct {
		AccountType AccountType `json:"accountType,omitempty"`
	} `json:"properties,omitempty"`
}

// The access keys for the storage account.
type StorageAccountKeys struct {
	autorest.Response `json:"-"`
	Key1              string `json:"key1,omitempty"`
	Key2              string `json:"key2,omitempty"`
}

// The list storage accounts operation response.
type StorageAccountListResult struct {
	autorest.Response `json:"-"`
	Value             []StorageAccount `json:"value,omitempty"`
}

type StorageAccountRegenerateKeyParameters struct {
	KeyName KeyName `json:"keyName,omitempty"`
}

// The parameters to update on the account.
type StorageAccountUpdateParameters struct {
	Id         string            `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Type       string            `json:"type,omitempty"`
	Location   string            `json:"location,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
	Properties struct {
		AccountType  AccountType `json:"accountType,omitempty"`
		CustomDomain struct {
			Name         string `json:"name,omitempty"`
			UseSubDomain bool   `json:"useSubDomain,omitempty"`
		} `json:"customDomain,omitempty"`
	} `json:"properties,omitempty"`
}

// Describes Storage Resource Usage.
type Usage struct {
	Unit         UsageUnit `json:"unit,omitempty"`
	CurrentValue int       `json:"currentValue,omitempty"`
	Limit        int       `json:"limit,omitempty"`
	Name         struct {
		Value          string `json:"value,omitempty"`
		LocalizedValue string `json:"localizedValue,omitempty"`
	} `json:"name,omitempty"`
}

// The List Usages operation response.
type UsageListResult struct {
	autorest.Response `json:"-"`
	Value             []Usage `json:"value,omitempty"`
}
