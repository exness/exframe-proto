# exframe-proto

Protobuf extensions used by the exFrame Go framework.

## exframe.auth.v1

`exframe/auth/v1/auth.proto` declares the method option `(exframe.auth.v1.method)`, the access policy of an RPC method:

- `scopes`: access is granted when the caller has any of the scopes;
- `public`: requests that fail authorization are let through and counted.

A method without the option is closed for every caller.

```proto
import "exframe/auth/v1/auth.proto";

service ContactService {
  rpc GetUserPhoneV1(GetUserPhoneV1Request) returns (GetUserPhoneV1Response) {
    option (exframe.auth.v1.method) = {
      scopes: ["contact-service.phone.read", "contact-service.phone.read-write"]
    };
  }
}
```

| `public` | `scopes` | Caller without a required scope | Caller with a required scope |
|---|---|---|---|
| false | set | denied | allowed |
| false | empty | denied | denied |
| true | set | allowed and counted | allowed |
| true | empty | allowed | allowed |

## Usage

buf: add the module to `deps` in `buf.yaml` and run `buf dep update`.

```yaml
deps:
  - buf.build/exness/exframe-proto
```

Go: the generated code imports `github.com/exness/exframe-proto/exframe/auth/v1`, no `go_package` override is needed.

```sh
go get github.com/exness/exframe-proto
```

## Development

```sh
buf lint
buf format -d --exit-code
buf generate
go test ./...
```

The extension number `52008` and the field numbers of `MethodAuth` never change.
