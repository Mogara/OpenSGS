# OpenSGS

## Get Started
### Requirements
- Go 1.17+
- Go modules enabled
### Run Locally
1. Clone the repository
    ```bash
    git clone https://github.com/Mogara/OpenSGS.git
    ```
2. Install dependencies
    ```bash
    cd OpenSGS
    go mod download
    ```
3. Run the server
    ```bash
    go run cmd/server/main.go
    ```

## Development
1. Setup Go environment([v1.17](https://go.dev/doc/install))
2. Install [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
3. Install [pre-commit](https://pre-commit.com/#install)
4. Install `pre-commit` hooks
    ```bash
    # install commit-msg hook
    pre-commit install --hook-type commit-msg 
    pre-commit install --hook-type pre-push
    # install pre-commit hook
    pre-commit install
    # test hooks
    pre-commit run --all-files
    ```