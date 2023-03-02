resource "aws_s3_bucket" "growatt-to-iot" {
  bucket = "growatt-to-iot"
}

resource "aws_s3_bucket_acl" "growatt-to-iot_acl" {
  bucket = aws_s3_bucket.growatt-to-iot.id
  acl    = "private"
}