import json 
import sdl2
import sdl2.ext
import constants
import processJSON
import threading
import time
from collections import deque
colour = {
    0 :constants.backgroundColour,
    1 : constants.wallColour,
    2:constants.exitColour,
    3:constants.validPathColour,
    4:constants.zombiePathColour
}

def initWindow():
    sdl2.ext.init()
    window=sdl2.ext.Window("Simulation",(constants.width + 2*constants.borderSize,constants.height + 2*constants.borderSize))
    window.show()
    return window

def CreateMaze(renderer,jsonData):
    #render borders
    canvasWidth = constants.width + 2*constants.xScale
    canvasHeight = constants.height + 2*constants.yScale
    topBorder = (0,0,canvasWidth,constants.yScale)
    bottomBorder = (0,constants.height + constants.yScale,canvasWidth,constants.yScale)
    leftBorder = (0,0,constants.xScale,canvasHeight)
    rightBorder = (constants.width + constants.xScale,0,constants.xScale,canvasHeight)
    renderer.fill(topBorder,constants.wallColour)
    renderer.fill(bottomBorder,constants.wallColour)
    renderer.fill(leftBorder,constants.wallColour)
    renderer.fill(rightBorder,constants.wallColour)
    for x in range(len(jsonData['Maze'])):
        for y in  range(len(jsonData['Maze'][0])):
            block = (constants.xScale*x + constants.borderSize,constants.yScale*y + constants.borderSize,constants.xScale,constants.yScale)
            renderer.fill(block,colour[jsonData['Maze'][x][y]])

def addAgents(renderer,dynamicData):
    for humanLocation in dynamicData["HumanPositions"]:
        human = (constants.xScale*humanLocation["X"] + constants.borderSize,constants.yScale*humanLocation["Y"]+ constants.borderSize,constants.xScale,constants.yScale)
        renderer.fill(human,constants.humanColour)
    for zombieLocation in dynamicData["ZombiePositions"]:
        zombie = (constants.xScale*zombieLocation["X"] + constants.borderSize,constants.yScale*zombieLocation["Y"]+ constants.borderSize,constants.xScale,constants.yScale)
        renderer.fill(zombie,constants.zombieColour)

def cleanupFrame(renderer,dynamicData):
    for humanLocation in dynamicData["HumanPositions"]:
        human = (constants.xScale*humanLocation["X"] + constants.borderSize,constants.yScale*humanLocation["Y"]+ constants.borderSize,constants.xScale,constants.yScale)
        renderer.fill(human,constants.backgroundColour)
    for zombieLocation in dynamicData["ZombiePositions"]:
        zombie = (constants.xScale*zombieLocation["X"] + constants.borderSize,constants.yScale*zombieLocation["Y"]+ constants.borderSize,constants.xScale,constants.yScale)
        renderer.fill(zombie,constants.backgroundColour)

def getData(lock,socket,jsonBuffer):
    while True:
        jsonData = processJSON.receiveData(socket)
        if jsonData == False:
            print("breaking json receiver loop")
            break
        jsonBuffer.appendleft(jsonData)

def run():
    lock = threading.Lock()
    jsonBuffer = deque()
    socket = processJSON.connectToBackend(8080)
    # jsonReceiver = threading.Thread(target=getData,args=(lock,socket,jsonBuffer))
    # jsonReceiver.start()
    getData(lock,socket,jsonBuffer)
    print("got data")
    window = initWindow()
    renderer = sdl2.ext.Renderer(window)
    jsonData = jsonBuffer.pop()
    CreateMaze(renderer,jsonData)
    renderer.present()
    waiting = True
    while waiting:
        for event in sdl2.ext.get_events():
            if event.type == sdl2.SDL_MOUSEBUTTONDOWN:
                waiting = False
    while jsonBuffer:
        jsonData = jsonBuffer.pop()
        addAgents(renderer,jsonData)
        renderer.present()
        eventsList = sdl2.ext.get_events()
        for event in eventsList:
            if event.type == sdl2.SDL_QUIT:
                break
        #window.refresh()
        cleanupFrame(renderer,jsonData)
        time.sleep(0.03)
run()