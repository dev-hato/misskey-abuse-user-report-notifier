//go:build tools
// +build tools

package tools

import (
	_ "github.com/air-verse/air"         //nolint:typecheck
	_ "golang.org/x/tools/cmd/goimports" //nolint:typecheck
)
