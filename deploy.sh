#!/usr/bin/env bash

rm -rf public 

hugo

cd public

git init
git add -A
git commit -m "deploy"
git push -f https://github.com/zrcoder/leetcode.git master:gh-pages

cd ..
