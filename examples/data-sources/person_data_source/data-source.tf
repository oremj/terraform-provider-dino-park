data "dinopark_person" "example" {
  email = "oremj@mozilla.com"
}


terraform {
  required_providers {
    dinopark = {
      version = "~> 0.1"
      source  = "github.com/oremj/dinopark"
    }
  }
}
