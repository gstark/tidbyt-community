// Package usyieldcurve provides details for the US Yield Curve applet.
package usyieldcurve

import (
	_ "embed"

	"tidbyt.dev/community/apps"
	"tidbyt.dev/community/apps/manifest"
)

//go:embed us_yield_curve.star
var source []byte

func init() {
	apps.Manifests = append(apps.Manifests, New())
}

// New creates a new instance of the US Yield Curve applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "us-yield-curve",
		Name:        "US Yield Curve",
		Author:      "Rob Kimball",
		Summary:     "Plots the US yield curve",
		Desc:        "Track changes to the yield curve over different US Treasury maturities.",
		FileName:    "us_yield_curve.star",
		PackageName: "usyieldcurve",
		Source:      source,
	}
}
