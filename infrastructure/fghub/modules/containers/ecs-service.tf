resource "aws_ecs_service" "fghub-web" {
  name            = "fghub-web-${var.ENV}"
  cluster         = aws_ecs_cluster.fghub-cluster.id
  task_definition = aws_ecs_task_definition.fghub-web.arn
  #  execution_role_arn = aws_iam_role.ecs-task-execution-role.arn
  #  task_role_arn      = aws_iam_role.ecs-demo-task-role.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets = var.PUBLIC_SUBNETS
  }
  #  iam_role        = aws_iam_role.foo.arn
  #    depends_on      = [aws_iam_role_policy.foo]

  load_balancer {
    target_group_arn = aws_alb_target_group.fghub-web-app.arn
    container_name   = "fghub-web-app-${var.ENV}"
    container_port   = 80
  }
  #
  #  placement_constraints {
  #    type       = "memberOf"
  #    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  #  }
}