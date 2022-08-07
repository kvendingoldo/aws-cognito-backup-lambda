resource "aws_iam_role" "main" {
  name               = var.blank_name
  description        = "IAM role for for Lambda ${var.blank_name}"
  tags               = var.tags
  assume_role_policy = <<-POLICY
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  POLICY
}

#
# VPC permissions
#
resource "aws_iam_role_policy_attachment" "vpc_permissions" {
  count      = length(var.subnet_ids) != 0 ? 1 : 0
  role       = aws_iam_role.main.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

#
# Logging policy
#
resource "aws_iam_policy" "logging" {
  name        = format("%s-%s", var.blank_name, "logging")
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<-POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
POLICY
}
resource "aws_iam_role_policy_attachment" "logging" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.logging.arn
}

#
# ACM policy
#
resource "aws_iam_policy" "cognito" {
  name        = format("%s-%s", var.blank_name, "cognito")
  path        = "/"
  description = "IAM policy for working with Cognito from a lambda"

  policy = <<-POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "cognito-identity:*",
        "cognito-idp:*",
        "cognito-sync:*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
POLICY
}
resource "aws_iam_role_policy_attachment" "cognito" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.cognito.arn
}

#
# Route53 policy
#
resource "aws_iam_policy" "s3" {
  name        = format("%s-%s", var.blank_name, "s3")
  path        = "/"
  description = "IAM policy for working with S3 from a lambda"

  policy = <<-POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
POLICY
}
resource "aws_iam_role_policy_attachment" "route53" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.s3.arn
}


