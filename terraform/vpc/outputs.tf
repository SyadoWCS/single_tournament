# vpc_id
output "vpc_id" {
    description = "ID of project VPC"
    value       = "${aws_vpc.vpc.id}"
}
# public_subnets
/*output "public_subnets" {
    description = "List of IDs of public subnets"
    value       = "${aws_vpc.vpc.public_subnets}"
}*/
# public_subnet_1a
/*output "subnet_public_1a_id" {
    description = "List of IDs of public subnets 1a"
    value       = "${aws_subnet.public_1a.id}"
}*/
# public_subnet_1c
/*output "subnet_public_1c_id" {
    description = "List of IDs of public subnet 1c"
    value       = "${aws_subnet.public_1c.id}"
}*/