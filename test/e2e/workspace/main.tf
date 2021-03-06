terraform {
  backend "gcs" {
    bucket = "automatize-tfstate"
    prefix = "e2e"
  }
}

variable "suffix" {}

resource "random_id" "test" {
  byte_length = 2
}

output "random_string" {
  value = "${random_id.test.hex}-${var.suffix}"
}
