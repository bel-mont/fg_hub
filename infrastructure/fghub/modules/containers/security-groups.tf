resource "aws_security_group" "fghub-alb-securitygroup" {
  vpc_id      = var.VPC_ID
  name        = "fghub-alb-securitygroup-${var.ENV}"
  description = "Manages access to the FGHub load balancer"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"
  }
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "fghub-ecs-tasks-web-security-group" {
  vpc_id      = var.VPC_ID
  name        = "fghub-ecs-task-web-securitygroup-${var.ENV}"
  description = "Manage access into the ECS tasks"
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port       = 80
    to_port         = 80
    protocol        = "tcp"
    security_groups = [aws_security_group.fghub-alb-securitygroup.id]
  }
}