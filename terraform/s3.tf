resource "aws_s3_bucket" "growatt-to-iot" {
  bucket = "growatt-to-iot"
  acl    = "private"
}