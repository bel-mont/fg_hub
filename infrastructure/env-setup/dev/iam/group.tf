resource "aws_iam_group" "fghub-devops-dev" {
  name = "fghub-devops-dev"
}

// Might be better off creating individual policies with granular permissions.
// For now, using these to speed things up.
resource "aws_iam_group_policy_attachment" "fghub-devops-dev-vpc" {
  group     = aws_iam_group.fghub-devops-dev.name

  policy_arn = "arn:aws:iam::aws:policy/AmazonVPCFullAccess"
}

resource "aws_iam_group_policy_attachment" "fghub-devops-dev-ecs" {
  group     = aws_iam_group.fghub-devops-dev.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonECS_FullAccess"
}
