// package cloudflare

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// TODO: Some self delcared variables to get the ball rolling - now obtained from JSON
// var (
// 	// expectedPagesProject = &cloudflare.PagesProject{
// 	// 	// SubDomain: "dummynextjs.pages.dev",
// 	// 	Name:    "dummy-nextjs",
// 	// 	Domains: []string{
// 	// 		// "testdomain.com",
// 	// 		// "testdomain.org",
// 	// 	},
// 	// 	CanonicalDeployment: *expectedPagesProjectDeployment,
// 	// 	BuildConfig:         *expectedPagesProjectBuildConfig,
// 	// 	DeploymentConfigs:   *expectedPagesProjectDeploymentConfigs,
// 	// 	Source:              *expectedPagesProjectSource,
// 	// 	LatestDeployment:    *expectedPagesProjectDeployment,
// 	// }

// 	expectedPagesProjectDeployment = &cloudflare.PagesProjectDeployment{
// 		ProjectName: "dummy-nextjs",
// 		Environment: "preview",
// 		EnvVars: map[string]map[string]string{
// 			"NEXT_PUBLIC_SOME_ENV_VAR": {
// 				"value": "PREVIEW",
// 			},
// 			"ENV": {
// 				"value": "STAGING",
// 			},
// 		},
// 		BuildConfig: *expectedPagesProjectBuildConfig,
// 		Source:      *expectedPagesProjectSource,
// 	}

// 	expectedPagesProjectBuildConfig = &cloudflare.PagesProjectBuildConfig{
// 		BuildCommand:   "npm run build",
// 		DestinationDir: "build",
// 		RootDir:        "",
// 	}

// 	expectedPagesProjectSource = &cloudflare.PagesProjectSource{
// 		Type:   "github",
// 		Config: expectedPagesProjectSourceConfig,
// 	}

// 	expectedPagesProjectSourceConfig = &cloudflare.PagesProjectSourceConfig{
// 		Owner:              "adamsuk",
// 		RepoName:           "dummy-nextjs",
// 		ProductionBranch:   "main",
// 		PRCommentsEnabled:  true,
// 		DeploymentsEnabled: true,
// 	}

// 	expectedPagesProjectDeploymentConfigs = &cloudflare.PagesProjectDeploymentConfigs{
// 		Preview:    *expectedPagesProjectDeploymentConfigPreview,
// 		Production: *expectedPagesProjectDeploymentConfigProduction,
// 	}

// 	expectedPagesProjectDeploymentConfigPreview = &cloudflare.PagesProjectDeploymentConfigEnvironment{
// 		EnvVars: map[string]cloudflare.PagesProjectDeploymentVar{
// 			"BUILD_VERSION": {
// 				Value: "1.2",
// 			},
// 			"NEXT_PUBLIC_SOME_ENV_VAR": {
// 				Value: "PREVIEW",
// 			},
// 		},
// 	}

// 	expectedPagesProjectDeploymentConfigProduction = &cloudflare.PagesProjectDeploymentConfigEnvironment{
// 		EnvVars: map[string]cloudflare.PagesProjectDeploymentVar{
// 			"BUILD_VERSION": {
// 				Value: "1.2",
// 			},
// 			"NEXT_PUBLIC_SOME_ENV_VAR": {
// 				Value: "PRODUCTION",
// 			},
// 		},
// 	}
// )

func projectInState(project string, projects []cloudflare.PagesProject) bool {
	for _, p := range projects {
		if p.Name == project {
			return true
		}
	}
	return false
}

func writeProjectToFile(project cloudflare.PagesProject, filename string) {
	file, _ := json.MarshalIndent(project, "", "\t")
	_ = ioutil.WriteFile(filename, file, 0777)
}

