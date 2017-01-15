# jpg2go

A very easy tool for resizing jpg files. 

# Program logic

Target of this small project should be to practice the `decoupling` William Kennedy is talking about in [his great talk](https://www.goinggo.net/2016/12/developing-a-design-philosophy-in-go.html).

Therefore all the logic for resizing a list of pictures is moved into several interfaces. 

* FilePahtsGetter: Gets the filepaths, when a root is given
* Imager: Returns an image.Image, when a filepath is given
* Resizer: Resizes the image
* Writer: Writes the image to a destination

All interfaces are bundled into the type `FilesComponents`. 

The resize package just uses that interface to resize all the given files.

The file cmd/cli/main.go brings all the impementations and the resize package together. 

* osfiles - FilePathsGetter
* jpg - Imager and Writer
* nfntresizer - Resizer

