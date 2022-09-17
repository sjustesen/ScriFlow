#!/bin/sh

PROJECTDIR="$(pwd)"

echo Building server and running...
cd "server/golang/src/" && ./scriflow
cd $PROJECTDIR


echo Starting client... 
cd "$PROJECTDIR/client" && npm start

# cd back to the initial dir
cd $PROJECTDIR
