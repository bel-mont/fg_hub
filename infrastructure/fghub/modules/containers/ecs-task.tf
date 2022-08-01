resource "aws_ecs_task_definition" "fghub-web" {
  family                   = "fghub-web-${var.ENV}"
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  # TODO: Add task roles
  #  task_role_arn         = ""
  requires_compatibilities = [
    "FARGATE"
  ]
  container_definitions = jsonencode([
    {
      name         = "client-fgapp"
      image        = "nginx:1.23.0-alpine" # TODO: replace with webapp
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

