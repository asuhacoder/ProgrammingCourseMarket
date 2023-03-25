output "s3-name" {
  description = "S3 bucket name for React app"
  value       = aws_s3_bucket.bucket.id
}