## go work init ./hello

creates a workspsce for the module in the ./hello directory

## go work use ./example/hello

it is a method to use packages outside a module similar to go go mod edit -replace example.com.greetings=../greetings and go mod tidy command but this method is not commanly used.
