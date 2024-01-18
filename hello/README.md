## go mod init example.com/hello

used to initialize a module named hello with example.com as a published location

## go mod edit -replace example.com.greetings=../greetings

It is used to replace the package path with the local location of the package as it is only in local environment and not publish

## go mod tidy

this command synchronize the module with all the dependencies required by the packages, but not yet tracked by the module.

In case of referencing a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.

## require example.com/greetings v1.1.0
