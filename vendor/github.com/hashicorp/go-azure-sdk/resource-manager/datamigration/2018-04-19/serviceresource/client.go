package serviceresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceClient struct {
	Client *resourcemanager.Client
}

func NewServiceResourceClientWithBaseURI(api environments.Api) (*ServiceResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serviceresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceResourceClient: %+v", err)
	}

	return &ServiceResourceClient{
		Client: client,
	}, nil
}
