/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "sigs.k8s.io/node-feature-discovery/api/generated/clientset/versioned"
	internalinterfaces "sigs.k8s.io/node-feature-discovery/api/generated/informers/externalversions/internalinterfaces"
	nfdv1alpha1 "sigs.k8s.io/node-feature-discovery/api/generated/listers/nfd/v1alpha1"
	apinfdv1alpha1 "sigs.k8s.io/node-feature-discovery/api/nfd/v1alpha1"
)

// NodeFeatureRuleInformer provides access to a shared informer and lister for
// NodeFeatureRules.
type NodeFeatureRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() nfdv1alpha1.NodeFeatureRuleLister
}

type nodeFeatureRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeFeatureRuleInformer constructs a new informer for NodeFeatureRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeFeatureRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeFeatureRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeFeatureRuleInformer constructs a new informer for NodeFeatureRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeFeatureRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NfdV1alpha1().NodeFeatureRules().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NfdV1alpha1().NodeFeatureRules().Watch(context.TODO(), options)
			},
		},
		&apinfdv1alpha1.NodeFeatureRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeFeatureRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeFeatureRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeFeatureRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apinfdv1alpha1.NodeFeatureRule{}, f.defaultInformer)
}

func (f *nodeFeatureRuleInformer) Lister() nfdv1alpha1.NodeFeatureRuleLister {
	return nfdv1alpha1.NewNodeFeatureRuleLister(f.Informer().GetIndexer())
}
