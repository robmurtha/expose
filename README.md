# expose
Expose provides a Go package and command tool for extracting and parsing [XMP](http://www.adobe.com/products/xmp.html), [EXIF](http://www.exif.org/) and [IPTC](https://iptc.org/) metadata from image and video files using the exiv2 binary tool.
The output data is grouped by type in JSON format.

This example was created for the Feb 2017 [GoLangPhilly](https://www.meetup.com/GoLangPhilly/) meetup presentation on Processing Image and Video Metadata by Rob Murtha.

## Exiv2 Dependency
This package requires the exiv2 binary (see instructions below).
[exiv2 website](http://www.exiv2.org/index.html)
[exiv2 download / source](http://www.exiv2.org/download.html)


## Install go package and tools
```
go get github.com/robmurtha/expose
```

## Run
```
Usage: expose [-flags] filename
flags:
  -p    pretty JSON
  -v    verbose logging
```

## Organization
This project includes a main.go for running from the command line and extracting all possible information.
 An exiv2 parsing package is included and can be used independently. 

```
    // parse XMP tags
    // with output from exiv2 -PXgnycv filename (see main)
	expose := exiv2.New(output)
	fields, _ := expose.Fields()
	jsonBytes, _ := json.Marshal(fields)
	
```

## Installing
### mac-os (brew)
```
brew install exiv2
```

## Compiling on mac with brew
The default brew version has video support disabled, if you need video
support follow these instructions.

```
wget http://exiv2.org/exiv2-0.25.tar.gz
tar -xzf exiv2-0.25.tar.gz
cd exiv2-0.25
brew install exiv2 --only-dependencies
./configure --enable-video
make install
```

### apt-get
```
apt-get install exiv2
```

## Sample JSON
The JSON records are organized by metadata type as arrays of Field. Organization by group, type and additional parsing of values is up to the user.
```
{
 "expose": {
  "xmp": [
   {
    "group": "video",
    "name": "FileSize",
    "type": "XmpText",
    "count": 9,
    "value": "0.0515432"
   },
   {
    "group": "video",
    "name": "FileName",
    "type": "XmpText",
    "count": 27,
    "value": "testdata/samsung_galaxy.mp4"
   },
   {
    "group": "video",
    "name": "MimeType",
    "type": "XmpText",
    "count": 15,
    "value": "video/quicktime"
...
```
