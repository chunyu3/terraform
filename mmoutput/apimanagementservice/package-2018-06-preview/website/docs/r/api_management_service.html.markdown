---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file at
#     https://github.com/Azure/magic-module-specs
#
# ----------------------------------------------------------------------------
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_api_management_service"
sidebar_current: "docs-azurerm-resource-api-management-service"
description: |-
  Manage Azure ApiManagementService instance.
---

# azurerm_api_management_service

Manage Azure ApiManagementService instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the API Management service. Changing this forces a new resource to be created.

* `resource_group` - (Required) The name of the resource group. Changing this forces a new resource to be created.

* `location` - (Required) Resource location. Changing this forces a new resource to be created.

* `additional_locations` - (Optional) One or more `additional_location` block defined below.

* `certificates` - (Optional) One or more `certificate` block defined below.

* `custom_properties` - (Optional) Custom properties of the API Management service. Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2). Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1 and setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.

* `hostname_configurations` - (Optional) One or more `hostname_configuration` block defined below.

* `identity` - (Optional) One `identity` block defined below.

* `notification_sender_email` - (Optional) Email address from which the notification will be sent.

* `publisher_email` - (Optional) Publisher email.

* `publisher_name` - (Optional) Publisher name.

* `sku` - (Optional) One `sku` block defined below.

* `virtual_network_configuration` - (Optional) One `virtual_network_configuration` block defined below.

* `virtual_network_type` - (Optional) The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only. Defaults to `None`.

* `tags` - (Optional) Resource tags. Changing this forces a new resource to be created.

---

The `additional_location` block supports the following:

* `location` - (Required) The location name of the additional region among Azure Data center regions. Changing this forces a new resource to be created.

* `sku` - (Required) One `sku` block defined below.

* `virtual_network_configuration` - (Optional) One `virtual_network_configuration` block defined below.


---

The `sku` block supports the following:

* `name` - (Required) Name of the Sku.

* `capacity` - (Optional) Capacity of the SKU (number of deployed units of the SKU). The default value is 1.

---

The `virtual_network_configuration` block supports the following:

* `subnet_resource_id` - (Optional) The full resource ID of a subnet in a virtual network to deploy the API Management service in.

---

The `certificate` block supports the following:

* `encoded_certificate` - (Optional) Base64 Encoded certificate.

* `certificate_password` - (Optional) Certificate Password.

* `store_name` - (Required) The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and CertificateAuthority are valid locations.

* `certificate` - (Optional) One `certificate` block defined below.


---

The `certificate` block supports the following:

* `expiry` - (Required) Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.

* `thumbprint` - (Required) Thumbprint of the certificate.

* `subject` - (Required) Subject of the certificate.

---

The `hostname_configuration` block supports the following:

* `type` - (Required) Hostname type.

* `host_name` - (Required) Hostname to configure on the Api Management service.

* `key_vault_id` - (Optional) Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided, auto-update of ssl certificate will not work. This requires Api Management service to be configured with MSI. The secret should be of type *application/x-pkcs12*

* `encoded_certificate` - (Optional) Base64 Encoded certificate.

* `certificate_password` - (Optional) Certificate Password.

* `default_ssl_binding` - (Optional) Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate. If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The setting only applied to Proxy Hostname Type.

* `negotiate_client_certificate` - (Optional) Specify true to always negotiate client certificate on the hostname. Default Value is false.

* `certificate` - (Optional) One `certificate` block defined below.


---

The `certificate` block supports the following:

* `expiry` - (Required) Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.

* `thumbprint` - (Required) Thumbprint of the certificate.

* `subject` - (Required) Subject of the certificate.

---

The `identity` block supports the following:

* `type` - (Required) The identity type. Currently the only supported type is 'SystemAssigned'.

---

The `sku` block supports the following:

* `name` - (Required) Name of the Sku.

* `capacity` - (Optional) Capacity of the SKU (number of deployed units of the SKU). The default value is 1.

---

The `virtual_network_configuration` block supports the following:

* `subnet_resource_id` - (Optional) The full resource ID of a subnet in a virtual network to deploy the API Management service in.

## Attributes Reference

The following attributes are exported:

* `provisioning_state` - The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.

* `target_provisioning_state` - The provisioning state of the API Management service, which is targeted by the long running operation started on the service.

* `created_at_utc` - Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.

* `gateway_url` - Gateway URL of the API Management service.

* `gateway_regional_url` - Gateway URL of the API Management service in the Default Region.

* `portal_url` - Publisher portal endpoint Url of the API Management service.

* `management_api_url` - Management API endpoint URL of the API Management service.

* `scm_url` - SCM endpoint URL of the API Management service.

* `public_ip_addresses` - Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.

* `private_ip_addresses` - Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.

* `id` - Resource ID.

* `name` - Resource name.

* `type` - Resource type for API Management resource is set to Microsoft.ApiManagement.

* `etag` - ETag of the resource.
