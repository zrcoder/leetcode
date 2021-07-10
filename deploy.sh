#!/usr/bin/env bash

rm -rf _book

gitbook build .

cd ./_book

git init
git add -A
git commit -m "deploy"
git push -f https://github.com/zrcoder/leetcode.git master:gh-pages

cd ..
