package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/razonyang/opencc"
	"github.com/urfave/cli/v2"
)

const conversionUsage = `
	s2t Simplified Chinese to Traditional Chinese 簡體到繁體
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
	tw2t Traditional Chinese (Taiwan standard) to Traditional Chinese 臺灣正體到繁體（OpenCC 標準）`

func New() *cli.App {
	return &cli.App{
		Name:  "gopencc",
		Usage: "Open Chinese Convert Conversions between Traditional Chinese, Simplified Chinese",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conversion",
				Aliases:  []string{"c"},
				Usage:    conversionUsage,
				Required: true,
			},
			&cli.StringFlag{
				Name:     "input",
				Usage:    "Input string or file",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output file",
				Aliases: []string{"o"},
			},
		},
		Action: func(c *cli.Context) error {
			return convert(c.App, c.String("conversion"), c.String("input"), c.String("output"))
		},
	}
}

func convert(app *cli.App, conversion, input, out string) error {
	converter, err := getConverter(app, conversion)
	if err != nil {
		return err
	}

	if _, err := os.Stat(input); err == nil {
		inputBytes, err := os.ReadFile(input)
		if err != nil {
			return err
		}
		input = string(inputBytes)
	}

	data, err := converter.Convert(input)
	if err != nil {
		return err
	}
	if out == "" {
		fmt.Println(data)
		return nil
	}

	// write output to file.
	return saveFile(data, out)
}

func getConverter(app *cli.App, conversion string) (*opencc.OpenCC, error) {
	if _, ok := app.Metadata["converter"]; !ok {
		converter, err := opencc.New(conversion)
		if err != nil {
			return nil, err
		}
		app.Metadata["converter"] = converter
	}

	return app.Metadata["converter"].(*opencc.OpenCC), nil
}

func saveFile(data, output string) error {
	dir := path.Dir(output)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}

	return os.WriteFile(output, []byte(data), 0600)
}
