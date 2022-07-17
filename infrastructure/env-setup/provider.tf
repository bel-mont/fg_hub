provider "aws" {
  region = var.AWS_REGION
  default_tags {
    tags = {
      Terraform = true
    }
  }
}
