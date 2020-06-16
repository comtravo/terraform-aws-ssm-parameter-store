# Terraform AWS module for AWS SSM parameter store data module

## Introduction

This module fetches the arn of the secret in AWS SSM parameter store.

## Usage  
Checkout [examples](./examples) for how to use this module

## Authors

Module managed by [Comtravo](https://github.com/comtravo).

## License

MIT Licensed. See [LICENSE](LICENSE) for full details.

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.12 |
| aws | ~> 2.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 2.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| parameter | SSM parameter for which the ARN needs to be fetched | `string` | n/a | yes |
| enable | Enable this module | `bool` | `true` | no |
| prefix | SSM parameter prefix | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| arn | SSM parameter ARN |

