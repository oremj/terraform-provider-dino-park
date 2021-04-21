data "dinopark_group" "example" {
}


terraform {
  required_providers {
    dinopark = {
      version = "~> 0.1"
      source  = "hashicorp.com/oremj/dino-park"
    }
  }
}
