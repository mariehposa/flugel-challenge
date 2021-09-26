# Flugel-challenge

A Terraform project to create an AWS EC2 instance and S3 bucket, tested with Terratest to check that both resources are tagged properly.

* Github Actions after creating a pull request
![image](https://user-images.githubusercontent.com/33374159/134812944-a718145f-c8bf-4382-ac85-5944d7e94c24.png)

* Github Actions after Merging to **main** branch
![image](https://user-images.githubusercontent.com/33374159/134813134-de0afa61-1514-40c3-9e26-eab69f72c1c2.png)

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

* Clone the repository, open up the project directory and `cd` into `flugel-challenge/test`.
* Then run `go mod init github.com/<github-username>/<github-repo-name>`.
* Now you can run the test with `go test -count=1 -v ./...`. The `-count=1` flag is used to avoid Go test caching and `-v` stands for verbose.
* At the end of the test run, your test should pass and you should see something similar to the image below.

![flugel-test](https://user-images.githubusercontent.com/33374159/134811887-24cda8bf-f28c-4f23-9c13-8c10d407cc6c.png)

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