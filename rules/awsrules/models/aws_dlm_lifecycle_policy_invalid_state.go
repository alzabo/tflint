// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsDlmLifecyclePolicyInvalidStateRule checks the pattern is valid
type AwsDlmLifecyclePolicyInvalidStateRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDlmLifecyclePolicyInvalidStateRule returns new rule with default attributes
func NewAwsDlmLifecyclePolicyInvalidStateRule() *AwsDlmLifecyclePolicyInvalidStateRule {
	return &AwsDlmLifecyclePolicyInvalidStateRule{
		resourceType:  "aws_dlm_lifecycle_policy",
		attributeName: "state",
		enum: []string{
			"ENABLED",
			"DISABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Name() string {
	return "aws_dlm_lifecycle_policy_invalid_state"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDlmLifecyclePolicyInvalidStateRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`state is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
