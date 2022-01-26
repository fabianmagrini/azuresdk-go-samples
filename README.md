# Azure SDK for Go

References:

* <https://github.com/Azure/azure-sdk-for-go>
* <https://docs.microsoft.com/en-us/learn/paths/go-first-steps/>
* <https://docs.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code>
* <https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication?tabs=bash>
* <https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication-service-principal?tabs=azure-cli>
* <https://docs.microsoft.com/en-us/azure/developer/go/manage-resource-groups?tabs=bash%2Cazure-portal>
* <https://github.com/Azure-Samples/azure-sdk-for-go-samples>

## Prerequisites

* Azure subscription
* Go installed

## Authenticate with Azure

The Azure Identity module is used to authenticate to Azure. To support local development, the DefaultAzureCredential can authenticate as the user signed into the Azure CLI.

Run the following command to sign into the Azure CL

```sh
az login
```

## Download necessary Azure SDK for Go modules

```sh
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
go get github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions
```

## Run

```sh
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
go get github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions
```
