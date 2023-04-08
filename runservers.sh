#!/bin/sh

PROJECTDIR="$(pwd)"

echo Starting server and running...
cd "server/build/" && ./scriflow
cd $PROJECTDIR

echo Starting client... 
cd "$PROJECTDIR/client"
npm start

# cd back to the initial dir
cd $PROJECTDIR
