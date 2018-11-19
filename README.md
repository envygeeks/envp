[![Code Climate](https://img.shields.io/codeclimate/maintainability/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/maintainability)
[![Code Climate](https://img.shields.io/codeclimate/c/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/coverage)
[![Travis CI](https://img.shields.io/travis/com/envygeeks/envp/master.svg?style=for-the-badge)](https://travis-ci.com/envygeeks/envp)
[![GitHub release](https://img.shields.io/github/release/envygeeks/envp.svg?style=for-the-badge)](http://github.com/envygeeks/envp/releases/latest)

# EnvP

EnvP is a simple CLI util that passes your file through Go-Template with your environment, allowing you to do more advanced configurations in things like Docker without much effort.  It also provides several helps that will aid you in this task, and make your life generally easy.

## Usage

| Flag | Type | Description |
|------|------|-------------|
| -glob   | bool   | search, and use a dir full of `*.gohtml` |
| -stdout | bool   | Print to stdout, instead of write |
| -output | string | the file to output to |
| -file   | string | the file, or dir |

## Helpers

| Helper | Description |
| ------ | ----------- |
| reindent | Reindent like `<<~` in Ruby |
| trimEdges | Trim "\n" or "\r\n" from the edges |
| trim | Trim a string of left, and right whitespace |
| indentedTemplate | Pull a template, and reindent it |
| trimmedTemplate | runs `trimEmpty`, and `trimEdges` on your template |
| trimEmpty | Trims empty lines with nothing but space to only `^$\n` |
| boolEnv | Pull an env var as a bool: 1/true, 0/false |
| templateString | fetch a template to a string |
| templateExists | Check if a template exists |
| envExists | Check if an env var exists |
| env | Pull an env var as a string |

## An Example

Given you did

```bash
export GHOST_PORT=8080
export GHOST_ENV=production
export CADDY_TLS_EMAIL=user@example.com
export CADDY_TLS=true

envp \
  -prefix=ghost \
  -file=ghost.gohtml \
  -stdout
```

And `ghost.gohtml` was

```gohtml
{{- define "hostnames" -}}
  {{- if eq (env "ghost_env") "development" -}}
    http://localhost
  {{- else -}}
    {{ $g := env "ghost_hostname" }}
    {{- if boolEnv "caddy_tls" -}}
      http://{{$g}} https://{{$g}}
    {{- else -}}
      http://{{$g}}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- define "tls" -}}
  {{- if and (ne (env "ghost_env") "development") (boolEnv "caddy_tls") -}}
    tls {{ env "caddy_tls_email" }}
  {{- end -}}
{{- end -}}

{{ template "tls" }}
{{ template "hostnames" }}
root /srv/caddy/ghost
ext .html .htm

{{ if and (envExists "ghost_port") (ne (env "ghost_port") "") }}
proxy localhost:{{ env "ghost_port" }} {
  transparent
  websocket
}
{{ end }}
```

You would get

```
tls user@example.com
http://example.com
root /srv/caddy/ghost
ext .html .htm

proxy localhost:8080 {
  transparent
  websocket
}
```
