package model

type ExifItem struct {
	Key   string            `json:"key"`
	Label string            `json:"label"`
	Value string            `json:"value"`
	Enums map[string]string `json:"enums"`
}

var SupportedExifFormats = map[string]ExifItem{
	"ApertureValue":   {Key: "ApertureValue", Label: "光圈值"},
	"BitsPerSample":   {Key: "BitsPerSample", Label: "采样输出位数"},
	"BrightnessValue": {Key: "BrightnessValue", Label: "亮度值"},
	"ColorSpace":      {Key: "ColorSpace", Label: "色彩空间"},
	"PhotometricInterpretation": {
		Key:   "PhotometricInterpretation",
		Label: "光度",
		Enums: map[string]string{
			"1": "单色",
			"2": "RGB",
			"6": "YCbCr",
		},
	},
	"ComponentsConfiguration": {Key: "ComponentsConfiguration", Label: "组件配置"},
	"Compression": {
		Key:   "Compression",
		Label: "压缩",
		Enums: map[string]string{
			"1": "未压缩",
			"6": "JPEG",
		},
	},
	"DateTime":          {Key: "DateTime", Label: "日期时间"},
	"DateTimeDigitized": {Key: "DateTimeDigitized", Label: "日期时间数字化"},
	"DateTimeOriginal":  {Key: "DateTimeOriginal", Label: "日期时间原始"},
	"ExifTag":           {Key: "ExifTag", Label: "EXIF标签"},
	"ExifVersion":       {Key: "ExifVersion", Label: "Exif版本"},
	"ExposureBiasValue": {Key: "ExposureBiasValue", Label: "曝光补偿"},
	"ExposureMode":      {Key: "ExposureMode", Label: "曝光模式"},
	"ExposureProgram": {
		Key:   "ExposureProgram",
		Label: "曝光模式",
		Enums: map[string]string{
			"1": "手动曝光",
			"2": "程序曝光",
			"3": "光圈优先",
			"4": "快门优先",
			"5": "创意程序(慢速程序)",
			"6": "动作程序(高速程序)",
			"7": "肖像模式",
			"8": "风景模式",
		}},
	"ExposureTime": {Key: "ExposureTime", Label: "曝光时间"},
	"FNumber":      {Key: "FNumber", Label: "光圈"},
	"Flash": {Key: "Flash", Label: "闪光灯", Enums: map[string]string{
		"0": "无",
		"1": "闪光",
		"5": "闪光但未检测反射光",
		"7": "闪光且检测反射光",
	}},
	"FlashpixVersion":             {Key: "FlashpixVersion", Label: "Flashpix版本"},
	"FocalLength":                 {Key: "FocalLength", Label: "焦距"},
	"FocalLengthIn35mmFilm":       {Key: "FocalLengthIn35mmFilm", Label: "焦距35mm胶片"},
	"GPSAltitude":                 {Key: "GPSAltitude", Label: "GPS海拔"},
	"GPSAltitudeRef":              {Key: "GPSAltitudeRef", Label: "GPS海拔参考"},
	"GPSDateStamp":                {Key: "GPSDateStamp", Label: "GPS日期戳"},
	"GPSLatitude":                 {Key: "GPSLatitude", Label: "纬度"},
	"GPSLatitudeRef":              {Key: "GPSLatitudeRef", Label: "纬度参考"},
	"GPSLongitude":                {Key: "GPSLongitude", Label: "经度"},
	"GPSLongitudeRef":             {Key: "GPSLongitudeRef", Label: "经度参考"},
	"GPSProcessingMethod":         {Key: "GPSProcessingMethod", Label: "GPSProcessingMethod"},
	"GPSTag":                      {Key: "GPSTag", Label: "GPS标签"},
	"GPSTimeStamp":                {Key: "GPSTimeStamp", Label: "GPS时间戳"},
	"ISOSpeedRatings":             {Key: "ISOSpeedRatings", Label: "感光度"},
	"ImageLength":                 {Key: "ImageLength", Label: "图片高度"},
	"ImageWidth":                  {Key: "ImageWidth", Label: "图片宽度"},
	"InteroperabilityIndex":       {Key: "InteroperabilityIndex", Label: "互操作性指数"},
	"InteroperabilityTag":         {Key: "InteroperabilityTag", Label: "互操作性标签"},
	"InteroperabilityVersion":     {Key: "InteroperabilityVersion", Label: "互操作性版本"},
	"JPEGInterchangeFormat":       {Key: "JPEGInterchangeFormat", Label: "JPEG交换格式"},
	"JPEGInterchangeFormatLength": {Key: "JPEGInterchangeFormatLength", Label: "JPEGInterchangeFormatLength"},
	"LightSource": {
		Key:   "LightSource",
		Label: "光源",
		Enums: map[string]string{
			"0": "未知",
			"1": "日光", "2": "荧光灯",
			"3":   "白炽灯(钨丝)",
			"10":  "闪光灯",
			"17":  "标准光A",
			"18":  "标准光B",
			"19":  "标准光C",
			"20":  "D55",
			"21":  "D65",
			"22":  "D75",
			"255": "其他",
		},
	},
	"Make":             {Key: "Make", Label: "品牌"},
	"Model":            {Key: "Model", Label: "设备型号"},
	"MaxApertureValue": {Key: "MaxApertureValue", Label: "最大光圈值"},
	"MeteringMode": {
		Key:   "MeteringMode",
		Label: "测光模式",
		Enums: map[string]string{
			"0":   "未知",
			"1":   "平均测光",
			"2":   "中央重点测光",
			"3":   "点测光",
			"4":   "多点测光",
			"5":   "多区域测光",
			"6":   "部分测光",
			"255": "其他",
		},
	},
	"Orientation":         {Key: "Orientation", Label: "方向"},
	"PixelXDimension":     {Key: "PixelXDimension", Label: "像素X尺寸"},
	"PixelYDimension":     {Key: "PixelYDimension", Label: "像素Y尺寸"},
	"ResolutionUnit":      {Key: "ResolutionUnit", Label: "分辨率单位"},
	"SceneCaptureType":    {Key: "SceneCaptureType", Label: "场景捕获类型"},
	"SensingMethod":       {Key: "SensingMethod", Label: "传感方法"},
	"ShutterSpeedValue":   {Key: "ShutterSpeedValue", Label: "快门速度"},
	"SubSecTime":          {Key: "SubSecTime", Label: "亚秒时间"},
	"SubSecTimeDigitized": {Key: "SubSecTimeDigitized", Label: "亚秒时间数字化"},
	"SubSecTimeOriginal":  {Key: "SubSecTimeOriginal", Label: "亚秒时间原始"},
	"WhiteBalance":        {Key: "WhiteBalance", Label: "白平衡"},
	"XResolution":         {Key: "XResolution", Label: "X分辨率"},
	"YCbCrPositioning":    {Key: "YCbCrPositioning", Label: "YCBCR定位"},
	"YResolution":         {Key: "YResolution", Label: "Y分辨率"},
	"Software":            {Key: "Software", Label: "Software"},
	"Longitude":           {Key: "Longitude", Label: "经度"},
	"Latitude":            {Key: "Latitude", Label: "纬度"},
}