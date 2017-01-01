#!/bin/bash

#The zip must 

echo "Building JavaWrapper"
go build
echo "zipping install directory"
zip -r0 installer.zip ./installer
echo "appending the wrapper"
cat JavaWrapper >> installer.zip
echo "Moving JavaWrapper to a test direcory"
mkdir test
mv JavaWrapper test/JavaWrapper
cd test
./JavaWrapper



