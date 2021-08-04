# aws ecr frontend-repository-url
output "aws-ecr-frontend-repository-url" {
  value = module.ecr.frontend-repository-url
}
# aws ecr backend-repository-url
output "aws-ecr-backend-repository-url" {
  value = module.ecr.backend-repository-url
}
output "aws-alb-frontend-dns-name" {
  value = module.alb.alb-frontend-dns-name
}
output "aws-alb-frontend-zone-id" {
  value = module.alb.alb-frontend-zone-id
}
output "aws-alb-frontend-ecs-tg" {
  value = module.alb.alb-frontend-ecs-tg
}
output "aws-alb-backend-dns-name" {
  value = module.alb.alb-backend-dns-name
}
output "aws-alb-backend-zone-id" {
  value = module.alb.alb-backend-zone-id
}
output "aws-alb-backend-ecs-tg" {
  value = module.alb.alb-backend-ecs-tg
}