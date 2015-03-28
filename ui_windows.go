// +build !ui

package main

import (
	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
)

type MainWindow struct {
	*gform.Form
}

func NewWindow() *MainWindow {

	gform.Init()
	return &MainWindow{gform.NewForm(nil)}
}

func ShowGUI() (err error) {

	mw := NewWindow()
	mw.Center()
	mw.SetSize(440, 100)
	mw.EnableSizable(false)
	mw.EnableMaxButton(false)
	mw.SetCaption("Empires Launcher")
	mw.Bind(w32.WM_CLOSE, mw.btnClose_onclick)
	mw.Show()

	btnVanilla := gform.NewPushButton(mw)
	btnVanilla.SetPos(10, 15)
	btnVanilla.SetSize(125, 25)
	btnVanilla.SetCaption("Empires Vanilla")
	btnVanilla.OnLBUp().Bind(mw.btnVanilla_onclick)

	btnCommunity := gform.NewPushButton(mw)
	btnCommunity.SetPos(145, 15)
	btnCommunity.SetSize(125, 25)
	btnCommunity.SetCaption("Community Scripts")
	btnCommunity.OnLBUp().Bind(mw.btnCommunity_onclick)

	btnClose := gform.NewPushButton(mw)
	btnClose.SetPos(280, 15)
	btnClose.SetSize(125, 25)
	btnClose.SetCaption("Close")
	btnClose.OnLBUp().Bind(mw.btnClose_onclick)

	gform.RunMainLoop()

	return
}

func (mw *MainWindow) btnVanilla_onclick(arg *gform.EventArg) {

	if err := ApplyAndLaunch("EmpiresVanilla"); err != nil {

		gform.MsgBox(arg.Sender().Parent(), "Fatal Error", err.Error(), w32.MB_OK|w32.MB_ICONERROR)
		return
	}

	mw.Form.ControlBase.Close()
	gform.Exit()
}

func (mw *MainWindow) btnCommunity_onclick(arg *gform.EventArg) {

	if err := ApplyAndLaunch("CommunityScripts"); err != nil {

		gform.MsgBox(arg.Sender().Parent(), "Fatal Error", err.Error(), w32.MB_OK|w32.MB_ICONERROR)
		return
	}

	mw.Form.ControlBase.Close()
	gform.Exit()
}

func (mw *MainWindow) btnClose_onclick(arg *gform.EventArg) {

	mw.Form.ControlBase.Close()
	gform.Exit()
}
