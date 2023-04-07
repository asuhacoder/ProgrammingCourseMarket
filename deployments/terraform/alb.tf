resource "aws_s3_bucket" "lb_logs" {
  bucket = "alb-${var.product_name}-20230405"
}

resource "aws_lb" "application_load_balancer" {
  name               = "${var.product_name}-${var.app_environment}-alb"
  internal           = false
  load_balancer_type = "application"
  subnets            = [aws_subnet.public.id, aws_subnet.public_dummy.id]
  security_groups    = [aws_security_group.load_balancer_security_group.id]

  access_logs {
    bucket  = aws_s3_bucket.lb_logs.id
    prefix  = "log-lb"
    enabled = true
  }

  tags = {
    Name        = "${var.product_name}-alb"
    Environment = var.app_environment
  }
}

resource "aws_security_group" "load_balancer_security_group" {
  vpc_id = aws_vpc.aws-vpc.id

  ingress {
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name        = "${var.product_name}-sg"
    Environment = var.app_environment
  }
}

resource "aws_lb_target_group" "target_group" {
  name        = "${var.product_name}-${var.app_environment}-tg"
  port        = 8080
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.aws-vpc.id

  health_check {
    interval            = 30
    path                = "/health"
    port                = 8080
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2
    matcher             = 200
  }

  tags = {
    Name        = "${var.product_name}-lb-tg"
    Environment = var.app_environment
  }

  lifecycle {
    create_before_destroy = true
  }
}

# resource "aws_lb_target_group_attachment" "gateway" {
#   target_group_arn = aws_lb_target_group.target_group.arn
#   target_id        = aws_ecs_task_definition.aws-ecs-task.container_definitions[2].arn
#   port             = 8080
# }

resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_lb.application_load_balancer.id
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.id
  }

  depends_on = [
    aws_acm_certificate_validation.cert
  ]
}

data "aws_iam_policy_document" "lb_logs" {
  statement {
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::127311923021:root"]
    }
    actions   = ["s3:PutObject"]
    resources = ["${aws_s3_bucket.lb_logs.arn}/*"]
  }
}

resource "aws_s3_bucket_policy" "tf_bucket_policy_lb_logs" {
  bucket = aws_s3_bucket.lb_logs.id
  policy = data.aws_iam_policy_document.lb_logs.json
}