terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.59.0"
    }
  }

  backend "s3" {
    bucket = "skhole-terraform"
    key    = "terraform.tfstate"
    region = "us-east-2"
  }
}

locals {
  app_name    = var.product_name
  host_domain = "${var.product_name}.club"
}

provider "aws" {
  region = var.region
}