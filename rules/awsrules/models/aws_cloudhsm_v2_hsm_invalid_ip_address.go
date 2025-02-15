// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudhsmV2HsmInvalidIPAddressRule checks the pattern is valid
type AwsCloudhsmV2HsmInvalidIPAddressRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2HsmInvalidIPAddressRule returns new rule with default attributes
func NewAwsCloudhsmV2HsmInvalidIPAddressRule() *AwsCloudhsmV2HsmInvalidIPAddressRule {
	return &AwsCloudhsmV2HsmInvalidIPAddressRule{
		resourceType:  "aws_cloudhsm_v2_hsm",
		attributeName: "ip_address",
		pattern:       regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Name() string {
	return "aws_cloudhsm_v2_hsm_invalid_ip_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2HsmInvalidIPAddressRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`ip_address does not match valid pattern ^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
