variable "region" {
  default = "ap-south-1"
}

variable "app_image" {
  description = "Docker image"
}

variable "vpc_cidr" {
  default = "10.0.0.0/16"
}