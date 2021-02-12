---
layout: default
title: Usage
nav_order: 2
---
<link rel="stylesheet" href="../../assets/css/custom.css">

# orion-cli

orion-cli is the command line interface for orion. Since the tool is still in progress
new command will be available in short time. So far, we can use `orion-cli lint` and `orion-cli lint`.   

## orion-cli help

Help about any command. Additionally, we can show the help of a specific command:

- `orion-cli help lint`
- `orion-cli help run`

## orion-cli lint

Verify the content of the input file is correctly written.

Available flags:
- `--input`: Path of input file. 

## orion-cli run

Execute the scenarios in the provided input file.

Available flags:
- `--input`: path of the input file.
- `--vars`: path of the variables file.
- `--verbose`: log level. Supported values are: 'DEBUG','INFO','WARN','ERROR' (default "INFO").