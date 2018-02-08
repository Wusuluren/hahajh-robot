Dim objShell

Set objShell = CreateObject("Wscript.Shell")
While objShell.AppActivate("hahajh-robot")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 1000
objShell.SendKeys("(^k)")

While objShell.AppActivate("Commit Changes")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 1000
objShell.SendKeys("{DOWN}")
Wscript.Sleep 1000
objShell.SendKeys("{TAB 6}")
Wscript.Sleep 1000
objShell.SendKeys("{ENTER}")
'objShell.SendKeys("%{F4}")

While objShell.AppActivate("hahajh-robot")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 1000
objShell.SendKeys("(^+K)")

While objShell.AppActivate("Push Commits")=False
	Wscript.Sleep 100
Wend
Wscript.Sleep 1000
objShell.SendKeys("{TAB 4}")
Wscript.Sleep 1000
objShell.SendKeys("{ENTER}")