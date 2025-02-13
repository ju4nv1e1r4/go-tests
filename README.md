## Golang tests

This repo is my first contact with golang tests.

### Content

- Unit tests;
- Integration tests;
- "go test" command and "testing" package;
- Iteration in various test scenarios, using slice structures.

### Observations

**Output on Terminal**

In the screenshot below we have two scenarios: a failed test and a successful test, as well as the syntax in the linux/MacOS CLI.

![alt text](img/image.png)

**Cached**

![alt text](img/image1.png)

When you perform a test, the time is calculated and shown on the screen. When you run another identical test, without changing anything in the code of the previously run test, the return will be “cached”. When you run a Go test and see the message “(cached)”, it means that Go is reusing the results of a previous test run. In other words, the test wasn't actually run again, because Go assumed that the result would be the same.

How does Go know if the result will be the same?

Go has a cache system for tests. It stores information about previously executed tests, such as the code files that were compiled, the test results and the execution time. When you run a test again, Go checks to see if there have been any changes to the relevant files. If there are no changes, it simply loads the results from the cache, saving time and resources.

The test cache is invalidated when:

- Test files are modified: If you make changes to the code of your tests or to the files that the tests depend on, the cache will be invalidated and the tests will be run again;
- Dependencies are updated: If you update your project's dependencies, the cache will also be invalidated;
- You run the command go clean -testcache: This command manually clears the test cache.

**Command Line Tips**

Run all tests on package:
```bash
go test ./...
```

Run test on verbose mode:
```bash
go test -v
```

Run the percentual of your package is covered:
```bash
go test --cover
```

Take off a complete report of your tests:
```bash
go test --coverprofile result.txt

# Bring details about report
go tool cover --func=result.txt

# Bring a temporary HTML file explaining more
go tool cover --html=result.txt
```

![alt text](img/image2.png)