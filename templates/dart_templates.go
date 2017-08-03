package templates

// DartPubSpec is the pubspec.yaml file required for Dart projects.
const DartPubSpec = `name: dslink_{{.Name}}
description: {{.Node}} DSLink
environment:
	sdk: ">=1.15.0 <2.0.0"
dependencies:
	dslink:
		git: https://github.com/IOT-DSA/sdk-dslink-dart.git

`

// Dart run is the template for the bin/run.dart file.
const DartRun = `import 'dart:async';

import 'package:dslink/dslink.dart';

import 'package:dslink_{{.Name}}/{{.Name}}.dart';

Future<Null> main(List<String> args) async {
	LinkProvider link;

	link = new LinkProvider(args, "{{.Node}}-", autoInitialize: false, profiles: {

	}, defaultNodes: {

	});

	link.init();

	await link.connect();
}


`