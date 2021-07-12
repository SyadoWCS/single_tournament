variable "aws_access_key" {}
variable "aws_secret_key" {}

provider "aws" {
  region     = "ap-northeast-1"
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

resource "aws_vpc" "myVPC" {
    cidr_block = "192.168.1.0/24"
    instance_tenancy = "default"
    enable_dns_support = "true"
    enable_dns_hostnames = "false"
}