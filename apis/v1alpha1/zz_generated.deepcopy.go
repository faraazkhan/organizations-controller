// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.IAMUserAccessToBilling != nil {
		in, out := &in.IAMUserAccessToBilling, &out.IAMUserAccessToBilling
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.GovCloudAccountID != nil {
		in, out := &in.GovCloudAccountID, &out.GovCloudAccountID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RequestedTimestamp != nil {
		in, out := &in.RequestedTimestamp, &out.RequestedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account_SDK) DeepCopyInto(out *Account_SDK) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.JoinedMethod != nil {
		in, out := &in.JoinedMethod, &out.JoinedMethod
		*out = new(string)
		**out = **in
	}
	if in.JoinedTimestamp != nil {
		in, out := &in.JoinedTimestamp, &out.JoinedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account_SDK.
func (in *Account_SDK) DeepCopy() *Account_SDK {
	if in == nil {
		return nil
	}
	out := new(Account_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateAccountStatus) DeepCopyInto(out *CreateAccountStatus) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.GovCloudAccountID != nil {
		in, out := &in.GovCloudAccountID, &out.GovCloudAccountID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RequestedTimestamp != nil {
		in, out := &in.RequestedTimestamp, &out.RequestedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateAccountStatus.
func (in *CreateAccountStatus) DeepCopy() *CreateAccountStatus {
	if in == nil {
		return nil
	}
	out := new(CreateAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegatedAdministrator) DeepCopyInto(out *DelegatedAdministrator) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.DelegationEnabledDate != nil {
		in, out := &in.DelegationEnabledDate, &out.DelegationEnabledDate
		*out = (*in).DeepCopy()
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.JoinedMethod != nil {
		in, out := &in.JoinedMethod, &out.JoinedMethod
		*out = new(string)
		**out = **in
	}
	if in.JoinedTimestamp != nil {
		in, out := &in.JoinedTimestamp, &out.JoinedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegatedAdministrator.
func (in *DelegatedAdministrator) DeepCopy() *DelegatedAdministrator {
	if in == nil {
		return nil
	}
	out := new(DelegatedAdministrator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegatedService) DeepCopyInto(out *DelegatedService) {
	*out = *in
	if in.DelegationEnabledDate != nil {
		in, out := &in.DelegationEnabledDate, &out.DelegationEnabledDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegatedService.
func (in *DelegatedService) DeepCopy() *DelegatedService {
	if in == nil {
		return nil
	}
	out := new(DelegatedService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectivePolicy) DeepCopyInto(out *EffectivePolicy) {
	*out = *in
	if in.LastUpdatedTimestamp != nil {
		in, out := &in.LastUpdatedTimestamp, &out.LastUpdatedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.PolicyContent != nil {
		in, out := &in.PolicyContent, &out.PolicyContent
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectivePolicy.
func (in *EffectivePolicy) DeepCopy() *EffectivePolicy {
	if in == nil {
		return nil
	}
	out := new(EffectivePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnabledServicePrincipal) DeepCopyInto(out *EnabledServicePrincipal) {
	*out = *in
	if in.DateEnabled != nil {
		in, out := &in.DateEnabled, &out.DateEnabled
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnabledServicePrincipal.
func (in *EnabledServicePrincipal) DeepCopy() *EnabledServicePrincipal {
	if in == nil {
		return nil
	}
	out := new(EnabledServicePrincipal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovCloudAccount) DeepCopyInto(out *GovCloudAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovCloudAccount.
func (in *GovCloudAccount) DeepCopy() *GovCloudAccount {
	if in == nil {
		return nil
	}
	out := new(GovCloudAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GovCloudAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovCloudAccountList) DeepCopyInto(out *GovCloudAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GovCloudAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovCloudAccountList.
func (in *GovCloudAccountList) DeepCopy() *GovCloudAccountList {
	if in == nil {
		return nil
	}
	out := new(GovCloudAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GovCloudAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovCloudAccountSpec) DeepCopyInto(out *GovCloudAccountSpec) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.IAMUserAccessToBilling != nil {
		in, out := &in.IAMUserAccessToBilling, &out.IAMUserAccessToBilling
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovCloudAccountSpec.
func (in *GovCloudAccountSpec) DeepCopy() *GovCloudAccountSpec {
	if in == nil {
		return nil
	}
	out := new(GovCloudAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovCloudAccountStatus) DeepCopyInto(out *GovCloudAccountStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.GovCloudAccountID != nil {
		in, out := &in.GovCloudAccountID, &out.GovCloudAccountID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RequestedTimestamp != nil {
		in, out := &in.RequestedTimestamp, &out.RequestedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovCloudAccountStatus.
func (in *GovCloudAccountStatus) DeepCopy() *GovCloudAccountStatus {
	if in == nil {
		return nil
	}
	out := new(GovCloudAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handshake) DeepCopyInto(out *Handshake) {
	*out = *in
	if in.ExpirationTimestamp != nil {
		in, out := &in.ExpirationTimestamp, &out.ExpirationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.RequestedTimestamp != nil {
		in, out := &in.RequestedTimestamp, &out.RequestedTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handshake.
func (in *Handshake) DeepCopy() *Handshake {
	if in == nil {
		return nil
	}
	out := new(Handshake)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Organization) DeepCopyInto(out *Organization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Organization.
func (in *Organization) DeepCopy() *Organization {
	if in == nil {
		return nil
	}
	out := new(Organization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Organization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationList) DeepCopyInto(out *OrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Organization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationList.
func (in *OrganizationList) DeepCopy() *OrganizationList {
	if in == nil {
		return nil
	}
	out := new(OrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpec) DeepCopyInto(out *OrganizationSpec) {
	*out = *in
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpec.
func (in *OrganizationSpec) DeepCopy() *OrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationStatus) DeepCopyInto(out *OrganizationStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailablePolicyTypes != nil {
		in, out := &in.AvailablePolicyTypes, &out.AvailablePolicyTypes
		*out = make([]*PolicyTypeSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyTypeSummary)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountARN != nil {
		in, out := &in.MasterAccountARN, &out.MasterAccountARN
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountEmail != nil {
		in, out := &in.MasterAccountEmail, &out.MasterAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountID != nil {
		in, out := &in.MasterAccountID, &out.MasterAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationStatus.
func (in *OrganizationStatus) DeepCopy() *OrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Organization_SDK) DeepCopyInto(out *Organization_SDK) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.AvailablePolicyTypes != nil {
		in, out := &in.AvailablePolicyTypes, &out.AvailablePolicyTypes
		*out = make([]*PolicyTypeSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyTypeSummary)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountARN != nil {
		in, out := &in.MasterAccountARN, &out.MasterAccountARN
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountEmail != nil {
		in, out := &in.MasterAccountEmail, &out.MasterAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.MasterAccountID != nil {
		in, out := &in.MasterAccountID, &out.MasterAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Organization_SDK.
func (in *Organization_SDK) DeepCopy() *Organization_SDK {
	if in == nil {
		return nil
	}
	out := new(Organization_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationalUnit) DeepCopyInto(out *OrganizationalUnit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationalUnit.
func (in *OrganizationalUnit) DeepCopy() *OrganizationalUnit {
	if in == nil {
		return nil
	}
	out := new(OrganizationalUnit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationalUnit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationalUnitList) DeepCopyInto(out *OrganizationalUnitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationalUnit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationalUnitList.
func (in *OrganizationalUnitList) DeepCopy() *OrganizationalUnitList {
	if in == nil {
		return nil
	}
	out := new(OrganizationalUnitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationalUnitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationalUnitSpec) DeepCopyInto(out *OrganizationalUnitSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParentID != nil {
		in, out := &in.ParentID, &out.ParentID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationalUnitSpec.
func (in *OrganizationalUnitSpec) DeepCopy() *OrganizationalUnitSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationalUnitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationalUnitStatus) DeepCopyInto(out *OrganizationalUnitStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationalUnitStatus.
func (in *OrganizationalUnitStatus) DeepCopy() *OrganizationalUnitStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationalUnitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationalUnit_SDK) DeepCopyInto(out *OrganizationalUnit_SDK) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationalUnit_SDK.
func (in *OrganizationalUnit_SDK) DeepCopy() *OrganizationalUnit_SDK {
	if in == nil {
		return nil
	}
	out := new(OrganizationalUnit_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.PolicySummary != nil {
		in, out := &in.PolicySummary, &out.PolicySummary
		*out = new(PolicySummary)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySummary) DeepCopyInto(out *PolicySummary) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.AWSManaged != nil {
		in, out := &in.AWSManaged, &out.AWSManaged
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySummary.
func (in *PolicySummary) DeepCopy() *PolicySummary {
	if in == nil {
		return nil
	}
	out := new(PolicySummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyTypeSummary) DeepCopyInto(out *PolicyTypeSummary) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyTypeSummary.
func (in *PolicyTypeSummary) DeepCopy() *PolicyTypeSummary {
	if in == nil {
		return nil
	}
	out := new(PolicyTypeSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy_SDK) DeepCopyInto(out *Policy_SDK) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.PolicySummary != nil {
		in, out := &in.PolicySummary, &out.PolicySummary
		*out = new(PolicySummary)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy_SDK.
func (in *Policy_SDK) DeepCopy() *Policy_SDK {
	if in == nil {
		return nil
	}
	out := new(Policy_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Root) DeepCopyInto(out *Root) {
	*out = *in
	if in.PolicyTypes != nil {
		in, out := &in.PolicyTypes, &out.PolicyTypes
		*out = make([]*PolicyTypeSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyTypeSummary)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Root.
func (in *Root) DeepCopy() *Root {
	if in == nil {
		return nil
	}
	out := new(Root)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}