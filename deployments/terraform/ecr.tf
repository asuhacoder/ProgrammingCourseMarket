resource "aws_ecr_repository" "user" {
  name                 = "${var.product_name}_user"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "user" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.user.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.user.repository_url} user"
  }
}

resource "aws_ecr_repository" "auth" {
  name                 = "${var.product_name}_auth"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "auth" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.auth.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.auth.repository_url} auth"
  }
}

resource "aws_ecr_repository" "gateway" {
  name                 = "${var.product_name}_gateway"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "gateway" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.gateway.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.gateway.repository_url} gateway"
  }
}

resource "aws_ecr_repository" "course" {
  name                 = "${var.product_name}_course"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "course" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.course.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.course.repository_url} course"
  }
}

resource "aws_ecr_repository" "lesson" {
  name                 = "${var.product_name}_lesson"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "lesson" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.lesson.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.lesson.repository_url} lesson"
  }
}

resource "aws_ecr_repository" "test_case" {
  name                 = "${var.product_name}_test_case"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "test_case" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.test_case.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.test_case.repository_url} test_case"
  }
}

resource "aws_ecr_repository" "runner" {
  name                 = "${var.product_name}_runner"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "null_resource" "runner" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ${var.region} | docker login --username AWS --password-stdin ${aws_ecr_repository.runner.repository_url}"
  }

  provisioner "local-exec" {
    command = "${var.script_path} ${aws_ecr_repository.runner.repository_url} runner"
  }
}