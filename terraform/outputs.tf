output "instance_id" {
  value = aws_instance.oneinstance.id
}

output "bucket_id" {
  value = aws_s3_bucket.onebucket.id
}