// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferSSHKeyInvalidServerIDRule checks the pattern is valid
type AwsTransferSSHKeyInvalidServerIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsTransferSSHKeyInvalidServerIDRule returns new rule with default attributes
func NewAwsTransferSSHKeyInvalidServerIDRule() *AwsTransferSSHKeyInvalidServerIDRule {
	return &AwsTransferSSHKeyInvalidServerIDRule{
		resourceType:  "aws_transfer_ssh_key",
		attributeName: "server_id",
		pattern:       regexp.MustCompile(`^s-([0-9a-f]{17})$`),
	}
}

// Name returns the rule name
func (r *AwsTransferSSHKeyInvalidServerIDRule) Name() string {
	return "aws_transfer_ssh_key_invalid_server_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferSSHKeyInvalidServerIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsTransferSSHKeyInvalidServerIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferSSHKeyInvalidServerIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferSSHKeyInvalidServerIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`server_id does not match valid pattern ^s-([0-9a-f]{17})$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
