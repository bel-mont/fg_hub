variable "AMI" {
  type    = map(map(string))
  default = {
    "ap-northeast-1" = {
      "aws-linux" = "ami-004332b441f90509b"
    }
  }
}

variable "AWS_LINUX" {
  default = "aws-linux"
}