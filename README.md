Rascal, n: Synonym for rogue.

Rascal is a very simple roguelike game. I like to re-implement it in
different languages, as an exercise in learning a new language more
deeply.

This is the go version.

BUILD & INSTALL

Install go 1.3 on your system. Get a source checkout; set GOPATH to that folder.

Then execute the following:
<pre>
PKG_CONFIG_PATH=~/local-pkg-config/ go get -v code.google.com/p/goncurses
go install rascal
./bin/rascal
</pre>

