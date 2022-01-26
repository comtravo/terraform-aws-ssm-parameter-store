variable "secret_name" {
  type        = string
  description = "Secret's ARN to fetch"
}

provider "aws" {
}

module "fetch_secret_arn" {
  source = "../../"

  parameter = var.secret_name
}

output "arn" {
  value = module.fetch_secret_arn.arn
}
