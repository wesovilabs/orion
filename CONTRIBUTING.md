# Contributing

Visit [the open issues](https://github.com/wesovilabs/orion/issues) and feel free to take one of this.

This is only the beginning! By the way, If you missed any action,  or you found a bug, please [create a new issue on Github or vote the existing ones](https://github.com/wesovilabs/orion/issues)!


## Pull Request Checklist

Before sending your pull requests, make sure you followed this list.

- Read [contributing guidelines](CONTRIBUTING.md).
- Read [Code of Conduct](CODE_OF_CONDUCT.md).
- A pull request must b
- Run unit tests (`make test`)
- Run linter checks ( `make lint`)
- Code coverage must be equal or higher than the existing one. 
 
Before starting to code,  we recommend you to execute command `make setup`. This command downloads the dependencies, but it also
creates a set of useful git hooks.   
 
Keep in mind the below considerations:

- Commits must be descriptive and starts by one of the following prefixes `feat`, `fix`, `refactor`,  `test` and`docs`
- Directory `vendor` mustn't be pushed to repository.
- Write code thinking other people have to read it. 
- Update documentation when required. Documentation can be found in [docs/](/docs).


From the code, you can run any example with command

```bash
go run ./cmd/orion/orion.go run --input <INPUT_PATH>
go run ./cmd/orion/orion.go run --input <INPUT_PATH> --vars <VARIABLES_PATH>
```

## Others

Additionally, you can contribute to Orion by reviewing documentation and making it clearer for others  
or writing `Guides & Tutorials`. 

In case of you like Orion, please click on like button or sharing the project with your networks. 
