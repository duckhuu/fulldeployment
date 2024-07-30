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

resource "aws_security_group" "security-group" {
    vpc_id = aws_vpc.tf-testing-vpc.id
    name = "securitygroup-testing"
    ingress {
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        from_port = 80
        to_port = 80
    }
    egress {
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        from_port = 80
        to_port = 80
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

resource "aws_instance" "vm" {
    subnet_id = aws_subnet.testing-subnet.id
    ami = var.ami
    instance_type = var.instance-type
    vpc_security_group_ids = [ aws_security_group.security-group.id ]
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



//resource "aws_internet_gateway_attachment" "internet-gateway-attachment" {
//    internet_gateway_id = aws_internet_gateway.internet-gateway.id
//    vpc_id = aws_vpc.tf-testing-vpc.id
//}
