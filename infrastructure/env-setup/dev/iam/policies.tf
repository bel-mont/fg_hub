resource "aws_iam_policy" "dev-ecr-access" {
  name        = "FGHubDevECRAccess"
  path        = "/"
  description = "Grants ECR Repository create / delete access on dev resources."

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ecr:CreateRepository",
          "ecr:DeleteRepository",
          "ecr:DescribeRepositories",
          "ecr:ListTagsForResource",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:ecr:*:${var.AWS_ACC}:repository/fghub-*-${var.ENV}"
      },
    ]
  })
}