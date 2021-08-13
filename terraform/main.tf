terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.50"
    }
  }
  required_version = ">= 1.0.2"
}

provider "aws" {
  region  = var.aws_region
}