[![Code Climate](https://img.shields.io/codeclimate/maintainability/envygeeks/jekyll-assets.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/maintainability)
[![Code Climate](https://img.shields.io/codeclimate/c/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/coverage)
[![Travis CI](https://img.shields.io/travis/com/envygeeks/envp/master.svg?style=for-the-badge)](https://travis-ci.org/envygeeks/jekyll-assets)
[![GitHub release](https://img.shields.io/github/release/envygeeks/envp.svg?style=for-the-badge)](http://github.com/envygeeks/envp/releases/latest)

# EnvP

EnvP is a simple CLI util that passes your file through Go-Template with your environment, allowing you to do more advanced configurations in things like Docker without much effort.  It also provides several helps that will aid you in this task, and make your life generally easy.

## Usage

```
Usage of envp:
  -file string the file, or dir
  -glob search, and use a dir full of *.gohtml
  -output string the file to write to
  -stdout print to stdout
```

## Helpers

* `eBool` - Pull an env var as a bool: 1/true, 0/false
* `trimStr` - Trim a string of left, and right whitespace
* `templateExists` - Check if a template exists
* `eExists` - Check if an env var exists
* `eStr` - Pull an env var as a string

## An Example

```gohtml
{{- define "hostnames" -}}
  {{- if eq (eStr "GHOST_ENV") "development" -}}
    http://localhost
  {{- else -}}
    {{ $g := eStr "GHOST_HOSTNAME" }}
    {{- if eBool "CADDY_SSL" -}}
      http://{{$g}} https://{{$g}}
    {{- else -}}
      http://{{$g}}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- define "tls" -}}
  {{- if and (ne (eStr "GHOST_ENV") "development") (eBool "CADDY_SSL") -}}
    {{- if templateExists "ssl.gohtml" -}}
      {{ template "ssl.gohtml" }}
    {{- else -}}
      tls {{ eStr "CADDY_SSL_EMAIL" }}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{ template "tls" }}
{{ template "hostnames" }}
root /srv/caddy/ghost
ext .html .htm

{{ if and (eExists "GHOST_PORT") (ne (eStr "GHOST_PORT") "") }}
proxy localhost:{{ eStr "GHOST_PORT" }} {
  transparent
  websocket
}
{{ end }}
```

With:

```
export GHOST_PORT=8080
export GHOST_ENV=production
export CADDY_SSL_EMAIL=user@example.com
export CADDY_SSL=true
./envp -file=./test \
  -stdout
```

Results in

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
