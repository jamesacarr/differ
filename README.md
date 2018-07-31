# Differ

## Description

Differ is a command line application that will find and open all of the modified files from a specific GitLab merge request.

## Usage

Before running Differ as an executable, it will need to be built and installed:

```sh
go install
```

```plain
Usage: bin/differ OPTIONS

OPTIONS

-e, --editor <application>  editor to open changed files [default: vscode]
-h, --help                  print help and exit
-m, --merge-id <id>         ID for the merge request
-p, --project-id <id>       ID for the GitLab project containing merge request
-t, --token <token>         your GitLab personal access token
-v, --version               print differ version
```
