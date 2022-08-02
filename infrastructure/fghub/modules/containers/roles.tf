resource "aws_iam_role" "fghub-ecs-task-execution-role" {
  name               = "fghub-ecs-task-execution-role-${var.ENV}"
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "ecs-tasks.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      }
    ]
  })
}


resource "aws_iam_role_policy" "fghub-ecs-task-execution-role" {
  name = "fghub-ecs-task-execution-role-policy-${var.ENV}"
  role = aws_iam_role.fghub-ecs-task-execution-role.id

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          #          "logs:CreateLogStream",
          #          "logs:PutLogEvents",
          #          "ssm:GetParameters",
          #          "ssm:GetParameter"
        ],
        "Resource" : "*" # TODO: Limit to var.ENV resources
      }
    ]
  })
}

# Web application role
# TODO: Rest of containers
resource "aws_iam_role" "fghub-ecs-task-web-role" {
  name = "fghub-ecs-task-web-role-${var.ENV}"

  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "",
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "ecs-tasks.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      }
    ]
  })
}
