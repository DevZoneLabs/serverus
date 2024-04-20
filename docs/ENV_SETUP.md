# Environment Setup

## General Configuration

1. [Install ASDF for Managing Versions (Mac / Linux)](https://asdf-vm.com/guide/getting-started.html)

    It is preferred to do the installation via curl.

    ```
    git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.0
    ```

    Add the following to the top of your ~/.zshrc 
    ```
    # ASDF
    . "$HOME/.asdf/asdf.sh"
    # append completions to fpath
    fpath=(${ASDF_DIR}/completions $fpath)
    # initialise completions with ZSH's compinit
    autoload -Uz compinit && compinit
    ```

    Apply changes
    ```
    source ~/.zshrc
    ```

    Installing GoLang through ASDF

    ```
    asdf plugin add golang

    asdf install golang 1.22.2

    asdf global golang 1.22.2

    # Verifying Installation
    go version
    ```

2. [Install Golang for Windows Version 1.22.2](https://go.dev/doc/install)

2. [Installing Docker](https://docs.docker.com/engine/install/)


## Running a Service