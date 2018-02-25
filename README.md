# zencli
A small utility for interacting with ZenHub boards through a simple command line interface.

## Want to contribute?

 - Awesome. Contributions are welcome!
 - Please fork and submit a pull request with any changes that you think would be useful.

## Setup

`zen` expects the following environment variables to be set:
 - ZENCLI_GITHUBAUTHTOKEN - https://github.com/settings/tokens Must have repo and user access.
 - ZENCLI_ZENHUBAUTHTOKEN - https://dashboard.zenhub.io/#/settings
 - ZENCLI_REPOOWNER - The name of the organization that owns the repo (i.e. eltorocorp).
 - ZENCLI_REPONAME - The name of the default repo you are targetting. (i.e. zencli)
 
## To build and install from source:

1. Clone the source to your working directory in your go path.
1. $ cd [...]/zencli/zen
1. $ go install
1. $ zen help
```
NAME
    zen -- a small CLI for interacting with zenhub/github

SYNOPSIS
    zen <command> [parameters]

DESCRIPTION
    zen is a small utility for interacting with ZenHub boards through a simple command line interface.

COMMANDS
    ...
    ...
    ...
```

## yeah, I know
 - I know about the `flag` package. I wrote the custom parser for this just for the hell of it.
 - I know there are github API wrappers out there already for Go. I wanted to keep things simple and avoid vendored dependencies.
