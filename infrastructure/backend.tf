terraform {
  backend "s3" {
    bucket = "fghub-terraform-state"
    key = "terraform/state"
    region = var.AWS_REGION
  }
}