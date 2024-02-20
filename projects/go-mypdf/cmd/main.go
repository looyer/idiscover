package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"

	"go-mypdf/src/conf"
)

//./cmd.exe -unicode b53da9190c3841a24842dab33af2c30579e2b84a20293efcb093b640a13d5ff1 -input ./Algebra_Artin_clear8.pdf -outer ./Algebra-Artin-InvertColor.pdf

func main() {
	fmt.Println("~~~~~~~~~  go-pdf  ~~~~~~~~~")

	conf.CommandLine()

	err := license.SetMeteredKey(conf.UniCloudCode)
	if err != nil {
		panic(err)
	}

	//适合小型PDF，一次性转换，比如100page以内
	// smallPDF()

	//适合大型型PDF，比如100page以上
	biggerPDF()
}

func smallPDF() {

	startTime := time.Now()
	myreader, f, err := model.NewPdfReaderFromFile(conf.InputPDF, nil)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numPages, err := myreader.GetNumPages()
	if err != nil {
		panic(err)
	}
	if numPages <= 0 {
		fmt.Println("page-num is zero!")
		return
	}

	listimages := make([]*image.RGBA, numPages)
	wg := sync.WaitGroup{}

	fmt.Printf("total-page:%v\n", numPages)
	for i := 0; i < numPages; i++ {
		page, err := myreader.GetPage(i + 1)
		if err != nil {
			panic(err)
		}

		expor, err := extractor.New(page)
		if err != nil {
			panic(err)
		}
		images, err := expor.ExtractPageImages(nil)
		if err != nil {
			panic(err)
		}
		num := len(images.Images)
		if num == 0 {
			fmt.Printf("page:%v no image! skip!\n", i+1)
			continue
		}
		if num > 1 {
			fmt.Printf("page:%v image num>1! only use 1-image!\n", i+1)
		}
		gimg, _ := images.Images[0].Image.ToGoImage()

		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			listimages[k] = inverseColor(k, gimg)
		}(i)

		if (i > 0) && (i%20 == 0) {
			fmt.Printf("...page:%v...", i)
		}
	}
	wg.Wait()
	fmt.Printf("page change done.\n")

	fmt.Printf("start write new pdf...\n")
	mywriter := creator.New()
	for i := 0; i < numPages; i++ {
		if listimages[i] == nil {
			continue
		}
		cpimage, _ := mywriter.NewImageFromGoImage(listimages[i])
		cpimage.ScaleToWidth(612.0)
		height := 612.0 * cpimage.Height() / cpimage.Width()

		mywriter.SetPageSize(creator.PageSize{612.0, height})
		mywriter.NewPage()

		cpimage.SetPos(0, 0)
		mywriter.Draw(cpimage)

		if (i > 0) && (i%50) == 0 {
			fmt.Printf("...writerpage:%v...", i)
		}
	}
	err = mywriter.WriteToFile(conf.OuterPDF)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wirter done.\n\n")

	fmt.Printf("use time:%.2f(s)\n", time.Since(startTime).Seconds())
	fmt.Printf("change success!\n")
}

func biggerPDF() {
	startTime := time.Now()

	year, month, day := time.Now().Date()
	tempdirname := fmt.Sprintf("%04d-%02d-%02d-%v", year, int(month), day, time.Now().Unix())
	//创建临时文件夹，用于存放中间生成的图片
	os.MkdirAll(tempdirname, os.ModePerm)

	myreader, f, err := model.NewPdfReaderFromFile(conf.InputPDF, nil)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numPages, err := myreader.GetNumPages()
	if err != nil {
		panic(err)
	}
	if numPages <= 0 {
		fmt.Println("page-num is zero!")
		return
	}

	listnames := make([]string, numPages)
	wg := sync.WaitGroup{}

	fmt.Printf("total-pages:%v\n", numPages)
	for i := 0; i < numPages; i++ {
		page, err := myreader.GetPage(i + 1)
		if err != nil {
			panic(err)
		}

		expor, err := extractor.New(page)
		if err != nil {
			panic(err)
		}
		images, err := expor.ExtractPageImages(nil)
		if err != nil {
			panic(err)
		}
		num := len(images.Images)
		if num == 0 {
			fmt.Printf("page:%v no image! skip!\n", i+1)
			continue
		}
		if num > 1 {
			fmt.Printf("page:%v image num>1! only use 1-image!\n", i+1)
		}
		gimg, _ := images.Images[0].Image.ToGoImage()

		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			imgout := inverseColor(k, gimg)

			fname := fmt.Sprintf("./%v/image-%04d.png", tempdirname, k)
			fout, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				return
			}
			defer fout.Close()

			if err := png.Encode(fout, imgout); err != nil {
				fmt.Printf("png-%04d is export err:%v\n", k, err)
			}
			imgout = nil
			listnames[k] = fname
		}(i)

		if (i > 0) && (i%50 == 0) {
			fmt.Printf("...page:%v...", i)
			runtime.GC()
		}
	}
	wg.Wait()
	fmt.Printf("page change done.\n")

	runtime.GC()

	fmt.Printf("start write new pdf...\n")
	mywriter := creator.New()
	for i := 0; i < numPages; i++ {
		if listnames[i] == "" {
			continue
		}
		fpng, err := os.Open(listnames[i])
		if err != nil {
			fmt.Printf("read png:%v err:%v\n", listnames[i], err)
			continue
		}
		pngimg, err := png.Decode(fpng)
		if err != nil {
			fmt.Printf("read png:%v err:%v\n", listnames[i], err)
			continue
		}
		fpng.Close()

		cpimage, _ := mywriter.NewImageFromGoImage(pngimg)
		cpimage.ScaleToWidth(612.0)
		height := 612.0 * cpimage.Height() / cpimage.Width()

		mywriter.SetPageSize(creator.PageSize{612.0, height})
		mywriter.NewPage()

		cpimage.SetPos(0, 0)
		mywriter.Draw(cpimage)

		if (i > 0) && (i%50) == 0 {
			fmt.Printf("...writerpage:%v...", i)
		}
	}
	err = mywriter.WriteToFile(conf.OuterPDF)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wirter done.\n\n")

	fmt.Printf("use time:%.2f(s)\n", time.Since(startTime).Seconds())
	fmt.Printf("change success!\n")
}

func inverseColor(k int, gimg image.Image) *image.RGBA {
	rect := gimg.Bounds()
	newimg := image.NewRGBA(rect)
	for x := 0; x < rect.Max.X; x++ {
		for y := 0; y < rect.Max.Y; y++ {
			c := gimg.At(x, y)
			r, g, b, a := c.RGBA()
			r0, g0, b0, a0 := uint8(255-r), uint8(255-g), uint8(255-b), uint8(a) //黑白反转
			if r0 < 31 && g0 < 31 && b0 < 31 {                                   //纯黑调色
				r0, g0, b0 = 31, 31, 31
			}
			newimg.SetRGBA(x, y, color.RGBA{R: r0, G: g0, B: b0, A: a0})
		}
	}
	return newimg
}
