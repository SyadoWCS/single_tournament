variable aws_access_key {}
variable aws_secret_key {}
variable aws_region {}
variable aws_vpc_cidr {}
variable aws_vpc_tags_name {}

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

output "single-tournament-vpc" {
  value = "${module.vpc.vpc_id}"
}