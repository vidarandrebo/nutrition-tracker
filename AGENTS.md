# Repository guidance

This repository contains a Go API in `api/` and a client in `client/`. The Go module is rooted at `api/`; run Go commands from there.

## Go error handling

- Check and return errors explicitly. Do not silently discard database, scan, iteration, or cleanup errors; do not use `panic` for expected failures.
- Document the errors each error-returning method can return in its method comment. Name a specific database error type or sentinel when it is known from the database API call, rather than only saying "database error"; do not claim a particular type for driver-dependent failures. Say whether errors are propagated or translated, and document intentional no-match behavior.
- `database/sql`'s `QueryRow(...).Scan(...)` returns `sql.ErrNoRows` when no row is returned; mention it in the method comment wherever applicable, including `INSERT ... RETURNING` calls. List queries with no matching rows normally return an empty result, not `sql.ErrNoRows`; deletes with no matching row normally return no error. Use `errors.Is` to recognize wrapped sentinel errors.
- Return database errors directly from the repository layer to the service layer, without translating them into internal errors. If adding context in a repository, wrap the original error with `%w` so the service can inspect it with `errors.Is` or `errors.As`.
- At the service boundary, log the original database error with relevant, non-sensitive context before translating it to a client-safe predefined internal error (`utils.ErrNotFound`, `utils.ErrForbidden`, `utils.ErrUnauthorized`, or `utils.ErrUnknown`). Do not propagate database errors or implementation details beyond the service layer, and do not treat unexpected database failures as missing resources.
- Ownership checks belong in the service layer: a missing resource maps to `utils.ErrNotFound`, another owner's resource to `utils.ErrForbidden`, and an unexpected lookup failure to `utils.ErrUnknown`.
- Keep authentication and authorization distinct: `utils.ErrUnauthorized` means valid authentication credentials are required (HTTP 401 semantics); `utils.ErrForbidden` means access is denied despite identification (HTTP 403 semantics). Do not expose internal database details in HTTP responses.