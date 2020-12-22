/*


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
	"github.com/vaibhavjainwiz/kogito-operator/community-kogito-operator/core/kogitoruntime"
	"github.com/vaibhavjainwiz/kogito-operator/product-kogito-operator/api/v1beta1"
	"github.com/vaibhavjainwiz/kogito-operator/product-kogito-operator/internal"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KogitoRuntimeReconciler reconciles a KogitoRuntime object
type KogitoRuntimeReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=app.vajain.com,resources=kogitoruntimes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.vajain.com,resources=kogitoruntimes/status,verbs=get;update;patch

func (r *KogitoRuntimeReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
	_ = context.Background()
	log := r.Log.WithValues("kogitoruntime", req.NamespacedName)

	// Step 1: Fetch KogitoRuntime instance
	instance, err := internal.FetchKogitoRuntimeService(r.Client, req.Name, req.Namespace, log)
	if err != nil {
		return
	}
	if instance == nil {
		log.Info("Instance not found", "KogitoRuntime", req.Name, "Namespace", req.Namespace)
		return
	}

	// Step 2: Setup RBAC
	rbacService := kogitoruntime.RBACService{
		Log:    log,
		Client: r.Client,
	}
	err = rbacService.SetupRBAC(req.Namespace)
	if err != nil {
		return
	}

	return ctrl.Result{}, nil
}

func (r *KogitoRuntimeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.KogitoRuntime{}).
		Complete(r)
}
