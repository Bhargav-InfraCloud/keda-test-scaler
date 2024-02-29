/*
Copyright 2024.

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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	scalerv1alpha1 "github.com/Bhargav-InfraCloud/keda-test-scaler/api/v1alpha1"
)

// TestScalerReconciler reconciles a TestScaler object
type TestScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=scaler.test.keda.sh,resources=testscalers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=scaler.test.keda.sh,resources=testscalers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=scaler.test.keda.sh,resources=testscalers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *TestScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var testScaler scalerv1alpha1.TestScaler
	if err := r.Get(ctx, req.NamespacedName, &testScaler); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Println(*testScaler.Spec.Replicas)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&scalerv1alpha1.TestScaler{}).
		Complete(r)
}
