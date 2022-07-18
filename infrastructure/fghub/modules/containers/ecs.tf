# cluster
resource "aws_ecs_cluster" "fghub-cluster-dev" {
  name = "fghub-cluster-${var.ENV}"
}

#resource "aws_launch_configuration" "fgapp-ecs-launchconfig" {
#  name_prefix          = "ecs-launchconfig"
#  image_id             = var.AMI[var.AWS_REGION]
#  instance_type        = var. INSTANCE_TYPE
#  key_name             = aws_key_pair.fgapp-key.key_name
#  iam_instance_profile = aws_iam_instance_profile.ecs-ec2-role.id
#  security_groups      = [aws_security_group.ecs-securitygroup.id]
#  user_data            = "#!/bin/bash\necho 'ECS_CLUSTER=example-cluster' > /etc/ecs/ecs.config\nstart ecs"
#  lifecycle {
#    create_before_destroy = true
#  }
#}
#
#resource "aws_autoscaling_group" "ecs-example-autoscaling" {
#  name                 = "ecs-example-autoscaling"
#  vpc_zone_identifier  = [aws_subnet.fghub-public-1.id, aws_subnet.fghub-public-2.id]
#  launch_configuration = aws_launch_configuration.fgapp-ecs-launchconfig.name
#  min_size             = 1
#  max_size             = 1
#  tag {
#    key                 = "Name"
#    value               = "ecs-ec2-container"
#    propagate_at_launch = true
#  }
#}
#
