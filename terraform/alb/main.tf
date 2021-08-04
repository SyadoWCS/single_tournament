// ALB作成(フロントエンド)
resource "aws_lb" "single-tournament-frontend-alb" {
  name                       = "single-tournament-frontend-alb"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = false

  subnets = [
    var.vpc_subnet_1a,
    var.vpc_subnet_1c
  ]

  security_groups = [
    var.alb_front_security_group
  ]

  tags = {
    Name = "single-tournament-frontend-alb"
  }
}
// ターゲットグループ作成(フロントエンド)
resource "aws_lb_target_group" "single-tournament-front-ecs-tg" {
    name        = "single-tournament-front-ecs-tg"
    port        = 80
    protocol    = "HTTP"
    target_type = "ip"
    vpc_id      = var.vpc_network

    health_check {
        path = "/"
        port = 80
        protocol = "HTTP"
    }
}
// ALBリスナー作成(フロントエンド)
resource "aws_lb_listener" "single-tournament-front-alb-listener" {
    # ルールを追加するリスナー
    load_balancer_arn = aws_lb.single-tournament-frontend-alb.arn
    port              = "80"
    protocol          = "HTTP"

    default_action {
        target_group_arn = aws_lb_target_group.single-tournament-front-ecs-tg.arn
        type             = "forward"
    }
}
// ALB作成(バックエンド)
resource "aws_lb" "single-tournament-backend-alb" {
  name                       = "single-tournament-backend-alb"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = false

  subnets = [
    var.vpc_subnet_1a,
    var.vpc_subnet_1c
  ]

  security_groups = [
    var.alb_back_security_group
  ]

  tags = {
    Name = "single-tournament-backend-alb"
  }
}
// ターゲットグループ作成(バックエンド)
resource "aws_lb_target_group" "single-tournament-back-ecs-tg" {
    name        = "single-tournament-back-ecs-tg"
    port        = 80
    protocol    = "HTTP"
    target_type = "ip"
    vpc_id      = var.vpc_network

    health_check {
        path = "/"
        port = 80
        protocol = "HTTP"
    }
}
// ALBリスナー作成(バックエンド)
resource "aws_lb_listener" "single-tournament-back-alb-listener" {
    # ルールを追加するリスナー
    load_balancer_arn = aws_lb.single-tournament-backend-alb.arn
    port              = "80"
    protocol          = "HTTP"

    default_action {
        target_group_arn = aws_lb_target_group.single-tournament-back-ecs-tg.arn
        type             = "forward"
    }
}