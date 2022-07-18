resource "aws_ecr_repository" "fghub-web-dev" {
  name = "fghub-web-${var.ENV}"
  // TODO: Add repository config. TTL etc.
}
