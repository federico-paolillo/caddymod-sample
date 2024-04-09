# Sample Caddy module

A Caddy module that provides a `my_mod` directive that adds `X-Hello-World` header to the response with a value specified in configuration.

## Directive

```
my_mod <header_value>

```

where `<header_value>` is the value of the `X-Hello-World` header.

To use the directive in a Caddyfile you have to specify the order of the directive in the global options, for example:

```
{
  order my_mod before respond
}

http://localhost:8080 {
  my_mod "Header Value!"
  respond "Ok!"
}

```

alternatively you can use a `route` directive:

```
http://localhost:65535 {
	route {
		my_mod "My-Mod-Header"
		respond "Hello, world!"
	}
}
```

## Build

You can build the Dockerfile to get a Caddy server version 2.7.6 with the `my_mod` directive embedded

## Contribute

To quickly try changes made to the module source code you can run `run.sh` which will invoke `xcaddy` and run a Caddy server instance with the configuration under `testdata/Caddyfile`.
