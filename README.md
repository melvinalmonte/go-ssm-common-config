# AWS Reusable client example

This documentation provides an example code that demonstrates how to load the AWS config and SSM client only once and use them to write and read data to and from AWS SSM.

## Prerequisites

To use the code provided in this documentation, you need the following:

- An AWS account
- AWS CLI installed
- AWS credentials set up on your machine
- Go language environment installed

## Code Overview

The code consists of three Go files:

### common.go

This file provides the `getAwsConfig()` function that initializes and returns the AWS configuration. The `GetSSMClient()` function creates an SSM client only once and returns it.

### services.go

This file provides two functions: `WriteToSSM()` and `ReadFromSSM()`. These functions respectively write a parameter to SSM and read a parameter from SSM.

### main.go

Our driver code, it calls the `WriteToSSM()` and `ReadFromSSM()` functions.

## How to Use the Code

1. Clone the repository to your local machine.
2. Navigate to the project root directory in the terminal.
3. Run the following command to install the project dependencies:

    ```bash
    go mod tidy
    ```

4. Run the following command to execute the program:

    ```bash
    go run main.go
    ```

The `main()` function writes a parameter to SSM and then reads the same parameter from SSM. You can modify the `paramStorePath` and `paramStoreValue` variables to write and read any parameter to and from SSM.

Main takeaway, using the `sync.once` mutex we are able to efficiently run our code once helping us load our AWS and SSM config only once when needed.
