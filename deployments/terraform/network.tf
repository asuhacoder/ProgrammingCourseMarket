resource "aws_vpc" "aws_vpc" {
  cidr_block           = "10.10.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Name        = "${var.product_name}-vpc"
    Environment = var.app_environment
  }
}

resource "aws_internet_gateway" "aws_igw" {
  vpc_id = aws_vpc.aws_vpc.id
  tags = {
    Name        = "${var.product_name}-igw"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.aws_vpc.id
  cidr_block        = var.private_subnet
  availability_zone = var.availability_zone

  tags = {
    Name        = "${var.product_name}-private-subnet"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "private-ecs" {
  vpc_id            = aws_vpc.aws_vpc.id
  cidr_block        = var.private_subnet_ecs
  availability_zone = var.availability_zone

  tags = {
    Name        = "${var.product_name}-private-subnet-ecs"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "dummy" {
  vpc_id            = aws_vpc.aws_vpc.id
  cidr_block        = var.private_subnet_dummy
  availability_zone = var.availability_zone_dummy

  tags = {
    Name        = "${var.product_name}-private-subnet-dummy"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.aws_vpc.id
  cidr_block              = var.public_subnet
  availability_zone       = var.availability_zone
  map_public_ip_on_launch = true

  tags = {
    Name        = "${var.product_name}-public-subnet"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "public_dummy" {
  vpc_id                  = aws_vpc.aws_vpc.id
  cidr_block              = var.public_dummy_subnet
  availability_zone       = var.availability_zone_dummy
  map_public_ip_on_launch = true

  tags = {
    Name        = "${var.product_name}-public-dummy-subnet"
    Environment = var.app_environment
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.aws_vpc.id

  tags = {
    Name        = "${var.product_name}-routing-table-public"
    Environment = var.app_environment
  }
}

resource "aws_route" "public" {
  route_table_id         = aws_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.aws_igw.id
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_dummy" {
  subnet_id      = aws_subnet.public_dummy.id
  route_table_id = aws_route_table.public.id
}

resource "aws_security_group" "service_security_group" {
  vpc_id = aws_vpc.aws_vpc.id

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = [aws_security_group.load_balancer_security_group.id]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name        = "${var.product_name}-service-sg"
    Environment = var.app_environment
  }
}

resource "aws_route_table" "ecs_private" {
  vpc_id = aws_vpc.aws_vpc.id
}

resource "aws_route_table_association" "ecs_private" {
  subnet_id      = aws_subnet.private-ecs.id
  route_table_id = aws_route_table.ecs_private.id
}

resource "aws_vpc_endpoint" "s3" {
  vpc_id            = aws_vpc.aws_vpc.id
  service_name      = "com.amazonaws.${var.region}.s3"
  vpc_endpoint_type = "Gateway"
}

resource "aws_vpc_endpoint_route_table_association" "private_s3" {
  vpc_endpoint_id = aws_vpc_endpoint.s3.id
  route_table_id  = aws_route_table.ecs_private.id
}

resource "aws_security_group" "vpc_endpoint" {
  name   = "vpc_endpoint_sg"
  vpc_id = aws_vpc.aws_vpc.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.aws_vpc.cidr_block]
  }

  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.aws_vpc.cidr_block]
  }
}

resource "aws_vpc_endpoint" "ecr_dkr" {
  vpc_id              = aws_vpc.aws_vpc.id
  service_name        = "com.amazonaws.${var.region}.ecr.dkr"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private-ecs.id, aws_subnet.dummy.id]
  security_group_ids  = [aws_security_group.vpc_endpoint.id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id              = aws_vpc.aws_vpc.id
  service_name        = "com.amazonaws.${var.region}.ecr.api"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private-ecs.id, aws_subnet.dummy.id]
  security_group_ids  = [aws_security_group.vpc_endpoint.id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "logs" {
  vpc_id              = aws_vpc.aws_vpc.id
  service_name        = "com.amazonaws.${var.region}.logs"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private-ecs.id, aws_subnet.dummy.id]
  security_group_ids  = [aws_security_group.vpc_endpoint.id]
  private_dns_enabled = true
}

resource "aws_vpc_endpoint" "secretsmanager" {
  vpc_id              = aws_vpc.aws_vpc.id
  service_name        = "com.amazonaws.${var.region}.secretsmanager"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private-ecs.id]
  security_group_ids  = [aws_security_group.vpc_endpoint.id]
  private_dns_enabled = true
}
