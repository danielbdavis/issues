package main

import (
	"context"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
)

//import (
//	"bufio"
//	"context"
//	"fmt"
//	"os"
//	"strings"
//	"syscall"
//
//	"github.com/google/go-github/github"
//	"golang.org/x/crypto/ssh/terminal"
//)
//
//func main() {
//	r := bufio.NewReader(os.Stdin)
//	fmt.Print("GitHub Username: ")
//	username, _ := r.ReadString('\n')
//
//	fmt.Print("GitHub Password: ")
//	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
//	password := string(bytePassword)
//
//	tp := github.BasicAuthTransport{
//		Username: strings.TrimSpace(username),
//		Password: strings.TrimSpace(password),
//	}
//
//	client := github.NewClient(tp.Client())
//	ctx := context.Background()
//	user, _, err := client.Users.Get(ctx, "")
//
//	// Is this a two-factor auth error? If so, prompt for OTP and try again.
//	if _, ok := err.(*github.TwoFactorAuthError); ok {
//		fmt.Print("\nGitHub OTP: ")
//		otp, _ := r.ReadString('\n')
//		tp.OTP = strings.TrimSpace(otp)
//		user, _, err = client.Users.Get(ctx, "")
//	}
//
//	if err != nil {
//		fmt.Printf("\nerror: %v\n", err)
//		return
//	}
//
//	fmt.Printf("\n%v\n", github.Stringify(user))
//}

var client *github.Client

func main() {

	// auth
	client = authenticateGithub()

	organization := selectOrganization()

	repository := selectRepositoryByOrganization(organization)

	listIssues(repository)


	//r := bufio.NewReader(os.Stdin)
	//fmt.Print("GitHub Username: ")
	//username, _ := r.ReadString('\n')
	//
	//fmt.Print("GitHub Password: ")
	//bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	//password := string(bytePassword)
	//
	//tp := github.BasicAuthTransport{
	//	Username: strings.TrimSpace(username),
	//	Password: strings.TrimSpace(password),
	//}
	//
	//client := github.NewClient(tp.Client())
	//ctx := context.Background()
	//user, _, err := client.Users.Get(ctx, "")
	//
	//// Is this a two-factor auth error? If so, prompt for OTP and try again.
	//if _, ok := err.(*github.TwoFactorAuthError); ok {
	//	fmt.Print("\nGitHub OTP: ")
	//	otp, _ := r.ReadString('\n')
	//	tp.OTP = strings.TrimSpace(otp)
	//	user, _, err = client.Users.Get(ctx, "")
	//}
	//
	//if err != nil {
	//	fmt.Printf("\nerror: %v\n", err)
	//	return
	//}
	//
	//fmt.Printf("\n%v\n", github.Stringify(user))

	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: ""},
	//	)
	//tc := oauth2.NewClient(ctx, ts)
	//client := github.NewClient(tc)
	//
	////// setup our GitHub client
	////client := github.NewClient(nil)
	//
	//// setup our CLI
	//app := cli.NewApp()
	//app.EnableBashCompletion = true
	//app.Name = "issues"
	//app.Usage = "issues"
	//app.Version = "0.0.1"
	//
	//// add commands
	//app.Commands = []cli.Command{
	//	{
	//		Name:	"orgs",
	//		Usage:	"fetch all organizations",
	//		Action: func(c *cli.Context) error {
	//			// username:
	//			var username string
	//			fmt.Print("Enter GitHub username: ")
	//			fmt.Scanf("%s", &username)
	//
	//			// fetch organizations
	//			organizations, err := fetchOrganizations(username)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//
	//			// print
	//			for i, organization := range organizations {
	//				log.Printf("%v. %v\n", i+1, organization.GetLogin())
	//			}
	//
	//			return nil
	//		},
	//	},
	//
	//	{
	//		Name:	"all",
	//		Aliases:[]string{"a"},
	//		Usage:"fetching all issues",
	//		Action: func(c *cli.Context) error {
	//			log.Println("fetch all issues")
	//
	//			// lookup options
	//			//opt := &github.RepositoryListByOrgOptions{Type: "public"}
	//
	//			opt := &github.ListOptions{}
	//
	//			// list organizations
	//			organizations, _, err := client.Organizations.List(context.Background(), "", opt)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//
	//			log.Println(organizations)
	//
	//			for _, organization := range organizations {
	//				print(organization)
	//				spew.Dump(organization)
	//			}
	//
	//			//// list organizations
	//			//for _, organization := range client.Organizations.Get(context.Background(), "github") {
	//			//	spew.Dump(organization)
	//			//}
	//
	//			////
	//			//repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	//			//if err != nil {
	//			//	log.Fatal(err)
	//			//}
	//			//
	//			//// print to screen
	//			//for _, repo := range repos {
	//			//	log.Println("")
	//			//	spew.Dump(repo)
	//			//}
	//
	//			return nil
	//		},
	//	},
	//}
	//
	//// run our CLI
	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

