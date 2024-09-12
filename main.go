package main

import (
	"context"
	"fmt"
	"custom-cloud-provider/mycloud"
)

func main() {
	// Initialize the custom cloud provider.
	provider := mycloud.NewCustomProvider()

	// Use the provider to interact with resources.
	ctx := context.Background()

	// Instances
	instance, err := provider.Instances().CreateInstance(ctx, "my-instance")
	if err != nil {
		fmt.Println("Error creating instance:", err)
	} else {
		fmt.Println("Created instance:", instance.ID)
	}

	// Storage
	bucket, err := provider.Storage().CreateBucket(ctx, "my-bucket")
	if err != nil {
		fmt.Println("Error creating bucket:", err)
	} else {
		fmt.Println("Created bucket:", bucket.Name)
	}

	// Networking
	network, err := provider.Networking().CreateNetwork(ctx, "my-network", "172.16.0.0/16")
	if err != nil {
		fmt.Println("Error creating network:", err)
	} else {
		fmt.Println("Created network:", network.ID)
	}
}
