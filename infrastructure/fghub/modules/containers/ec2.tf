variable "INSTANCE_TYPE" {
  default = "t4g.nano"
}

variable "EC2ECSServiceRole" {
  default = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}