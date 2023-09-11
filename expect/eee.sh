#!/bin/bash

expect <<!
source ~/.bash_profile
psql -h 42.159.87.142 -p 5432 -d postgres -U oushu
expect "*Password*"
send "test"
expect "*ostgres*"
send "\q"
expect eof
!
