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

// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	externalversions "knative.dev/super-controller/pkg/client/informers/externalversions"
	fake "knative.dev/super-controller/pkg/client/injection/client/fake"
	factory "knative.dev/super-controller/pkg/client/injection/informers/factory"
)

var Get = factory.Get

func init() {
	injection.Fake.RegisterInformerFactory(withInformerFactory)
}

func withInformerFactory(ctx context.Context) context.Context {
	c := fake.Get(ctx)
	opts := make([]externalversions.SharedInformerOption, 0, 1)
	if injection.HasNamespaceScope(ctx) {
		opts = append(opts, externalversions.WithNamespace(injection.GetNamespaceScope(ctx)))
	}
	return context.WithValue(ctx, factory.Key{},
		externalversions.NewSharedInformerFactoryWithOptions(c, controller.GetResyncPeriod(ctx), opts...))
}
