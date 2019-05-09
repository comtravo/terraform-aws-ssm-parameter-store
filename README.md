# terraform-aws-ssm-parameter-store
Terraform module to fetch ARNs of secrets from SSM parameter store

Usage:

```hcl
module "metabase_encryption_secret_key" {
  source = "github.com/comtravo/terraform-aws-ssm-parameter-store"

  parameter = "MB_ENCRYPTION_SECRET_KEY"
}

```
