## AWS IOT Infrastructure
### AWS IOT Authentication
An x.509 certificate is required to authenticate with AWS IOT service via TLS SNI before the MQTT
handshake.

We won't use Terraform to create the certificate because it could expose the
private key. Better that we create the certificate with AWS via the CLI
then associate it with the TF infrastructure.

### Create The Keys
Set the region to the region where the infrastructure will be installed. Take note of the
certificate ARN. It is needed as an infrastructure variable.
```shell
myAwsRegion=us-east-1

aws iot create-keys-and-certificate \
    --certificate-pem-outfile "growatt-to-iot.cert.pem" \
    --public-key-outfile "growatt-to-iot.public.key" \
    --private-key-outfile "growatt-to-iot.private.key" \
    --set-as-active \
    --region $myAwsRegion
```
Be careful not to commit certificate files to git.

### Terraform Auto TFvars
Add the `"certificateArn"` from the output above into a `terraform.auto.tfvars` file:
```terraform
aws_iot_certificate_arn = "arn:aws:iot:us-east-1:12345678910:cert/47747474747474747474747474"
```
You can also set the `aws_region` here instead of the default `us-east-1`. However, note that
some infrastructure may not be available in all regions. Also ensure you follow the same region
that you used to generate the certificates.

### Terraform Apply
You can either go with local state or organise a `backend.tf` like:
```terraform
terraform {
  backend "remote" {
    organization = "my-org"
    workspaces {
      name = "growatt-you-like"
    }
  }

  required_version = ">= 1.0.2"
}
```

Then:
```shell
terraform init
terraform apply
```

### Get The AWS CA Certificate
We'll use the `-ats` endpoint. The endpoint is output to the console when the Terraform
build has been completed. The `-ats` endpoint requires the AWS Root CA certificate from here:

https://www.amazontrust.com/repository/AmazonRootCA1.pem