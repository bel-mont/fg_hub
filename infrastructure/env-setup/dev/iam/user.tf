resource "aws_iam_user" "fghub-dev-bel-mont" {
  name = "fghub-dev-bel-mont"
}

resource "aws_iam_group_membership" "devops-dev-users" {
  name  = "devops-dev-users"
  users = [
    aws_iam_user.fghub-dev-bel-mont.name,
  ]
  group = aws_iam_group.fghub-devops-dev.name
}
