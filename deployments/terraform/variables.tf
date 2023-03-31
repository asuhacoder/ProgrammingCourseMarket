variable "product_name" {
  type    = string
  default = "skhole"
}

variable "region" {
  type    = string
  default = "us-east-1"
}

variable "user_image_name" {
  type    = string
  default = "skhole-user"
}

variable "user_docker_path" {
  type    = string
  default = "../../build/package/docker/user/Dockerfile"
}

variable "script_path" {
  type    = string
  default = "../../scripts/push-image.sh"
}

variable "serveice_names" {
  type    = set(string)
  default = ["user", "auth", "gateway", "course", "lesson", "test_case", "runner"]
}

variable "app_environment" {
  type    = string
  default = "production"
}

variable "private_subnet" {
  type    = string
  default = "10.10.0.0/24"
}

variable "private_subnet_dummy" {
  type    = string
  default = "10.10.10.0/24"
}

variable "public_subnet" {
  type    = string
  default = "10.10.100.0/24"
}

variable "public_dummy_subnet" {
  type    = string
  default = "10.10.90.0/24"
}

variable "availability_zone" {
  type    = string
  default = "us-east-1a"
}

variable "availability_zone_dummy" {
  type    = string
  default = "us-east-1b"
}

variable "ecs_cpu" {
  type    = number
  default = 30
}

variable "ecs_memory" {
  type    = number
  default = 512
}

variable "jdoodle_client_id" {
  type = string
}

variable "jdoodle_client_secret" {
  type = string
}

variable "db_engine" {
  type    = string
  default = "postgres"
}