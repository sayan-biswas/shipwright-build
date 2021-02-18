// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package clusterbuildstrategy

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/build/pkg/config"
	"github.com/shipwright-io/build/pkg/ctxlog"
)

// Add creates a new ClusterBuildStrategy Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(ctx context.Context, c *config.Config, mgr manager.Manager) error {
	ctx = ctxlog.NewContext(ctx, "clusterbuildstrategy-controller")
	return add(ctx, mgr, NewReconciler(ctx, c, mgr))
}

// NewReconciler returns a new reconcile.Reconciler
func NewReconciler(ctx context.Context, c *config.Config, mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileClusterBuildStrategy{
		ctx:    ctx,
		config: c,
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(ctx context.Context, mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("clusterbuildstrategy-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ClusterBuildStrategy
	err = c.Watch(&source.Kind{Type: &buildv1alpha1.ClusterBuildStrategy{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileClusterBuildStrategy implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileClusterBuildStrategy{}

// ReconcileClusterBuildStrategy reconciles a ClusterBuildStrategy object
type ReconcileClusterBuildStrategy struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	ctx    context.Context
	config *config.Config
	client client.Client
	scheme *runtime.Scheme
}