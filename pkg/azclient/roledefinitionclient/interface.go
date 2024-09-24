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

// +azure:enableclientgen:=true
package roledefinitionclient

import (
	"context"

	armauthorization "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// +azure:client:verbs=,resource=RoleDefinition,packageName=github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2,packageAlias=armauthorization,clientName=RoleDefinitionsClient,expand=false,outOfSubscriptionScope=true
type Interface interface {
	List(ctx context.Context, scope string, option *armauthorization.RoleDefinitionsClientListOptions) (result []*armauthorization.RoleDefinition, rerr error)
}
