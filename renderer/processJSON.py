import json
import socket

def connectToBackend(address):
    print("waiting for connection")
    client = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    client.connect(("localhost",address))
    print("connected succesfully")
    return client

def receiveData(socket):
    try:
        jsonData = socket.recv(10*1024).decode('utf8')
        jsonData = json.loads(jsonData)
        return jsonData
    except ConnectionResetError:
        print("all data sent")
        return False

