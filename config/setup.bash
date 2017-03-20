#!/bin/bash

# Setup snippets in part from:
# https://rossfairbanks.com/2015/03/31/hello-world-in-ec2-container-service.html

set -e
set -x

AMI_ID='ami-af8b30cf'
AWS_ACCOUNT_ID=''

# Need to accept agreenment for these images before you can launch them,
# see http://aws.amazon.com/marketplace/pp?sku=4jvb72q6a56js2x7jzd24jar5
get_most_recent_container_image() {
    aws ec2 describe-images \
        --filters "Name=name,Values=amzn-ami*amazon-ecs-optimized*" \
        --owner aws-marketplace \
        --query 'Images[*].[ImageId,CreationDate]' \
        --output text | \
    sort -k2 -r  | \
    head -n1 | \
    cut -f 1
}

# Requires IAM read-only access on the user
get_aws_account_id() {
    aws iam get-user --output text | cut -d\: -f 5
}

if [ -z "$AMI_ID" ]; then
    AMI_ID="$(get_most_recent_container_image)"
fi

if [ -z "$AWS_ACCOUNT_ID" ]; then
    AWS_ACCOUNT_ID="$(get_aws_account_id)"
fi

# TODO: automate the creation of these beasts
key_pair_name="ecs-instances"
ecs_instance_role="ecsInstanceRole"

# This requires permission for ec2:RunInstances Action
launch_container_instances() {
    aws ec2 run-instances \
        --instance-type "t2.micro" \
        --count "2" \
        --image-id "$AMI_ID" \
        --iam-instance-profile "Name=$ecs_instance_role" \
        --associate-public-ip-address \
        --key-name "$key_pair_name"
}

# launch_container_instances

aws ecs register-task-definition \
    --cli-input-json file://documents-server-task.json

aws ecs run-task \
    --task-definition 'documents:1' \
    --count 1

aws ecs create-service \
    --service-name documents \
    --task-definition 'documents:1' \
    --desired-count 1

