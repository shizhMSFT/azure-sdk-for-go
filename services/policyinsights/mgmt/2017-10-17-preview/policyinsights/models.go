package policyinsights

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// PolicyStatesResource enumerates the values for policy states resource.
type PolicyStatesResource string

const (
	// Default ...
	Default PolicyStatesResource = "default"
	// Latest ...
	Latest PolicyStatesResource = "latest"
)

// Operation operation definition.
type Operation struct {
	// Name - Operation name.
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name.
	Operation *string `json:"operation,omitempty"`
	// Description - Operation description.
	Description *string `json:"description,omitempty"`
}

// Operations list of available operations.
type Operations struct {
	autorest.Response `json:"-"`
	// Value - List of available operations.
	Value *[]Operation `json:"value,omitempty"`
}

// PolicyEvent policy event record.
type PolicyEvent struct {
	// OdataID - OData entity ID; always set to null since policy event records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy event record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
	// TenantID - Tenant ID for the policy event record.
	TenantID *string `json:"tenantId,omitempty"`
	// PrincipalOid - Principal object ID for the user who initiated the resource operation that triggered the policy event.
	PrincipalOid *string `json:"principalOid,omitempty"`
}

// PolicyEventsQueryResults query results.
type PolicyEventsQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy event records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyEvent `json:"value,omitempty"`
}

// PolicyState policy state record.
type PolicyState struct {
	// OdataID - OData entity ID; always set to null since policy state records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy state record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
}

// PolicyStatesQueryResults query results.
type PolicyStatesQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy state records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyState `json:"value,omitempty"`
}

// QueryFailure error response.
type QueryFailure struct {
	// Error - Error definition.
	Error *QueryFailureError `json:"error,omitempty"`
}

// QueryFailureError error definition.
type QueryFailureError struct {
	// Code - Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
}

// String ...
type String struct {
	autorest.Response `json:"-"`
	Value             *string `json:"value,omitempty"`
}
