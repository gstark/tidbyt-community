// Package nightscout provides details for the Nightscout applet.
package nightscout

import (
	_ "embed"

	"tidbyt.dev/community/apps"
	"tidbyt.dev/community/apps/manifest"
)

//go:embed nightscout.star
var source []byte

func init() {
	apps.Manifests = append(apps.Manifests, New())
}

// New creates a new instance of the Nightscout applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "nightscout",
		Name:        "Nightscout",
		Author:      "Jeremy Tavener",
		Summary:     "Shows Nightscout CGM Data",
		Desc:        "Displays Continuous Glucose Monitoring (CGM) data from the Nightscout Open Source project (https://nightscout.github.io/).",
		FileName:    "nightscout.star",
		PackageName: "nightscout",
		Source:      source,
	}
}
