resource "aws_secretsmanager_secret" "jdoodle_client_id" {
  name                           = "jdoodle-secrets-client-id"
  recovery_window_in_days        = 0
  force_overwrite_replica_secret = true
}

resource "aws_secretsmanager_secret_version" "jdoodle_client_id" {
  secret_id     = aws_secretsmanager_secret.jdoodle_client_id.id
  secret_string = var.jdoodle_client_id
}

resource "aws_secretsmanager_secret" "jdoodle_client_secret" {
  name                           = "jdoodle-secrets-client-secret"
  recovery_window_in_days        = 0
  force_overwrite_replica_secret = true
}

resource "aws_secretsmanager_secret_version" "jdoodle_client_secret" {
  secret_id     = aws_secretsmanager_secret.jdoodle_client_secret.id
  secret_string = var.jdoodle_client_secret
}