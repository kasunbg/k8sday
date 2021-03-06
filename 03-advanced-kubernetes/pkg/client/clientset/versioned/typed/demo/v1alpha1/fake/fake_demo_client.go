// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/isurulucky/k8sday/03-advanced-kubernetes/pkg/client/clientset/versioned/typed/demo/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDemoV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDemoV1alpha1) Hellos(namespace string) v1alpha1.HelloInterface {
	return &FakeHellos{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDemoV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
