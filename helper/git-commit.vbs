Dim objShell

Set objShell = CreateObject("Wscript.Shell")
While objShell.AppActivate("hahajh-robot")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 100
objShell.SendKeys("(^K)")

While objShell.AppActivate("Commit Changes")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 100
objShell.SendKeys("{DOWN}")
Wscript.Sleep 100
objShell.SendKeys("{TAB 6}")
Wscript.Sleep 100
objShell.SendKeys("{ENTER}")
'objShell.SendKeys("%{F4}")