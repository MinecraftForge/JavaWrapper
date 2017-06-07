#!/usr/bin/python

import json
import datetime as dt

tstamp = dt.datetime.now().strftime('%m-%d-%Y_%H:%M:%S')
mirrorfile = "example.mirrors.json"

print("Now updating the mirror timestamp", tstamp)

#with open(mirrorfile, "r") as jfile:
#    data = json.load(jfile)

with open(mirrorfile, "r+") as jFile:
    data = json.load(jFile)

    tmp = data["stamp"]
    data["stamp"] = tstamp
    jFile.seek(0)
    json.dump(data, jFile, indent=4, sort_keys=True)
    jFile.truncate()


print(data)
