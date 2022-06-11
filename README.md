# go-binary-invokation-unit-testing
Figuring out best way to unit test functions which execute binaries in go

Wanted to test out various ways to test go functions which use `exec.Cmd` as that creates external dependancy on the os system
but still needs testing to check behaviors around failure conditions etc. the example function doesn't do anything specail other 
than expect to return bytes containing the output of the command and the error if one occurs when running. 

The TLDR pass a "context" which will be the interface of `exec.Cmd` the commands run in the function will be created off of this context so.
  
  real world -> exec.Command
  
  tests -> faked context to hyjact the creation of the executable =  `exec.Cmd`
  
  The next key is that you can have a test run another test as a sub process using `os.Args[0]` as during testing that will be the go unit 
  testing tool being run. So the general idea is to run the testing tool again targeting a single test that will then be written to either pass
  or fail depending on what you would like to test. Do you want to test what happens when your function runs and the executable it invokes fails
  great make the single test that it runs fail to emulate that. See [aws_test.go](aws/aws_test.go)
