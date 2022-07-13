# Internet VPC
resource "aws_vpc" "fghub" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = "true"
  enable_dns_hostnames = "true"
  enable_classiclink   = "false"
  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}"
  }
}

# Subnets
resource "aws_subnet" "fghub-public-1" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = "true"
  availability_zone       = var.AZ.tokyo.a

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-public-1"
  }
}

resource "aws_subnet" "fghub-public-2" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.2.0/24"
  map_public_ip_on_launch = "true"
  availability_zone       = var.AZ.tokyo.c

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-public-2"
  }
}

resource "aws_subnet" "fghub-public-3" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.3.0/24"
  map_public_ip_on_launch = "true"
  availability_zone       = var.AZ.tokyo.d

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-public-3"
  }
}

resource "aws_subnet" "fghub-private-1" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.4.0/24"
  map_public_ip_on_launch = "false"
  availability_zone       = var.AZ.tokyo.a

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-private-1"
  }
}

resource "aws_subnet" "fghub-private-2" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.5.0/24"
  map_public_ip_on_launch = "false"
  availability_zone       = var.AZ.tokyo.c

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-private-2"
  }
}

resource "aws_subnet" "fghub-private-3" {
  vpc_id                  = aws_vpc.fghub.id
  cidr_block              = "10.0.6.0/24"
  map_public_ip_on_launch = "false"
  availability_zone       = var.AZ.tokyo.d

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-private-3"
  }
}

# Internet GW
resource "aws_internet_gateway" "fghub-main-gw" {
  vpc_id = aws_vpc.fghub.id

  tags = {
    Name = var.PLAT_NAME
  }
}

# route tables
resource "aws_route_table" "fghub-public" {
  vpc_id = aws_vpc.fghub.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.fghub-main-gw.id
  }

  tags = {
    Name = "${var.PLAT_NAME}-${var.ENV}-public-1"
  }
}

# route associations public
resource "aws_route_table_association" "fghub-public-1-a" {
  subnet_id      = aws_subnet.fghub-public-1.id
  route_table_id = aws_route_table.fghub-public.id
}

resource "aws_route_table_association" "fghub-public-2-a" {
  subnet_id      = aws_subnet.fghub-public-2.id
  route_table_id = aws_route_table.fghub-public.id
}

resource "aws_route_table_association" "fghub-public-3-a" {
  subnet_id      = aws_subnet.fghub-public-3.id
  route_table_id = aws_route_table.fghub-public.id
}

