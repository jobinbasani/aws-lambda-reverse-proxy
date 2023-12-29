# AWS Lambda Reverse Proxy
A quick on-demand reverse proxy using AWS Lambda.
This makes use of the Function URL offered by AWS Lambda.
The host to be proxied is set as an environment variable.

Infrastructure is created using CDK.

To deploy this function,

    cd infra
    cdk deploy
