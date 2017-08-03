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
dslink-cmd help <command>
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

_Note: It does not currently support symbolic link files._


### dslink init

Generates project scaffolding for a new dslink repository.

```
dslink-cmd init <projectLanguage> <projectName> [--node=<nodeName>] [--dir=<directory>] 
```

`projectLanguage` is the programming language that the project will be in. Supported values are
`dart`, `java`, `javascript`, `python`, `dotnet`, `c`, and `scala`. _Note: Currently only Dart is
implemented._

`projectName` is the name for the project sources. This will automatically
create a project with the name *dslink_projectname* and generate several files
by the name of *projectname.dart*.

Optional `nodeName` is the default name for the nodes which will appear in DSA/DGLux for
this project. This is often a capitalized version of the project name.

Optional `directory` will initialize the project in the specified directory. This will override
the default project directory of `dslink-<language>-<projectName>`

This tool will generate the following files for Dart. _Please be sure to check all files as some
may require additional information before working correctly_:

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