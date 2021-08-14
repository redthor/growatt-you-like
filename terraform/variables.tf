variable "aws_iot_certificate_arn" {
    description = "The ARN of the certificate to associate with AWS IOT core connections"
    type        = string
}

variable "aws_region" {
    description = "The AWS region"
    type        = string
    default     = "us-east-1"
}

variable "sumologic_time_zone" {
    description = "Your timezone for sumo logic triggers"
    type        = string
    default     = "AEST"
}

variable "sumologic_trigger_email" {
    description = "Email address to send monitoring alerts"
    type        = string
}