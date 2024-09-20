// Generated automatically.  DO NOT HAND-EDIT.

package rxvt

import "github.com/ales999/tcell/v2/terminfo"

func init() {

	// rxvt terminal emulator (X Window System)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "rxvt",
		Aliases:      []string{"rxvt-color"},
		Columns:      80,
		Lines:        24,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[2J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b=",
		ExitKeypad:   "\x1b>",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:    "\x1b[39;49m",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b(B\x1b)0",
		Mouse:        "\x1b[M",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\x7f",
		KeyHome:      "\x1b[7~",
		KeyEnd:       "\x1b[8~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[11~",
		KeyF2:        "\x1b[12~",
		KeyF3:        "\x1b[13~",
		KeyF4:        "\x1b[14~",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyF21:       "\x1b[23$",
		KeyF22:       "\x1b[24$",
		KeyF23:       "\x1b[11^",
		KeyF24:       "\x1b[12^",
		KeyF25:       "\x1b[13^",
		KeyF26:       "\x1b[14^",
		KeyF27:       "\x1b[15^",
		KeyF28:       "\x1b[17^",
		KeyF29:       "\x1b[18^",
		KeyF30:       "\x1b[19^",
		KeyF31:       "\x1b[20^",
		KeyF32:       "\x1b[21^",
		KeyF33:       "\x1b[23^",
		KeyF34:       "\x1b[24^",
		KeyF35:       "\x1b[25^",
		KeyF36:       "\x1b[26^",
		KeyF37:       "\x1b[28^",
		KeyF38:       "\x1b[29^",
		KeyF39:       "\x1b[31^",
		KeyF40:       "\x1b[32^",
		KeyF41:       "\x1b[33^",
		KeyF42:       "\x1b[34^",
		KeyF43:       "\x1b[23@",
		KeyF44:       "\x1b[24@",
		KeyBacktab:   "\x1b[Z",
		KeyShfLeft:   "\x1b[d",
		KeyShfRight:  "\x1b[c",
		KeyShfUp:     "\x1b[a",
		KeyShfDown:   "\x1b[b",
		KeyShfHome:   "\x1b[7$",
		KeyShfEnd:    "\x1b[8$",
		KeyShfInsert: "\x1b[2$",
		KeyShfDelete: "\x1b[3$",
		KeyCtrlUp:    "\x1b[Oa",
		KeyCtrlDown:  "\x1b[Ob",
		KeyCtrlRight: "\x1b[Oc",
		KeyCtrlLeft:  "\x1b[Od",
		KeyCtrlHome:  "\x1b[7^",
		KeyCtrlEnd:   "\x1b[8^",
		AutoMargin:   true,
		XTermLike:    true,
	})

	// rxvt 2.7.9 with xterm 256-colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "rxvt-256color",
		Columns:      80,
		Lines:        24,
		Colors:       256,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[2J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b=",
		ExitKeypad:   "\x1b>",
		SetFg:        "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:        "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:      "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		ResetFgBg:    "\x1b[39;49m",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b(B\x1b)0",
		Mouse:        "\x1b[M",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\x7f",
		KeyHome:      "\x1b[7~",
		KeyEnd:       "\x1b[8~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[11~",
		KeyF2:        "\x1b[12~",
		KeyF3:        "\x1b[13~",
		KeyF4:        "\x1b[14~",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyF21:       "\x1b[23$",
		KeyF22:       "\x1b[24$",
		KeyF23:       "\x1b[11^",
		KeyF24:       "\x1b[12^",
		KeyF25:       "\x1b[13^",
		KeyF26:       "\x1b[14^",
		KeyF27:       "\x1b[15^",
		KeyF28:       "\x1b[17^",
		KeyF29:       "\x1b[18^",
		KeyF30:       "\x1b[19^",
		KeyF31:       "\x1b[20^",
		KeyF32:       "\x1b[21^",
		KeyF33:       "\x1b[23^",
		KeyF34:       "\x1b[24^",
		KeyF35:       "\x1b[25^",
		KeyF36:       "\x1b[26^",
		KeyF37:       "\x1b[28^",
		KeyF38:       "\x1b[29^",
		KeyF39:       "\x1b[31^",
		KeyF40:       "\x1b[32^",
		KeyF41:       "\x1b[33^",
		KeyF42:       "\x1b[34^",
		KeyF43:       "\x1b[23@",
		KeyF44:       "\x1b[24@",
		KeyBacktab:   "\x1b[Z",
		KeyShfLeft:   "\x1b[d",
		KeyShfRight:  "\x1b[c",
		KeyShfUp:     "\x1b[a",
		KeyShfDown:   "\x1b[b",
		KeyShfHome:   "\x1b[7$",
		KeyShfEnd:    "\x1b[8$",
		KeyShfInsert: "\x1b[2$",
		KeyShfDelete: "\x1b[3$",
		KeyCtrlUp:    "\x1b[Oa",
		KeyCtrlDown:  "\x1b[Ob",
		KeyCtrlRight: "\x1b[Oc",
		KeyCtrlLeft:  "\x1b[Od",
		KeyCtrlHome:  "\x1b[7^",
		KeyCtrlEnd:   "\x1b[8^",
		AutoMargin:   true,
		XTermLike:    true,
	})

	// rxvt 2.7.9 with xterm 88-colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "rxvt-88color",
		Columns:      80,
		Lines:        24,
		Colors:       88,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[2J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		EnterKeypad:  "\x1b=",
		ExitKeypad:   "\x1b>",
		SetFg:        "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:        "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:      "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		ResetFgBg:    "\x1b[39;49m",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b(B\x1b)0",
		Mouse:        "\x1b[M",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\x7f",
		KeyHome:      "\x1b[7~",
		KeyEnd:       "\x1b[8~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[11~",
		KeyF2:        "\x1b[12~",
		KeyF3:        "\x1b[13~",
		KeyF4:        "\x1b[14~",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyF21:       "\x1b[23$",
		KeyF22:       "\x1b[24$",
		KeyF23:       "\x1b[11^",
		KeyF24:       "\x1b[12^",
		KeyF25:       "\x1b[13^",
		KeyF26:       "\x1b[14^",
		KeyF27:       "\x1b[15^",
		KeyF28:       "\x1b[17^",
		KeyF29:       "\x1b[18^",
		KeyF30:       "\x1b[19^",
		KeyF31:       "\x1b[20^",
		KeyF32:       "\x1b[21^",
		KeyF33:       "\x1b[23^",
		KeyF34:       "\x1b[24^",
		KeyF35:       "\x1b[25^",
		KeyF36:       "\x1b[26^",
		KeyF37:       "\x1b[28^",
		KeyF38:       "\x1b[29^",
		KeyF39:       "\x1b[31^",
		KeyF40:       "\x1b[32^",
		KeyF41:       "\x1b[33^",
		KeyF42:       "\x1b[34^",
		KeyF43:       "\x1b[23@",
		KeyF44:       "\x1b[24@",
		KeyBacktab:   "\x1b[Z",
		KeyShfLeft:   "\x1b[d",
		KeyShfRight:  "\x1b[c",
		KeyShfUp:     "\x1b[a",
		KeyShfDown:   "\x1b[b",
		KeyShfHome:   "\x1b[7$",
		KeyShfEnd:    "\x1b[8$",
		KeyShfInsert: "\x1b[2$",
		KeyShfDelete: "\x1b[3$",
		KeyCtrlUp:    "\x1b[Oa",
		KeyCtrlDown:  "\x1b[Ob",
		KeyCtrlRight: "\x1b[Oc",
		KeyCtrlLeft:  "\x1b[Od",
		KeyCtrlHome:  "\x1b[7^",
		KeyCtrlEnd:   "\x1b[8^",
		AutoMargin:   true,
		XTermLike:    true,
	})

	// rxvt-unicode terminal (X Window System)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:              "rxvt-unicode",
		Columns:           80,
		Lines:             24,
		Colors:            88,
		Bell:              "\a",
		Clear:             "\x1b[H\x1b[2J",
		EnterCA:           "\x1b[?1049h",
		ExitCA:            "\x1b[r\x1b[?1049l",
		ShowCursor:        "\x1b[?12l\x1b[?25h",
		HideCursor:        "\x1b[?25l",
		AttrOff:           "\x1b[m\x1b(B",
		Underline:         "\x1b[4m",
		Bold:              "\x1b[1m",
		Italic:            "\x1b[3m",
		Blink:             "\x1b[5m",
		Reverse:           "\x1b[7m",
		EnterKeypad:       "\x1b=",
		ExitKeypad:        "\x1b>",
		SetFg:             "\x1b[38;5;%p1%dm",
		SetBg:             "\x1b[48;5;%p1%dm",
		SetFgBg:           "\x1b[38;5;%p1%d;48;5;%p2%dm",
		ResetFgBg:         "\x1b[39;49m",
		AltChars:          "+C,D-A.B0E``aaffgghFiGjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:          "\x1b(0",
		ExitAcs:           "\x1b(B",
		EnableAutoMargin:  "\x1b[?7h",
		DisableAutoMargin: "\x1b[?7l",
		Mouse:             "\x1b[M",
		SetCursor:         "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:       "\b",
		CursorUp1:         "\x1b[A",
		KeyUp:             "\x1b[A",
		KeyDown:           "\x1b[B",
		KeyRight:          "\x1b[C",
		KeyLeft:           "\x1b[D",
		KeyInsert:         "\x1b[2~",
		KeyDelete:         "\x1b[3~",
		KeyBackspace:      "\x7f",
		KeyHome:           "\x1b[7~",
		KeyEnd:            "\x1b[8~",
		KeyPgUp:           "\x1b[5~",
		KeyPgDn:           "\x1b[6~",
		KeyF1:             "\x1b[11~",
		KeyF2:             "\x1b[12~",
		KeyF3:             "\x1b[13~",
		KeyF4:             "\x1b[14~",
		KeyF5:             "\x1b[15~",
		KeyF6:             "\x1b[17~",
		KeyF7:             "\x1b[18~",
		KeyF8:             "\x1b[19~",
		KeyF9:             "\x1b[20~",
		KeyF10:            "\x1b[21~",
		KeyF11:            "\x1b[23~",
		KeyF12:            "\x1b[24~",
		KeyF13:            "\x1b[25~",
		KeyF14:            "\x1b[26~",
		KeyF15:            "\x1b[28~",
		KeyF16:            "\x1b[29~",
		KeyF17:            "\x1b[31~",
		KeyF18:            "\x1b[32~",
		KeyF19:            "\x1b[33~",
		KeyF20:            "\x1b[34~",
		KeyBacktab:        "\x1b[Z",
		KeyShfLeft:        "\x1b[d",
		KeyShfRight:       "\x1b[c",
		KeyShfUp:          "\x1b[a",
		KeyShfDown:        "\x1b[b",
		KeyShfHome:        "\x1b[7$",
		KeyShfEnd:         "\x1b[8$",
		KeyShfInsert:      "\x1b[2$",
		KeyShfDelete:      "\x1b[3$",
		KeyCtrlUp:         "\x1b[Oa",
		KeyCtrlDown:       "\x1b[Ob",
		KeyCtrlRight:      "\x1b[Oc",
		KeyCtrlLeft:       "\x1b[Od",
		KeyCtrlHome:       "\x1b[7^",
		KeyCtrlEnd:        "\x1b[8^",
		AutoMargin:        true,
		InsertChar:        "\x1b[@",
	})

	// rxvt-unicode terminal with 256 colors (X Window System)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:              "rxvt-unicode-256color",
		Columns:           80,
		Lines:             24,
		Colors:            256,
		Bell:              "\a",
		Clear:             "\x1b[H\x1b[2J",
		EnterCA:           "\x1b[?1049h",
		ExitCA:            "\x1b[r\x1b[?1049l",
		ShowCursor:        "\x1b[?12l\x1b[?25h",
		HideCursor:        "\x1b[?25l",
		AttrOff:           "\x1b[m\x1b(B",
		Underline:         "\x1b[4m",
		Bold:              "\x1b[1m",
		Italic:            "\x1b[3m",
		Blink:             "\x1b[5m",
		Reverse:           "\x1b[7m",
		EnterKeypad:       "\x1b=",
		ExitKeypad:        "\x1b>",
		SetFg:             "\x1b[38;5;%p1%dm",
		SetBg:             "\x1b[48;5;%p1%dm",
		SetFgBg:           "\x1b[38;5;%p1%d;48;5;%p2%dm",
		ResetFgBg:         "\x1b[39;49m",
		AltChars:          "+C,D-A.B0E``aaffgghFiGjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:          "\x1b(0",
		ExitAcs:           "\x1b(B",
		EnableAutoMargin:  "\x1b[?7h",
		DisableAutoMargin: "\x1b[?7l",
		Mouse:             "\x1b[M",
		SetCursor:         "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:       "\b",
		CursorUp1:         "\x1b[A",
		KeyUp:             "\x1b[A",
		KeyDown:           "\x1b[B",
		KeyRight:          "\x1b[C",
		KeyLeft:           "\x1b[D",
		KeyInsert:         "\x1b[2~",
		KeyDelete:         "\x1b[3~",
		KeyBackspace:      "\x7f",
		KeyHome:           "\x1b[7~",
		KeyEnd:            "\x1b[8~",
		KeyPgUp:           "\x1b[5~",
		KeyPgDn:           "\x1b[6~",
		KeyF1:             "\x1b[11~",
		KeyF2:             "\x1b[12~",
		KeyF3:             "\x1b[13~",
		KeyF4:             "\x1b[14~",
		KeyF5:             "\x1b[15~",
		KeyF6:             "\x1b[17~",
		KeyF7:             "\x1b[18~",
		KeyF8:             "\x1b[19~",
		KeyF9:             "\x1b[20~",
		KeyF10:            "\x1b[21~",
		KeyF11:            "\x1b[23~",
		KeyF12:            "\x1b[24~",
		KeyF13:            "\x1b[25~",
		KeyF14:            "\x1b[26~",
		KeyF15:            "\x1b[28~",
		KeyF16:            "\x1b[29~",
		KeyF17:            "\x1b[31~",
		KeyF18:            "\x1b[32~",
		KeyF19:            "\x1b[33~",
		KeyF20:            "\x1b[34~",
		KeyBacktab:        "\x1b[Z",
		KeyShfLeft:        "\x1b[d",
		KeyShfRight:       "\x1b[c",
		KeyShfUp:          "\x1b[a",
		KeyShfDown:        "\x1b[b",
		KeyShfHome:        "\x1b[7$",
		KeyShfEnd:         "\x1b[8$",
		KeyShfInsert:      "\x1b[2$",
		KeyShfDelete:      "\x1b[3$",
		KeyCtrlUp:         "\x1b[Oa",
		KeyCtrlDown:       "\x1b[Ob",
		KeyCtrlRight:      "\x1b[Oc",
		KeyCtrlLeft:       "\x1b[Od",
		KeyCtrlHome:       "\x1b[7^",
		KeyCtrlEnd:        "\x1b[8^",
		AutoMargin:        true,
		InsertChar:        "\x1b[@",
	})
}
