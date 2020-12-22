package internal

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/vaibhavjainwiz/kogito-operator/community-kogito-operator/api/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FetchKogitoRuntimeService provide KogitoRuntime instance for given name and namespace
func FetchKogitoRuntimeService(client client.Client, name string, namespace string, log logr.Logger) (*v1beta1.KogitoRuntime, error) {
	log.Info("going to fetch deployed kogito runtime service", "name", name, "namespace", namespace)
	instance := &v1beta1.KogitoRuntime{}
	if resultErr := client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, instance); resultErr != nil {
		log.Error(resultErr, "Error occurs while fetching deployed kogito runtime service instance", "name", name)
		return nil, resultErr
	}
	log.Info("Successfully fetch deployed kogito runtime reference", "name", name)
	return instance, nil
}
