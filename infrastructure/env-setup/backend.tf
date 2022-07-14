terraform {
  backend "s3" {
    bucket = "fghub-tf-state"
    key = "terraform/env-setup/state"
    region = "ap-northeast-1"
  }
}