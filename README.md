# GitHub Actions Crash Couse - Play ground

Welcome to the github action crash course playgroung! The aim of this repository is to learn how to use reusable workflow using 





## Challenges

### Run a workflow locally
1. Install act
2. Clone github-actions-call-reusable-workflow
3. Run locally the workflow from `/.github/workflows/challenge_1.yml`

### Create a composite action 
1. Modify workflow from `/.github/workflows/challenge_2.yml` to call composite action `hello_world_composite.yml`
2. Run succesfully `challenge_2.yml` workflow using act `act -j "hello_world_composite"`

See https://docs.github.com/en/actions/creating-actions/about-custom-actions#composite-actions

### Create a simple reusable workflow
The aim of this challenge is to implement a workflow that make use of the workflow defined in:
https://github.com/XoCedillo/crashcourse-reusable-workflows/

1. Clone https://github.com/XoCedillo/crashcourse-reusable-workflows repository
2. Read https://github.com/XoCedillo/crashcourse-reusable-workflows/blob/main/README.md to understand the structure of the reusable workflow
3. Modify the `github-actions-call-reusable-workflow/.github/workflows/challenge_3.yml` so it can work
4. Run succesfully `challenge_3.yml` workflow using act `act -j "go-workflow"`



## How to run workflow locally

Act is a tool that allows you to run Github actions locally. To install act, go to the repository and follow the installation steps:
https://github.com/nektos/act#installation


Yo can find the corresponding documentation in:
https://nektosact.com/beginner/index.html



To run an specific job you can try:
```act -j "<job_name>"  ```