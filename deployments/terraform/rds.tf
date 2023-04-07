resource "aws_db_subnet_group" "db-subnet" {
  name       = "db-subnet"
  subnet_ids = [aws_subnet.private.id, aws_subnet.dummy.id]
}

resource "aws_db_parameter_group" "pg" {
  name   = "my-pg"
  family = "postgres11"

  parameter {
    name  = "log_connections"
    value = "1"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_db_instance" "user" {
  allocated_storage      = 10
  db_name                = "gorm"
  engine                 = var.db_engine
  engine_version         = "11.18"
  instance_class         = var.db_instance_class
  username               = "gorm"
  password               = "gormpassword"
  parameter_group_name   = aws_db_parameter_group.pg.name
  skip_final_snapshot    = true
  db_subnet_group_name   = aws_db_subnet_group.db-subnet.name
  vpc_security_group_ids = ["${aws_security_group.private-db-sg-user.id}"]
}

resource "aws_security_group" "private-db-sg-user" {
  name   = "private-db-sg-user"
  vpc_id = aws_vpc.aws-vpc.id
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.public_subnet]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "public-db-sg"
  }
}

resource "aws_db_instance" "course" {
  allocated_storage      = 10
  db_name                = "gorm"
  engine                 = var.db_engine
  engine_version         = "11.18"
  instance_class         = var.db_instance_class
  username               = "gorm"
  password               = "gormpassword"
  parameter_group_name   = aws_db_parameter_group.pg.name
  skip_final_snapshot    = true
  db_subnet_group_name   = aws_db_subnet_group.db-subnet.name
  vpc_security_group_ids = ["${aws_security_group.private-db-sg-course.id}"]
}

resource "aws_security_group" "private-db-sg-course" {
  name   = "private-db-sg-course"
  vpc_id = aws_vpc.aws-vpc.id
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.public_subnet]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "public-db-sg"
  }
}

resource "aws_db_instance" "lesson" {
  allocated_storage      = 10
  db_name                = "gorm"
  engine                 = var.db_engine
  engine_version         = "11.18"
  instance_class         = var.db_instance_class
  username               = "gorm"
  password               = "gormpassword"
  parameter_group_name   = aws_db_parameter_group.pg.name
  skip_final_snapshot    = true
  db_subnet_group_name   = aws_db_subnet_group.db-subnet.name
  vpc_security_group_ids = ["${aws_security_group.private-db-sg-lesson.id}"]
}

resource "aws_security_group" "private-db-sg-lesson" {
  name   = "private-db-sg-lesson"
  vpc_id = aws_vpc.aws-vpc.id
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.public_subnet]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "public-db-sg"
  }
}

resource "aws_db_instance" "test_case" {
  allocated_storage      = 10
  db_name                = "gorm"
  engine                 = var.db_engine
  engine_version         = "11.18"
  instance_class         = var.db_instance_class
  username               = "gorm"
  password               = "gormpassword"
  parameter_group_name   = aws_db_parameter_group.pg.name
  skip_final_snapshot    = true
  db_subnet_group_name   = aws_db_subnet_group.db-subnet.name
  vpc_security_group_ids = ["${aws_security_group.private-db-sg-test_case.id}"]
}

resource "aws_security_group" "private-db-sg-test_case" {
  name   = "private-db-sg-test_case"
  vpc_id = aws_vpc.aws-vpc.id
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.public_subnet]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "public-db-sg"
  }
}