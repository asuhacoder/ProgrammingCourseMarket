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

resource "aws_cloudwatch_log_group" "aws_route53_skhole" {
  provider = aws.us-east-1

  name              = "/aws/route53/${data.aws_route53_zone.host_domain.name}"
  retention_in_days = 30
}

data "aws_iam_policy_document" "route53-query-logging-policy" {
  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["arn:aws:logs:*:*:log-group:/aws/route53/*"]

    principals {
      identifiers = ["route53.amazonaws.com"]
      type        = "Service"
    }
  }
}

resource "aws_cloudwatch_log_resource_policy" "route53-query-logging-policy" {
  provider = aws.us-east-1

  policy_document = data.aws_iam_policy_document.route53-query-logging-policy.json
  policy_name     = "route53-query-logging-policy"
}

resource "aws_route53_query_log" "skhole" {
  depends_on = [aws_cloudwatch_log_resource_policy.route53-query-logging-policy]

  cloudwatch_log_group_arn = aws_cloudwatch_log_group.aws_route53_skhole.arn
  zone_id                  = data.aws_route53_zone.host_domain.zone_id
}
