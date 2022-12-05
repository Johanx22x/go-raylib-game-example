package main

import "github.com/gen2brain/raylib-go/raylib"

// Game constants
const (
    screenWidth  = 1280
    screenHeight = 720
)

// Game variables
var (
    running = true
    bgColor = rl.NewColor(147, 211, 196, 255)
    
    grassSprite rl.Texture2D
    playerSprite rl.Texture2D

    playerSrc rl.Rectangle
    playerDest rl.Rectangle

    playerSpeed float32 = 3

    musicPaused bool
    music rl.Music
)

// This function draws a sprite to the screen 
func drawScene() {
    rl.DrawTexture(grassSprite, 100, 50, rl.White)
    rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

// This function is called every frame
// It is used to get user input
func input() {
    if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
        playerDest.Y -= playerSpeed
    }

    if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
        playerDest.Y += playerSpeed
    }

    if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
        playerDest.X -= playerSpeed
    }

    if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
        playerDest.X += playerSpeed
    }

    if rl.IsKeyPressed(rl.KeyQ) {
        musicPaused = !musicPaused // Toggle musicPaused
    }
}

// Update function is called every frame
// It is used to update the game logic and variables
// Also, it is used to check if the game should be closed
func update() {
    running = !rl.WindowShouldClose()

    rl.UpdateMusicStream(music)
    if musicPaused {
        rl.PauseMusicStream(music)
    } else {
        rl.ResumeMusicStream(music)
    }
}

// This function is called every frame
// It is used to render the scene to the screen 
func render() {
    rl.BeginDrawing()

    rl.ClearBackground(bgColor)

    drawScene()

    rl.EndDrawing()
}

// init function is called before main
// It is used to initialize the game window and other properties
func init() {
    rl.InitWindow(screenWidth, screenHeight, "Sprout Game Example")
    rl.SetExitKey(0)
	rl.SetTargetFPS(60)

    grassSprite = rl.LoadTexture("res/tilesets/grass.png")
    playerSprite = rl.LoadTexture("res/characters/basicCharacterSpriteSheet.png")

    playerSrc = rl.NewRectangle(0, 0, 48, 48)
    playerDest = rl.NewRectangle(200, 200, 100, 100)

    rl.InitAudioDevice()
    music = rl.LoadMusicStream("res/music.mp3")
    musicPaused = false
    rl.PlayMusicStream(music)
}

// This is the main function
// It is used to run the game 
func main() {
    // Main game loop
	for running {
        // Get user input
        input()
        // Update game logic
        update()
        // Render the scene
        render()
	}

    // Free up resources
    exit()
}

// This function is called when the game is closed
// It is used to free up resources
func exit() {
    rl.UnloadTexture(grassSprite)
    rl.UnloadTexture(playerSprite)
    rl.UnloadMusicStream(music)
    rl.CloseAudioDevice()
    rl.CloseWindow()
}
