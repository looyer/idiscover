package conf

import "flag"

var (
	UniCloudCode string = "" //https://cloud.unidoc.io 上申请的app-secret
	InputPDF     string = "C:/Users/looye/Desktop/mathematics/test.pdf"
	OuterPDF     string = "华章数学译丛55-代数 原书第2版 反色.pdf"
	LogDir       string = "./logs/" //日志文件所在
)

func CommandLine() {
	flag.StringVar(&UniCloudCode, "unicode", UniCloudCode, "cloud.unidoc.io上的App-Secret")
	flag.StringVar(&InputPDF, "input", InputPDF, "pdf-文件路径")
	flag.StringVar(&OuterPDF, "outer", OuterPDF, "输出文件路径")
	flag.StringVar(&LogDir, "logdir", LogDir, "日志文件夹路径")

	flag.Parse()
}
