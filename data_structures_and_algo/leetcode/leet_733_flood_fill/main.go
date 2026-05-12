package main

func main() {
    image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
    floodFill(image, 1, 1, 2)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if color == image[sr][sc] {
        return image
    }

    paint(image, sr, sc, image[sr][sc], color)
    return image
}

func paint(image [][]int, sr int, sc int, oldColor int, newColor int) {
    if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
        return
    }

    if image[sr][sc] == oldColor {
        image[sr][sc] = newColor

        paint(image, sr-1, sc, oldColor, newColor)
        paint(image, sr+1, sc, oldColor, newColor)
        paint(image, sr, sc-1, oldColor, newColor)
        paint(image, sr, sc+1, oldColor, newColor)
    }
}