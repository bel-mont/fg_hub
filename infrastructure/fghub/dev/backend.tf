terraform {
  backend "s3" {
    bucket = "fghub-tf-state"
    key = "terraform/dev/state"
    region = "ap-northeast-1"
  }
}