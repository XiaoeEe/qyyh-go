package robotService

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/utils"
	"strings"
	"time"
)

func GetGGData(name string) (gg model.GG, err error) {
	err = utils.Get("https://api.maplestory.gg/v2/public/character/gms/"+url.QueryEscape(name), nil, &gg)
	if gg.CharacterData.Name == "" {
		err = errors.New("gg查询错误")
	}
	return
}

func myGG(message model.CQMessage) {
	my := table.GetMainPlayerByQQ(message.UserId)
	if my.Name != "" {
		GGForName(message, my.Name)
	} else {
		SendGroupMsg(message, "未查询到你的主号")
	}
}

func GGForName(message model.CQMessage, name string) {
	name = strings.ToLower(name)

	if !utils.FileIsExisted(fmt.Sprintf("./file/gg/%s.png", name)) {
		gg, err := GetGGData(name)
		if err != nil {
			SendGroupMsg(message, err.Error())
			return
		}
		makeGGImg(gg)
	}
	file, err := os.ReadFile(fmt.Sprintf("./file/gg/%s.png", name))
	if err != nil {
		return
	}
	baseCode := base64.StdEncoding.EncodeToString(file)
	SendGroupMsg(message, fmt.Sprintf("[CQ:image,file=base64://%s]", baseCode))
}

func makeGGImg(gg model.GG) {
	name := strings.ToLower(gg.CharacterData.Name)
	img := image.NewRGBA(image.Rect(0, 0, 1553, 474))

	avatarF, _ := http.Get(gg.CharacterData.CharacterImageURL)
	backgroundF, _ := os.Open("./assets/img/background.png")

	backgroundImg, _, _ := image.Decode(backgroundF)

	avatar, _, _ := image.Decode(avatarF.Body)
	draw.Draw(img, img.Bounds(), backgroundImg, image.Pt(0, 0), draw.Over)
	if avatar != nil {
		avatar = resize.Resize(144, 144, avatar, resize.Bilinear)
		draw.Draw(img, img.Bounds(), avatar, image.Pt(-40, -40), draw.Over)
	}

	content := freetype.NewContext()
	content.SetClip(img.Bounds())
	content.SetDst(img)
	content.SetDPI(72)
	content.SetFont(getGGFont())
	content.SetSrc(image.White)

	content.SetFontSize(70)
	s := gg.CharacterData.Name
	_, _ = content.DrawString(s, freetype.Pt(220, 100))

	content.SetFontSize(21)
	s = fmt.Sprintf("Level %d (%.2f%%) | %s | %s | MapleStory Global", gg.CharacterData.Level, gg.CharacterData.EXPPercent, gg.CharacterData.Class, gg.CharacterData.Server)
	_, _ = content.DrawString(s, freetype.Pt(220, 140))

	content.SetFontSize(14)
	s = fmt.Sprintf("全服排名 %d | %s 排名 %d | %s 排名 %d | %s %s 排名 %d", gg.CharacterData.GlobalRanking, gg.CharacterData.Class, gg.CharacterData.ClassRank, gg.CharacterData.Server, gg.CharacterData.ServerRank, gg.CharacterData.Server, gg.CharacterData.Class, gg.CharacterData.ServerClassRanking)
	_, _ = content.DrawString(s, freetype.Pt(220, 170))

	if gg.CharacterData.LegionLevel != 0 {
		s = fmt.Sprintf("成就点数 %d(排名: %d)", gg.CharacterData.AchievementPoints, gg.CharacterData.AchievementRank)
		_, _ = content.DrawString(s, freetype.Pt(1000, 100))
		s = fmt.Sprintf("联盟等级 %d(排名: %d)", gg.CharacterData.LegionLevel, gg.CharacterData.LegionRank)
		_, _ = content.DrawString(s, freetype.Pt(1000, 130))
		s = fmt.Sprintf("联盟战力 %d(每日联盟币: %d)", gg.CharacterData.LegionPower, gg.CharacterData.LegionCoinsPerDay)
		_, _ = content.DrawString(s, freetype.Pt(1000, 160))
	}

	if len(gg.CharacterData.GraphData) > 2 {
		var (
			c1   []chart.Value
			c2   []chart.Series
			c3   []chart.Series
			c23x []time.Time
			c2y  []float64
			c3y  []float64
		)
		for _, item := range gg.CharacterData.GraphData {
			date, _ := time.Parse("2006-01-02", item.DateLabel)
			c1 = append(c1, chart.Value{
				Label: date.Add(time.Hour * 24).Format("02"),
				Value: float64(item.EXPDifference) / 1000000000,
			})
			c23x = append(c23x, date)
			c2y = append(c2y, float64(item.TotalOverallEXP)/1000000000)
			c3y = append(c3y, float64(item.Level))
			c2 = []chart.Series{
				chart.TimeSeries{
					Style: chart.Style{
						StrokeColor: drawing.Color{R: 255, G: 255, B: 0, A: 255},
					},
					XValues: c23x,
					YValues: c2y,
				},
			}
			c3 = []chart.Series{
				chart.TimeSeries{
					Style: chart.Style{
						StrokeColor: drawing.Color{R: 255, G: 255, B: 0, A: 255},
					},
					XValues: c23x,
					YValues: c3y,
				},
			}
		}
		c1 = c1[:len(c1)-1]
		drawGGBarChart(c1, name)
		drawGGLineChart(c2, name, "2")
		drawGGLineChart(c3, name, "3")

		content.SetFontSize(21)

		_, _ = content.DrawString("每日经验获取", freetype.Pt(40, 234))
		tmp := readImg(fmt.Sprintf("./file/gg/tmp/%s/chart1.png", name))
		if tmp != nil {
			draw.Draw(img, img.Bounds(), tmp, image.Pt(0, -234), draw.Over)
		}

		_, _ = content.DrawString("总经验曲线", freetype.Pt(557, 234))
		tmp = readImg(fmt.Sprintf("./file/gg/tmp/%s/chart2.png", name))
		if tmp != nil {
			draw.Draw(img, img.Bounds(), tmp, image.Pt(-517, -234), draw.Over)
		}

		_, _ = content.DrawString("等级曲线", freetype.Pt(1074, 234))
		tmp = readImg(fmt.Sprintf("./file/gg/tmp/%s/chart3.png", name))
		if tmp != nil {
			draw.Draw(img, img.Bounds(), tmp, image.Pt(-1034, -234), draw.Over)
		}
	}
	file, _ := os.Create(fmt.Sprintf("./file/gg/%s.png", name))
	defer file.Close()
	_ = png.Encode(file, img)
	_ = os.RemoveAll("./file/gg/tmp/" + name)
}

