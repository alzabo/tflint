# terraform_dash_in_resource_name

Disallow dashes (-) in `resource` names without.

## Example

```hcl
resource "aws_eip" "dash-name" {
}

resource "aws_eip" "no_dash_name" {
}
```

```
$ tflint
resources.tf
        NOTICE:1 `dash-name` resource name has a dash (terraform_dash_in_resource_name)
```

## Why

Naming conventions are optional, so it is not necessary to follow this. But this rule is useful if you want to force the following naming conventions in line with the [Terraform Plugin Naming Best Practices](https://www.terraform.io/docs/extend/best-practices/naming.html).

## How To Fix

Use underscores (_) instead of dashes (-).
