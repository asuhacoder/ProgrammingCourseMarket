resource "aws_secretsmanager_secret" "jdoodle" {
  name                           = "jdoodle-secrets"
  recovery_window_in_days        = 0
  force_overwrite_replica_secret = true
}

resource "aws_secretsmanager_secret_version" "jdoodle" {
  secret_id = aws_secretsmanager_secret.jdoodle.id
  secret_string = jsonencode({
    jdoodle_client_id     = var.jdoodle_client_id,
    jdoodle_client_secret = var.jdoodle_client_secret
  })
}
