---
layout: default
title: Usage
nav_order: 2
---
<link rel="stylesheet" href="../../assets/css/custom.css">

# orion

orion is the command line interface for orion. Since the tool is still in progress
new command will be available in short time. So far, we can use `orion lint` and `orion lint`.   

## orion help

Help about any command. Additionally, we can show the help of a specific command:

- `orion help lint`
- `orion help run`

## orion lint

Verify the content of the input file is correctly written.

Available flags:
- `--input`: Path of input file. 

## orion run

Execute the scenarios in the provided input file.

Available flags:
- `--input`: path of the input file.
- `--vars`: path of the variables file.
- `--tags`: comma separated list of tag names. Scenarios containing any of the listed tags will be executed.
- `--verbose`: log level. Supported values are: 'DEBUG','INFO','WARN','ERROR' (default "INFO").
