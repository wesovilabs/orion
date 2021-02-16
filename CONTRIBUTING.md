# Contributing


## Pull Request Checklist

Before sending your pull requests, make sure you followed this list.

- Read [contributing guidelines](CONTRIBUTING.md).
- Read [Code of Conduct](CODE_OF_CONDUCT.md).
- Run unit tests (`make test`)
- Run linter checks ( `make lint`)
 
 
Before starting to code we recommend you to execute command `make setup`. This command downloads the dependencies, but it also
creates a set of useful git hooks.   
 
Keep in mind the below considerations:

- Commits must be descriptive and starts by one of the following prefixes `feat`, `fix`, `refactor`,  `test` and`docs`
- Directory `vendor` mustn't be pushed to repository.
- Write code thinking other people have to read it. 
- Update documentation when required. Documentation can be found in [docs/](/docs).



You can run any example with command

```bash
go run ./cmd/orion/orion.go run --input <INPUT_PATH>
go run ./cmd/orion/orion.go run --input <INPUT_PATH> --vars <VARIABLES_PATH>
```
