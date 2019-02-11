package main

import (
	"context"
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/serviceusage/v1"
)

func EnableServices(ctx context.Context, projectNumber int64) error {
	hc, err := google.DefaultClient(ctx, serviceusage.CloudPlatformScope)
	if err != nil {
		return err
	}
	client, err := serviceusage.New(hc)
	if err != nil {
		return err
	}

	o, err := client.Services.BatchEnable(fmt.Sprintf("projects/%d", projectNumber), &serviceusage.BatchEnableServicesRequest{
		ServiceIds: []string{
			"compute.googleapis.com",
			"datastore.googleapis.com", // TODO datastore api enalbeにしてもFirestoreは大丈夫か？
			"bigquery-json.googleapis.com",
		},
	}).Do()
	if err != nil {
		return err
	}
	pp.Println(o)

	for {
		o, err := client.Operations.Get(o.Name).Do()
		if err != nil {
			return err
		}
		if o.Done {
			pp.Println(o)
			break
		}
		fmt.Println("zzzz.......")
		time.Sleep(10 * time.Second)
	}

	return nil
}
