resource "aws_ecs_service" "fghub-web-dev" {
  name            = "fghub-web-dev"
  cluster         = aws_ecs_cluster.fghub-cluster-dev.id
  task_definition = aws_ecs_task_definition.fghub-web.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets = var.SUBNETS
  }
  #  iam_role        = aws_iam_role.foo.arn
  #  depends_on      = [aws_iam_role_policy]

  #
  #  load_balancer {
  #    target_group_arn = aws_lb_target_group.foo.arn
  #    container_name   = "mongo"
  #    container_port   = 8080
  #  }
  #
  #  placement_constraints {
  #    type       = "memberOf"
  #    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  #  }
}