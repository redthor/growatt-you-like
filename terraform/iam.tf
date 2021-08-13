resource "aws_iam_role" "growatt-to-iot-assume-role" {
  name = "growatt-to-iot-assume-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement":[{
      "Effect": "Allow",
      "Principal": {
        "Service": "iot.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
  }]
}
EOF
}

resource "aws_iam_policy" "growatt-to-iot-s3-write-only-policy" {
  name        = "growatt-to-iot-s3-write-only-policy"
  description = "IAM policy for growatt-to-iot to access S3"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:ListBucket"
      ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::${aws_s3_bucket.growatt-to-iot.id}"
      ]
    },
    {
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::${aws_s3_bucket.growatt-to-iot.id}/*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "grant-s3" {
  role       = aws_iam_role.growatt-to-iot-assume-role.name
  policy_arn = aws_iam_policy.growatt-to-iot-s3-write-only-policy.arn
}

resource "aws_iam_policy" "growatt-to-iot-cloudwatch-policy" {
  name        = "growatt-to-iot-cloudwatch-policy"
  description = "IAM policy for growatt-to-iot to write logs to Cloudwatch"

  policy = <<EOF
{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                    "logs:PutMetricFilter",
                    "logs:PutRetentionPolicy"
                 ],
                "Resource": [
                    "*"
                ]
            }
        ]
    }
EOF
}

resource "aws_iam_role_policy_attachment" "grant-cw" {
  role       = aws_iam_role.growatt-to-iot-assume-role.name
  policy_arn = aws_iam_policy.growatt-to-iot-cloudwatch-policy.arn
}