variable "product_name" {
  type    = string
  default = "skhole"
}

variable "region" {
  type    = string
  default = "us-east-2"
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
  type = set(string)
  default = ["user", "auth", "gateway", "course", "lesson", "test_case", "runner"]
}