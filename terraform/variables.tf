variable "aws_iot_certificate_arn" {
    description = "The ARN of the certificate to associate with AWS IOT core connections"
    type        = string
}

variable "aws_region" {
    description = "The AWS region"
    type        = string
    default     = "us-east-1"
}