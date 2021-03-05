// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

type Account_SDK struct {
	ARN *string `json:"arn,omitempty"`
	Email *string `json:"email,omitempty"`
	ID *string `json:"id,omitempty"`
	JoinedMethod *string `json:"joinedMethod,omitempty"`
	JoinedTimestamp *metav1.Time `json:"joinedTimestamp,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

type CreateAccountStatus struct {
	AccountID *string `json:"accountID,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	CompletedTimestamp *metav1.Time `json:"completedTimestamp,omitempty"`
	FailureReason *string `json:"failureReason,omitempty"`
	GovCloudAccountID *string `json:"govCloudAccountID,omitempty"`
	ID *string `json:"id,omitempty"`
	RequestedTimestamp *metav1.Time `json:"requestedTimestamp,omitempty"`
	State *string `json:"state,omitempty"`
}

type DelegatedAdministrator struct {
	ARN *string `json:"arn,omitempty"`
	DelegationEnabledDate *metav1.Time `json:"delegationEnabledDate,omitempty"`
	Email *string `json:"email,omitempty"`
	ID *string `json:"id,omitempty"`
	JoinedMethod *string `json:"joinedMethod,omitempty"`
	JoinedTimestamp *metav1.Time `json:"joinedTimestamp,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

type DelegatedService struct {
	DelegationEnabledDate *metav1.Time `json:"delegationEnabledDate,omitempty"`
}

type EffectivePolicy struct {
	LastUpdatedTimestamp *metav1.Time `json:"lastUpdatedTimestamp,omitempty"`
	PolicyContent *string `json:"policyContent,omitempty"`
}

type EnabledServicePrincipal struct {
	DateEnabled *metav1.Time `json:"dateEnabled,omitempty"`
}

type Handshake struct {
	ExpirationTimestamp *metav1.Time `json:"expirationTimestamp,omitempty"`
	RequestedTimestamp *metav1.Time `json:"requestedTimestamp,omitempty"`
}

type Organization_SDK struct {
	ARN *string `json:"arn,omitempty"`
	AvailablePolicyTypes []*PolicyTypeSummary `json:"availablePolicyTypes,omitempty"`
	FeatureSet *string `json:"featureSet,omitempty"`
	ID *string `json:"id,omitempty"`
	MasterAccountARN *string `json:"masterAccountARN,omitempty"`
	MasterAccountEmail *string `json:"masterAccountEmail,omitempty"`
	MasterAccountID *string `json:"masterAccountID,omitempty"`
}

type OrganizationalUnit_SDK struct {
	ARN *string `json:"arn,omitempty"`
	ID *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Parent struct {
	ID *string `json:"id,omitempty"`
}

type PolicySummary struct {
	ARN *string `json:"arn,omitempty"`
	AWSManaged *bool `json:"awsManaged,omitempty"`
	Description *string `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type_,omitempty"`
}

type PolicyTypeSummary struct {
	Status *string `json:"status,omitempty"`
	Type *string `json:"type_,omitempty"`
}

type Policy_SDK struct {
	Content *string `json:"content,omitempty"`
	PolicySummary *PolicySummary `json:"policySummary,omitempty"`
}

type Root struct {
	PolicyTypes []*PolicyTypeSummary `json:"policyTypes,omitempty"`
}

type Tag struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}