// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gsuiteaddons_v1_generated_GSuiteAddOns_GetDeployment_sync]

package main

import (
	"context"

	gsuiteaddons "cloud.google.com/go/gsuiteaddons/apiv1"
	gsuiteaddonspb "google.golang.org/genproto/googleapis/cloud/gsuiteaddons/v1"
)

func main() {
	ctx := context.Background()
	c, err := gsuiteaddons.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gsuiteaddonspb.GetDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END gsuiteaddons_v1_generated_GSuiteAddOns_GetDeployment_sync]
