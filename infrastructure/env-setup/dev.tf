# group definition
resource "aws_iam_group" "fghub-devops-dev" {
  name = "fghub-devops-dev"
}

resource "aws_iam_policy_attachment" "fghub-devops-dev-attach" {
  name       = "fghub-devops-dev-attach"
  groups     = [aws_iam_group.fghub-devops-dev.name]
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess" # TODO: change to more granular permissions
}

# user
resource "aws_iam_user" "fghub-dev-bel-mont" {
  name = "fghub-dev-bel-mont"
}

resource "aws_iam_group_membership" "devops-dev-users" {
  name = "devops-dev-users"
  users = [
    aws_iam_user.fghub-dev-bel-mont.name,
  ]
  group = aws_iam_group.fghub-devops-dev.name
}
