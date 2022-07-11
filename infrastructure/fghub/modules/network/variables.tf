variable "PLAT_NAME" {
  default = "fghub"
}

variable "AZ" {
  default = {
    "tokyo" = {
      "a" = "ap-northeast-1a",
      "c" = "ap-northeast-1d",
      "d" = "ap-northeast-1c",
    }
  }
}