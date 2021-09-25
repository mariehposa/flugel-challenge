terraform {
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
    ami = "ami-087c17d1fe0178315"
    instance_type = "t2.micro"
    key_name = "flugeladmin"

    tags = {
        Name = "Flugel"
        Owner = "InfraTeam"
    }
}

resource "aws_s3_bucket" "onebucket" {
    bucket = "buc955815154140"
    acl = "public-read"

    versioning {
      enabled = true
   }

    tags = {
        Name = "Flugel"
        Owner = "InfraTeam"
    }
}