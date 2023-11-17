# Package: internal

Internal package contains the internal code of the application and also implements the interfaces defined in the domain package.

## Code Convention

- Every domain should have its own sub-packages: `internal/<domain>/...`
- In each domain sub-package, there should be a `doc.md` file that describes the domain.
- Sub-packages consist of the following:
  - `internal/<domain>/repo`: contains the implementations of repository interfaces.
  - `internal/<domain>/ucase`: contains the implementations of use case interfaces.
  - `internal/<domain>/delivery`: contains the delivery mechanisms of the domain such as REST API, gRPC, etc.
- Every implementation has its own test file
- Test file should be named as `<filename>_test.go`
