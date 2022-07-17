resource "aws_iam_group" "fghub-devops-dev" {
  name = "fghub-devops-dev"
}

resource "aws_iam_policy_attachment" "fghub-devops-dev-attach" {
  name       = "fghub-devops-dev-attach"
  groups     = [aws_iam_group.fghub-devops-dev.name]
  policy_arn = aws_iam_policy.dev-vpc-access.arn
}
