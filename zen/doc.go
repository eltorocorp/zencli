package main

const usage = `NAME
    zen -- a small CLI for interacting with zenhub/github

SYNOPSIS
    zen <command> [parameters]

DESCRIPTION
    zen is a small utility for interacting with ZenHub boards through a simple command line interface.

COMMANDS
    close <issue>                    Changes the status of the specified issue to closed.
    create <title> as <pipeline>     Creates a new issue in the specified pipeline.
    drop <issue>                     Removes you as an assignee on the specified issue.
    list [parameters]                Lists all of the pipelines and issues for the current repository.
        parameters:
        [backlog]                    The backlog pipeline is omitted from results unless "backlog" is supplied.
        [only <login>|me]            The list of issues will be filtered to only include the specified github login.
                                     When this option is supplied, unassigned issues are still displayed.
                                     If "me" is supplied as the login, the current authenticated user's login is used.
    move <issue> [to] <pipeline>     Moves the specified issue from its current pipeline to the specified pipeline.
    open <issue>                     Changes the status of the specified issue to open.
    pick up <issue>                  Adds you as the assignee on the specified issue and removes any other assignees.

EXAMPLES
    To close an issue number 123:
        
        $ zen close 123

    To create a new issue in the 'prioritized' pipeline:

        $ zen create "There's clearly a bug in this code" as prioritized

    To create a new issue in the 'in progress' pipeline:

        $ zen create "This is another issue." as "in progress"

    To list only only my issues:

        $ zen list only me

    To move issue 999 to the "in progress" pipeline:

        $ zen move 999 to "in progress"
`
