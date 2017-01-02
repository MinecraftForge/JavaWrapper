#!/bin/bash

#The zip must 

echo "Building JavaWrapper"
go build
echo "zipping install directory"
zip  installer.zip installer
echo "appending the wrapper"
cat installer.zip JavaWrapper
echo "Moving JavaWrapper to a test direcory"
mkdir test
mv JavaWrapper test/JavaWrapper
cd test
./JavaWrapper



