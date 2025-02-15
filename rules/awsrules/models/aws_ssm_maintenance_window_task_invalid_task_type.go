// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule checks the pattern is valid
type AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmMaintenanceWindowTaskInvalidTaskTypeRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTaskInvalidTaskTypeRule() *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule {
	return &AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule{
		resourceType:  "aws_ssm_maintenance_window_task",
		attributeName: "task_type",
		enum: []string{
			"RUN_COMMAND",
			"AUTOMATION",
			"STEP_FUNCTIONS",
			"LAMBDA",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule) Name() string {
	return "aws_ssm_maintenance_window_task_invalid_task_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskTypeRule) Check(runner *tflint.Runner) error {
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
					`task_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
