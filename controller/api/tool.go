package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/internal/util"
	"github.com/yeyudekuangxiang/imagedraw"
	"image/color"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Code struct {
	Canvas Canvas `json:"canvas"`
	Object []Data `json:"object"`
}
type Canvas struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Src    string `json:"src"`
}
type Data struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Src       string `json:"src"`
	Shape     string `json:"shape"`
	Content   string `json:"content"`
	FontSize  int    `json:"fontSize"`
	TextAlign string `json:"textAlign"`
	Area      struct {
		X int `json:"x"`
		Y int `json:"y"`
		W int `json:"w"`
		H int `json:"h"`
	} `json:"area"`
	Border struct {
		LT int `json:"lt"`
		RT int `json:"rt"`
		LB int `json:"lb"`
		RB int `json:"rb"`
	} `json:"border"`

	MaxLineNum     int    `json:"maxLineNum"`
	OutStr         string `json:"outStr"`
	OutStrPosition string `json:"outStrPosition"`
	Font           string `json:"font"`
	Bold           bool   `json:"bold"`
	LineHeight     int    `json:"lineHeight"`
	Color          struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
		A int `json:"a"`
	} `json:"color"`
	AutoLine   bool `json:"autoLine"`
	OverHidden bool `json:"overHidden"`
}

func GetCode(c *gin.Context) (gin.H, error) {
	var code Code
	if err := util.BindForm(c, &code); err != nil {
		return nil, err
	}
	return gin.H{
		"code": getCode(code.Canvas, code.Object),
	}, nil
}

