package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Cloud Resource Manager API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/cloudresourcemanager
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
        "fmt"
        "log"

        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/cloudresourcemanager/v1"
)

func main() {
        ctx := context.Background()

        c, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
        if err != nil {
                log.Fatal(err)
        }

        cloudresourcemanagerService, err := cloudresourcemanager.New(c)
        if err != nil {
                log.Fatal(err)
        }

        // The Project ID (for example, `my-project-123`).
        // Required.
        projectId := "arctic-keyword-205205" // TODO: Update placeholder value.

        resp, err := cloudresourcemanagerService.Projects.Get(projectId).Context(ctx).Do()
        if err != nil {
                log.Fatal(err)
        }

        // TODO: Change code below to process the `resp` object:
        fmt.Printf("%#v\n", resp)
}