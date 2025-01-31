/*
Copyright 2020 The Knative Authors

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
)

var functionCondSet = apis.NewLivingConditionSet()

// GetGroupVersionKind implements kmeta.OwnerRefable
func (*Function) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Function")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (d *Function) GetConditionSet() apis.ConditionSet {
	return functionCondSet
}

// InitializeConditions sets the initial values to the conditions.
func (ds *FunctionStatus) InitializeConditions() {
	functionCondSet.Manage(ds).InitializeConditions()
}

// MarkPodsNotReady makes the Function be not ready.
func (ds *FunctionStatus) MarkFunctionNotReady(reason string) {
	functionCondSet.Manage(ds).MarkFalse(
		FunctionConditionReady,
		"FunctionNotReady",
		"Function are not ready yet,%s", reason)
}

// MarkFunctionReady makes the Function be ready.
func (ds *FunctionStatus) MarkFunctionReady() {
	functionCondSet.Manage(ds).MarkTrue(FunctionConditionReady)
}
