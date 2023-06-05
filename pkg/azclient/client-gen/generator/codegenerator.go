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

// Package generator
package generator

import (
	"go/ast"
	"os/exec"

	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	FuncCreateOrUpdate = "CreateOrUpdate"
	FuncGet            = "Get"
	FuncDelete         = "Delete"
	FuncListByRG       = "ListByRG"
	FuncList           = "List"
)

// clientGenMarker s a marker for generating client code for azure services.
var (
	enableClientGenMarker = markers.Must(markers.MakeDefinition("azure:enableclientgen", markers.DescribesPackage, markers.RawArguments(nil)))
	clientGenMarker       = markers.Must(markers.MakeDefinition("azure:client", markers.DescribesType, ClientGenConfig{}))
)

// +controllertools:marker:generateHelp
// Generator generates client code for azure services.
type Generator struct {
	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	return markers.RegisterAll(into, clientGenMarker, enableClientGenMarker)
}

func (g Generator) Generate(ctx *genall.GenerationContext) error {
	if err := exec.Command("go", "get", "github.com/golang/mock/mockgen/model").Run(); err != nil {
		return err
	}
	var headerText string

	if g.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(g.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}

	for _, root := range ctx.Roots {
		pkgMakers, err := markers.PackageMarkers(ctx.Collector, root)
		if err != nil {
			root.AddError(err)
			return err
		}
		if _, markedForGeneration := pkgMakers[enableClientGenMarker.Name]; !markedForGeneration {
			continue
		}

		//check for syntax error
		ctx.Checker.Check(root)

		//visit each type
		root.NeedTypesInfo()

		if err := generateClient(ctx, root, headerText); err != nil {
			root.AddError(err)
			return err
		}
		if err := generateMock(ctx, root, headerText); err != nil {
			root.AddError(err)
			return err
		}
		if err := generateTest(ctx, root, headerText); err != nil {
			root.AddError(err)
			return err
		}
	}
	//nolint:gosec // G204 ignore this!
	if err := exec.Command("goimports", "-local", "sigs.k8s.io/cloud-provider-azure/pkg/azclient", "-w", ".").Run(); err != nil {
		return err
	}
	//nolint:gosec // G204 ignore this!
	return exec.Command("go", "test", "./...").Run()
}

func (Generator) CheckFilter() loader.NodeFilter {
	return func(node ast.Node) bool {
		// ignore structs
		_, isIface := node.(*ast.InterfaceType)
		return isIface
	}
}
