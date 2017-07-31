# dslink-cmd

DSLink command is a program to manage dslink repositories. This includes
the ability to generate scaffolding for DSA DSLink projects, and build a
distributable zip file. This tool is still very much in early development.
At the moment it only supports generating the scaffolding for Dart based
projects.

## Usage

### dslink help

Display help on the available commands.

```
dslink help <command>
```

### dslink build

Generate a .zip file for the dslink project. 
```
dslink-cmd build <language> [-o=project.zip]
```

The `language` must be a valid language for the dslink to build. Currently only `dart` is
supported. More languages will be available in the future.

The `-o` argument is optional and if omitted, the default build file will be
`build.zip`.

_Note: It does not currently support symlinked files._


_The following is currently unsupported_
### dslink init

Generates project scaffolding for a new dslink repository.

```
dslink init -project=<projectName> -node=<nodeName>
```

`projectName` is the name for the project sources. This will automatically
create a project with the name *dslink_projectname* and generate several files
by the name of *projectname.dart*.

`nodeName` is the default name for the nodes which will appear in DSA/DGLux for
this project.

This tool will generate the following files:

```
.gitignore
README.md
dslink.json
pubspec.yaml
bin\
bin\run.dart
lib\
lib\models.dart
lib\<projectName>.dart
lib\src\
lib\src\nodes\
lib\src\models\
```