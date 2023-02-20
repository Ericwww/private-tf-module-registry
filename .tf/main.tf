terraform {
  required_providers {
    echo = {
      source = "echo/echo"
      version = "0.0.1"
    }
  }
}

module "consul_aws" {
  source = "127.0.0.1:6966/hashicorp/consul/aws"
  version = "1.0.0"
}

resource "echo" "echo1" {
  create_failed = true
}