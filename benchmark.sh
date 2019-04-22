#!/bin/sh

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

go tool pprof -svg cpu.prof
mv profile001.svg cpu.svg

go tool pprof -svg mem.prof
mv profile001.svg mem.svg

rm *.prof
rm *.test
