
[![license](https://img.shields.io/github/license/cgrotz/terraform-provider-bis.svg)](https://github.com/cgrotz/terraform-provider-bis/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/cgrotz/terraform-provider-bis.svg)](https://github.com/cgrotz/terraform-provider-bis/releases/latest)
[![Build](https://github.com/cgrotz/terraform-provider-bis/workflows/Go/badge.svg?branch=master&event=push)](https://github.com/cgrotz/terraform-provider-bis/workflows/Go/badge.svg?branch=master&event=push)
[![CodeFactor](https://www.codefactor.io/repository/github/cgrotz/terraform-provider-bis/badge)](https://www.codefactor.io/repository/github/cgrotz/terraform-provider-bis)

Bosch IoT Suite Terraform Provider (unofficial)
================

## Introduction
This Terraform provider allows you to provision resources on the Bosch IoT Suite using Terraform.

At the moment due to limitations on the Bosch IoT Suite the setup of the provider is very inconvenient. I am sorry about that.

It's currently not possible to subscribe to services. The service instances already need to be created. The IDs of those service instances need to be extracted manually from the IoT Suite Portal (please me mindful to copy the right IDs).

## Setup of the provider
You can simple add the provider to your Terraform script with the following code snippet:
```
provider "bis" {
  client_id = "<OAuth Client ID>"
  client_secret = "<OAuth Client Secret>"
  things_solution_id = "<iot_thing_solution_id" // Optional; set this value if you want to provision resources on IoT Things
}
```

### OAuth Client Creation
You can simply create an OAuth Client using the Bosch IoT Suite portal. If you want to use the Client to manage resources on IoT Things, you will also need to manually add the OAuth client to the IoT Things solution policy.

Add a subject following the template:
```
"iot-suite:service:iot-things-eu-1:<solution_id>/full-access@<client_id>": {
  "type": "suite-auth"
},
```
to the solution policy. It should have a generous set of access.
```
"resources": {
        "solution:/": {
          "grant": [
            "READ",
            "WRITE"
          ],
          "revoke": []
        },
        "policy:/": {
          "grant": [
            "READ",
            "WRITE"
          ],
          "revoke": []
        }
      }
```