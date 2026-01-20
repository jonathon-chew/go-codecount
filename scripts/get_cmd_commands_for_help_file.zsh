#!/usr/bin/env bash

# Copy to clipboard the arguments that can be found in the cmd file in order to help build out the help function
grep "case" cmd/cmd.go | sed -E 's/.*case(.*):.*/aphrodite.PrintInfo(\1)/;s/,/+/' | pbcopy
echo "Placed on to the clipboard"
