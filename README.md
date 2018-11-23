[![Code Climate](https://img.shields.io/codeclimate/maintainability/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/maintainability)
[![Code Climate](https://img.shields.io/codeclimate/c/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/coverage)
[![Travis CI](https://img.shields.io/travis/com/envygeeks/envp/master.svg?style=for-the-badge)](https://travis-ci.com/envygeeks/envp)
[![GitHub release](https://img.shields.io/github/release/envygeeks/envp.svg?style=for-the-badge)](http://github.com/envygeeks/envp/releases/latest)

# EnvP

EnvP is a simple CLI util that passes your file through Go-Template with your environment, allowing you to do more advanced configurations in things like Docker without much effort.  It also provides several helps that will aid you in this task, and make your life generally easy.

## Usage

| Flag | Type | Description |
|------|------|-------------|
| -output | string | the file to output to |
| -file   | string | the file, or dir |

## Helpers
### reindent

*Reindent like `<<~` in Ruby or `String#strip_heredoc` in Rails.  Reindent will strip the shortest indentation across all lines, bringing your text to the edge, while keeping sub-indentation. This function will also run `trimEdges`, and `trimEmpty` to ensure a clean indent.*

```
{{ reindent $myStr }}
```

### trimEdges

*Strip `\r\n`, `\n`, `\t`, `\s` from the edges of a string (the top, and the bottom (multi-line), or left, and right (single line)) leaving a clean string to work with, without all the nonsense spacing.*

```
{{ trimEdges $myStr }}
```

### indentedTemplate

*Pulls a template, and runs `reindent` on it, returning the cleaned up template for your golden template to use. **Since this is not a builtin you can also capture this to a variable***

```
{{- define myTemplate -}}
  1
    2
    3
  4
{{- end }}
```

```
{{ indentedTemplate "myTemplate" }}
```

```
1
  2
  3
4
```

### trimmedTemplate

*Pulls a template, and runs `trimEdges`, and `trimEmpty` on it, returning the cleaned up template for your golden template to use. **Since this is not a builtin you can also capture this to a variable***

```
{{- define myTemplate -}}


  1
    2
    3
  4


{{- end }}
```

```
{{ trimmedTemplate "myTemplate" }}
```

```
  1
    2
    3
  4
```

### trimEmpty

*Trim a string's empty lines of space, and only of space, leaving just a truly blank `\n` for you to work with, this is particularly useful for reindenting, where we need to strip that so it doesn't affect how we detect indentation.*

```
{{ myStr := "Hello\n  \nWorld"
{{ trimEmpty $myStr }}
```

```
Hello

World
```

### indent

*Strip all indentation to the edge, and then indent to n<int> you send to us, allowing you to deeply indent within define, or in configuration files in a `{}` or otherwise.*

```
{{ define "myTemplate" }}
  1
    2
  3
{{ end }}
```

```
{{ indent (templateString "myTemplate") 4 }}
```

```
    1
      2
    3
```

### boolEnv
### templateString
### templateExists
### envExists
### env

### Stdlib

* `trim` -> https://golang.org/pkg/strings/#Trim

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
