resource "aws_iot_topic_rule" "growatt-to-iot-to-s3-rule" {
  name        = "growatt_to_iot_to_s3_rule"
  description = "Send growatt data into S3"
  enabled     = true
  sql         = "SELECT * FROM '+/inbound-raw'"
  sql_version = "2016-03-23"

  # TODO - is timestamp() in milliseconds? Or will we possibly overwrite files?
  s3 {
    bucket_name = aws_s3_bucket.growatt-to-iot.id
    key = "$${topic()}/$${timestamp()}.growatt.bin"
    role_arn = aws_iam_role.growatt-to-iot-assume-role.arn
  }
}