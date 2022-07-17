resource "aws_iam_policy" "dev-vpc-access" {
  name        = "FGHubDevVPCAccess"
  path        = "/"
  description = "Grants Full control on Dev VPC resources."

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = var.VPC_ACTIONS
        Effect   = "Allow"
        Resource = "arn:aws:vpc:*:${var.AWS_ACC}:fghub-*-${var.ENV}"
      },
    ]
  })
}