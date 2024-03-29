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

resource "aws_iot_thing_group" "growatt-to-iot-group" {
  name = "growatt-to-iot-group"
}

resource "aws_iot_thing_group_membership" "growatt-to-iot-group-membership" {
  thing_name       = aws_iot_thing.growatt-to-iot.name
  thing_group_name = aws_iot_thing_group.growatt-to-iot-group.name
}

resource "aws_iot_logging_options" "growatt-to-iot-logging" {
  default_log_level = "WARN"
  role_arn          = aws_iam_role.growatt-to-iot-assume-role.arn
}