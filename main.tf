terraform {
    backend "s3" {
      bucket = "dc1testing-tfstate-s3"
      key = "tfstate/terraform.tfstate"
      region = "ap-southeast-2"
      encrypt = true
#      profile = "dc1testing"
    }
}

provider "aws" {
    region =  var.aws-region

#    profile = "dc1testing"
}


resource "aws_vpc" "tf-testing-vpc" {
    cidr_block = var.cidr-block
    tags = {
        vpc = "dc1testing-vpc"
    }
}

resource "aws_internet_gateway" "internet-gateway" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    tags = {
        internet-gateway = "dc1testing-internet-gateway"
    }
}

resource "aws_route_table" "vpc-route" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    route {
        cidr_block = "12.12.0.0/16"
        gateway_id = "local"
    }
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.internet-gateway.id
    }
}

resource "aws_route_table_association" "vpc-route-associate" {
    subnet_id = aws_subnet.testing-subnet.id
    route_table_id = aws_route_table.vpc-route.id
}

resource "aws_route_table_association" "vpc-route-associate-2" {
    subnet_id = aws_subnet.testing-subnet-2.id
    route_table_id = aws_route_table.vpc-route.id
}

resource "aws_security_group" "security-group" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    name = "securitygroup-testing"
    ingress {
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        from_port = 80
        to_port = 80
    }
    ingress {
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        from_port = 22
        to_port = 22
    }
    egress {
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        from_port = 0
        to_port = 65535
    }
    tags = {
        security-group = "dc1-security-group"
    }
    
}

resource "aws_vpc_security_group_ingress_rule" "ingress-allow-all" {
    security_group_id = aws_security_group.security-group.id
    ip_protocol = "-1"
    cidr_ipv4 = "0.0.0.0/0"
#    from_port = 80
#    to_port = 80
  
}

resource "aws_vpc_security_group_egress_rule" "egress-allow-all" {
    security_group_id = aws_security_group.security-group.id
    cidr_ipv4 = "0.0.0.0/0"
    ip_protocol = "-1"
}

resource "aws_subnet" "testing-subnet" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    cidr_block = "12.12.12.0/24"
    availability_zone = "ap-southeast-2a"
    tags = {
        subnet = "dc1-testing-subnet2a-web-server"
    }

}

resource "aws_subnet" "testing-subnet-2" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    cidr_block = "12.12.13.0/24"
    availability_zone = "ap-southeast-2b"
    tags = {
        subnet = "dc1-testing-subnet2b-web-server"
    }
  
}

resource "aws_eip" "elastic-ip" {
    domain = "vpc"
    instance = aws_instance.vm.id 
}

resource "aws_key_pair" "demoenvkey" {
  key_name   = "demoenvkey"
  public_key = var.public-key
}

resource "aws_instance" "vm" {
    subnet_id = aws_subnet.testing-subnet.id
    ami = var.ami
    instance_type = var.instance-type
    vpc_security_group_ids = [ aws_security_group.security-group.id ]
    key_name = aws_key_pair.demoenvkey.id
    user_data = "${file("install_nginx.sh")}"
    tags = {
        server = "dc1-webserver"
    }

}

resource "aws_db_subnet_group" "testing-dbsubnet" {
    subnet_ids = [ aws_subnet.testing-subnet.id, aws_subnet.testing-subnet-2.id ]
    name = "testing-db-subnet"
    tags = {
        dbsubnet = "testing-dbsubnet"
    }
}

resource "aws_db_instance" "database" {
    db_subnet_group_name = aws_db_subnet_group.testing-dbsubnet.id
    allocated_storage = 10
    db_name = "dbtesting"
    engine = "mysql"
    engine_version = "8.0"
    instance_class = "db.t3.micro"
    username = "dc1testing"
    password = "dc1testing"
    parameter_group_name = "default.mysql8.0"
    skip_final_snapshot = true
    vpc_security_group_ids = [ aws_security_group.security-group.id ]
    tags = {
        database = "dc1-database-testing"
    }

}

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "lambda.js"
  output_path = "lambda_function_payload.zip"
}

resource "aws_lambda_function" "lambda_function" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
    filename      = "lambda_function_payload.zip"
    function_name = "lambda_function_name"
    role          = "arn:aws:iam::236983045839:role/lambda-role"

    handler       = "index.test"

    source_code_hash = data.archive_file.lambda.output_base64sha256

    runtime = "nodejs20.x"

    environment {
      variables = {
        foo = "test"
      }
    }
}

//resource "aws_internet_gateway_attachment" "internet-gateway-attachment" {
//    internet_gateway_id = aws_internet_gateway.internet-gateway.id
//    vpc_id = aws_vpc.tf-testing-vpc.id
//}
