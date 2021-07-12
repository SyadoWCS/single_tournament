variable "vpc" {
    type = map(string)

    default = {
        "default.cidr" = "192.168.1.0/24"
    }
}