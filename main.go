package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
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

	for i := param.Start; i <= param.End; i++ {
		ctx := context.Background()
		projectID := fmt.Sprintf("%s%03d", param.ProjectIDPrefix, i)
		fmt.Printf("Try Create %s Project\n", projectID)
		//if err := CreateProject(ctx, projectID, param.ParentID); err != nil {
		//		//	fmt.Printf("failed CreateProject %+v\n", err)
		//		//	os.Exit(1)
		//		//}
		//		//if err := SetBillingAccount(ctx, projectID, param.BillingAccountID); err != nil {
		//		//	fmt.Printf("failed SetBillingAccount %+v\n", err)
		//		//	os.Exit(1)
		//		//}

		//if err := CreateAppEngineApp(ctx, projectID); err != nil {
		//	if gerr, ok := err.(*googleapi.Error); ok {
		//		if gerr.Code == 409 {
		//			fmt.Printf("%s is App Engine Alredy exits\n", projectID)
		//		}
		//	} else {
		//		log.Fatalf("fatal create AppEngine. err=%+v\n", err)
		//	}
		//}

		//p, err := GetProject(ctx, projectID)
		//if err != nil {
		//	log.Fatalf("fatal get project %s\n", projectID)
		//}
		//if err := EnableServices(ctx, p.ProjectNumber); err != nil {
		//	log.Fatalf("fatal enable services %+v\n", err)
		//}

		if err := DeleteProject(ctx, projectID); err != nil {
			log.Fatalf("fatal delete project. project id =%s. err=%+v\n", projectID, err)
		}
	}

	fmt.Println("Finish !!")
	os.Exit(0)
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
