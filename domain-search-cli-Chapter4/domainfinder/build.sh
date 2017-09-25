#!/bin/bash
echo Building domainfinder...
go build -o domainfinder.o
echo Building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms.o
echo Building available...
cd ../available
go build -o ../domainfinder/lib/available.o
echo Building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle.o
echo Building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify.o
echo Building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify.o
echo Done.
