variable "secret_name" {
  type        = string
  description = "Secret's ARN to fetch"
}

variable "prefix" {
  type        = string
  description = "Custom secretprefix"
}

provider "aws" {
  s3_force_path_style         = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  access_key                  = "This is not an actual access key."
  secret_key                  = "This is not an actual secret key."

  endpoints {
    iam = "http://localstack:4593"
    ssm = "http://localstack:4583"
    sts = "http://localstack:4592"
  }
}

module "fetch_secret_arn" {
  source    = "../../"
  prefix    = var.prefix
  parameter = var.secret_name
}

output "arn" {
  value = module.fetch_secret_arn.arn
}
