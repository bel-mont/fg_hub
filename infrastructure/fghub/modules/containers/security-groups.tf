resource "aws_security_group" "fghub-web-securitygroup" {
  vpc_id      = var.VPC_ID
  name        = "fghub-web-securitygroup-${var.ENV}"
  description = "Manages access to any ec2 instance used by the main FGHub web application"
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
    #    security_groups = [aws_security_group.myapp-elb-securitygroup.id]
  }
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}