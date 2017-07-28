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

### dslink build

Generate a .zip file for the dslink project. This command requires that
dart2js is in the `$PATH` variable

```
dslink build -o=project.zip
```

The `-o` argument is optional and if omitted, the default build file will be
`build.zip`.

Build will also copy any files and directories in the `data/` directory. _Note:
It does not currently support symlinked files._