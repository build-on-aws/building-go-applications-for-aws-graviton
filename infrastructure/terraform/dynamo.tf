resource "aws_dynamodb_table" "go-link-shortener" {
  name         = "goUrlShortener"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "shortURL"

  attribute {
    name = "shortURL"
    type = "S"
  }

  tags = {
    Name = "goUrlShortener"
  }
}
