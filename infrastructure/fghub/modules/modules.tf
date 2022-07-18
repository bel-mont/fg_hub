module "storage" {
  source = "./storage"
  #  ENV = var.ENV
}

module "network" {
  source = "./network"
  ENV    = var.ENV
}

module "containers" {
  source  = "./containers"
  ENV     = var.ENV
  SUBNETS = [
    module.network.fghub-public-1, module.network.fghub-public-2, module.network.fghub-public-3,
    module.network.fghub-private-1, module.network.fghub-private-2, module.network.fghub-private-3
  ]
}