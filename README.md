> ## Documentation
> 
* testify https://github.com/stretchr/testify
* viper https://github.com/spf13/viper 

## Detail repository
 
* setting defaults
* reading from JSON, ENV
* live watching and re-reading of config files (optional)
* reading from environment variables
* reading from remote config systems (etcd or Consul), and watching changes
* reading from command line flags
* reading from buffer
* setting explicit values 

## Install

**Note:** Install Golang https://go.dev/doc/install

```shell
go get github.com/stretchr/testify 
go get github.com/spf13/viper

```

## Structure Project 
 
 - project_root
 - configuration
    - .env
    - config.json
    - configuration.go
    - configuration_test.go


