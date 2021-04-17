terraform {
  required_version = "0.15.0"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.36.0"
    }
  }
}

provider "aws" {
  region = "eu-west-2"
}

variable "bucket_name" {
  description = "The name of the bucket"
  default     = "-example"
}

resource "aws_s3_bucket" "terratest_bucket" {
  bucket = "terratest${var.bucket_name}"
  versioning {
    enabled = false
  }
}

output "bucket_id" {
  value = aws_s3_bucket.terratest_bucket.id
}
