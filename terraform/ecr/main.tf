# ECR create repository(frontend)
/*resource "aws_ecr_repository" "single-tournament-frontend" {
    name                 = ${var.front_image_name}
    image_tag_mutability = "MUTABLE"

    image_scanning_configuration {
        scan_on_push = true
    }
}*/

# ECR docker image push(frontend)
resource "null_resource" "single-tournament-frontend_image_push" {
    provisioner "local-exec" {
        command = "$(aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${var.repository_url})"
    }

    provisioner "local-exec" {
        command = "docker build -t ${var.front_image_name} ."
    }

    provisioner "local-exec" {
        command = "docker tag ${var.front_image_name}:latest ${var.front_repository_uri}:latest"
        //command = "docker tag ${var.image_name}:latest ${aws_ecr_repository.single-tournament-frontend.repository_url}"
    }

    provisioner "local-exec" {
        command = "docker push ${var.front_repository_uri}:latest"
        //command = "docker push ${aws_ecr_repository.single-tournament-frontend.repository_url}"
    }
}

# ECR create repository(backend)
/*resource "aws_ecr_repository" "single-tournament-backend" {
    name                 = ${var.back_image_name}
    image_tag_mutability = "MUTABLE"

    image_scanning_configuration {
        scan_on_push = true
    }
}*/

# ECR docker image push(backend)
resource "null_resource" "single-tournament-backend_image_push" {
    provisioner "local-exec" {
        command = "$(aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${var.repository_url})"
    }

    provisioner "local-exec" {
        command = "docker build -t ${var.back_image_name} ."
    }

    provisioner "local-exec" {
        command = "docker tag ${var.back_image_name}:latest ${var.back_repository_uri}:latest"
        //command = "docker tag ${var.image_name}:latest ${aws_ecr_repository.single-tournament-backend.repository_url}"
    }

    provisioner "local-exec" {
        command = "docker push ${var.back_repository_uri}:latest"
        //command = "docker push ${aws_ecr_repository.single-tournament-backend.repository_url}"
    }
}