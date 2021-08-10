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
        "s3:ListBucket",
        "s3:PutObject"
    ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::${aws_s3_bucket.growatt-to-iot.id}"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "grant" {
  role       = aws_iam_role.growatt-to-iot-assume-role.name
  policy_arn = aws_iam_policy.growatt-to-iot-s3-write-only-policy.arn
}