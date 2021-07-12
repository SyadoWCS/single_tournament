resource "aws_vpc" "single-tournament" {
    cidr_block = var.vpc["default.cidr"]
    instance_tenancy = "default"
    enable_dns_support = "true"
    enable_dns_hostnames = "false"
}