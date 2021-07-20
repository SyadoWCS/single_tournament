# aws ecr frontend repository
output "frontend-repository-url" {
    description = "ecr of frontend repository url"
    value       = "${aws_ecr_repository.single-tournament-frontend.repository_url}"
}
# aws ecr backend repository
output "backend-repository-url" {
    description = "ecr of backend repository url"
    value       = "${aws_ecr_repository.single-tournament-backend.repository_url}"
}