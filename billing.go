package main

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbilling/v1"
)

func SetBillingAccount(ctx context.Context, projectID string, billingID string) error {
	hc, err := google.DefaultClient(ctx, cloudbilling.CloudPlatformScope)
	if err != nil {
		return err
	}
	client, err := cloudbilling.New(hc)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("projects/%s", projectID)

	resp, err := client.Projects.UpdateBillingInfo(name, &cloudbilling.ProjectBillingInfo{
		// TG
		BillingAccountName: fmt.Sprintf("billingAccounts/%s", billingID),

		// sinmetal.jp ハンズオン
		// BillingAccountName: fmt.Sprintf("billingAccounts/%s", "011F4B-52726C-BD1E83"),
	}).Context(ctx).Do()
	if err != nil {
		return err
	}
	pp.Println(resp)

	return nil
}