func main() {
	// Useful links:
	// https://go.dev/blog/using-go-modules
	// https://golangexample.com/a-go-library-for-interacting-with-cloudflares-api-v4/
	// https://github.com/cloudflare/cloudflare-go/pull/724/files

	fmt.Println("Welcome to cloudflare-go CI'd!")

	// Construct a new API object
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)

	// Fetch the zone ID
	id, err := api.ZoneIDByName("sradams.co.uk") // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)

	// Fetch Pages Projects
	projects, _, err := api.ListPagesProjects(ctx, os.Getenv("CF_ACCOUNT_ID"), cloudflare.PaginationOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// Print projects details
	s_projects, _ := json.MarshalIndent(projects, "", "\t")
	fmt.Println(string(s_projects))

	// Open our jsonFile
	jsonFile, err := os.Open("pages.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened pages.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var project cloudflare.PagesProject
	json.Unmarshal(byteValue, &project)

	// TODO: Some params with different errors - either nothing at all (GREAT) or internal service error :(
	// var expectedPagesProjectDeploymentConfig = &cloudflare.PagesProjectDeploymentConfigEnvironment{
	// 	EnvVars: map[string]cloudflare.PagesProjectDeploymentVar{
	// 		"BUILD_VERSION": {
	// 			Value: "1.2",
	// 		},
	// 		"NEXT_PUBLIC_SOME_ENV_VAR": {
	// 			Value: "PREVIEW",
	// 		},
	// 		"SOME_OTHER_VAR": {
	// 			Value: "ERRMMMMMM",
	// 		},
	// 	},
	// }

	// var expectedPagesProjectDeploymentConfigs = &cloudflare.PagesProjectDeploymentConfigs{
	// 	Preview:    *expectedPagesProjectDeploymentConfig,
	// 	Production: *expectedPagesProjectDeploymentConfig,
	// }
	// project := cloudflare.PagesProject{
	// 	Name: "updated-dummy-nextjs",
	// 	// DeploymentConfigs: *expectedPagesProjectDeploymentConfigs,
	// }

	// Print the expected Pages Project object
	s, _ := json.MarshalIndent(project, "", "\t")
	fmt.Println(string(s))

	// // Fetch project from Cloudflare
	_, err = api.PagesProject(ctx, os.Getenv("CF_ACCOUNT_ID"), project.Name)
	if err != nil && err.Error() == "Project not found (8000007)" {
		// if it errors it's because the project doesn't exist - needs to be created
		fmt.Println("Project not found ... creating") // Create a new page project
		createdProject, err := api.CreatePagesProject(ctx, os.Getenv("CF_ACCOUNT_ID"), project)
		if err != nil {
			log.Fatal(err)
		}
		// Write project details to file
		fmt.Println(createdProject)
		writeProjectToFile(createdProject, "pages.json")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Project already exists ... updating")
		// for whatever reason the name field cannot be present when updating as the API will error on project name already being used
		project.Name = ""
		// Update a new page project
		updatedProject, err := api.UpdatePagesProject(ctx, os.Getenv("CF_ACCOUNT_ID"), project.Name, project)
		if err != nil {
			log.Fatal(err)
		}
		// Write project details to file
		fmt.Println(updatedProject)
		writeProjectToFile(updatedProject, "pages.json")
	}

	// TODO: it'd be nicer to fetch all projects and filter but the update function breaks if project.Name is defined!! :(
	// // Check if a page exists in projects
	// if projectInState(project.Name, projects) {
	// 	fmt.Println("Project already exists ... updating")
	// 	// Update a new page project
	// 	updatedProject, err := api.UpdatePagesProject(ctx, os.Getenv("CF_ACCOUNT_ID"), project.Name, project)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// Write project details to file
	// 	fmt.Println(updatedProject)
	// 	writeProjectToFile(updatedProject, "pages.json")
	// } else {
	// 	fmt.Println("Project does not exist ... creating")
	// 	// Create a new page project
	// 	createdProject, err := api.CreatePagesProject(ctx, os.Getenv("CF_ACCOUNT_ID"), project)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// Write project details to file
	// 	fmt.Println(createdProject)
	// 	writeProjectToFile(createdProject, "pages.json")
	// }
}