func drawGGBarChart(value []chart.Value, user string) {
	font := getGGFont()
	c := drawing.Color{R: 255, G: 255, B: 0, A: 255}
	for i, _ := range value {
		value[i].Style = chart.Style{
			StrokeColor: c,
			FillColor:   c,
		}
	}
	graph := chart.BarChart{
		Width:  516,
		Height: 280,
		XAxis: chart.Style{
			FontColor: c,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				FontColor: c,
			},
			ValueFormatter: func(v interface{}) string {
				return fmt.Sprintf("%.2fB", v.(float64))
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:   40,
				Left:  20,
				Right: 20,
			},
		},
		BarWidth: 20,
		Font:     font,
		Bars:     value,
	}
	tmpPath := "./file/gg/tmp/" + user
	_ = os.MkdirAll(tmpPath, 0700)
	f, _ := os.Create(filepath.Join(tmpPath, "chart1.png"))
	defer f.Close()
	_ = graph.Render(chart.PNG, f)
}

func drawGGLineChart(value []chart.Series, user string, flag string) {
	font := getGGFont()
	c := drawing.Color{R: 255, G: 255, B: 0, A: 255}
	graph := chart.Chart{
		Width:  516,
		Height: 210,
		Series: value,
		XAxis: chart.XAxis{
			Style: chart.Style{
				FontColor:   c,
				FillColor:   c,
				DotColor:    c,
				StrokeColor: c,
			},
			GridMajorStyle: chart.Style{
				FontColor:   c,
				FillColor:   c,
				DotColor:    c,
				StrokeColor: c,
			},
			GridMinorStyle: chart.Style{
				FontColor:   c,
				FillColor:   c,
				DotColor:    c,
				StrokeColor: c,
			},
			ValueFormatter: func(v interface{}) string {
				return time.Unix(int64(v.(float64)/1000000000), 0).Format("02")
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				FontColor:   c,
				FillColor:   c,
				DotColor:    c,
				StrokeColor: c,
			},
			ValueFormatter: func(v interface{}) string {
				if flag == "2" {
					return fmt.Sprintf("%.2fB", v.(float64))
				}
				return fmt.Sprintf("%.0f级", v.(float64))
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:   40,
				Left:  20,
				Right: 20,
			},
		},
		Font: font,
	}
	tmpPath := "./file/gg/tmp/" + user
	_ = os.MkdirAll(tmpPath, 0700)
	f, _ := os.Create(filepath.Join(tmpPath, fmt.Sprintf("chart%s.png", flag)))
	defer f.Close()
	_ = graph.Render(chart.PNG, f)
}

func readImg(src string) *image.RGBA {
	f, _ := os.Open(src)
	old, _, err := image.Decode(f)
	if err != nil {
		return nil
	}
	newImg := image.NewRGBA(old.Bounds())
	draw.Draw(newImg, newImg.Bounds(), old, image.Pt(0, 0), draw.Over)
	for x := 0; x < newImg.Bounds().Dx(); x++ {
		for y := 0; y < newImg.Bounds().Dy(); y++ {
			if likeWhite(newImg.At(x, y)) {
				newImg.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 0})
			}
		}
	}
	return newImg
}

func likeWhite(point color.Color) bool {
	s := fmt.Sprintf("%+v", point)

	r, g, b, _ := point.RGBA()
	r = r >> 8
	g = g >> 8
	b = b >> 8

	if r < 255 && g < 255 && b >= 1 || s == "{R:255 G:255 B:255 A:255}" {
		return true
	}
	return false
}

func getGGFont() (font *truetype.Font) {
	bf, err := os.ReadFile("./assets/font/msyh.ttc")
	if err != nil {
		fmt.Println("获取字体失败")
	}
	font, _ = freetype.ParseFont(bf)
	return
}
