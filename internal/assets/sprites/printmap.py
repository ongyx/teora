import sys

from PIL import Image


def main():
    if len(sys.argv) < 3:
        print(f"usage: {sys.argv[0]} image tile_size")
        sys.exit(1)

    img = Image.open(sys.argv[1])
    tile_size = int(sys.argv[2])

    tiles_x = img.width // tile_size
    tiles_y = img.height // tile_size
    tiles = tiles_x * tiles_y

    if tiles_x < 1 or tiles_y < 1:
        raise ValueError("image smaller than tile size")

    print(f"image size: {img.size}, tile size: {tile_size}, total tiles: {tiles}")

    pad = len(str(tiles))
    for y in range(tiles_y):
        for x in range(tiles_x):
            index = (tiles_x * y) + x
            print(f"{index: >{pad}}", end=" ")
        print()


if __name__ == "__main__":
    main()
