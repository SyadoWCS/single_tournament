output "vpc" {
    value = {
        "vpc_id" = aws_vpc.vpc.id
    }
}