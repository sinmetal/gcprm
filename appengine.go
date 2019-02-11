package main

import (
	"context"
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/appengine/v1"
)

func CreateAppEngineApp(ctx context.Context, projectID string) error {
	hc, err := google.DefaultClient(ctx, appengine.CloudPlatformScope)
	if err != nil {
		return err
	}
	client, err := appengine.New(hc)
	if err != nil {
		return err
	}
	o, err := client.Apps.Create(&appengine.Application{
		Id:         projectID,
		LocationId: "asia-northeast1",
	}).Do()
	if err != nil {
		return err
	}
	pp.Println(o)

	for {
		o, err := client.Apps.Operations.Get(projectID, o.Name).Do()
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
