Dim objShell

Set objShell = CreateObject("Wscript.Shell")
objShell.AppActivate("hahajh-robot")
objShell.SendKeys("(^+K)")

While objShell.AppActivate("Push Commits")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 100
objShell.SendKeys("{TAB 4}")
Wscript.Sleep 100
objShell.SendKeys("{ENTER}")
'objShell.SendKeys("%{F4}")