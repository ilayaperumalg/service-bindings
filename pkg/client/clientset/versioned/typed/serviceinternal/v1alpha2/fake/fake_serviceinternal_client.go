/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/serviceinternal/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeInternalV1alpha2 struct {
	*testing.Fake
}

func (c *FakeInternalV1alpha2) ServiceBindingProjections(namespace string) v1alpha2.ServiceBindingProjectionInterface {
	return &FakeServiceBindingProjections{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeInternalV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}