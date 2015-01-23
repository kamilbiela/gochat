#!/bin/bash

fswatch -0 path | while read -d "" event 
do 
   go build && ./gochat
hr
done