func GetText(c *gin.Context) (gin.H, error) {
	var text Data
	if err := util.BindForm(c, &text); err != nil {
		return nil, err
	}
	var t *imagedraw.Text
	if strings.Contains(text.Content, "\n") {
		t = imagedraw.NewLineText(strings.Split(text.Content, "\n"))
	} else {
		t = imagedraw.NewText(text.Content)
	}
	t.SetArea(text.Area.X, text.Area.Y, text.Area.W, text.Area.H)
	t.SetTextAlign(text.TextAlign)
	t.SetAutoLine(text.AutoLine)
	switch text.Font {
	case "siyuanheiti":
		if text.Bold {
			t.SetFont(imagedraw.SiYuanHeiYiBold())
		} else {
			t.SetFont(imagedraw.SiYuanHeiYi())
		}
	}

	t.SetColor(color.RGBA{
		R: uint8(text.Color.R),
		G: uint8(text.Color.G),
		B: uint8(text.Color.B),
		A: uint8(text.Color.A),
	})
	t.SetOutStrPosition(text.OutStrPosition)
	t.SetMaxLineNum(text.MaxLineNum)
	t.SetOutStr(text.OutStr)
	t.SetFontSize(text.FontSize)
	t.SetLineHeight(text.LineHeight)
	t.SetOverHidden(text.OverHidden)
	result, err := t.Calc()

	if err != nil {
		return nil, err
	}

	return gin.H{
		"lineHeight": result.LineHeight,
		"height":     result.Height,
		"width":      result.Width,
		"texts":      result.SplitTextList,
	}, nil
}
func getCode(canvas Canvas, list []Data) string {
	code := ""
	for i, item := range list {
		if item.Type == "text" {
			code += getTextCode(item)
		} else {
			code += getImageCode(item)
		}
		if i != len(list)-1 {
			code += "\n\n"
		}
	}
	replacer := strings.NewReplacer(
		"{{src}}", canvas.Src,
		"{{width}}", strconv.Itoa(canvas.Width),
		"{{height}}", strconv.Itoa(canvas.Height),
		"{{imageCode}}", code,
	)
	if canvas.Src == "" {
		return replacer.Replace(`package main

import (
	"github.com/yeyudekuangxiang/imagedraw"
	"image/color"
	"log"
)
func main()  {
	img,err:=create()
	if err!=nil{
		log.Fatal(err)
	}
	log.Println(img.SaveAs("img.png"))
}
func create() (*imagedraw.Image,error) {
	base := imagedraw.NewBaseImage({{width}}, {{height}})

	{{imageCode}}
	
	return base, nil
}
`)
	}
	return replacer.Replace(`package main

import (
	"github.com/yeyudekuangxiang/imagedraw"
	"image/color"
	"log"
)
func main()  {
	img,err:=create()
	if err!=nil{
		log.Fatal(err)
	}
	log.Println(img.SaveAs("img.png"))
}
func create() (*imagedraw.Image,error) {
	base, err := imagedraw.LoadImageFromUrl("{{src}}")
	if err != nil {
		return nil, err
	}
	base = base.Resize({{width}}, {{height}})
	
	{{imageCode}}
	
	return base, nil
}
`)
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
func getImageCode(data Data) string {
	if data.Name == "" {
		data.Name = fmt.Sprintf("image%d", rand.Int())
	}

	replace := strings.NewReplacer(
		"{{image}}", data.Name,
		"{{src}}", data.Src,
		"{{width}}", strconv.Itoa(data.Area.W),
		"{{height}}", strconv.Itoa(data.Area.H),
		"{{startX}}", strconv.Itoa(data.Area.X),
		"{{startY}}", strconv.Itoa(data.Area.Y),
		"{{borderRadius}}", fmt.Sprintf("%d,%d,%d,%d", data.Border.LT, data.Border.RT, data.Border.RB, data.Border.LB),
	)

	code := `{{image}}, err := imagedraw.LoadImageFromUrl("{{src}}")
			if err != nil {
				return nil, err
			}
			{{image}} = {{image}}.Resize({{width}}, {{height}})
`
	if data.Shape == "cycle" {
		code += `{{image}} = {{image}}.Ellipse({{image}}.Width()/2, {{image}}.Height()/2, {{image}}.Width()/2, {{image}}.Height()/2)
`
	}

	if (data.Border.LT != 0 || data.Border.RT != 0 || data.Border.RB != 0 || data.Border.LB != 0) && data.Shape != "cycle" {

		code += `{{image}} = {{image}}.BorderRadius({{borderRadius}})
`
	}

	code += `{{image}}.SetArea({{startX}}, {{startY}}, {{width}}, {{height}})
			base.Fill({{image}})
`
	return replace.Replace(code)
}
func getTextCode(data Data) string {
	if data.Name == "" {
		data.Name = fmt.Sprintf("text%d", rand.Int())
	}

	code := ""
	if strings.Contains(data.Content, "\n") {
		code += `{{variable}} = imagedraw.NewLineText([]string{"` + strings.ReplaceAll(data.Content, "\n", "\",\"") + `"}, "\n"))`
	} else {
		code += `{{variable}} = imagedraw.NewText("{{text}}")`
	}

	font := "imagedraw.SiYuanHeiYi()"
	switch data.Font {
	case "siyuanheiti":
		if data.Bold {
			font = "imagedraw.SiYuanHeiYiBold()"
		} else {
			font = "imagedraw.SiYuanHeiYi()"
		}
	}
	replace := strings.NewReplacer(
		"{{variable}}", data.Name,
		"{{text}}", data.Content,
		"{{startX}}", strconv.Itoa(data.Area.X),
		"{{startY}}", strconv.Itoa(data.Area.Y),
		"{{width}}", strconv.Itoa(data.Area.W),
		"{{height}}", strconv.Itoa(data.Area.H),
		"{{textAlign}}", data.TextAlign,
		"{{autoLine}}", fmt.Sprintf("%v", data.AutoLine),
		"{{R}}", strconv.Itoa(data.Color.R),
		"{{G}}", strconv.Itoa(data.Color.G),
		"{{B}}", strconv.Itoa(data.Color.B),
		"{{A}}", strconv.Itoa(data.Color.A),
		"{{outStrPosition}}", data.OutStrPosition,
		"{{maxLineNum}}", strconv.Itoa(data.MaxLineNum),
		"{{outStr}}", data.OutStr,
		"{{fontSize}}", strconv.Itoa(data.FontSize),
		"{{lineHeight}}", strconv.Itoa(data.LineHeight),
		"{{overHidden}}", fmt.Sprintf("%v", data.OverHidden),
		"{{font}}", font,
	)
	return replace.Replace(`var {{variable}} *imagedraw.Text
	` + code + `
	{{variable}}.SetArea({{startX}}, {{startY}}, {{width}}, {{height}})
	{{variable}}.SetTextAlign("{{textAlign}}")
	{{variable}}.SetAutoLine({{autoLine}})
	{{variable}}.SetColor(color.RGBA{R: {{R}},G: {{G}},B: {{B}},A: {{A}}})
	{{variable}}.SetOutStrPosition("{{outStrPosition}}")
	{{variable}}.SetMaxLineNum({{maxLineNum}})
	{{variable}}.SetOutStr("{{outStr}}")
	{{variable}}.SetFontSize({{fontSize}})
	{{variable}}.SetLineHeight({{lineHeight}})
	{{variable}}.SetOverHidden({{overHidden}})
	{{variable}}.SetFont({{font}})
	base.Fill({{variable}})`)
}
