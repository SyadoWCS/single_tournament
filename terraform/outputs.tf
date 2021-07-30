# aws ecr frontend-repository-url
output "aws-ecr-frontend-repository-url" {
  value = module.ecr.frontend-repository-url
}
# aws ecr backend-repository-url
output "aws-ecr-backend-repository-url" {
  value = module.ecr.backend-repository-url
}