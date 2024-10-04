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
        jsonData = socket.recv(5*1024).decode('utf8')
        jsonData = json.loads(jsonData)
        return jsonData
    except ConnectionResetError:
        print("failed to get data")
        return False

