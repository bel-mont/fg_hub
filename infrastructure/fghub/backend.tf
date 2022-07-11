terraform {
  backend "s3" {
    bucket = "fghub-tf-state"
    key = "terraform/state"
    region = "ap-northeast-1"
  }
}