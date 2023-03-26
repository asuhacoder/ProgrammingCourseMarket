output "s3-name" {
  description = "S3 bucket name for React app"
  value       = aws_s3_bucket.bucket.id
}

output "s3-bucket-policy" {
  description = "S3 bucket policy"
  value       = data.aws_iam_policy_document.allow_cloudfront_service_principal_s3_readonly.json
}
