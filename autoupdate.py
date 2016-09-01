#!/usr/bin/python
import subprocess, time
try:
    while True:
        subprocess.Popen(["git", "remote", "update"], stdout=subprocess.PIPE).communicate()[0]
        local = subprocess.Popen(["git", "rev-parse", r"@"], stdout = subprocess.PIPE).communicate()[0]
        remote = subprocess.Popen(["git", "rev-parse", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
        base = subprocess.Popen(["git", "merge-base", "@", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
        if local == remote:
            time.sleep(30)
            f = open("update.log", 'w')
            f.write("No updates.\n")
            f.close()
        elif local == base:
            f = open("update.log", 'w')
            f.write("Pulling changes.\n")
            f.write(subprocess.Popen(["git", "pull"], stdout = subprocess.PIPE).communicate()[0] + "\n")
            f.close()
except:
    f = open("update.log", 'w')
    f.write("Error: " + sys.exc_info()[0] + "\n")
    f.write("Exiting.\n")
    f.close()
