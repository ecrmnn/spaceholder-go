# spaceholder-go
> Go implementation of [spaceholder](https://github.com/ecrmnn/spaceholder)

## Work in progress
This project is experimental and under development. If you want to use it's functionality I recommend you install the [original project](https://github.com/ecrmnn/spaceholder)

### What is it?
Spaceholder makes it easy to download placeholder images when you need them.
Images are downloaded from LoremPixel.com, Placehold.it, PlaceIMG.com, Dummyimage.com, Unsplash.it.

### Installation
Clone, build ``go build spaceholder-go`` and move to ``/usr/local/bin``

### Usage
```bash
spaceholder
# Downloads 1 image (1024x768px) into current directory

spaceholder -number 100
# Downloads 100 images into current directory

spaceholder -size 800x600
# Downloads 1 image (800x600px) into current directory

spaceholder -number 50 -size 800x600 -provider loremPixel
# Downloads 50 images (800x600px) from LoremPixel into current directory
# If no --provider is specified, each image is downloaded from a random provider
```

### Providers / Sources
- [Dummy Image](http://dummyimage.com)
- [FakeImg](fakeimg.pl)
- [Lorem Pixel](http://lorempixel.com)
- [PlaceholdIt](https://placehold.it)
- [PlaceImg](http://placeimg.com)
- [UnsplashIt](https://unsplash.it)

### License
MIT © [Daniel Eckermann](http://danieleckermann.com)