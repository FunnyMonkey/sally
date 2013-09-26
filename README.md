sally
=====

Open Content

Introduction
=======
This is the API portion of sally. This will consist of the methods and utilities
used to broker data.

Requirements
========
This project requires revel and mongodb in order to run.

Getting Started
========
1. It is expected that you have a MongoDB instance running on localhost.
1. Setup your Go environment http://golang.org/doc/, the main points to cover here are;
  1. Set your `GOPATH` environment variable. eg. `export GOPATH=/home/USERNAME/gocode` in your .bashrc or whatever file is appropriate.
  1. Add you `$GOPATH/bin` to your search path. eg. `export PATH=$PATH:$GOPATH/bin`
1. Install revel. `go get github.org/robfig/revel/revel`
1. `go get github.com/FunnyMonkey/sally` This will not work until this repository is public. In the meantime just use the following steps.

````
mkdir -p $GOPATH/src/github.com/FunnyMonkey &&
cd $GOPATH/src/github.com/FunnyMonkey &&
git clone git@github.com:FunnyMonkey/sally.git
````

Assuming everything has worked properly you should be able to start the app in dev mode `revel run github.com/FunnyMonkey/sally` this will start sally running on localhost on port 9000.

Usage
========

````
curl -H "Accept: text/plain" 127.0.0.1:9000/doc
curl -H "Accept: text/html" 127.0.0.1:9000/doc
curl -H "Accept: application/json" 127.0.0.1:9000/doc
````

[![Build Status](https://secure.travis-ci.org/FunnyMonkey/sally.png?branch=master)](http://travis-ci.org/FunnyMonkey/sally)
