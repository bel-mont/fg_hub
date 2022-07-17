provider "aws" {
  region = var.AWS_REGION
  profile = "fghub-bel-mont"
  default_tags {
    tags = {
      Terraform = true
    }
  }
}
