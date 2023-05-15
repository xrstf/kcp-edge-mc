//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
)

// SinglePlacementSliceClusterLister can list SinglePlacementSlices across all workspaces, or scope down to a SinglePlacementSliceLister for one workspace.
// All objects returned here must be treated as read-only.
type SinglePlacementSliceClusterLister interface {
	// List lists all SinglePlacementSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SinglePlacementSlice, err error)
	// Cluster returns a lister that can list and get SinglePlacementSlices in one workspace.
	Cluster(clusterName logicalcluster.Name) SinglePlacementSliceLister
	SinglePlacementSliceClusterListerExpansion
}

type singlePlacementSliceClusterLister struct {
	indexer cache.Indexer
}

// NewSinglePlacementSliceClusterLister returns a new SinglePlacementSliceClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewSinglePlacementSliceClusterLister(indexer cache.Indexer) *singlePlacementSliceClusterLister {
	return &singlePlacementSliceClusterLister{indexer: indexer}
}

// List lists all SinglePlacementSlices in the indexer across all workspaces.
func (s *singlePlacementSliceClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.SinglePlacementSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.SinglePlacementSlice))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get SinglePlacementSlices.
func (s *singlePlacementSliceClusterLister) Cluster(clusterName logicalcluster.Name) SinglePlacementSliceLister {
	return &singlePlacementSliceLister{indexer: s.indexer, clusterName: clusterName}
}

// SinglePlacementSliceLister can list all SinglePlacementSlices, or get one in particular.
// All objects returned here must be treated as read-only.
type SinglePlacementSliceLister interface {
	// List lists all SinglePlacementSlices in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SinglePlacementSlice, err error)
	// Get retrieves the SinglePlacementSlice from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.SinglePlacementSlice, error)
	SinglePlacementSliceListerExpansion
}

// singlePlacementSliceLister can list all SinglePlacementSlices inside a workspace.
type singlePlacementSliceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all SinglePlacementSlices in the indexer for a workspace.
func (s *singlePlacementSliceLister) List(selector labels.Selector) (ret []*edgev1alpha1.SinglePlacementSlice, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SinglePlacementSlice))
	})
	return ret, err
}

// Get retrieves the SinglePlacementSlice from the indexer for a given workspace and name.
func (s *singlePlacementSliceLister) Get(name string) (*edgev1alpha1.SinglePlacementSlice, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SinglePlacementSlice"), name)
	}
	return obj.(*edgev1alpha1.SinglePlacementSlice), nil
}

// NewSinglePlacementSliceLister returns a new SinglePlacementSliceLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewSinglePlacementSliceLister(indexer cache.Indexer) *singlePlacementSliceScopedLister {
	return &singlePlacementSliceScopedLister{indexer: indexer}
}

// singlePlacementSliceScopedLister can list all SinglePlacementSlices inside a workspace.
type singlePlacementSliceScopedLister struct {
	indexer cache.Indexer
}

// List lists all SinglePlacementSlices in the indexer for a workspace.
func (s *singlePlacementSliceScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.SinglePlacementSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SinglePlacementSlice))
	})
	return ret, err
}

// Get retrieves the SinglePlacementSlice from the indexer for a given workspace and name.
func (s *singlePlacementSliceScopedLister) Get(name string) (*edgev1alpha1.SinglePlacementSlice, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SinglePlacementSlice"), name)
	}
	return obj.(*edgev1alpha1.SinglePlacementSlice), nil
}
