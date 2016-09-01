#!/usr/bin/python
import subprocess, time
while True:
    subprocess.Popen(["git", "remote", "update"], stdout=subprocess.PIPE).communicate()[0]
    local = subprocess.Popen(["git", "rev-parse", r"@"], stdout = subprocess.PIPE).communicate()[0]
    remote = subprocess.Popen(["git", "rev-parse", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
    base = subprocess.Popen(["git", "merge-base", "@", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
    if local == remote:
        time.sleep(30)
        break
    elif local == base:
        print "Pulling changes."
        print subprocess.Popen(["git", "pull"], stdout = subprocess.PIPE).communicate()[0]