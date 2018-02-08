Dim objShell

Set objShell = CreateObject("Wscript.Shell")
While objShell.AppActivate("hahajh-robot")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 100
objShell.SendKeys("(^k)")

While objShell.AppActivate("Commit Changes")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 100
objShell.SendKeys("(%i)")
Wscript.Sleep 100
objShell.SendKeys("%{UP 2}")
Wscript.Sleep 100
objShell.SendKeys("{ENTER}")
Wscript.Sleep 100
objShell.SendKeys("(%p)")
Wscript.Sleep 100
objShell.SendKeys("{ENTER}")
'objShell.SendKeys("%{F4}")