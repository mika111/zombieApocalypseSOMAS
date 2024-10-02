import json

file = open("mazeData.json","r")
jsondata = json.load(file)
print(len(jsondata["mazeData"]))