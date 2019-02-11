package main

import (
	"context"
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
)

func CreateProject(ctx context.Context, projectID string, parentID string) error {
	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		return err
	}
	crm, err := cloudresourcemanager.New(client)
	if err != nil {
		return err
	}

	o, err := crm.Projects.Create(&cloudresourcemanager.Project{
		ProjectId: projectID,
		Name:      projectID,
		Parent: &cloudresourcemanager.ResourceId{
			Id:   parentID,
			Type: "folder",
		},
	}).Do()
	if err != nil {
		return err
	}
	pp.Println(o)

	for {
		o, err := crm.Operations.Get(o.Name).Do()
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

func GetProject(ctx context.Context, projectID string) (*cloudresourcemanager.Project, error) {
	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	crm, err := cloudresourcemanager.New(client)
	if err != nil {
		return nil, err
	}

	resp, err := crm.Projects.Get(projectID).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteProject(ctx context.Context, projectID string) error {
	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		return err
	}
	crm, err := cloudresourcemanager.New(client)
	if err != nil {
		return err
	}

	_, err = crm.Projects.Delete(projectID).Do()
	if err != nil {
		return err
	}

	return nil
}
