provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "3.49.0"
    }
  }
}

module "ecr" {
  source   = "./ecr"
  account_id = var.aws_account_id
  front_repository_name = var.aws_ecr_front_repository_name
  front_image_name = var.aws_ecr_front_image_name
  front_repository_uri = var.aws_ecr_front_repository_uri
  back_repository_name = var.aws_ecr_back_repository_name
  back_image_name = var.aws_ecr_back_image_name
  back_repository_uri = var.aws_ecr_back_repository_uri
}