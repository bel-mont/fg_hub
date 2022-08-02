resource "aws_alb" "fghub-web-alb" {
  name                             = "fghub-web-alb-${var.ENV}"
  subnets                          = var.PUBLIC_SUBNETS
  load_balancer_type               = "application"
  enable_cross_zone_load_balancing = true
  security_groups                  = [aws_security_group.fghub-alb-securitygroup.id]
}

resource "aws_alb_listener" "fghub-web-app" {
  load_balancer_arn = aws_alb.fghub-web-alb.arn
  port              = "80"
  protocol          = "TCP"

  default_action {
    target_group_arn = aws_alb_target_group.fghub-web-app.id
    type             = "forward"
  }
  lifecycle {
    ignore_changes = [
      default_action,
    ]
  }
}

resource "aws_alb_target_group" "fghub-web-app" {
  name                 = "fghub-web-app-${var.ENV}"
  port                 = "80"
  protocol             = "TCP"
  target_type          = "ip"
  vpc_id               = var.VPC_ID
  deregistration_delay = "30"

  health_check {
    healthy_threshold   = 2
    unhealthy_threshold = 2
    protocol            = "TCP"
    interval            = 30
  }
}