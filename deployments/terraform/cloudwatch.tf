resource "aws_cloudwatch_log_group" "user" {
  name              = "/ecs/project/dev/user"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "auth" {
  name              = "/ecs/project/dev/auth"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "gateway" {
  name              = "/ecs/project/dev/gateway"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "course" {
  name              = "/ecs/project/dev/course"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "lesson" {
  name              = "/ecs/project/dev/lesson"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "test_case" {
  name              = "/ecs/project/dev/test_case"
  retention_in_days = 30
}

resource "aws_cloudwatch_log_group" "runner" {
  name              = "/ecs/project/dev/runner"
  retention_in_days = 30
}