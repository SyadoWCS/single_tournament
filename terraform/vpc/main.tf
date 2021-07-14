# VPC
resource "aws_vpc" "vpc" {
    cidr_block = var.vpc_cidr
    instance_tenancy = "default"
    enable_dns_support = "true"
    enable_dns_hostnames = "false"
    tags = {
        Name = var.vpc_tags_name
    }
}
# Subnet(public_1a)
/*resource "aws_subnet" "public_1a" {
    vpc_id = aws_vpc.vpc.id
    cidr_block = lookup(
        var.vpc,
        "${terraform.workspace}.public_1a",
        var.vpc["default.public_1a"]
    )
    availability_zone = var.vpc["default.az_1a"]

    tags = {
        Name = "${terraform.workspace}-public-1a"
    }
}*/
# Subnet(public_1c)
/*resource "aws_subnet" "public_1c" {
    vpc_id = aws_vpc.vpc.id
    cidr_block = lookup(
        var.vpc,
        "${terraform.workspace}.public_1c",
        var.vpc["default.public_1c"]
    )
    availability_zone = var.vpc["default.az_1c"]

    tags = {
        Name = "${terraform.workspace}-public-1c"
    }
}*/
# Igw
/*resource "aws_route_table" "public" {
    vpc_id = aws_vpc.vpc.id

    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.igw.id
    }

    tags = {
        Name = "${terraform.workspace}-public-rt"
    }
}*/
# Route Table(public_1a)
/*resource "aws_route_table_association" "public_1a" {
    route_table_id = aws_route_table.public.id
    subnet_id      = aws_subnet.public_1a.id
}*/
# Route Table(public_1c)
/*resource "aws_route_table_association" "public_1c" {
    route_table_id = aws_route_table.public.id
    subnet_id      = aws_subnet.public_1c.id
}*/
