variable "parameter" {
  type        = "string"
  description = "parameter for which the ARN needs to be fetched"
}

locals {
  prefix = "${upper(terraform.workspace)}_"
}

data "aws_ssm_parameter" "parameter" {
  name            = "${local.prefix}${var.parameter}"
  with_decryption = false
}

output arn {
  description = "SSM parameter ARN"
  value       = "${data.aws_ssm_parameter.parameter.arn}"
}
