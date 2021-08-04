provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.49.0"
    }
  }
}

module "alb" {
  source                   = "./alb"
  vpc_network              = var.aws_vpc_network
  vpc_subnet_1a            = var.aws_vpc_subnet_1a
  vpc_subnet_1c            = var.aws_vpc_subnet_1c
  alb_front_security_group = var.aws_alb_front_security_group
  alb_back_security_group  = var.aws_alb_back_security_group
  acm_cert                 = var.aws_acm_cert
}
module "ecr" {
  source                = "./ecr"
  account_id            = var.aws_account_id
  front_repository_name = var.aws_ecr_front_repository_name
  front_image_name      = var.aws_ecr_front_image_name
  front_repository_uri  = var.aws_ecr_front_repository_uri
  back_repository_name  = var.aws_ecr_back_repository_name
  back_image_name       = var.aws_ecr_back_image_name
  back_repository_uri   = var.aws_ecr_back_repository_uri
}
module "ecs" {
  source                   = "./ecs"
  account_id               = var.aws_account_id
  db_user                  = var.aws_db_user
  db_password              = var.aws_db_password
  db_endpoint              = var.aws_db_endpoint
  db_database              = var.aws_db_database
  vpc_subnet_1a            = var.aws_vpc_subnet_1a
  vpc_subnet_1c            = var.aws_vpc_subnet_1c
  ecs_front_security_group = var.aws_ecs_front_security_group
  ecs_back_security_group  = var.aws_ecs_back_security_group
  ecs_front_target_group   = module.alb.alb-frontend-ecs-tg
  ecs_back_target_group    = module.alb.alb-backend-ecs-tg
}
module "route53" {
  source             = "./route53"
  zone_id            = var.aws_zone_id
  front_domain_name  = var.aws_front_domain_name
  alb_front_dns_name = module.alb.alb-frontend-dns-name
  alb_front_zone_id  = module.alb.alb-frontend-zone-id
  back_domain_name   = var.aws_back_domain_name
  alb_back_dns_name  = module.alb.alb-backend-dns-name
  alb_back_zone_id   = module.alb.alb-backend-zone-id
}