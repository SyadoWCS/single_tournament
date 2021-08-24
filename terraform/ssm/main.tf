resource "aws_ssm_parameter" "list" {
  for_each = var.list

  name  = "${each.key}"
  value = each.value
  type  = "String"
}