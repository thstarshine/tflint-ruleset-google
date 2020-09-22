// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeRouterPeerInvalidAdvertiseModeRule checks the pattern is valid
type GoogleComputeRouterPeerInvalidAdvertiseModeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeRouterPeerInvalidAdvertiseModeRule returns new rule with default attributes
func NewGoogleComputeRouterPeerInvalidAdvertiseModeRule() *GoogleComputeRouterPeerInvalidAdvertiseModeRule {
	return &GoogleComputeRouterPeerInvalidAdvertiseModeRule{
		resourceType:  "google_compute_router_peer",
		attributeName: "advertise_mode",
	}
}

// Name returns the rule name
func (r *GoogleComputeRouterPeerInvalidAdvertiseModeRule) Name() string {
	return "google_compute_router_peer_invalid_advertise_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeRouterPeerInvalidAdvertiseModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeRouterPeerInvalidAdvertiseModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeRouterPeerInvalidAdvertiseModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeRouterPeerInvalidAdvertiseModeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"DEFAULT", "CUSTOM", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
