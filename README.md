# Mastering Go

Go is a ....

## Noteworthy resources

- [Security Best Practices for Go Developers](https://go.dev/security/best-practices): Covers topics like fuzzy testing to increase resilience in your code by using auto generated input for tests. Use govulncheck to scan your code for known security vulnerabilities, which can also be included into the CI/CD pipeline. Use the race detector to find race conditions, ie when two or more goroutines access the same resource at once, which can lead to weird issues. Use the vet command to find potential issues like unused variables, unreachable code or common mistakes around goroutines.
