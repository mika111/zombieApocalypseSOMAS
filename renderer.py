import json
import pygame
jsonFile = open("state.json","r")
jsonData = json.load(jsonFile)
scaleX,scaleY = 1,1 #used to scale the display wrt to the size of the simulation map. 
borderSize = jsonData["BorderSize"] #thickness of border around map. 
exitColour = (255,255,255) 
zombieColour = (255,0,0) 
humanColour = (255,255,0) 
wallColour = (0,0,255)
def initialiseDisplay(stateData):
     width = borderSize+scaleX*stateData["MapSize"]["X"]
     height = borderSize+scaleY*stateData["MapSize"]["Y"]
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
     return screen,clock



def createWall(wall):
    height = scaleY*(wall["TopLeftCorner"]["Y"] - wall["BottomRightCorner"]["Y"])
    width = scaleX*(wall["BottomRightCorner"]["X"] - wall["TopLeftCorner"]["X"])
    return pygame.Rect(scaleX*wall["TopLeftCorner"]["X"]+borderSize,scaleY*wall["BottomRightCorner"]["Y"]+borderSize,width,height)

def createExit(exit):
    pointA = (exit["PointA"]["X"]+ borderSize,exit["PointA"]["Y"]+borderSize)
    pointB = (exit["PointB"]["X"]+borderSize,exit["PointB"]["Y"]+borderSize)
    #print(pointA,pointB)
    width = max(1,abs(pointA[0]-pointB[0]))    # In 2D space an exit is a line segment and therefore has only 1 dimension
    height = max(1,abs(pointA[1]-pointB[1]))     # Therefore either height or width will be 0 and we set it to 1 to render to screen
    top = min(pointA[1],pointB[1])
    left = min(pointA[0],pointB[0])
    #print(top,left,width,height)
    return pygame.Rect(left,top,width,height)

def generateFrame(screen,jsonData):
    for wall in jsonData["Walls"]:
        wallRect = createWall(wall)
        pygame.draw.rect(screen,wallColour,wallRect)
    for exit in jsonData["Exits"]:
        exitRect = createExit(exit)
        pygame.draw.rect(screen,exitColour,exitRect)
    for humanLocation in jsonData["HumanPositions"]:
        location = (scaleX*humanLocation["X"] + borderSize,scaleY*humanLocation["Y"]+ borderSize)
        pygame.draw.circle(screen,humanColour,location,1)
    for zombieLocation in jsonData["ZombiePositions"]:
        location = (scaleX*zombieLocation["X"]+borderSize,scaleY*zombieLocation["Y"]+borderSize)
        pygame.draw.circle(screen,zombieColour,location,1)
    pygame.display.flip()
    
screen,clock = initialiseDisplay(jsonData)
while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT: #quit on clicking X
            pygame.quit()
            break
    generateFrame(screen,jsonData)
    clock.tick(60) #lim to 60 fps
