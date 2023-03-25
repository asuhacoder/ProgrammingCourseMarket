resource "aws_vpc" "aws-vpc" {
  cidr_block           = "10.10.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Name        = "${var.product_name}-vpc"
    Environment = var.app_environment
  }
}

resource "aws_internet_gateway" "aws-igw" {
  vpc_id = aws_vpc.aws-vpc.id
  tags = {
    Name        = "${var.product_name}-igw"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.aws-vpc.id
  cidr_block        = var.private_subnet
  availability_zone = var.availability_zone

  tags = {
    Name        = "${var.product_name}-dummy-subnet"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "dummy" {
  vpc_id            = aws_vpc.aws-vpc.id
  cidr_block        = var.private_subnet_dummy
  availability_zone = var.availability_zone_dummy

  tags = {
    Name        = "${var.product_name}-private-subnet"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.aws-vpc.id
  cidr_block              = var.public_subnet
  availability_zone       = var.availability_zone
  map_public_ip_on_launch = true

  tags = {
    Name        = "${var.product_name}-public-subnet"
    Environment = var.app_environment
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.aws-vpc.id

  tags = {
    Name        = "${var.product_name}-routing-table-public"
    Environment = var.app_environment
  }
}

resource "aws_route" "public" {
  route_table_id         = aws_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.aws-igw.id
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.public.id
}

resource "aws_security_group" "service_security_group" {
  vpc_id = aws_vpc.aws-vpc.id

  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    # security_groups = [aws_security_group.load_balancer_security_group.id]
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

# resource "aws_alb" "application_load_balancer" {
#   name               = "${var.product_name}-${var.app_environment}-alb"
#   internal           = false
#   load_balancer_type = "application"
#   subnets            = aws_subnet.public.*.id
#   security_groups    = [aws_security_group.load_balancer_security_group.id]

#   tags = {
#     Name        = "${var.product_name}-alb"
#     Environment = var.app_environment
#   }
# }

# resource "aws_security_group" "load_balancer_security_group" {
#   vpc_id = aws_vpc.aws-vpc.id

#   ingress {
#     from_port        = 80
#     to_port          = 80
#     protocol         = "tcp"
#     cidr_blocks      = ["0.0.0.0/0"]
#     ipv6_cidr_blocks = ["::/0"]
#   }

#   egress {
#     from_port        = 0
#     to_port          = 0
#     protocol         = "-1"
#     cidr_blocks      = ["0.0.0.0/0"]
#     ipv6_cidr_blocks = ["::/0"]
#   }
#   tags = {
#     Name        = "${var.product_name}-sg"
#     Environment = var.app_environment
#   }
# }

# resource "aws_lb_target_group" "target_group" {
#   name        = "${var.product_name}-${var.app_environment}-tg"
#   port        = 80
#   protocol    = "HTTP"
#   target_type = "ip"
#   vpc_id      = aws_vpc.aws-vpc.id

#   health_check {
#     healthy_threshold   = "3"
#     interval            = "300"
#     protocol            = "HTTP"
#     matcher             = "200"
#     timeout             = "3"
#     path                = "/v1/status"
#     unhealthy_threshold = "2"
#   }

#   tags = {
#     Name        = "${var.product_name}-lb-tg"
#     Environment = var.app_environment
#   }
# }

# resource "aws_lb_listener" "listener" {
#   load_balancer_arn = aws_alb.application_load_balancer.id
#   port              = "80"
#   protocol          = "HTTP"
#   certificate_arn   = aws_acm_certificate.cert.arn

#   default_action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.target_group.id
#   }
# }