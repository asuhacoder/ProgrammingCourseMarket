resource "aws_ecs_cluster" "cluster" {
  name = "${var.product_name}-ecs-cluster"
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = file("files/assume_role_policy/ecs-task.json")
}

resource "aws_iam_role_policy" "ecs_task_execution_role_policy" {
  name   = "iam_role_policy"
  role   = aws_iam_role.ecs_task_execution_role.id
  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource": "*"
      },
      {
        "Effect": "Allow",
        "Action": [
          "secretsmanager:GetSecretValue"
        ],
        "Resource": [
          "${aws_secretsmanager_secret.jdoodle_client_id.arn}",
          "${aws_secretsmanager_secret.jdoodle_client_secret.arn}"
        ]
      }
    ]
  }
  EOF
}
# {
#   "Sid": "AllowDescribeRepoImage",
#   "Effect": "Allow",
#   "Action": [
#       "ecr:DescribeImages",
#       "ecr:DescribeRepositories"
#   ],
#   "Resource": [
#     "${aws_ecr_repository.user.arn}",
#     "${aws_ecr_repository.auth.arn}",
#     "${aws_ecr_repository.gateway.arn}",
#     "${aws_ecr_repository.course.arn}",
#     "${aws_ecr_repository.lesson.arn}",
#     "${aws_ecr_repository.test_case.arn}",
#     "${aws_ecr_repository.runner.arn}"
#   ]
# }



resource "aws_iam_instance_profile" "ecs_task_execution_profile" {
  name = "profile"
  role = aws_iam_role.ecs_task_execution_role.name
}



