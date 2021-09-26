terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.26.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.0.1"
    }
  }
  required_version = ">= 0.14"

  backend "remote" {
    organization = "mariam-demo"

    workspaces {
      name = "mariam-demo"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "oneinstance" {
  ami           = "ami-087c17d1fe0178315"
  instance_type = "t2.micro"
  key_name      = "flugeladmin"

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}

resource "aws_s3_bucket" "onebucket" {
  bucket = "buc955815154140"
  acl    = "public-read"

  versioning {
    enabled = true
  }

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}