variable "ENV" {}

variable "SUBNETS" {
  type = list(string)
}

variable "VPC_ID" {
  type = string
}

variable "REGION" {
  type = string
}
