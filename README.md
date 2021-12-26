# Terraform Provider Mattermost

The Terraform Mattermost provider is a plugin for Terraform that allows for the full lifecycle management of Mattermost resources.
This provider is currently early alpha quality and incomplete.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.17

## Using the provider

Documentation is available [here](https://registry.terraform.io/providers/ndrpnt/mattermost/latest/docs).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`.
This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.
Acceptance tests requires a real Mattermost instance and create real resources.
To spin up a local testing [Docker Compose](https://docs.docker.com/compose/) based
environment, run `make up`.

