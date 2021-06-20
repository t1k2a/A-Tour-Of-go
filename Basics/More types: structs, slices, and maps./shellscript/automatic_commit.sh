#!/bin/bash
#--------------------------------------------------------------
#GitHubにコミット&プッシュ
#--------------------------------------------------------------
git add .
git commit --allow-empty-message -m ''
git push origin master
