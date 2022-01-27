package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Define key global variables.
var (
	subscriptionId = os.Getenv("AZURE_SUBSCRIPTION_ID")
)

// List all the resource groups of an Azure subscription.
func listResourceGroups(ctx context.Context, credential azcore.TokenCredential) ([]*armresources.ResourceGroup, error) {
	rgClient := armresources.NewResourceGroupsClient(subscriptionId, credential, nil)

	pager := rgClient.List(nil)

	var resourceGroups []*armresources.ResourceGroup
	for pager.NextPage(ctx) {
		resp := pager.PageResponse()
		if resp.ResourceGroupListResult.Value != nil {
			resourceGroups = append(resourceGroups, resp.ResourceGroupListResult.Value...)
		}
	}
	return resourceGroups, pager.Err()
}

// Define the standard 'main' function for an app that is called from the command line.
func main() {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()

	rgList, err := listResourceGroups(ctx, cred)
	if err != nil {
		log.Fatalf("Listing of resource groups failed: %+v", err)
	}
	log.Printf("Your Azure subscription has a total of %d resource groups", len(rgList))

}
