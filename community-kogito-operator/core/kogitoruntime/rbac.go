// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kogitoruntime

import (
	"context"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceAccountName = "kogito-service-viewer"
	roleName           = "kogito-service-viewer"
	roleBindingName    = "kogito-service-viewer"
	roleAPIGroup       = "rbac.authorization.k8s.io"
)

var serviceViewerRoleVerbs = []string{"list", "get", "watch", "update", "patch"}
var serviceViewerRoleAPIGroups = []string{""}
var serviceViewerRoleResources = []string{"services", "configmaps"}

type RBACService struct {
	Client client.Client
	Log    logr.Logger
}

func (r *RBACService) SetupRBAC(namespace string) (err error) {
	// create service viewer role
	if err = r.Client.Create(context.TODO(), getServiceViewerRole(namespace)); err != nil {
		r.Log.Error(err, "Fail to create role for service viewer")
		return
	}

	// create service viewer service account
	if err = r.Client.Create(context.TODO(), getServiceViewerServiceAccount(namespace)); err != nil {
		r.Log.Error(err, "Fail to create service account for service viewer")
		return
	}

	// create service viewer rolebinding
	if err = r.Client.Create(context.TODO(), getServiceViewerRoleBinding(namespace)); err != nil {
		r.Log.Error(err, "Fail to create role binding for service viewer")
		return
	}
	return
}

func getServiceViewerServiceAccount(namespace string) runtime.Object {
	return &v1.ServiceAccount{
		ObjectMeta: v12.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: namespace,
		},
	}
}

func getServiceViewerRole(namespace string) runtime.Object {
	return &rbac.Role{
		ObjectMeta: v12.ObjectMeta{
			Name:      roleName,
			Namespace: namespace,
		},
		Rules: []rbac.PolicyRule{
			{
				Verbs:     serviceViewerRoleVerbs,
				APIGroups: serviceViewerRoleAPIGroups,
				Resources: serviceViewerRoleResources,
			},
		},
	}
}
func getServiceViewerRoleBinding(namespace string) runtime.Object {
	return &rbac.RoleBinding{
		ObjectMeta: v12.ObjectMeta{
			Name:      roleBindingName,
			Namespace: namespace,
		},
		Subjects: []rbac.Subject{
			{
				Kind: "ServiceAccount",
				Name: serviceAccountName,
			},
		},
		RoleRef: rbac.RoleRef{
			APIGroup: roleAPIGroup,
			Name:     roleName,
			Kind:     "Role",
		},
	}
}
