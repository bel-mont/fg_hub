module "storage" {
  source = "./storage"
#  ENV = var.ENV
}

module "network" {
  source = "./network"
  ENV = var.ENV
}
