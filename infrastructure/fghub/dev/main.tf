module "local_modules" {
  source = "../modules"
  ENV    = var.ENV
  REGION = var.REGION
}