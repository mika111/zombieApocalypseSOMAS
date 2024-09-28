import json
import pygame
jsonFile = open("state.json","r")
jsonData = json.load(jsonFile)
scaleX,scaleY = 1,1
exitColour = (255,255,255) #white
zombieColour = (255,0,0) #green
humanColour = (255,255,0) #
wallColour = (0,0,255)

def createWall(wall):
    height = scaleY*(wall["TopLeftCorner"]["Y"] - wall["BottomRightCorner"]["Y"])
    width = scaleX*(wall["BottomRightCorner"]["X"] - wall["TopLeftCorner"]["X"])
    return pygame.Rect(scaleX*wall["TopLeftCorner"]["X"],scaleY*wall["TopLeftCorner"]["Y"],width,height)

def createExit(exit):
    pointA = (exit["PointA"]["X"],exit["PointA"]["Y"])
    pointB = (exit["PointB"]["X"],exit["PointB"]["Y"])
    height = max(1,abs(pointA[0]-pointB[0]))    # In 2D space an exit is a line segment and therefore has only 1 dimension
    width = max(1,abs(pointA[1]-pointB[1]))     # Therefore either height or width will be 0 and we set it to 1 to render to screen
    top = max(pointA[1],pointB[1])
    left = min(pointA[0],pointB[0])
    return pygame.Rect(left,top,width,height)

def generateFrame(screen,jsonData):
    for wall in jsonData["Walls"]:
        wallRect = createWall(wall)
        pygame.draw.rect(screen,wallColour,wallRect)
    for exit in jsonData["Exits"]:
        exitRect = createExit(exit)
        pygame.draw.rect(screen,exitColour,exitRect)
    for humanLocation in jsonData["HumanPositions"]:
        location = (humanLocation["X"],humanLocation["Y"])
        pygame.draw.circle(screen,humanColour,location,1)
    for zombieLocation in jsonData["ZombiePositions"]:
        location = (zombieLocation["X"],zombieLocation["Y"])
        pygame.draw.circle(screen,zombieColour,location,1)
    pygame.display.flip()
    
#main render loop
pygame.init()
clock = pygame.time.Clock()
screen = pygame.display.set_mode([scaleX*jsonData["MapSize"]["X"],scaleY*jsonData["MapSize"]["Y"]])
while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT: #quit on clicking X
            pygame.quit()
            break
    generateFrame(screen,jsonData)
    clock.tick(60) #lim to 60 fps