// ---------------------------------------------------------------------------------------------------------------------
// auth
// ---------------------------------------------------------------------------------------------------------------------

func authenticateGithub() *github.Client {
	// fetch username
	var username string
	fmt.Print("GitHub username: ")
	fmt.Scanf("%s", &username)

	// fetch access token
	var accessToken string
	fmt.Print("GitHub Access Token: ")
	fmt.Scanf("%s", &accessToken)

	context := context.Background()
	token := oauth2.Token{AccessToken: accessToken}
	tokenSource := oauth2.StaticTokenSource(&token)
	tokenClient := oauth2.NewClient(context, tokenSource)
	client := github.NewClient(tokenClient)

	return client
}

// ---------------------------------------------------------------------------------------------------------------------
// organizations
// ---------------------------------------------------------------------------------------------------------------------

func selectOrganization() (organization *github.Organization) {
	// fetch organizations
	organizations, _, err := client.Organizations.List(context.Background(), "", nil)
	if err != nil {
		log.Fatal(err)
	}

	// print numbered list
	for i, organization := range organizations {
		fmt.Printf("%3v. %v\n", i, *organization.Login)
	}

	// prompt
	fmt.Printf("\n Select an organization: ")

	// scan
	var organizationIndex int
	fmt.Scanf("%d", &organizationIndex)

	organization = organizations[organizationIndex]
	return organization
}

// ---------------------------------------------------------------------------------------------------------------------
// repositories
// ---------------------------------------------------------------------------------------------------------------------

func selectRepositoryByOrganization(organization *github.Organization) (repository *github.Repository) {
	// fetch repos for this org
	repositories, _, err := client.Repositories.ListByOrg(context.Background(), *organization.Login, nil)
	if err != nil {
		log.Fatal(err)
	}

	// print numbered list
	for i, repository := range repositories {
		fmt.Printf("%3v. %v\n", i, *repository.Name)
	}

	// prompt
	fmt.Printf("\n Select a repository: ")

	// scan
	var repositoryIndex int
	fmt.Scanf("%d", &repositoryIndex)

	repository = repositories[repositoryIndex]
	return repository
}

// ---------------------------------------------------------------------------------------------------------------------
// issues
// ---------------------------------------------------------------------------------------------------------------------

func listIssues(repository *github.Repository) {
	// fetch issues
	issues, _, err := client.Issues.ListByRepo(context.Background(), "", *repository.Name, nil)
	if err != nil {
		log.Fatal(err)
	}

	// print numbered list
	for i, issue := range issues {
		fmt.Printf("%3v. %v\n", i, *issue.Title)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// unused
// ---------------------------------------------------------------------------------------------------------------------

func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func listRepositories() {
	// fetch repos
	repositories, _, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		log.Fatal(err)
	}

	for i, repository := range repositories {
		fmt.Printf("%3v. %v\n", i, *repository.Name)
	}
}

func listAllRepositories() {
	// fetch all repos
	repositories, _, err := client.Repositories.ListAll(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for i, repository := range repositories {
		fmt.Printf("%3v. %v\n", i, *repository.Name)
	}
}
