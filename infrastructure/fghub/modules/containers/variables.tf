variable "ENV" {}

variable "PUBLIC_SUBNETS" {
  type = list(string)
}

variable "PRIVATE_SUBNETS" {
  type = list(string)
}

variable "VPC_ID" {
  type = string
}

variable "REGION" {
  type = string
}
