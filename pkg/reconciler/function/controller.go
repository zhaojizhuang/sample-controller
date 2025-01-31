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


package function

import (
	"context"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"time"

	kubeclient "knative.dev/pkg/client/injection/kube/client"
	podinformer "knative.dev/pkg/client/injection/kube/informers/core/v1/pod"
	functioninformer "knative.dev/super-controller/pkg/client/injection/informers/function/v1beta1/function"
	functionreconciler "knative.dev/super-controller/pkg/client/injection/reconciler/function/v1beta1/function"
)

// NewController creates a Reconciler and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	// Obtain an informer to both the main and child resources. These will be started by
	// the injection framework automatically. They'll keep a cached representation of the
	// cluster's state of the respective resource at all times.
	functionInformer := functioninformer.Get(ctx)
	podInformer := podinformer.Get(ctx)

	r := &Reconciler{
		// The client will be needed to create/delete Pods via the API.
		kubeclient: kubeclient.Get(ctx),
		// A lister allows read-only access to the informer's cache, allowing us to cheaply
		// read pod data.
		podLister: podInformer.Lister(),
	}
	impl := functionreconciler.NewImpl(ctx, r)

	logger.Info("Setting up event handlers.")

	// Listen for events on the main resource and enqueue themselves.
	functionInformer.Informer().AddEventHandlerWithResyncPeriod(controller.HandleAll(impl.Enqueue),10*time.Second)

	// Listen for events on the child resources and enqueue the owner of them.
	podInformer.Informer().AddEventHandler(controller.HandleAll(impl.EnqueueControllerOf))

	return impl
}
