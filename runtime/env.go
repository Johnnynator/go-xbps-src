package runtime

import (
	"mvdan.cc/sh/v3/expand"
)

// Env returns the xbps-src environment
func (r *Runtime) Env(arch, cross string) Environ {
	var m, t expand.Variable
	m = expand.Variable{
		Exported: true,
		ReadOnly: true,
		Kind: expand.String,
		Str:    arch,
	}
	t = m

	if cross != "" {
		t = expand.Variable{
			Exported: true,
			ReadOnly: true,
			Kind: expand.String,
			Str:    cross,
		}
	}

	return Environ{
		"XBPS_UHELPER_CMD": expand.Variable{
			Exported: true,
			ReadOnly: true,
			Kind: expand.String,
			Str:    "xbps-uhelper",
		},
		"XBPS_MACHINE":        m,
		"XBPS_TARGET_MACHINE": t,
	}
}
