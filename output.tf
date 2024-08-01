output "eip-public-ip" {
    value = aws_eip.elastic-ip.public-ip
}

output "eip-public-dns" {
    value = aws_eip.elastic-ip.public-dns
}
