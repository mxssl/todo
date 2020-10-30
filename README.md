# todo

My example of a todo app.

## Development

### Task

Install [task](https://taskfile.dev). Use `task -l` to see all possible commands.

```text
task: Available tasks for this project:
* dc-build:             Run docker-compose build
* dc-up:                Run docker-compose up
* docker-lint:          Run linters in docker containers
* docker-release:       Build and push docker image
* github-release:       Goreleaser
* github-release-dry:   Goreleaser test
* lint:                 Run linters
* run:                  Run server for development
* swagger:              Run docker-compose up swagger
* swagger-generate:     Generate server based on a swagger file
* swagger-validate:     Validate swagger file
* test:                 Run unit tests
```

### Swagger

#### Validate swagger file

```sh
task swagger-validate
```

#### Generate code based on a swagger spec file

```sh
task swagger-generate
```
