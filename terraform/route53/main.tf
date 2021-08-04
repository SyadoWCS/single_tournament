// ドメイン紐付け(フロントエンド)
resource "aws_route53_record" "single-tournament-front-zone-record" {
  zone_id = var.zone_id
  name = var.front_domain_name
  type = "A"

  alias {
    name = var.alb_front_dns_name
    zone_id = var.alb_front_zone_id
    evaluate_target_health = true
  }
}
// ドメイン紐付け(バックエンド)
resource "aws_route53_record" "single-tournament-back-zone-record" {
  zone_id = var.zone_id
  name = var.back_domain_name
  type = "A"

  alias {
    name = var.alb_back_dns_name
    zone_id = var.alb_back_zone_id
    evaluate_target_health = true
  }
}