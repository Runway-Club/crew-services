# Package: Service

A place to put 3rd party services that are not part of the core application such as postgres db, redis, etc.

## Code Convention

- Use PascalCase for struct name, interface name, and public method name.
- Use camelCase for private method name.
- Use camelCase for variable name except for struct fields.
- Service does not contain any business logic. It is just a wrapper for the 3rd party service.
