variable "secret_name" {
  type        = string
  description = "Secret's ARN to fetch"
}

variable "prefix" {
  type        = string
  description = "Custom secretprefix"
}

provider "aws" {
}

module "fetch_secret_arn" {
  source    = "../../"
  prefix    = var.prefix
  parameter = var.secret_name
}

output "arn" {
  value = module.fetch_secret_arn.arn
}
