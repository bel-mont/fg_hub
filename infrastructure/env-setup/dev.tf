module "dev" {
  source = "./dev"
  VPC_ACTIONS = var.VPC_ACTIONS
  AWS_ACC = var.AWS_ACC
}
