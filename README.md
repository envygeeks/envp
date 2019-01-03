[![Code Climate](https://img.shields.io/codeclimate/maintainability/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/maintainability)
[![Code Climate](https://img.shields.io/codeclimate/c/envygeeks/envp.svg?style=for-the-badge)](https://codeclimate.com/github/envygeeks/envp/coverage)
[![Travis CI](https://img.shields.io/travis/com/envygeeks/envp/master.svg?style=for-the-badge)](https://travis-ci.com/envygeeks/envp)
[![GitHub release](https://img.shields.io/github/release/envygeeks/envp.svg?style=for-the-badge)](http://github.com/envygeeks/envp/releases/latest)

# EnvP

EnvP is a simple CLI util that passes your file through [golang/Template](https://golang.org/pkg/text/template) with your environment, allowing you to do more advanced configurations in things like Docker without very much effort.  It also provides several helpers that will aid in this task, and make your life generally easy.

## Usage

| Flag | Type | Description | Repeatable (Array) |
|------|------|-------------| ------------------ |
| --write-to | string | the file to output to | `false`
| --file   | string | the file, or dir | `true`

```
Build your configuration files with helpers, and access
to the current env, so that you can shim configuration files
in a Docker image when they do no support such mechanisms.

Usage:
  envp [flags]
  envp [command]

Available Commands:
  info        Show build info

Flags:
      --debug              verbose debug output
      --file stringArray   files to read in as templates
  -h, --help               help for envp
      --version            the current app version
      --write-to string    write to (stdout)
```

*Leaving `--file` empty will print the final result to stdout, this is really meant for testing before you make commits but can be used any way you wish.  As well, if you set `--file` to a directory, it will glob for `.gohtml` (even if it's gotxt)*

## Helpers
### split

```
{{ split [string] [delim] }}
```

```
{{ myStr := "1,2,3" }}
{{ range $_, $e := (split $myStr ",") }}
  {{ $e }}
{{ end }}
```

### chomp

```
{{ chomp [string] [cutset] }}
```

```
{{ myStr := "hello" }}
{{ chomp $myStr "o" }}
```

### indent

*Strips indentation to the edge at the smallest length, and reindents to the specified length.*

```
{{ indent [string] [n] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }
{{ end }}
{{ str := templateString "myTemplate" }}
{{ indent $str 6 }}
```

```
      hello {
        world
      }
```

### addSpace

*Add space to the beginning of a string*

```
{{ addSpace [string] [int] }}
```

### templateString

*Get a template as a string that can be manipulated*

```
{{ templateString [template] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }
{{ end }}
{{ $template := (templateString "myTemplate") }}
{{ indent $template 8 }}
```

### strippedTemplate

*Strips lines of empty space, and removes all edge space.*

```
{{ strippedTemplate [template] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }



{{ end }}
{{ strippedTemplate "myTemplate" }}
```

```
  hello {
    world
  }
```

### strip

*Strips lines of empty space, and removes all edge space.*

```
{{ strip [string] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }



{{ end }}
{{ $myVar := (templateString "myTemplate" }}
{{ strip $myVar }}
```

```
  hello {
    world
  }
```

### fixIndentedTemplate

*Strips indentation to the edge like `String#strip_heredoc` or `<<~STR` in Ruby.*

```
{{ fixIndentedTemplate [template] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }
{{ end }}
{{ fixIndentedTemplate "myTemplate" }}
```

```
hello {
  world
}
```

### templateExists

*Allows you to perform booleans on a template*

```
{{ templateExists [template] }}
```

```
{{ if (templateExists "myTemplate") }}
  {{ template "myTemplate" }}
{{ end }}
```

### fixIndentation

*Strips indentation to the edge like String#strip_heredoc or <<~STR in Ruby.*

```
{{ fixIndentation [string] }}
```

```
{{ define "myTemplate" }}
  hello {
    world
  }
{{ end }}
{{ myVar := (templateString "myTemplate") }}
{{ fixIndentation $myVar }}
```

```
hello {
  world
}
```

### envExists

*Lets you check if an environment variable exists.*

```
{{ envExists [key] }}
```

```
{{ if (envExists "key") }}
  Do Work
{{ end }}
```

### boolEnv

*Extracts an environment variable as a boolean.*

```
{{ boolEnv [key] }}
```

```
{{ if (boolEnv "key") }}
  Do Work
{{ end }}
```

### env

*Extracts an environment variable as a string.*

```
{{ env [key] }}
```

### randomPassword

*Generate an alphanumeric password using cryptographically derived random numbers.*

```
{{ randomPassword [length] }}
```

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
