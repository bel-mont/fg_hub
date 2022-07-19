resource "tls_private_key" "fghub-web-keypair" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "fghub-web-keypair" {
  key_name   = "fghub-web-keypair-${var.ENV}"
  public_key = tls_private_key.fghub-web-keypair.public_key_openssh

  provisioner "local-exec" {
    command = "echo '${tls_private_key.fghub-web-keypair.private_key_pem}' > ./fghub-web-keypair.pem"
  }
}