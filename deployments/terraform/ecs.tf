resource "aws_ecs_cluster" "cluster" {
  name = "${var.product_name}-ecs-cluster"
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = file("files/assume_role_policy/ecs-task.json")
}

resource "aws_ecs_task_definition" "aws-ecs-task" {
  family                   = "${var.product_name}_task_definition"
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "1024"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name          = "user-task-definition"
      image         = aws_ecr_repository.user.repository_url
      cpu           = var.ecs_cpu
      memory        = var.ecs_memory
      port_mappings = [50051]
    },
    {
      name  = "auth-task-definition"
      image = aws_ecr_repository.auth.repository_url

      port_mappings = [50052]
    },
    {
      name          = "gatewaay-task-definition"
      image         = aws_ecr_repository.gateway.repository_url
      port_mappings = [8080]
    },
    {
      name          = "course-task-definition"
      image         = aws_ecr_repository.auth.repository_url
      port_mappings = [50053]
    },
    {
      name          = "lesson-task-definition"
      image         = aws_ecr_repository.auth.repository_url
      port_mappings = [50054]
    },
    {
      name          = "test-case-task-definition"
      image         = aws_ecr_repository.auth.repository_url
      port_mappings = [50056]
    },
    {
      name          = "runner-task-definition"
      image         = aws_ecr_repository.auth.repository_url
      port_mappings = [50055]
      secrets = [
        {
          "name" : "CLIENT_ID",
          "valueFrom" : "${aws_secretsmanager_secret.jdoodle.arn}:${var.jdoodle_client_id}::"
        },
        {
          "name" : "CLIENT_SECRET",
          "valueFrom" : "${aws_secretsmanager_secret.jdoodle.arn}:${var.jdoodle_client_secret}::"
        }
      ]
    },
  ])
}

resource "aws_ecs_service" "aws-ecs-service" {
  name                 = "${var.product_name}-${var.app_environment}-ecs-service"
  cluster              = aws_ecs_cluster.cluster.id
  task_definition      = aws_ecs_task_definition.aws-ecs-task.arn
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 1
  force_new_deployment = true

  network_configuration {
    subnets          = aws_subnet.private.*.id
    assign_public_ip = false
    security_groups = [
      aws_security_group.service_security_group.id,
      # aws_security_group.load_balancer_security_group.id
    ]
  }

  # load_balancer {
  #   target_group_arn = aws_lb_target_group.target_group.arn
  #   container_name   = "${var.product_name}-${var.app_environment}-container"
  #   container_port   = 8000
  # }

  # depends_on = [aws_lb_listener.listener]
}
