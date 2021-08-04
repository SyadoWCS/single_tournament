output "alb-frontend-dns-name" {
  description = "フロントエンドのALBのDNS名"
  value       = aws_lb.single-tournament-frontend-alb.dns_name
}
output "alb-frontend-zone-id" {
  description = "フロントエンドのALBのゾーンID"
  value       = aws_lb.single-tournament-frontend-alb.zone_id
}
output "alb-frontend-ecs-tg" {
  description = "フロントエンドのALBのターゲットグループ"
  value       = aws_lb_target_group.single-tournament-front-ecs-tg.arn
}
output "alb-backend-dns-name" {
  description = "バックエンドのALBのDNS名"
  value       = aws_lb.single-tournament-backend-alb.dns_name
}
output "alb-backend-zone-id" {
  description = "バックエンドのALBのゾーンID"
  value       = aws_lb.single-tournament-backend-alb.zone_id
}
output "alb-backend-ecs-tg" {
  description = "バックエンドのALBのターゲットグループ"
  value       = aws_lb_target_group.single-tournament-back-ecs-tg.arn
}