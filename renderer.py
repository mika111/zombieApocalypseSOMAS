import json
import pygame
jsonFile = open("state.json","r")
jsonData = json.load(jsonFile)
scaleX,scaleY = 5,5

def createWall(wall):
    height = scaleY*(wall["TopLeftCorner"]["Y"] - wall["BottomRightCorner"]["Y"])
    width = scaleX*(wall["BottomRightCorner"]["X"] - wall["TopLeftCorner"]["X"])
    return pygame.Rect(scaleX*wall["TopLeftCorner"]["X"],scaleY*wall["TopLeftCorner"]["Y"],width,height)

def generateFrame(screen,jsonData):
    for wall in jsonData["Walls"]:
        rect = createWall(wall)
        pygame.draw.rect(screen,(0,0,255),rect)
    for humanLocation in jsonData["HumanPositions"]:
        location = (humanLocation["X"],humanLocation["Y"])
        pygame.draw.circle(screen,(0,255,0),location,scaleX)
    for zombieLocation in jsonData["ZombiePositions"]:
        location = (zombieLocation["X"],zombieLocation["Y"])
        pygame.draw.circle(screen,(255,0,0),location,scaleX)
    pygame.display.flip()
    
pygame.init()
clock = pygame.time.Clock()
screen = pygame.display.set_mode([scaleX*jsonData["MapSize"]["X"],scaleY*jsonData["MapSize"]["Y"]])
while True:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            pygame.quit()
            break
    generateFrame(screen,jsonData)
    clock.tick(60)