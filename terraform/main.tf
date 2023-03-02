terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.0"
    }
    sumologic = {
      source = "SumoLogic/sumologic"
      version = "~> 2.21"
    }
  }
}

provider "aws" {
  region  = var.aws_region
}