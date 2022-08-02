resource "aws_ecs_task_definition" "fghub-web" {
  family                   = "fghub-web-${var.ENV}"
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  # The task execution role is the one that runs the ECS actions like pulling data off of ECR.
  # https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html
  execution_role_arn       = aws_iam_role.fghub-ecs-task-execution-role.arn
  # The task role is what is assigned to the container(s) running inside the task, for example
  # to store data in S3, etc.
  # https://stackoverflow.com/questions/48999472/difference-between-aws-elastic-container-services-ecs-executionrole-and-taskr
  task_role_arn            = aws_iam_role.fghub-ecs-task-web-role.arn
  requires_compatibilities = [
    "FARGATE"
  ]
  container_definitions = jsonencode([
    {
      name         = "client-fgapp"
      image        = "nginx:1.23.0-alpine" # TODO: replace with webapp ecr url
      essential    = true
      portMappings = [
        {
          containerPort = 80
          # hostPort      = 80 #not needed for fargate
        }
      ]
      # logConfiguration = {}
      # secrets = []
      # environment = []
    }
  ])
}

