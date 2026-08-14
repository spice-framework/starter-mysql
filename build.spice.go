//go:build spice_config

package spice

import "github.com/spice-framework/spice/project"

var Build = project.Build{
	Kind: project.StarterKind,
	Dependencies: project.Dependencies{
		project.Library("github.com/go-sql-driver/mysql", "v1.10.0"),
		project.BuildTool("github.com/spice-framework/toolchain", "v0.0.0-20260806133530-71211498297c"),
	},
}
