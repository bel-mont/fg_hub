# cluster
resource "aws_ecs_cluster" "fghub-cluster-dev" {
  name = "fghub-cluster-${var.ENV}"
}

resource "aws_launch_configuration" "fghub-web-launchconfig" {
  name_prefix     = "fghub-web-launchconfig-${var.ENV}"
  image_id        = var.AMI[var.REGION][var.AWS_LINUX]
  instance_type   = var.INSTANCE_TYPE
  key_name        = aws_key_pair.fghub-web-keypair.key_name
  #  iam_instance_profile = aws_iam_instance_profile.ecs-ec2-role.id
  security_groups = [aws_security_group.fghub-web-securitygroup.id]
  user_data       = "#!/bin/bash\necho ECS_CLUSTER=${aws_ecs_cluster.fghub-cluster-dev.name} >> /etc/ecs/ecs.config"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "fghub-web-autoscaling" {
  name                 = "fghub-web-autoscaling-${var.ENV}"
  vpc_zone_identifier  = var.SUBNETS
  launch_configuration = aws_launch_configuration.fghub-web-launchconfig.name
  min_size             = 1
  max_size             = 1
}

