terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 3.50"
    }
    sumologic = {
      source = "SumoLogic/sumologic"
      version = "2.9.8"
    }
  }
}

provider "aws" {
  region  = var.aws_region
}