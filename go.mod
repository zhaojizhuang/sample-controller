module knative.dev/super-controller

go 1.15

require (
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	k8s.io/code-generator v0.19.7
	k8s.io/kube-openapi v0.0.0-20210113233702-8566a335510f
	knative.dev/eventing v0.21.2
	knative.dev/hack v0.0.0-20210317214554-58edbdc42966
	knative.dev/hack/schema v0.0.0-20210309141825-9b73a256fd9a
	knative.dev/pkg v0.0.0-20210318052054-dfeeb1817679
	knative.dev/serving v0.21.0
)
