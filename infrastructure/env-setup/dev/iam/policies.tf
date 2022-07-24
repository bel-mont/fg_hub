resource "aws_iam_policy" "dev-ecr-access" {
  name        = "FGHubDevECRAccess"
  path        = "/"
  description = "Grants ECR Repository create / delete access on dev resources."

  policy = jsonencode({
    Version   = "2012-10-17"
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

resource "aws_iam_policy" "dev-ec2-access" {
  name        = "FGHubDevEC2Access"
  path        = "/"
  description = "Grants EC2 create / delete access on dev resources."

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:ImportKeyPair",
          "ec2:DeleteKeyPair",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:ec2:*:${var.AWS_ACC}:key-pair/fghub-*-${var.ENV}"
      },
    ]
  })
}

resource "aws_iam_policy" "dev-ec2-autoscaling-access" {
  name        = "FGHubDevEC2AutoscalingAccess"
  path        = "/"
  description = "Grants EC2 Autoscaling create / delete access on dev resources."

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action = [
          "autoscaling:SetInstanceProtection",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:autoscaling:*:${var.AWS_ACC}:fghub-*-${var.ENV}"
      },
    ]
  })
}
