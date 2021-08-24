variable "list" {
  description = "AWS SSM パラメータストアに登録する名前と値のセット"
  type        = map(string)
  default = {
    "DB/HOST "    = "xxxxx"
    "DB/USER"     = "xxxxx"
    "DB/PASSWORD" = "xxxxx"
    "DB/DATABASE" = "xxxxx"
  }
}