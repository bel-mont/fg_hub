resource "aws_ecr_repository" "fgweb" {
  name = "fgweb-${var.ENV}"
}
