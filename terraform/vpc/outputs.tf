output "vpc" {
    value = {
        "vpc_id" = aws_vpc.single-tournament.id
    }
}