// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsAthenaNamedQueryInvalidDatabaseRule checks the pattern is valid
type AwsAthenaNamedQueryInvalidDatabaseRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAthenaNamedQueryInvalidDatabaseRule returns new rule with default attributes
func NewAwsAthenaNamedQueryInvalidDatabaseRule() *AwsAthenaNamedQueryInvalidDatabaseRule {
	return &AwsAthenaNamedQueryInvalidDatabaseRule{
		resourceType:  "aws_athena_named_query",
		attributeName: "database",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAthenaNamedQueryInvalidDatabaseRule) Name() string {
	return "aws_athena_named_query_invalid_database"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAthenaNamedQueryInvalidDatabaseRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsAthenaNamedQueryInvalidDatabaseRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsAthenaNamedQueryInvalidDatabaseRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAthenaNamedQueryInvalidDatabaseRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"database must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"database must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
