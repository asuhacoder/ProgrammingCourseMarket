output "s3-name" {
  description = "S3 bucket name for React app"
  value       = aws_s3_bucket.bucket.id
}

output "s3-bucket-policy" {
  description = "S3 bucket policy"
  value       = data.aws_iam_policy_document.allow_cloudfront_service_principal_s3_readonly.json
}

output "region" {
  description = "AWS region"
  value       = var.region
}

output "ecr-user-url" {
  description = "user ECR repository url"
  value       = aws_ecr_repository.user.repository_url
}

output "ecr-auth-url" {
  description = "auth ECR repository url"
  value       = aws_ecr_repository.auth.repository_url
}

output "ecr-course-url" {
  description = "course ECR repository url"
  value       = aws_ecr_repository.course.repository_url
}

output "ecr-lesson-url" {
  description = "lesson ECR repository url"
  value       = aws_ecr_repository.lesson.repository_url
}

output "ecr-test-case-url" {
  description = "test_case ECR repository url"
  value       = aws_ecr_repository.test_case.repository_url
}

output "ecr-gateway-url" {
  description = "gateway ECR repository url"
  value       = aws_ecr_repository.gateway.repository_url
}

output "ecr-runner-url" {
  description = "runner ECR repository url"
  value       = aws_ecr_repository.runner.repository_url
}
