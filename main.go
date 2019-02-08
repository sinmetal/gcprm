package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
)

type Param struct {
	ProjectIDPrefix  string
	Start            int
	End              int
	ParentID         string
	BillingAccountID string
}

func main() {
	param, err := getFlag()
	if err != nil {
		fmt.Printf("failed getFlag. err=%+v", err)
		os.Exit(1)
	}

	start := 3
	end := 40
	for i := start; i <= end; i++ {
		ctx := context.Background()
		projectID := fmt.Sprintf("%s%03d", param.ProjectIDPrefix, i)
		fmt.Printf("Try Create %s Project\n", projectID)
		if err := CreateProject(ctx, projectID, param.ParentID); err != nil {
			fmt.Printf("failed CreateProject %+v\n", err)
			os.Exit(1)
		}
		if err := SetBillingAccount(ctx, projectID, param.BillingAccountID); err != nil {
			fmt.Printf("failed SetBillingAccount %+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Finish !!")
	os.Exit(0)
}

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

func getFlag() (*Param, error) {
	var (
		projectPrefix = flag.String("project-prefix", "", "")
		start         = flag.Int("start", 1, "")
		end           = flag.Int("end", 5, "")
		parent        = flag.String("parent", "", "")
		billing       = flag.String("billing", "", "")
	)
	flag.Parse()

	var emsg string
	if len(*projectPrefix) < 1 {
		emsg += "project-prefix is required\n"
	}
	if *start == 0 {
		emsg += "start is required\n"
	}
	if *end == 0 {
		emsg += "end is required\n"
	}
	if len(*parent) < 1 {
		emsg += "parent is required\n"
	}
	if len(*billing) < 1 {
		emsg += "billing is required\n"
	}

	if len(emsg) > 0 {
		return nil, errors.New(emsg)
	}

	return &Param{
		ProjectIDPrefix:  *projectPrefix,
		Start:            *start,
		End:              *end,
		ParentID:         *parent,
		BillingAccountID: *billing,
	}, nil
}
