resource "aws_iam_group" "fghub-devops-dev" {
  name = "fghub-devops-dev"
}

resource "aws_iam_policy_attachment" "fghub-devops-dev-attach" {
  name       = "fghub-devops-dev-attach"
  groups     = [aws_iam_group.fghub-devops-dev.name]
  // Should it be more granular? Might take a lot of time to set it up properly.
  policy_arn = "arn:aws:iam::aws:policy/AmazonVPCFullAccess"
}
