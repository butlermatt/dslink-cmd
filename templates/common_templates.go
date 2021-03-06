package templates

// README is the contents of the README.md file
const README = `# dslink-{{.Lang}}-{{.Name}}
## {{.Node}} DSLink

A DSLink for <complete this>

`

const LICENSE = `   Copyright <Year> <Person/Company>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

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
  {{ else if eq "java" .Lang }}
  "main": "bin/dslink-{{.Lang}}-{{.Name}}",
  {{ else if eq "javascript" .Lang }}
  "main": "index.js",
  "engines": {
    "node" : ">=0.10.0"
  },
  "getDependencies": [
    "node install.js"
  ],
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
  },
  "license": "Apache"
}

`

const GitIgnore = `# Logs
logs
*.log

# Runtime data
pids
*.pid
*.pids
*.seed
{{ if eq "dart" .Lang }}
# Files and directories created by pub
.buildlog
.packages
.project
.pub/
build/
packages

# Files created by dart2js
*.dart.js
*.part.js
*.js.deps
*.js.map
#.info.json

# Directory created by dartdoc
doc/api/

# Should include pubspec lock file as its an application.
#pubspec.lock
{{ else if eq "javascript" .Lang }}
# Directory for instrumented libs generated by jscoverage/JSCover
lib-cov

# Coverage directory used by tools like istanbul
coverage

# Grunt intermediate storage
.grunt

# node-waf configuration
.lock-wscript

# Compiled binary addons
build/Release

# Dependency directory
node_modules
{{ end }}

# DSA Specific ignores
nodes.json
.key
.dslink.key
.dslink.commit

`