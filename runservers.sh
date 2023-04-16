#!/bin/sh

PROJECTDIR="$(pwd)"

echo Starting server in the background...
cd "server/build/" && ./scriflow &
cd $PROJECTDIR

echo Starting client... 
cd "$PROJECTDIR/client"
npm start

# cd back to the initial dir
cd $PROJECTDIR
