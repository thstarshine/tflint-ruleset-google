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

// GoogleDatastoreIndexInvalidAncestorRule checks the pattern is valid
type GoogleDatastoreIndexInvalidAncestorRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleDatastoreIndexInvalidAncestorRule returns new rule with default attributes
func NewGoogleDatastoreIndexInvalidAncestorRule() *GoogleDatastoreIndexInvalidAncestorRule {
	return &GoogleDatastoreIndexInvalidAncestorRule{
		resourceType:  "google_datastore_index",
		attributeName: "ancestor",
	}
}

// Name returns the rule name
func (r *GoogleDatastoreIndexInvalidAncestorRule) Name() string {
	return "google_datastore_index_invalid_ancestor"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDatastoreIndexInvalidAncestorRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDatastoreIndexInvalidAncestorRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDatastoreIndexInvalidAncestorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDatastoreIndexInvalidAncestorRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"NONE", "ALL_ANCESTORS", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
