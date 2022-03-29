# sqlc
## Description
This meeting was dedicated to sqlc tool. The tool generates fully type-safe idiomatic Go code from SQL. The workflow for development with that tool is the folowing:
1. Write SQL queries 
2. Run sqlc to generate Go code that presents type-safe interfaces to those queries
3. Write application code that calls the methods sqlc generated

Further information can be found in the [documentation](https://docs.sqlc.dev/en/latest/index.html).
## Key Takeaways
We concluded that the tool can be used in cases where SQL queries are known upfront, but it's not suitable for dealing with dynamic queries, e.g. when you need to search data by multiple attributes which can be optional. 
We also noted that there must be only one [TestMain](https://pkg.go.dev/testing#hdr-Main) function per package. And we discussed the ability to omit the call to `os.Exit` in the new versions of `testing` package, which allows you to use `defer` statement. This [issue](https://github.com/golang/go/issues/34129) provides more information about that.
