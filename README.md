# Flugel-challenge

A Terraform code to create an AWS EC2 instance and S3 bucket, test with Terratest to check if  both resources are tagged properly.

* Github Actions
![image](https://user-images.githubusercontent.com/33374159/115548743-b976c300-a29f-11eb-9479-1f477e131f9a.png)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development purpose.

## Prerequisites

To get started with this project you need a basic knowledge of the following.
```
Version Control (Git)
AWS
Terraform
Terratest
Go
```

## Testing with Terratest

* Clone the repository, in the project directory, `cd` into `flugel-challenge/test`.
* Then run `go mod init github.com/<github-username>/<github-repo-name>`.
* Now you can run the test with `go test -count=1 -v ./...`. `-count=1` is used to avoid Go test caching and `-v` stands for verbose.
* At the end of the test run, your test should pass and you should see something similar to the image below.
![image](https://user-images.githubusercontent.com/33374159/115548743-b976c300-a29f-11eb-9479-1f477e131f9a.png)

## Validating Terraform code with Github Actions

* Create a new workspace on Terraform Cloud.
* Then add `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` with their respective values as environment variables for the new workspace.
* Generate an API token named `GitHub Actions` from your Terraform Cloud User Settings, then add it to your Github repository as a secret. Name the secret `TF_API_TOKEN`.

* Clone and open your Github repo on your code editor.
* Checkout to a new branch with `git checkout -b <branch-name>`.
* Add your changes with `git add .` and commit the changes with a message using `git commit -m "<commit-message>"`.
* Then push your changes with `git push`.
* Go back to the Github repository and generate a pull request from the new branch. You can view the status of the job through the pull request created, Githbub Actions page or the Terraform Cloud organization.

## Author

* [Mariam Adedeji](https://github.com/mariehposa)

## Acknowledgments

* Flugel

terraform {
  backend "remote" {
    organization = "mariam-demo"

    workspaces {
      name = "mariam-demo"
    }
  }
}
