/**
* # Terraform AWS module for AWS SSM parameter store data module
*
* ## Introduction
*
* This module fetches the arn of the secret in AWS SSM parameter store.
*
* ## Usage
* Checkout [examples](./examples) for how to use this module
*
* ## Authors
*
* Module managed by [Comtravo](https://github.com/comtravo).
*
* ## License
*
* MIT Licensed. See [LICENSE](LICENSE) for full details.
*/

variable "parameter" {
  type        = string
  description = "parameter for which the ARN needs to be fetched"
}

variable "prefix" {
  type        = string
  description = "parameter for which the ARN needs to be fetched"
  default = null
}

locals {
  prefix = var.prefix != null ? var.prefix : "${upper(terraform.workspace)}_"
}

data "aws_ssm_parameter" "parameter" {
  name            = "${local.prefix}${var.parameter}"
  with_decryption = false
}

output "arn" {
  description = "SSM parameter ARN"
  value       = data.aws_ssm_parameter.parameter.arn
}
