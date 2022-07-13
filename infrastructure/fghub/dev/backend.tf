terraform {
  backend "s3" {
    bucket = "fghub-tf-state"
    key = "terraform/${var.ENV}/state"
    region = "ap-northeast-1"
  }
}