package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type TileOffset_ struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

type Terrain struct {
	Name       string     `xml:"name,attr"`
	Tile       int        `xml:"tile,attr"`
	Properties []Property `xml:"properties>property"`
}

type Image_ struct {
	//Format TODO
	Source string `xml:"source,attr"`
	Trans  string `xml:"trans,attr"` //To int?
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type Tile struct {
	Id      int    `xml:"id,attr"`
	Terrain string `xml:"terrain,attr"`
	//Probability int    `xml:"probability,attr"`
	Image      Image_     `xml:"image"`
	Properties []Property `xml:"properties>property"`
}

type TileSet struct {
	FirstGid   int         `xml:"firstgid,attr"`
	Source     string      `xml:"source,attr"`
	Name       string      `xml:"name,attr"`
	TileWidth  int         `xml:"tilewidth,attr"`
	TileHeight int         `xml:"tileheight,attr"`
	Spacing    int         `xml:"spacing,attr"`
	Margin     int         `xml:"margin,attr"`
	TileOffset TileOffset_ `xml:"tileoffset"`
	Image      Image_      `xml:"image"`
	Terrains   []Terrain   `xml:"terraintypes>terrain"`
	Tiles      []Tile      `xml:"tile"`
	Properties []Property  `xml:"properties>property"`
}

type Data_ struct {
	Encoding    string `xml:"encoding,attr"`
	Compression string `xml:"compression,attr"`
	Data        []byte `xml:",innerxml"`
	//Tile TODO
}

type Layer struct {
	Name       string     `xml:"name,attr"`
	Opacity    float32    `xml:"opacity,attr"`
	Visible    int        `xml:"visible,attr"`
	Data       Data_      `xml:"data"`
	Properties []Property `xml:"properties>property"`
}

type ImageLayer struct {
	Name       string     `xml:"name,attr"`
	Opacity    float32    `xml:"opacity,attr"`
	Visible    int        `xml:"visible,attr"`
	Image      Image_     `xml:"image"`
	Properties []Property `xml:"properties>property"`
}

type ObjectGroup_ struct {
	Name       string     `xml:"name,attr"`
	Color      string     `xml:"color,attr"`
	Opacity    float32    `xml:"opacity,attr"`
	Visible    int        `xml:"visible,attr"`
	Properties []Property `xml:"properties>property"`
	Objects    []Object   `xml:"object"`
}

type Ellipse_ struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type Polygon_ struct {
	Points string `xml:"points,attr"`
}

type Polyline_ struct {
	Points string `xml:"points,attr"`
}

type Object struct {
	Name       string     `xml:"name,attr"`
	Type       string     `xml:"type,attr"`
	X          int        `xml:"x,attr"`
	Y          int        `xml:"y,attr"`
	Width      int        `xml:"width,attr"`
	Height     int        `xml:"height,attr"`
	Rotation   int        `xml:"rotation,attr"`
	Gid        int        `xml:"gid,attr"`
	Visible    int        `xml:"visible,attr"`
	Ellipse    Ellipse_   `xml:"ellipse"`
	Polygon    Polygon_   `xml:"polygon"`
	Polyline   Polyline_  `xml:"polyline"`
	Properties []Property `xml:"properties>property"`
}

type TmxMap struct {
	Version         string       `xml:"version,attr"`
	Orientation     string       `xml:"orientation,attr"`
	Width           int          `xml:"width,attr"`
	Height          int          `xml:"height,attr"`
	TileWidth       int          `xml:"tilewidth,attr"`
	TileHeight      int          `xml:"tileheight,attr"`
	BackgroundColor string       `xml:"backgroundcolor,attr"` //To int?
	TileSets        []TileSet    `xml:"tileset"`
	Layers          []Layer      `xml:"layer"`
	ObjectGroup     ObjectGroup_ `xml:"objectgroup"`
	ImageLayers     []ImageLayer `xml:"imagelayer"`
	Properties      []Property   `xml:"properties>property"`
}

func main() {
	file, _ := ioutil.ReadFile("examples/isometric_grass_and_water.tmx")
	var t TmxMap
	xml.Unmarshal(file, &t)
	fmt.Println(t)
}
