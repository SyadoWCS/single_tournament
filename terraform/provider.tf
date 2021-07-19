variable aws_access_key {}
variable aws_secret_key {}
variable aws_region {}
variable aws_vpc_cidr {}
variable aws_vpc_tags_name {}
variable aws_ecr_repository_url{}
variable aws_ecr_front_image_name {}
variable aws_ecr_front_repository_uri {}
variable aws_ecr_back_image_name {}
variable aws_ecr_back_repository_uri {}
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
module "vpc" {
  source   = "./vpc"
  vpc_cidr = var.aws_vpc_cidr
  vpc_tags_name = var.aws_vpc_tags_name
}

module "ecr" {
  source   = "./ecr"
  repository_url = var.aws_ecr_repository_url
  front_image_name = var.aws_ecr_front_image_name
  front_repository_uri = var.aws_ecr_front_repository_uri
  back_image_name = var.aws_ecr_back_image_name
  back_repository_uri = var.aws_ecr_back_repository_uri
}

module "ecs-fargate-frontend" {
  source   = "./frontend"
}

module "ecs-fargate-backend" {
  source   = "./backend"
}

output "single-tournament-vpc" {
  value = "${module.vpc.vpc_id}"
}