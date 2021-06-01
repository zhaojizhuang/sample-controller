module knative.dev/sample-controller

go 1.15

require (
	go.uber.org/zap v1.17.0
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	k8s.io/code-generator v0.20.7
	k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd
	knative.dev/hack v0.0.0-20210614141220-66ab1a098940
	knative.dev/hack/schema v0.0.0-20210309141825-9b73a256fd9a
	knative.dev/pkg v0.0.0-20210318052054-dfeeb1817679
)

replace knative.dev/pkg => github.com/zhaojizhuang/pkg v0.0.0-20210621104351-3c56792f29bd
