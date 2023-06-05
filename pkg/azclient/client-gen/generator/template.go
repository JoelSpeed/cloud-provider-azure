/*
Copyright 2023 The Kubernetes Authors.

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

package generator

import "html/template"

type ClientGenConfig struct {
	Verbs        []string
	Resource     string
	SubResource  string `marker:"subResource,optional"`
	PackageName  string
	PackageAlias string
	ClientName   string
	Expand       bool `marker:"expand,optional"`
}

var ClientTemplate = template.Must(template.New("object-scaffolding-client-struct").Parse(`
type Client struct{
	*{{.PackageAlias}}.{{.ClientName}}
}
`))

var ClientFactoryTemplate = template.Must(template.New("object-scaffolding-factory").Parse(`
func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}

	client, err := {{.PackageAlias}}.New{{.ClientName}}(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
`))

const CreateOrUpdateFuncTemplateRaw = `
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
// CreateOrUpdate creates or updates a {{$resource}}.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string,{{with .SubResource}}parentResourceName string, {{end}} resource {{.PackageAlias}}.{{$resource}}) (*{{.PackageAlias}}.{{$resource}}, error) {
	resp, err := utils.NewPollerWrapper(client.{{.ClientName}}.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName,{{with .SubResource}}parentResourceName,{{end}} resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.{{$resource}}, nil
	}
	return nil, nil
}
`

var CreateOrUpdateFuncTemplate = template.Must(template.New("object-scaffolding-create-func").Parse(CreateOrUpdateFuncTemplateRaw))

const ListByRGFuncTemplateRaw = `
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
// List gets a list of {{$resource}} in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string{{with .SubResource}}, parentResourceName string{{end}}) (result []*{{.PackageAlias}}.{{$resource}}, rerr error) {
	pager := client.{{.ClientName}}.NewListByResourceGroupPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
`

var ListByRGFuncTemplate = template.Must(template.New("object-scaffolding-list-func").Parse(ListByRGFuncTemplateRaw))

const ListFuncTemplateRaw = `
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
// List gets a list of {{$resource}} in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string{{with .SubResource}}, parentResourceName string{{end}}) (result []*{{.PackageAlias}}.{{$resource}}, rerr error) {
	pager := client.{{.ClientName}}.NewListPager(resourceGroupName,{{with .SubResource}}parentResourceName,{{end}} nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
`

var ListFuncTemplate = template.Must(template.New("object-scaffolding-list-func").Parse(ListFuncTemplateRaw))

const DeleteFuncTemplateRaw = `
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
// Delete deletes a {{$resource}} by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, {{with .SubResource}} parentResourceName string, {{end}}resourceName string) error {
	_, err := utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName,{{with .SubResource}}parentResourceName,{{end}} resourceName, nil)).WaitforPollerResp(ctx)
	return err
}
`

var DeleteFuncTemplate = template.Must(template.New("object-scaffolding-delete-func").Parse(DeleteFuncTemplateRaw))

const GetFuncTemplateRaw = `
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
// Get gets the {{$resource}}
func (client *Client) Get(ctx context.Context, resourceGroupName string, {{with .SubResource}}parentResourceName string,{{end}} resourceName string{{if .Expand}}, expand *string{{end}}) (result *{{.PackageAlias}}.{{$resource}}, rerr error) {
	var ops *{{.PackageAlias}}.{{.ClientName}}GetOptions
	{{if .Expand}}if expand != nil {
		ops = &{{.PackageAlias}}.{{.ClientName}}GetOptions{ Expand: expand }
	}{{end}}

	resp, err := client.{{.ClientName}}.Get(ctx, resourceGroupName,{{with .SubResource}}parentResourceName,{{end}} resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.{{$resource}}, nil
}
`

var GetFuncTemplate = template.Must(template.New("object-scaffolding-get-func").Parse(GetFuncTemplateRaw))

var ImportTemplate = template.Must(template.New("import").Parse(`{{.Alias}} "{{.Package}}"
`))

type ImportStatement struct {
	Alias   string
	Package string
}

var TestSuiteTemplate = template.Must(template.New("object-scaffolding-test-suite").Parse(
	`
{{ $resource := .Resource}}
{{ if (gt (len .SubResource) 0) }}
{{ $resource = .SubResource}}
{{ end }}
func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var resourceGroupName = "aks-cit"
var resourceName = "testdisk"
var subscriptionID string
var location = "eastus"
var resourceGroupClient *armresources.ResourceGroupsClient
var err error
var recorder *recording.Recorder
var realClient Interface

var _ = BeforeSuite(func(ctx context.Context) {
	recorder, err = recording.NewRecorder("testdata/{{$resource}}")
	Expect(err).ToNot(HaveOccurred())
	subscriptionID = recorder.SubscriptionID()
	Expect(err).NotTo(HaveOccurred())
	cred := recorder.TokenCredential()
	resourceGroupClient, err = armresources.NewResourceGroupsClient(subscriptionID, cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Retry: policy.RetryOptions{
				MaxRetryDelay: 1 * time.Millisecond,
				RetryDelay:    1 * time.Millisecond,
			},
			Transport: recorder.HTTPClient(),
		},
	})
	Expect(err).NotTo(HaveOccurred())
	realClient, err = New(subscriptionID, recorder.TokenCredential(), &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Retry: policy.RetryOptions{
				MaxRetryDelay: 1 * time.Millisecond,
				RetryDelay:    1 * time.Millisecond,
			},
			Transport: recorder.HTTPClient(),
		},
	})
	Expect(err).NotTo(HaveOccurred())
	_, err = resourceGroupClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		armresources.ResourceGroup{
			Location: to.Ptr(location),
		},
		nil)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func(ctx context.Context) {
	pollerResp, err := resourceGroupClient.BeginDelete(ctx, resourceGroupName, nil)
	Expect(err).NotTo(HaveOccurred())
	_, err = pollerResp.PollUntilDone(ctx, nil)
	Expect(err).NotTo(HaveOccurred())

	err = recorder.Stop()
	Expect(err).ToNot(HaveOccurred())
})
`))
