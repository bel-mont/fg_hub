resource "aws_key_pair" "fgapp-key" {
  key_name   = "fgapp-key"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
  lifecycle {
    ignore_changes = [public_key]
  }
}

variable "PATH_TO_PUBLIC_KEY" {
  default = "fgapp-key.pub"
}