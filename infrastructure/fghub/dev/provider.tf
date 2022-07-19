provider "aws" {
  region  = var.REGION
  profile = "fghub-bel-mont"
  default_tags {
    tags = {
      Terraform = true
    }
  }
}
