/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validatingadmissionpolicy

import (
	"net/http"
	"time"

	"k8s.io/api/admissionregistration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PolicyDecisionAction string

const (
	ActionAdmit PolicyDecisionAction = "admit"
	ActionDeny  PolicyDecisionAction = "deny"
)

type PolicyDecisionEvaluation string

const (
	EvalAdmit PolicyDecisionEvaluation = "admit"
	EvalError PolicyDecisionEvaluation = "error"
	EvalDeny  PolicyDecisionEvaluation = "deny"
)

type PolicyDecision struct {
	Action     PolicyDecisionAction
	Evaluation PolicyDecisionEvaluation
	Message    string
	Reason     metav1.StatusReason
	Elapsed    time.Duration
}

type policyDecisionWithMetadata struct {
	PolicyDecision
	Definition *v1alpha1.ValidatingAdmissionPolicy
	Binding    *v1alpha1.ValidatingAdmissionPolicyBinding
}

func ReasonToCode(r metav1.StatusReason) int32 {
	switch r {
	case metav1.StatusReasonForbidden:
		return http.StatusForbidden
	case metav1.StatusReasonUnauthorized:
		return http.StatusUnauthorized
	case metav1.StatusReasonRequestEntityTooLarge:
		return http.StatusRequestEntityTooLarge
	case metav1.StatusReasonInvalid:
		return http.StatusUnprocessableEntity
	default:
		// It should not reach here since we only allow above reason to be set from API level
		return http.StatusUnprocessableEntity
	}
}
