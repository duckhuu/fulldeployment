variable "aws-region" {
    default = "ap-southeast-2"
}

variable "cidr-block" {
    default = "12.12.0.0/16"
}

variable "ami" {
    default = "ami-03f0544597f43a91d"
}

variable instance-type {
    default = "t3.micro"
}

variable public-key {
    default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDS5GQzu89D2Fncs6igs1GzhfgXaqnAXaF3j71Nmf0VG66jnXyQh4wbFL9XwWsZ/rDUPQmq2NI7MaElvfadNpPolTpyOiUIKHXrQA1JyYtNfaTmWE/kxoGw95n+oJk54GfEKOgyxy1UjRkBFnBE4UbTXO5LaRxDMgPeZdYV91+Uk5cvW8rDwjgXgiJOFiEZCadVN190qckW+Wmf1w4FOCyPpGt91jagGEU43b3/vDTKLPBnU4vPbWVM8zuEWFqt42kxpAs0Jyj1+L7LaUmjfO7JuESbwK6r3vrxKU4milkvbtVNVnL4j4lZcJji1rf6WsnPTkhDVVbEZWxWNnd/b1gX rsa-key-20240802"
}