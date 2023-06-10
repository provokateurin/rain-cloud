package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
)

//go:embed templates/skaffold.yaml.tmpl
var templateSkaffold string

//go:embed templates/ingress.yaml.tmpl
var templateIngress string

//go:embed templates/manifest.yaml.tmpl
var templateManifest string

//go:embed templates/include.go.tmpl
var templateInclude string

//go:embed templates/app.go.tmpl
var templateApp string

type App struct {
	ID              string
	KubernetesID    string
	InterfacePrefix string
	Spec            string
	PathPrefixes    []string
}

func main() {
	apps := []App{
		{
			ID:              "provisioning_api",
			KubernetesID:    "provisioning-api",
			InterfacePrefix: "ProvisioningApi",
			Spec:            "nextcloud/server/apps/provisioning_api/openapi.json",
			PathPrefixes: append(
				pathPrefixesForApp("apps/provisioning_api", false, true),
				"/ocs/v2.php/cloud/apps",
				"/ocs/v2.php/cloud/groups",
				"/ocs/v2.php/cloud/user",
				"/ocs/v2.php/cloud/users",
			),
		},
		{
			ID:              "theming",
			KubernetesID:    "theming",
			InterfacePrefix: "Theming",
			Spec:            "nextcloud/server/apps/theming/openapi.json",
			PathPrefixes:    pathPrefixesForApp("apps/theming", true, true),
		},
		{
			ID:              "user_status",
			KubernetesID:    "user-status",
			InterfacePrefix: "UserStatus",
			Spec:            "nextcloud/server/apps/user_status/openapi.json",
			PathPrefixes:    pathPrefixesForApp("apps/user_status", false, true),
		},
		// Core has to be last
		{
			ID:              "core",
			KubernetesID:    "core",
			InterfacePrefix: "Core",
			Spec:            "nextcloud/server/core/openapi.json",
			PathPrefixes: append(concatMultipleSlices([][]string{
				pathPrefixesForApp("avatar", true, false),
				pathPrefixesForApp("core", true, false),
				pathPrefixesForApp("login", true, false),
				pathPrefixesForApp("cloud", false, true),
				pathPrefixesForApp("collaboration", false, true),
				pathPrefixesForApp("core", false, true),
				pathPrefixesForApp("hovercard", false, true),
				pathPrefixesForApp("profile", false, true),
				pathPrefixesForApp("references", false, true),
				pathPrefixesForApp("search", false, true),
				pathPrefixesForApp("translation", false, true),
			}),
				"/status.php",
			),
		},
	}

	err := writeTemplate("skaffold.yaml", templateSkaffold, apps)
	if err != nil {
		panic(err)
	}

	err = writeTemplate("helm/rain-cloud/templates/ingress.yaml", templateIngress, apps)
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		err = writeTemplate(fmt.Sprintf("helm/rain-cloud/templates/app_%s.yaml", app.ID), templateManifest, app)
		if err != nil {
			panic(err)
		}

		err = writeTemplate(fmt.Sprintf("pkg/include/%s.gen.go", app.ID), templateInclude, app)
		if err != nil {
			panic(err)
		}

		err = writeTemplate(fmt.Sprintf("pkg/api/%s/app.gen.go", app.ID), templateApp, app)
		if err != nil {
			panic(err)
		}

		spec, err := openapi3.NewLoader().LoadFromFile(app.Spec)
		if err != nil {
			panic(err)
		}

		paths := openapi3.Paths{}
		for path, item := range spec.Paths {
			path = strings.Replace(path, "/index.php", "", 1)
			paths[path] = item
		}
		spec.Paths = paths

		code, err := codegen.Generate(spec, codegen.Configuration{
			PackageName: fmt.Sprintf("%sapi", app.ID),
			Generate: codegen.GenerateOptions{
				Models:    true,
				ChiServer: true,
				Strict:    true,
			},
			OutputOptions: codegen.OutputOptions{
				SkipPrune: true,
			},
		})
		if err != nil {
			panic(err)
		}
		//nolint:gosec
		err = os.WriteFile(fmt.Sprintf("pkg/api/%s/openapi.gen.go", app.ID), []byte(code), 0o644)
		if err != nil {
			panic(err)
		}
	}
}

func writeTemplate(path, t string, data any) error {
	tmpl, err := template.New("manifest").Parse(t)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	return nil
}

func pathPrefixesForApp(name string, index, ocs bool) []string {
	var prefixes []string

	if index {
		prefixes = append(
			prefixes,
			fmt.Sprintf("/%s", name),
			fmt.Sprintf("/index.php/%s", name),
		)
	}

	if ocs {
		prefixes = append(
			prefixes,
			fmt.Sprintf("/ocs/v2.php/%s", name),
		)
	}

	return prefixes
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
