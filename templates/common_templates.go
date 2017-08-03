package templates

// README is the contents of the README.md file
const README = `# dslink-{{.Lang}}-{{.Name}}
## {{.Node}} DSLink

A DSLink for <complete this>

`

// DSLinkJSON is the contents of dslink.json file
const DSLinkJson = `{
  "name": "dslink-{{.Lang}}-{{.Name}}",
  "version": "0.0.1",
  "description": "{{.Node}} DSLink",
  "license": "Apache",
  "author": {
    "name": "<Your name>",
    "email": "<Your email>"
  },
  {{ if eq "dart" .Lang }}
  "main": "bin/run.dart",
  "engines": {
    "dart": ">1.15.0"
  },
  {{ else if eq "java" .Lang}}
  "main": "bin/dslink-{{.Lang}}-{{.Name}}",
  {{ end }}
  "repository": {
    "type": "git",
    "url": "https://<your repo address>"
  },
  "bugs": {
    "url": "https://<your bug tracker address>"
  },
  "configs": {
    "name": {
      "type": "string",
      "default": "{{.Node}}"
    },
    "broker": {
      "type": "url"
    },
    "token": {
      "type": "string"
    },
    "nodes": {
      "type": "path",
      "default": "nodes.json"
    },
    "key": {
      "type": "path",
      "default": ".key"
    },
    "log": {
      "type": "enum",
      "default": "info"
    }
    {{ if eq .Lang "java" }}
    "handler_class": {
      "type": "string",
      "default": "<your handler class>"
    }
    {{ end }}
  }
}

`