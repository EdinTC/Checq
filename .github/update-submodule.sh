#!/bin/bash

#
# Use this script to update the submodule(s)
#
echo "Enter the main folder.."
cd ../
echo "Run update:"
git submodule foreach git pull origin master
