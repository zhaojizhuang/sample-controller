/*
Copyright 2021 The Knative Authors

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
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	eventingv1beta1 "knative.dev/eventing/pkg/apis/sources/v1beta1"
)

// Function is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
//
// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Function struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the Function (from the client).
	// +optional
	Spec FunctionSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the Function (from the controller).
	// +optional
	Status FunctionStatus `json:"status,omitempty"`
}

var (
	// Check that AddressableService can be validated and defaulted.
	_ apis.Validatable   = (*Function)(nil)
	_ apis.Defaultable   = (*Function)(nil)
	_ kmeta.OwnerRefable = (*Function)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*Function)(nil)
)

// FunctionSpec holds the desired state of the Function (from the client).
type FunctionSpec struct {

	// Instances specifies multi revisions you can release at one time.
	Instances []FunctionInstanceSpec `json:"instances"`

	// HttpTrigger specifies multi http ingress for your function.
	// +optional
	HttpTriggers []HttpTriggerHost `json:"http_triggers"`

	// PingSource specifies pingsource for your function
	// +optional
	PingSources []eventingv1beta1.PingSource `json:"ping_sources"`
	// Traffic specifies how to distribute traffic over a collection of
	// revisions and configurations.
	// +optional
	Traffic []TrafficTarget `json:"traffic,omitempty"`

	// +optional
	Ingresses []v1beta1.Ingress `json:"ingresses"`

	// +optional
	Configmaps []v1.ConfigMap `json:"configmaps"`

	// +optional
	Services []v1.Service `json:"services"`

	// +optional
	Secrets []v1.Secret `json:"secrets"`
}

// TrafficTarget holds a single entry of the routing table for a Route.
type TrafficTarget struct {

	// Tag is optionally used to expose a dedicated url for referencing
	// this target exclusively.
	// +optional
	VersionTag string `json:"version_tag"`

	// LatestRevision may be optionally provided to indicate that the latest
	// ready Revision of the Configuration should be used for this traffic
	// target.  When provided LatestRevision must be true if RevisionName is
	// empty; it must be false when RevisionName is non-empty.
	// +optional
	LatestRevision *bool `json:"latest_revision"`

	// Percent indicates that percentage based routing should be used and
	// the value indicates the percent of traffic that is be routed to this
	// Revision or Configuration. `0` (zero) mean no traffic, `100` means all
	// traffic.
	// When percentage based routing is being used the follow rules apply:
	// - the sum of all percent values must equal 100
	// - when not specified, the implied value for `percent` is zero for
	//   that particular Revision or Configuration
	// +optional
	Percent *int64 `json:"percent,omitempty"`
}

type FunctionInstanceSpec struct {
	servingv1.Service `json:",inline"`

	// VersionTag specifies the version of your function revision, which will generate specific revision name of knative,
	// such as aaaaa-version-bbbbb (aaaaa is functionName, and bbbbb is versionTag), and revision name will replace "." to "-"
	VersionTag string `json:"version_tag"`
}

type HttpTriggerHost struct {
	Host      string `json:"host"`
	ServePath string `json:"serve_path"`

	// tls  after base64 encode
	TlsCert string `json:"tls_cert"`
	TlsKey  string `json:"tls_key"`

	// rewrite config
	RewritePath string `json:"rewrite_path"`
	RewriteUrl  string `json:"rewrite_url"`
}

const (
	// FunctionConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	FunctionConditionReady = apis.ConditionReady
)

// FunctionStatus communicates the observed state of the Function (from the controller).
type FunctionStatus struct {
	duckv1.Status `json:",inline"`

	// KsvcStatus specific status of Ksvc
	KsvcStatus servingv1.ServiceStatus `json:"ksvc_status"`

	// specific to ResourceStatus.
	ResourceStatus []ResourceStatus `json:"resource_status"`
}

type ResourceStatus struct {
	Name  string `json:"name"`
	Kind  string `json:"kind"`
	Group string `json:"group"`
	Ready bool   `json:"ready"`
	// If Failed, Message tell you why.
	Message string `json:"message"`
}

// FunctionList is a list of AddressableService resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Function `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (f *Function) GetStatus() *duckv1.Status {
	return &f.Status.Status
}

func (f *Function) GetResourceStatus(name, kind string) ResourceStatus {

	for _, resource := range f.Status.ResourceStatus {
		if resource.Name == name && resource.Kind == kind {
			return resource
		}
	}
	return ResourceStatus{
		Name:    name,
		Kind:    kind,
		Group:   "",
		Ready:   false,
		Message: "Unkown Status",
	}
}
