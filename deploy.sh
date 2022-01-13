#!/usr/bin/env bash

remote=https://github.com/zrcoder/leetcode.git
if [ "$1" == "gitee" ]; then
remote=https://gitee.com/rdor/leetcode.git
fi

rm -rf public 

hugo

cd public

git init
git add -A
git commit -m "deploy"
git push -f ${remote} master:gh-pages

cd ..
