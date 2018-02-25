package main

const usage = `NAME
    zen -- a small CLI for interacting with zenhub/github

SYNOPSIS
    zen <command> [parameters]

DESCRIPTION
    zen is a small utility for interacting with ZenHub boards through a simple command line interface.

COMMANDS
    drop <issue>                Removes you as an assignee on the specified issue.
    list [parameters]           Lists all of the pipelines and issues for the current repository.
        parameters:
        [backlog]               The backlog pipeline is omitted from results unless "backlog" is supplied.
        [only <login>]          The list of issues will be filtered to only include the specified github login.
                                When this option is supplied, unassigned issues are still displayed.
    move <issue> [to] <pipeline>  Moves the specified issue from its current pipeline to the specified pipeline.
    pick up <issue>             Adds you as the assignee on the specified issue and removes any other assignees.
`
