// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsLambdaPermissionInvalidPrincipalRule checks the pattern is valid
type AwsLambdaPermissionInvalidPrincipalRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidPrincipalRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidPrincipalRule() *AwsLambdaPermissionInvalidPrincipalRule {
	return &AwsLambdaPermissionInvalidPrincipalRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "principal",
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidPrincipalRule) Name() string {
	return "aws_lambda_permission_invalid_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidPrincipalRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsLambdaPermissionInvalidPrincipalRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidPrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidPrincipalRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`principal does not match valid pattern ^.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
