#!/bin/sh

PROJECTDIR="$(pwd)"

echo $PROJECTDIR
echo Building server and running...
cd server/src && make build
cd $PROJECTDIR


echo Starting client... 
cd "$PROJECTDIR/client" && npm start

# cd back to the initial dir
cd $PROJECTDIR
