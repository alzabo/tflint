// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventTargetInvalidRoleArnRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchEventTargetInvalidRoleArnRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidRoleArnRule() *AwsCloudwatchEventTargetInvalidRoleArnRule {
	return &AwsCloudwatchEventTargetInvalidRoleArnRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "role_arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidRoleArnRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidRoleArnRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidRoleArnRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
