# Setup Name Vars

ROLE_NAME="go-link-shortener-ec2-role"
INSTANCE_PROFILE_NAME="go-link-shortener-ec2-role"
POLICY_NAME="go-link-shortener-dynamo-permissions"
TABLE_NAME="goUrlShortener"

## Create DynamoDB Table
echo "Creating DynamoDB Table..."
/usr/local/bin/aws dynamodb create-table --table-name $TABLE_NAME --key-schema AttributeName="shortURL",KeyType="HASH" --attribute-definitions AttributeName="shortURL",AttributeType="S" --billing-mode "PAY_PER_REQUEST" --no-cli-pager

## Create IAM Role
echo "Creating IAM Role..."
/usr/local/bin/aws iam create-role --role-name $ROLE_NAME --assume-role-policy-document file://./role-trust-policy.json --no-cli-pager

## Create Instance Profile
echo "Creating Instance Profile...."
/usr/local/bin/aws iam create-instance-profile --instance-profile-name $INSTANCE_PROFILE_NAME --no-cli-pager
/usr/local/bin/aws iam add-role-to-instance-profile --instance-profile-name $INSTANCE_PROFILE_NAME --role-name $ROLE_NAME --no-cli-pager

## Create IAM Policy
echo "Creating IAM Policy..."
/usr/local/bin/aws iam create-policy --policy-name $POLICY_NAME --policy-document file://./role-policy.json --no-cli-pager

## Get Policy ARN
POLICY_ARN=$(/usr/local/bin/aws iam list-policies --query 'Policies[?PolicyName==`go-link-shortener-dynamo-permissions`].Arn' --output text)

## Attach IAM Policy To Role
echo "Attaching IAM Policy to Role..."
/usr/local/bin/aws iam attach-role-policy --role-name $ROLE_NAME --policy-arn $POLICY_ARN --no-cli-pager