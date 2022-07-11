resource "aws_dynamodb_table" "fghub_tf_state_lock" {
  name           = "fghub_tf_state"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}