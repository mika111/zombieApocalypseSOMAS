import socket 
import json

def connectToBackend(address):
    print("waiting for connection")
    client = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    client.connect(("localhost",address))
    print("connected succesfully")
    return client

connection = connectToBackend(8080)

while True:
    jsonRaw = connection.recv(22*1100).decode('utf8')
    jsonData = json.loads(jsonRaw)
    print(jsonData)