resource "aws_ecs_task_definition" "aws-ecs-task-gateway" {
  family                   = "${var.product_name}_task_definition_gateway"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "gateway-task-definition"
      image = aws_ecr_repository.gateway.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "gateway"
          awslogs-group         = "/ecs/project/dev/gateway"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [
        {
          name  = "USER_HOST"
          value = "service-discovery-user.${var.product_name}.local"
        },
        {
          name  = "AUTH_HOST"
          value = "service-discovery-auth.${var.product_name}.local"
        },
        {
          name  = "COURSE_HOST"
          value = "service-discovery-course.${var.product_name}.local"
        },
        {
          name  = "LESSON_HOST"
          value = "service-discovery-lesson.${var.product_name}.local"
        },
        {
          name  = "TEST_CASE_HOST"
          value = "service-discovery-test-case.${var.product_name}.local"
        },
        {
          name  = "RUNNER_HOST"
          value = "service-discovery-runner.${var.product_name}.local"
        }
      ]
      portMappings = [{
        hostPort      = 8080
        protocol      = "tcp"
        containerPort = 8080
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-user" {
  family                   = "${var.product_name}_task_definition_user"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "user-task-definition"
      image = aws_ecr_repository.user.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "user"
          awslogs-group         = "/ecs/project/dev/user"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [{
        name  = "USER_DB_HOST"
        value = "${aws_db_instance.user.address}"
      }]
      portMappings = [{
        hostPort      = 50051
        protocol      = "tcp"
        containerPort = 50051
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-auth" {
  family                   = "${var.product_name}_task_definition_auth"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "auth-task-definition"
      image = aws_ecr_repository.auth.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "auth"
          awslogs-group         = "/ecs/project/dev/auth"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [{
        name  = "USER_DB_HOST"
        value = "${aws_db_instance.user.address}"
      }]
      portMappings = [{
        hostPort      = 50052
        protocol      = "tcp"
        containerPort = 50052
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-course" {
  family                   = "${var.product_name}_task_definition_course"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "course-task-definition"
      image = aws_ecr_repository.course.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "course"
          awslogs-group         = "/ecs/project/dev/course"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [
        {
          name  = "COURSE_DB_HOST"
          value = "${aws_db_instance.course.address}"
        },
        {
          name  = "USER_DB_HOST"
          value = "${aws_db_instance.user.address}"
        }
      ]
      portMappings = [{
        hostPort      = 50053
        protocol      = "tcp"
        containerPort = 50053
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-lesson" {
  family                   = "${var.product_name}_task_definition_lesson"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "lesson-task-definition"
      image = aws_ecr_repository.lesson.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "lesson"
          awslogs-group         = "/ecs/project/dev/lesson"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [
        {
          name  = "LESSON_DB_HOST"
          value = "${aws_db_instance.lesson.address}"
        },
        {
          name  = "USER_DB_HOST"
          value = "${aws_db_instance.user.address}"
        }
      ]
      portMappings = [{
        hostPort      = 50054
        protocol      = "tcp"
        containerPort = 50054
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-test-case" {
  family                   = "${var.product_name}_task_definition_test_case"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "test-case-task-definition"
      image = aws_ecr_repository.test_case.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "test_case"
          awslogs-group         = "/ecs/project/dev/test_case"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [
        {
          name  = "TEST_CASE_DB_HOST"
          value = "${aws_db_instance.test_case.address}"
        },
        {
          name  = "USER_DB_HOST"
          value = "${aws_db_instance.user.address}"
        }
      ]
      portMappings = [{
        hostPort      = 50056
        protocol      = "tcp"
        containerPort = 50056
      }]
    }
  ])
}

resource "aws_ecs_task_definition" "aws-ecs-task-runner" {
  family                   = "${var.product_name}_task_definition_runner"
  network_mode             = "awsvpc"
  cpu                      = var.ecs_task_definition_cpu
  memory                   = var.ecs_task_definition_memory
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "runner-task-definition"
      image = aws_ecr_repository.runner.repository_url
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.region
          awslogs-stream-prefix = "runner"
          awslogs-group         = "/ecs/project/dev/runner"
        }
      }
      cpu    = var.ecs_task_definition_cpu
      memory = var.ecs_task_definition_memory
      environment = [{
        name  = "USER_DB_HOST"
        value = "${aws_db_instance.user.address}"
      }]
      portMappings = [{
        hostPort      = 50055
        protocol      = "tcp"
        containerPort = 50055
      }]
      secrets = [
        {
          "name" : "CLIENT_ID",
          "valueFrom" : "${aws_secretsmanager_secret.jdoodle_client_id.arn}"
        },
        {
          "name" : "CLIENT_SECRET",
          "valueFrom" : "${aws_secretsmanager_secret.jdoodle_client_secret.arn}"
        }
      ]
    },
  ])
}

resource "aws_ecs_service" "aws-ecs-service-gateway" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-gateway"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-gateway.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
      aws_security_group.load_balancer_security_group.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn
    container_name   = "gateway-task-definition"
    container_port   = 8080
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["gateway"].arn
    port         = 8080
  }

  depends_on = [aws_lb_listener.listener]
}

resource "aws_ecs_service" "aws-ecs-service-user" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-user"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-user.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["user"].arn
    port         = 50051
  }
}

resource "aws_ecs_service" "aws-ecs-service-auth" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-auth"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-auth.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["auth"].arn
    port         = 50052
  }
}

resource "aws_ecs_service" "aws-ecs-service-course" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-course"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-course.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["course"].arn
    port         = 50053
  }
}

resource "aws_ecs_service" "aws-ecs-service-lesson" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-lesson"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-lesson.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["lesson"].arn
    port         = 50054
  }
}

resource "aws_ecs_service" "aws-ecs-service-test-case" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-test-case"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-test-case.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.private-ecs.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["test-case"].arn
    port         = 50056
  }
}

resource "aws_ecs_service" "aws-ecs-service-runner" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service-runner"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task-runner.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = [aws_subnet.public.id]
    assign_public_ip = true
    security_groups = [
      aws_security_group.service_security_group.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.skhole["runner"].arn
    port         = 50055
  }
}

resource "aws_service_discovery_service" "skhole" {
  for_each = toset(["gateway", "user", "auth", "course", "lesson", "test-case", "runner"])
  name     = "service-discovery-${each.key}"
  dns_config {
    namespace_id   = aws_service_discovery_private_dns_namespace.this.id
    routing_policy = "MULTIVALUE"
    dns_records {
      ttl  = 10
      type = "A"
    }
    dns_records {
      ttl  = 10
      type = "SRV"
    }
  }
  health_check_custom_config {
    failure_threshold = 1
  }
}