package mysql

import (
	"slices"
	"testing"

	spicestarter "github.com/spice-framework/spice/annotation/sdk/starter"
)

func TestManifestDeclaresStandaloneReviewedStarter(t *testing.T) {
	t.Parallel()

	manifest := Manifest()
	spec := manifest.Spec()
	if spec.Schema != spicestarter.Schema ||
		spec.ID != "github.com/spice-framework/starter-mysql" ||
		spec.Module != "github.com/spice-framework/starter-mysql" ||
		spec.Version != "0.1.0-dev" ||
		spec.SpiceAPI != spicestarter.APIVersion ||
		spec.MinimumGo != "1.26" ||
		spec.License != "Apache-2.0" ||
		spec.Review != "docs/dependency-review.md" ||
		spec.Activation.Mode != spicestarter.ActivationExplicitConstructor {
		t.Fatalf("Manifest() = %#v", spec)
	}
	wantCapabilities := []string{"data.mysql", "data.sql"}
	if !slices.Equal(spec.Capabilities, wantCapabilities) {
		t.Fatalf("capabilities = %v, want %v", spec.Capabilities, wantCapabilities)
	}
	if len(spec.Activation.EntryPoints) != 1 ||
		spec.Activation.EntryPoints[0] != (spicestarter.EntryPoint{
			Package: "github.com/spice-framework/starter-mysql",
			Symbol:  "Open",
		}) || len(spec.Dependencies) != 1 ||
		spec.Dependencies[0] != (spicestarter.Dependency{
			Module:  "github.com/go-sql-driver/mysql",
			Version: "v1.10.0",
			License: "MPL-2.0",
		}) {
		t.Fatalf("entrypoints=%#v dependencies=%#v", spec.Activation.EntryPoints, spec.Dependencies)
	}
	if err := manifest.Compatible(spicestarter.APIVersion, "go1.26.6"); err != nil {
		t.Fatalf("Compatible() error = %v", err)
	}
	content, err := manifest.JSON()
	if err != nil {
		t.Fatalf("JSON() error = %v", err)
	}
	parsed, err := spicestarter.Parse(content)
	if err != nil {
		t.Fatalf("Parse(JSON()) error = %v", err)
	}
	if parsed.Spec().ID != spec.ID {
		t.Fatalf("parsed ID = %q, want %q", parsed.Spec().ID, spec.ID)
	}
}
