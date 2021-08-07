terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.50"
    }
  }
  required_version = ">= 1.0.2"
}

provider "aws" {
  region  = var.aws_region
}

resource "aws_iot_thing" "growatt-to-iot" {
  name = "growatt-to-iot"
}

# Attach certificate to the Thing
resource "aws_iot_thing_principal_attachment" "att" {
  principal = var.aws_iot_certificate_arn
  thing     = aws_iot_thing.growatt-to-iot.name
}

resource "aws_iot_policy" "growatt-to-iot" {
  name = "growatt-to-iot"

  # We've connected the certificate to the Thing
  # See https://docs.aws.amazon.com/iot/latest/developerguide/thing-policy-examples.html
  policy = jsonencode({
    Version = "2012-10-17"
    "Statement":[
        {
            Effect:"Allow",
            Action:[
              "iot:Connect",
              "iot:Publish",
              "iot:Subscribe"
            ],
            Resource:[ "*" ],
            Condition: {
                Bool: {
                    "iot:Connection.Thing.IsAttached": ["true"]
                }
            }
        }
    ]
  })
}

# Attach certificate to Policy - now the cert is linked to the policy and the thing
resource "aws_iot_policy_attachment" "att" {
  policy = aws_iot_policy.growatt-to-iot.name
  target = var.aws_iot_certificate_arn
}