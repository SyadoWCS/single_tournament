# ECR create repository(frontend)
resource "aws_ecr_repository" "single-tournament-frontend" {
    name                 = var.front_repository_name
    image_tag_mutability = "MUTABLE"

    image_scanning_configuration {
        scan_on_push = true
    }
}

# ECR create repository(backend)
resource "aws_ecr_repository" "single-tournament-backend" {
    name                 = var.back_repository_name
    image_tag_mutability = "MUTABLE"

    image_scanning_configuration {
        scan_on_push = true
    }
}

# ECR docker image push(backend)
resource "null_resource" "single-tournament-image-push" {
    provisioner "local-exec" {
        command = "aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${var.account_id}.dkr.ecr.ap-northeast-1.amazonaws.com"
    }

    provisioner "local-exec" {
        command = "docker-compose build frontend"
    }

    /*provisioner "local-exec" {
        command = "docker-compose build backend"
    }*/

    provisioner "local-exec" {
        command = "docker tag ${var.front_image_name}:latest ${aws_ecr_repository.single-tournament-frontend.repository_url}"
    }

    /*provisioner "local-exec" {
        command = "docker tag ${var.back_image_name}:latest ${aws_ecr_repository.single-tournament-backend.repository_url}"
    }*/

    provisioner "local-exec" {
        command = "docker push ${aws_ecr_repository.single-tournament-frontend.repository_url}"
    }

    /*provisioner "local-exec" {
        command = "docker push ${aws_ecr_repository.single-tournament-backend.repository_url}"
    }*/
}