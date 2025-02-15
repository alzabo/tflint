// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsWorklinkFleetInvalidNameRule checks the pattern is valid
type AwsWorklinkFleetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWorklinkFleetInvalidNameRule returns new rule with default attributes
func NewAwsWorklinkFleetInvalidNameRule() *AwsWorklinkFleetInvalidNameRule {
	return &AwsWorklinkFleetInvalidNameRule{
		resourceType:  "aws_worklink_fleet",
		attributeName: "name",
		max:           48,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-z0-9](?:[a-z0-9\-]{0,46}[a-z0-9])?$`),
	}
}

// Name returns the rule name
func (r *AwsWorklinkFleetInvalidNameRule) Name() string {
	return "aws_worklink_fleet_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkFleetInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsWorklinkFleetInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkFleetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkFleetInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 48 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-z0-9](?:[a-z0-9\-]{0,46}[a-z0-9])?$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
