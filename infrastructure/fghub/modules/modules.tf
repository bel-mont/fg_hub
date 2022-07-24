module "storage" {
  source = "./storage"
  #  ENV = var.ENV
}

module "network" {
  source = "./network"
  ENV    = var.ENV
}

module "containers" {
  source         = "./containers"
  ENV            = var.ENV
  PUBLIC_SUBNETS = [
    module.network.fghub-public-1, module.network.fghub-public-2, module.network.fghub-public-3,
  ]
  VPC_ID = module.network.fghub-vpc-id
  REGION = var.REGION
}