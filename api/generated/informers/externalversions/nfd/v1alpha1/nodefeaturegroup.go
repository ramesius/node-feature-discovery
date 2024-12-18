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

// NodeFeatureGroupInformer provides access to a shared informer and lister for
// NodeFeatureGroups.
type NodeFeatureGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() nfdv1alpha1.NodeFeatureGroupLister
}

type nodeFeatureGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNodeFeatureGroupInformer constructs a new informer for NodeFeatureGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeFeatureGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeFeatureGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNodeFeatureGroupInformer constructs a new informer for NodeFeatureGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeFeatureGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NfdV1alpha1().NodeFeatureGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NfdV1alpha1().NodeFeatureGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&apinfdv1alpha1.NodeFeatureGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeFeatureGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeFeatureGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeFeatureGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apinfdv1alpha1.NodeFeatureGroup{}, f.defaultInformer)
}

func (f *nodeFeatureGroupInformer) Lister() nfdv1alpha1.NodeFeatureGroupLister {
	return nfdv1alpha1.NewNodeFeatureGroupLister(f.Informer().GetIndexer())
}
