// クラスタ作成
resource "aws_ecs_cluster" "single-tournament-ecs-cluster" {
  name               = "single-tournament-from-ecs-cluster"
  capacity_providers = ["FARGATE"]

  tags = {
    "Name" = "single-tournament-from-ecs-cluster"
  }
}
// タスク定義作成(フロントエンド)
resource "aws_ecs_task_definition" "single-tournament-frontend-task-definition" {
  family                   = "single-tournament-frontend-from-ecs-task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = "arn:aws:iam::${var.account_id}:role/ecsTaskExecutionRole"
  task_role_arn            = "arn:aws:iam::${var.account_id}:role/ecsTaskExecutionRole"
  container_definitions    = data.template_file.container_definitions_frontend.rendered

  tags = {
    "Name" = "single-tournament-frontend-from-ecs-task"
  }
}
// コンテナの定義ファイル(フロントエンド)
data "template_file" "container_definitions_frontend" {
  template = file("container_definitions/single_tournament_frontend_container_definitions.json")

  vars = {
    account_id = var.account_id
  }
}

// サービス作成(フロントエンド)
resource "aws_ecs_service" "single-tournament-frontend-service" {
  name                              = "single-tournament-frontend-from-ecs-service"
  cluster                           = aws_ecs_cluster.single-tournament-ecs-cluster.arn
  task_definition                   = aws_ecs_task_definition.single-tournament-frontend-task-definition.arn
  desired_count                     = 1
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  health_check_grace_period_seconds = 100
  deployment_maximum_percent        = 200

  network_configuration {
    assign_public_ip = true
    security_groups = [
      var.ecs_front_security_group
    ]
    subnets = [
      var.vpc_subnet_1a,
      var.vpc_subnet_1c,
    ]
  }

  load_balancer {
    target_group_arn = var.ecs_front_target_group
    container_name   = "single-tournament-frontend-container"
    container_port   = 8080
  }

  tags = {
    "Name" = "single-tournament-frontend-from-ecs-task"
  }
}

// タスク定義作成(バックエンド)
resource "aws_ecs_task_definition" "single-tournament-backend-task-definition" {
  family                   = "single-tournament-backend-from-ecs-task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = "arn:aws:iam::${var.account_id}:role/ecsTaskExecutionRole"
  task_role_arn            = "arn:aws:iam::${var.account_id}:role/ecsTaskExecutionRole"
  container_definitions    = data.template_file.container_definitions_backend.rendered


  tags = {
    "Name" = "single-tournament-backend-from-ecs-task"
  }
}
// コンテナの定義ファイル(バックエンド)
data "template_file" "container_definitions_backend" {
  template = file("container_definitions/single_tournament_backend_container_definitions.json")

  vars = {
    account_id  = var.account_id,
    db_user     = var.db_user,
    db_password = var.db_password,
    db_endpoint = var.db_endpoint,
    db_database = var.db_database,
  }
}
// サービス作成(バックエンド)
resource "aws_ecs_service" "single-tournament-backend-service" {
  name                              = "single-tournament-backend-from-ecs-service"
  cluster                           = aws_ecs_cluster.single-tournament-ecs-cluster.arn
  task_definition                   = aws_ecs_task_definition.single-tournament-backend-task-definition.arn
  desired_count                     = 1
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  health_check_grace_period_seconds = 100
  deployment_maximum_percent        = 200

  network_configuration {
    assign_public_ip = true
    security_groups = [
      var.ecs_back_security_group
    ]
    subnets = [
      var.vpc_subnet_1a,
      var.vpc_subnet_1c,
    ]
  }

  load_balancer {
    target_group_arn = var.ecs_back_target_group
    container_name   = "single-tournament-backend-container"
    container_port   = 80
  }

  tags = {
    "Name" = "single-tournament-backend-from-ecs-task"
  }
}