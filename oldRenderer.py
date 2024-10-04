import json
import pygame
import socket
import time
scaleX,scaleY = 10,10 #used to scale the display wrt to the size of the simulation map. 
borderSize = 10 #thickness of border around map. 
exitColour = (255,0,255) 
zombieColour = (100,100,100) 
validPathColour = (0,255,0)
humanColour = (255,255,0) 
wallColour = (255,255,255)
backgroundColour = (0,0,0)
zombiePathColour = (255,0,0)
initialState = 0
zombieSize = 5
humanSize = 5
jsonData = 0
color = {
    0 :backgroundColour,
    1 : wallColour,
    2:exitColour,
    3:validPathColour,
    4:zombiePathColour
}


def initialiseDisplay(initialState):
     
     width = 2*borderSize+scaleX*len(initialState["Maze"])
     height = 2*borderSize+scaleY*len(initialState["Maze"][0])
     pygame.init()
     clock = pygame.time.Clock()
     screen = pygame.display.set_mode([width,height])
     topBorder = pygame.Rect(0,0,width,borderSize)
     bottomBorder = pygame.Rect(0,height-borderSize,width,borderSize)
     rightBorder = pygame.Rect(width-borderSize,0,borderSize,height)
     leftBorder = pygame.Rect(0,0,borderSize,height)
     pygame.draw.rect(screen,wallColour,topBorder)
     pygame.draw.rect(screen,wallColour,bottomBorder)
     pygame.draw.rect(screen,wallColour,rightBorder)
     pygame.draw.rect(screen,wallColour,leftBorder)
     for x in range(len(jsonData['Maze'])):
        for y in range(len(jsonData['Maze'][0])):
            rect = pygame.Rect(x*scaleX+borderSize,y*scaleY+borderSize,scaleX,scaleY)
            pygame.draw.rect(screen,color[jsonData['Maze'][x][y]],rect)
     pygame.display.flip()
     return screen,clock

def generateFrame(screen,jsonData):
    for humanLocation in jsonData["HumanPositions"]:
        location = (scaleX*humanLocation["X"] + 0.5*scaleX + borderSize,scaleY*humanLocation["Y"]+0.5*scaleY + borderSize)
        pygame.draw.circle(screen,humanColour,location,humanSize)
    for zombieLocation in jsonData["ZombiePositions"]:
        location = (scaleX*zombieLocation["X"]+ 0.5*scaleX +borderSize,scaleY*zombieLocation["Y"]+ 0.5*scaleY +borderSize)
        pygame.draw.circle(screen,zombieColour,location,zombieSize)
    pygame.display.flip()

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

connection = connectToBackend(8080)
newJsonData = receiveData(connection)
if newJsonData != False:
    jsonData = newJsonData
screen,clock = initialiseDisplay(jsonData)
while True:
    newJsonData = receiveData(connection)
    if newJsonData != False:
        jsonData = newJsonData
    # for event in pygame.event.get():
    #     if event.type == pygame.QUIT: #quit on clicking X
    #         pygame.quit()
    #         break
    generateFrame(screen,jsonData)
    clock.tick(10) #lim to 60 fps
