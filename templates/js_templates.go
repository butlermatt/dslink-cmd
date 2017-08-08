package templates

const JsIndex = `var DS = require('dslink');

// creates a node with an action on it
var Increment = DS.createNode({
  onInvoke: function(columns) {
    // get current value of the link
    var previous = link.val('/counter');

    // set new value by adding an amount to the previous amount
    link.val('/counter', previous + parseInt(columns.amount));
  }
});

// Process the arguments and initializes the default nodes.
var link = new DS.LinkProvider(process.argv.slice(2), 'template-javascript-', {
  defaultNodes: {
    // counter is a value node, it holds the value of our counter
    counter: {
      $type: 'int',
      '?value': 0
    },
    // increment is an action node, it will increment /counter
    // by the specified amount
    increment: {
      // references the increment profile, which makes this node an instance of
      // our Increment class
      $is: 'increment',
      $invokable: 'write',
      // $params is the parameters that are passed to onInvoke
      $params: [
        {
          name: 'amount',
          type: 'int',
          default: 1
        }
      ]
    }
  },
  // register our custom node here as a profile
  // when we use $is with increment, it
  // creates our Increment node
  profiles: {
    increment: function(path, provider) {
      return new Increment(path, provider);
    }
  }
});

// Connect to the broker.
// link.connect() returns a Promise.
link.connect().catch(function(e) {
  console.log(e.stack);
});

`

const JsInstall = `var fs = require('fs'),
    path = require('path'),
    crypto = require('crypto'),
    child = require('child_process');

function npmInstall() {
  var MD5_PATH = path.join(__dirname, ".dslink.md5");

  var file = fs.readFileSync(path.join(__dirname, "package.json"));

  var md5 = "";
  if(fs.existsSync(MD5_PATH)) {
    md5 = fs.readFileSync(MD5_PATH).toString("utf8");
  }

  var hash = crypto.createHash("md5");
  hash.update(file);
  var base = hash.digest("base64");

  if(base !== md5) {
    fs.writeFileSync(MD5_PATH, base);

    var npm = child.exec("npm install --production");
    console.log("running npm install");
    npm.stdout.on('data', function(data) {
      console.log(data);
    });
  }
}

npmInstall();

`

const JsPackageJson = `{
  "name": "dslink-{{.Lang}}-{{.Name}}",
  "version": "0.0.1",
  "description": "A template to kickstart creating a DSLink using the JavaScript SDK.",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "repository": {
    "type": "git",
    "url": "https://github.com/IOT-DSA/dslink-javascript-template.git"
  },
  "author": "",
  "license": "Apache",
  "bugs": {
    "url": "<Bugs Url>"
  },
  "homepage": "<Homepage Url>",
  "dependencies": {
    "dslink": "^1.0.0"
  }
}

`