package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/sqweek/dialog"
)

var params = Params{
	DialogType: DialogTypeFile,
}

const (
	DialogTypeFile      = "file"
	DialogTypeDirectory = "dir"
	DialogTypeMessage   = "msg"
)

type Params struct {
	OutputType string   `json:"output"`
	DialogType string   `json:"type"`
	StartDir   string   `json:"start_dir"`
	StartFile  string   `json:"start_file"`
	Title      string   `json:"title"`
	FilterDesc string   `json:"filter_desc"`
	Filters    []string `json:"filters"`

	// for file
	FileDialogType string `json:"file_dialog_type"`

	// for message
	MessageDialogType string `json:"message_dialog_type"`
	Msg               string `json:"msg"`
}

func init() {
	pflag.StringVarP(&params.OutputType, "output_type", "o", "json", "json,text")
	pflag.StringVarP(&params.DialogType, "type", "t", DialogTypeFile, "file,dir,msg")
	pflag.StringVarP(&params.StartDir, "start_dir", "d", "", "start directory")
	pflag.StringVarP(&params.StartFile, "start_file", "", "", "start file")
	pflag.StringVarP(&params.Title, "title", "", "", "title")
	pflag.StringVarP(&params.FilterDesc, "filter_desc", "", "", "file filter description. example: pictures")
	pflag.StringSliceVar(&params.Filters, "filter", []string{}, "file filter. example: jpg,png,gif")
	pflag.StringVarP(&params.FileDialogType, "file_dialog_type", "", "load", "load,save")
	pflag.StringVarP(&params.MessageDialogType, "message_dialog_type", "", "info", "info,error,yes_or_no")
	pflag.StringVarP(&params.Msg, "msg", "m", "", "message info")
}

func main() {
	pflag.Parse()

	switch params.DialogType {
	case DialogTypeFile:
		v, err := dialogFile(params)
		showResult(params, v, err)
	case DialogTypeDirectory:
		v, err := dialogDirectory(params)
		showResult(params, v, err)
	case DialogTypeMessage:
		v, err := dialogMessage(params)
		showResult(params, v, err)
	}
}

func dialogFile(params Params) (string, error) {
	file := dialog.File()
	if params.StartDir != "" {
		file.SetStartDir(params.StartDir)
	}
	if params.StartFile != "" {
		file.SetStartFile(params.StartFile)
	}
	file.Title(params.Title)
	if len(params.Filters) > 0 {
		if params.FilterDesc == "" {
			params.FilterDesc = "files"
		}
		file.Filter(params.FilterDesc, params.Filters...)
	}
	if params.FileDialogType == "save" {
		return file.Save()
	}
	return file.Load()
}

func dialogDirectory(params Params) (string, error) {
	dir := dialog.Directory()
	if params.StartDir != "" {
		dir.SetStartDir(params.StartDir)
	}
	dir.Title(params.Title)
	return dir.Browse()
}

func dialogMessage(params Params) (interface{}, error) {
	msg := dialog.Message(params.Msg)
	msg.Title(params.Title)
	switch params.MessageDialogType {
	case "error":
		msg.Error()
	case "yes_or_no":
		return msg.YesNo(), nil
	default:
		msg.Info()
	}
	return "", nil
}

type Result struct {
	Value interface{} `json:"value"`
	Error string      `json:"error"`
}

func showResult(params Params, value interface{}, err error) {
	if params.OutputType == "json" {
		result := Result{
			Value: value,
		}
		if err != nil {
			result.Error = err.Error()
		}
		j, _ := json.Marshal(result)
		fmt.Println(string(j))
	} else {
		fmt.Println(value)
	}
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
