import pathlib
import string
import sys
import tempfile
from collections.abc import Iterator, Sequence
from typing import cast

import fontforge
import svg
from PIL import Image

SPRITESHEET = "teoran.png"

PIXEL = 200
CHARS = [
    *string.ascii_uppercase,
    *string.ascii_lowercase,
    *string.digits,
    # Special characters not present in the original font.
    "!",
    "?",
    ".",
    ",",
    ";",
    "/",
    "<",
    ">",
]

FONT_FAMILY = "Teoran"
FONT_NAME = f"{FONT_FAMILY}Standard"
EM = 1200


def sprites(sheet: Image.Image) -> Iterator[Image.Image]:
    # Since fonts are supposed to be grayscale, only the alpha channel is relevant.
    sheet = sheet.getchannel("A")

    width, height = sheet.size
    if height % width != 0:
        raise ValueError(
            f"sprites sheet must have a height divisible by width ({sheet.size =})"
        )

    for offset in range(0, height, width):
        yield sheet.crop((0, offset, width, offset + width))


def sprite_to_svg(sprite: Image.Image) -> svg.SVG:
    width, height = sprite.size
    if width != height:
        raise ValueError(f"sprite must be a fixed size ({sprite.size =})")

    # Assuming grayscale.
    pixels = cast(list[int], list(sprite.getdata()))

    # For each pixel, create an SVG rectangle for the corresponding area if the pixel is not blank.
    elements: list[svg.Element] = [
        svg.Rect(
            x=(i % width + 1) * PIXEL,
            y=(i // width) * PIXEL,
            width=PIXEL,
            height=PIXEL,
        )
        # Start from a 1 pixel offset for some space between the letters.
        for i, p in enumerate(pixels)
        if p > 0
    ]

    return svg.SVG(width=width * PIXEL, height=height * PIXEL, elements=elements)


def main():
    if len(sys.argv) != 2:
        print(f"usage: {__file__} <path to new font>")
        return 1

    font_path = pathlib.Path(sys.argv[1])

    # Stage 1: Create a new font using fontforge
    font = fontforge.font()

    font.familyname = FONT_FAMILY
    font.fullname = FONT_NAME
    font.fontname = FONT_NAME

    font.encoding = "UnicodeFull"
    font.em = EM

    with tempfile.TemporaryDirectory() as tempdir:
        tempdir = pathlib.Path(tempdir)

        # Stage 2: Generate an SVG for each sprite in the sheet
        with Image.open(SPRITESHEET) as sheet:
            for i, (char, sprite) in enumerate(zip(CHARS, sprites(sheet))):
                svg = sprite_to_svg(sprite)

                # Use char in hexadecimal as it may not be a valid filename.
                svg_path = tempdir / f"{i}.svg"
                with svg_path.open("w", encoding="utf-8") as f:
                    f.write(svg.as_str())

                # Stage 3: Map the SVG as a glyph into the font.
                glyph = font.createMappedChar(char)
                glyph.importOutlines(str(svg_path))
                glyph.width = EM

        # Add glyph for the spacebar.
        space = font.createMappedChar(" ")
        space.width = EM

        font.generate(str(font_path.with_suffix(".ttf")))

    return 0


if __name__ == "__main__":
    sys.exit(main())
