#!/bin/sh

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

go tool pprof -png cpu.prof
mv profile001.png cpu.png

go tool pprof -png mem.prof
mv profile001.png mem.png

rm *.prof
rm *.test
