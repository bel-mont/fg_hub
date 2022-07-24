resource "aws_ecr_repository" "fghub-web" {
  name = "fghub-web-${var.ENV}"
  // TODO: Add repository config. TTL etc.
}
