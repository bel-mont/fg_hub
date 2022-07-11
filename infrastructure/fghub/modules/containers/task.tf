## app
#
#data "template_file" "fghub-task-definition-template" {
#  template = file("templates/app.json.tpl")
#  vars = {
#    REPOSITORY_URL = replace(aws_ecr_repository.fgweb-prod.repository_url, "https://", "")
#  }
#}
#
#resource "aws_ecs_task_definition" "myapp-task-definition" {
#  family                = "fghub"
#  container_definitions = data.template_file.fghub-task-definition-template.rendered
#}
#
#resource "aws_elb" "fghub" {
#  name = "myapp-elb"
#
#  listener {
#    instance_port     = 3000
#    instance_protocol = "http"
#    lb_port           = 80
#    lb_protocol       = "http"
#  }
#
#  health_check {
#    healthy_threshold   = 3
#    unhealthy_threshold = 3
#    timeout             = 30
#    target              = "HTTP:3000/"
#    interval            = 60
#  }
#
#  cross_zone_load_balancing   = true
#  idle_timeout                = 400
#  connection_draining         = true
#  connection_draining_timeout = 400
#
#  subnets         = [aws_subnet.fghub-public-1.id, aws_subnet.fghub-public-2.id]
#  security_groups = [aws_security_group.fghub-elb-securitygroup.id]
#
#  tags = {
#    Name = "myapp-elb"
#  }
#}
#
#resource "aws_ecs_service" "myapp-service" {
#  name            = "myapp"
#  cluster         = aws_ecs_cluster.fgapp-cluster.id
#  task_definition = aws_ecs_task_definition.myapp-task-definition.arn
#  desired_count   = 1
#  iam_role        = aws_iam_role.ecs-service-role.arn
#  depends_on      = [aws_iam_policy_attachment.ecs-service-attach1]
#
#  load_balancer {
#    elb_name       = aws_elb.fghub.name
#    container_name = "myapp"
#    container_port = 3000
#  }
#  lifecycle {
#    ignore_changes = [task_definition]
#  }
#}
#
