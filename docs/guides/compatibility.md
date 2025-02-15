# Compatibility with Terraform

Some inspections implicitly assume the behavior of a specific version of provider plugins or Terraform. This always assumes the latest version and is as follows:

- Terraform v0.12.6
- AWS Provider v2.24.0

Of course, TFLint may work correctly if you run it on other versions. But, false positives/negatives can occur based on this assumption.

## Input Variables

Like Terraform, it supports the `--var`,` --var-file` options, automatic loading of variable definitions (`.tfvars`) files, and environment variables.

## Named Values

[Named values](https://www.terraform.io/docs/configuration/expressions.html#references-to-named-values) are supported only for [input variables](https://www.terraform.io/docs/configuration/variables.html) and [workspaces](https://www.terraform.io/docs/state/workspaces.html). Expressions that contain anything else are excluded from the  inspection.

## Built-in Functions

[Built-in Functions](https://www.terraform.io/docs/configuration/functions.html) are fully supported.

## Environment Variables

The following environment variables are supported:

- [TF_VAR_name](https://www.terraform.io/docs/commands/environment-variables.html#tf_var_name)
- [TF_DATA_DIR](https://www.terraform.io/docs/commands/environment-variables.html#tf_data_dir)
