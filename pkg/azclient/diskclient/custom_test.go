// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package diskclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	additionalTestCases = func() {
		When("the disk is patched with a new size", func() {
			It("", func(ctx context.Context) {
				_, err := realClient.Patch(ctx, resourceGroupName, resourceName, armcompute.DiskUpdate{
					Properties: &armcompute.DiskUpdateProperties{
						DiskSizeGB: to.Ptr[int32](300),
					},
				})
				Expect(err).To(HaveOccurred())
			})
		})
	}

	beforeAllFunc = func(ctx context.Context) {
		newResource = &armcompute.Disk{
			Location: to.Ptr(location),
			Properties: &armcompute.DiskProperties{
				CreationData: &armcompute.CreationData{
					CreateOption: to.Ptr(armcompute.DiskCreateOptionEmpty),
				},
				DiskSizeGB: to.Ptr[int32](200),
			},
		}
	}
	afterAllFunc = func(ctx context.Context) {
	}
}
