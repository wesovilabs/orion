# Orion Examples

*Please feel free to create a Pull Request with new examples*

## Preconditions

Orion must be installed in your computes. [Visti the documentation](../installation)

## Examples

### Sections: given, when, then

- [example001](https://github.com/wesovilabs-tools/orion-examples/tree/master/example001): It's a basic scenario with sections `given`, `when` and `then`.
```bash
orion run --input example001/feature.hcl
```

- [example002](https://github.com/wesovilabs-tools/orion-examples/tree/master/example002): Scenario with multiple `when`, `then` sections.
```bash
orion run --input example002/feature.hcl
```

- [example003](https://github.com/wesovilabs-tools/orion-examples/tree/master/example003): Scenario without `given` block
```bash
orion run --input example003/feature.hcl
```

- [example004](https://github.com/wesovilabs-tools/orion-examples/tree/master/example004): The feature contains multiple scenarios
```bash
orion run --input example004/feature.hcl
```

### Input variables

- [feature-vars/feature001.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/feature-vars/feature001.hcl):  Basic example of feature with two input variables.
```bash
orion run --input feature-vars/feature001.hcl
orion run --input feature-vars/feature001.hcl --varibales feature-vars/variables001.hcl
```

- [feature-vars/feature002.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/feature-vars/feature002.hcl): Scenario that sums all the element in a list.
```bash
orion run --input feature-vars/feature002.hcl
```

- [feature-vars/feature003.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/feature-vars/feature003.hcl): Scenario that modifies the elements in a list, and it returns a new list.
```bash
orion run --input feature-vars/feature003.hcl
```

- [feature-vars/feature004.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/feature-vars/feature004.hcl): Scenario that modifies partially the elements in a list.
```bash
orion run --input feature-vars/feature004.hcl
```

### Scenarios with examples 

- [scenario-examples/feature001.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/scenario-examples/feature001.hcl):  A couple of scenarios that make use of attribute `examples`.
```bash
orion run --input feature-vars/feature001.hcl
```

### Usage of hooks

- [hooks/feature001.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/hooks/feature001.hcl):  Feature with global hooks`before each` and `after each` .
```bash
orion run --input hooks/feature001.hcl
```

- [hooks/feature002.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/hooks/feature002.hcl):  Feature with global hooks`before each` and `after each` and  hooks for tags
```bash
orion run --input hooks/feature002.hcl
```

### Importing content from other files 

- [includes/feature001.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/includes):  A full example of making use of attribute `includes`.
```bash
orion run --input includes/feature001.hcl
```

### Execution continue when scenario fails

- [ignore_errors/feature001.hcl](https://github.com/wesovilabs-tools/orion-examples/tree/master/ignore_errors/feature001.hcl):  The execution of the features continue even though a scenario fails.
```bash
orion run --input ignore_errors/feature001.hcl
```

### All the examples in the site can be found here

- [site/**](https://github.com/wesovilabs-tools/orion-examples/tree/master/site)
```