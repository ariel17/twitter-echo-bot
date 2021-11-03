terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.63.0"
    }
  }

  required_version = "~> 1.0.5"
}

provider "aws" {
  profile = "ariel17"
  region  = "us-east-2"
}

resource "aws_instance" "bot_server" {
  ami           = "ami-0b59bfac6be064b78" # See available AMIs here: https://aws.amazon.com/amazon-linux-ami/
  instance_type = "t2.nano"
  count         = 1
  key_name      = "bot-server"

  provisioner "file" {
    source      = "../.env"
    destination = "/root/.env"

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file("~/.ssh/bot-server.pem")
      host        = self.public_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install docker -y",
      "sudo service docker start",
      "sudo docker run --env-file /root/.env -d ariel17/twitter-echo-bot:latest"
    ]

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file("~/.ssh/bot-server.pem")
      host        = self.public_ip
    }
  }

  tags = {
    name = "bot-server"
  }
}
