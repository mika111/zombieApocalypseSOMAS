import json
import pygame
jsonFile = open("state.json","r")
jsonData = json.load(jsonFile)
scaleX,scaleY = 5,5 #used to scale the display wrt to the size of the simulation map. 
borderSize = 5 #thickness of border around map. 
exitColour = (255,0,255) 
zombieColour = (255,0,0) 
humanColour = (255,255,0) 
wallColour = (255,255,255)
backgroundColour = (0,0,0)
def initialiseDisplay(stateData):
     width = 2*borderSize+scaleX*stateData["MapSize"]["X"]
     height = 2*borderSize+scaleY*stateData["MapSize"]["Y"]
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

def generateFrame(screen,jsonData):
    color = {
        0 :backgroundColour,
        1 : wallColour,
        2:exitColour,
    }
    for x in range(len(jsonData['Maze'])):
        for y in range(len(jsonData['Maze'][0])):
            rect = pygame.Rect(x*scaleX+borderSize,y*scaleY+borderSize,scaleX,scaleY)
            pygame.draw.rect(screen,color[jsonData['Maze'][x][y]],rect)

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
