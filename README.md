# Terratest examples

This repo houses code and tests as part of a set of posts about Terraform and testing.

- [Terraform knowledge to get you through the day](https://dev.to/benmatselby/terraform-knowledge-to-get-you-through-the-day-17kk)

To deploy the examples without Terratest:

```shell
cd aws
terraform init
terraform apply
```

To run the tests:

```shell
go get ./...
go test ./...
```
