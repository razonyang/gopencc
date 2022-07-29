# Go OpenCC

Go 语言编写的简繁体转换命令行工具，支持词汇级别的转换、异体字转换和地区习惯用词转换（中国大陆、臺湾、香港、日本新字体）。

## Installation

### Download

Download from [Releases](https://github.com/razonyang/gopencc/releases)

### Build

```go
$ go install github.com/razonyang/gopencc
```

## Usage

**翻译文本**

```bash
$ gopencc -c s2t -i "自然语言处理是人工智能领域中的一个重要方向。"
自然語言處理是人工智能領域中的一個重要方向。
```

**翻译文件**

```bash
$ echo "自然語言處理是人工智能領域中的一個重要方向。" > /tmp/gopencc.txt
$ gopencc -c t2s -i /tmp/gopencc.txt
自然语言处理是人工智能领域中的一个重要方向。
```

**写入文件**

```bash
$ gopencc -c s2t -i "自然语言处理是人工智能领域中的一个重要方向。" -o /tmp/gopencc.out
$ cat /tmp/gopencc.out 
自然語言處理是人工智能領域中的一個重要方向。
```

**Help**

```bash
$ gopencc -h
NAME:
   gopencc - Open Chinese Convert Conversions between Traditional Chinese, Simplified Chinese

USAGE:
   gopencc [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --conversion value, -c value  s2t Simplified Chinese to Traditional Chinese 簡體到繁體
                                 t2s Traditional Chinese to Simplified Chinese 繁體到簡體
                                 s2tw Simplified Chinese to Traditional Chinese (Taiwan Standard) 簡體到臺灣正體
                                 tw2s Traditional Chinese (Taiwan Standard) to Simplified Chinese 臺灣正體到簡體
                                 s2hk Simplified Chinese to Traditional Chinese (Hong Kong variant) 簡體到香港繁體
                                 hk2s Traditional Chinese (Hong Kong variant) to Simplified Chinese 香港繁體到簡體
                                 s2twp Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom 簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
                                 tw2sp Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom 繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙
                                 t2tw Traditional Chinese (OpenCC Standard) to Taiwan Standard 繁體（OpenCC 標準）到臺灣正體
                                 hk2t Traditional Chinese (Hong Kong variant) to Traditional Chinese 香港繁體到繁體（OpenCC 標準）
                                 t2hk Traditional Chinese (OpenCC Standard) to Hong Kong variant 繁體（OpenCC 標準）到香港繁體
                                 t2jp Traditional Chinese Characters (Kyūjitai) to New Japanese Kanji (Shinjitai) 繁體（OpenCC 標準，舊字體）到日文新字體
                                 jp2t New Japanese Kanji (Shinjitai) to Traditional Chinese Characters (Kyūjitai) 日文新字體到繁體（OpenCC 標準，舊字體）
                                 tw2t Traditional Chinese (Taiwan standard) to Traditional Chinese 臺灣正體到繁體（OpenCC 標準）
   --help, -h                    show help (default: false)
   --input value, -i value       Input string or file
   --output value, -o value      Output file
```


