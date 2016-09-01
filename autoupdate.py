#!/usr/bin/python
import subprocess, time, re, sys
from time import gmtime, strftime
if len(sys.argv) != 2:
    print sys.argv
    print "Please provide the amount of time (seconds) between each update."
    quit()
t = float(sys.argv[1])

def log(log_string):
    with open("update.log", 'a') as f:
        time = strftime("%Y-%m-%d %H:%M:%S", gmtime())
        indent = len(time + " - ") * " "
        strings = log_string.split("\n")
        f.write(time + " - " + re.sub(r'([^\n])$', r'\1\n', strings[0]))
        for s in strings[1:]:
            f.write(time + " - " + re.sub(r'([^\n])$', r'\1\n', s))

try:
    while True:
        subprocess.Popen(["git", "remote", "update"], stdout=subprocess.PIPE).communicate()[0]
        local = subprocess.Popen(["git", "rev-parse", r"@"], stdout = subprocess.PIPE).communicate()[0]
        remote = subprocess.Popen(["git", "rev-parse", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
        base = subprocess.Popen(["git", "merge-base", "@", r"@{u}"], stdout = subprocess.PIPE).communicate()[0]
        if local == remote:
            log("No updates.")
            time.sleep(t)
        elif local == base:
            log("Pulling changes.")
            log(subprocess.Popen(["git", "pull"], stdout = subprocess.PIPE).communicate()[0])
except Exception, e:
    print e
    log("Error: " + str(e) + "\n")
    log("Exiting.\n")
