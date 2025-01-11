## Postmark Terraform Provider

This is a terraform provider for Postmark Account Management such as servers, senders, domains, and templates.

This provider uses the Postmark API facilitated by the community library [Postmark Golang](https://github.com/mrz1836/postmark) by [Mr. Z](https://github.com/mrz1836).

> [!CAUTION]
> **It is a work in progress and is not yet ready for production use.** [!CAUTION]

## Developing Locally

### Prerequisites

* [Golang 1.18+](https://go.dev/dl)
* [Terraform 1.3.1+](https://www.terraform.io/downloads)

### ⚙️ Building the provider

Execute the build file of the provider.

```bash
make install
```

💡 A file named `~/.terraform.d/plugins/marcus.carey/terraform/postmark/1.0/${OS_ARCH}/terraform-provider-postmark` will be created. This is your custom provider.

### ⏯ Playing with the provider

1. Enter the `examples` directory.

```bash
cd examples
```

2. Create a `terraform.tfvars` file with the following content:

```hcl
postmark_account_token = "POSTMARK_ACCOUNT_TOKEN"
existing_server_id = "EXISTING_POSTMARK_SERVER_ID"
```

3. Initialize the provider plugins.

```bash
terraform init
```

4. Check the execution plan.

```bash
terraform plan
```

5.  🚀 Apply the changes.

```bash
terraform apply --auto-approve
```