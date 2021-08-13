data "aws_iot_endpoint" "growatt-to-iot" {
  endpoint_type = "iot:Data-ATS"
}

output "aws_iot_endpoint" {
  value = data.aws_iot_endpoint.growatt-to-iot.endpoint_address
}