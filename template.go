// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Some the original code has been edited by Walter Schulze

package main

import "text/template"

var head = template.Must(template.New("head").Parse(headText)) // HTML template
var style = template.Must(template.New("style").Parse(styleText))
var headFuncs = template.Must(template.New("headfuncs").Parse(headFuncsText))
var funcs = template.Must(template.New("funcs").Parse(funcsText))
var tailFuncs = template.Must(template.New("tailFuncs").Parse(tailFuncsText))
var table = template.Must(template.New("table").Parse(tableText))
var tail = template.Must(template.New("tail").Parse(tailText))

var headText = `<!doctype html>
<html>
<head>
<style>
pre, textarea {
	font-family: Monaco, 'Courier New', 'DejaVu Sans Mono', 'Bitstream Vera Sans Mono', monospace;
	font-size: 100%;
}
.hints {
	font-size: 0.8em;
	text-align: right;
}
`

var styleText = `
#{{.Name}}src, #{{.Name}}dst, #{{.Name}}err { width: 100%; text-align: left; }
#{{.Name}}src { height: 300px; }
#{{.Name}}dst { color: #00c; }
#{{.Name}}err { color: #c00; }
`

var headFuncsText = `
</style>
<script>

var xmlreq;

`

var funcsText = `
function insertTabs{{.Name}}(n) {
	// find the selection start and end
	var cont  = document.getElementById("{{.Name}}src");
	var start = cont.selectionStart;
	var end   = cont.selectionEnd;
	// split the textarea content into two, and insert n tabs
	var v = cont.value;
	var u = v.substr(0, start);
	for (var i=0; i<n; i++) {
		u += "\t";
	}
	u += v.substr(end);
	// set revised content
	cont.value = u;
	// reset caret position after inserted tabs
	cont.selectionStart = start+n;
	cont.selectionEnd = start+n;
}

function autoindent{{.Name}}(el) {
	var curpos = el.selectionStart;
	var tabs = 0;
	while (curpos > 0) {
		curpos--;
		if (el.value[curpos] == "\t") {
			tabs++;
		} else if (tabs > 0 || el.value[curpos] == "\n") {
			break;
		}
	}
	setTimeout(function() {
		insertTabs{{.Name}}(tabs);
	}, 1);
}

function preventDefault{{.Name}}(e) {
	if (e.preventDefault) {
		e.preventDefault();
	} else {
		e.cancelBubble = true;
	}
}

function keyHandler{{.Name}}(event) {
	var e = window.event || event;
	if (e.keyCode == 9) { // tab
		insertTabs{{.Name}}(1);
		preventDefault{{.Name}}(e);
		return false;
	}
	if (e.keyCode == 13) { // enter
		if (e.shiftKey) { // +shift
			compile{{.Name}}(e.target);
			preventDefault{{.Name}}(e);
			return false;
		} else {
			autoindent{{.Name}}(e.target);
		}
	}
	return true;
}

function auto{{.Name}}() {
	if(!document.getElementById("auto{{.Name}}").checked) {
		return;
	}
	compile{{.Name}}();
}

function compile{{.Name}}() {
	var prog = document.getElementById("{{.Name}}src").value;
	var req = new XMLHttpRequest();
	xmlreq = req;
	req.onreadystatechange = update{{.Name}};
	req.open("POST", "/func/{{.Name}}?path={{.Path}}", true);
	req.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
	req.send(prog);
}

function update{{.Name}}() {
		var req = xmlreq;
		if(!req || req.readyState != 4) {
			return;
		}
		if(req.status == 200) {
			document.getElementById("{{.Name}}dst").innerHTML = req.responseText;
			document.getElementById("{{.Name}}err").innerHTML = "";
		} else {
			document.getElementById("{{.Name}}err").innerHTML = req.responseText;
			document.getElementById("{{.Name}}dst").innerHTML = "";
		}
}`

var tailFuncsText = `
</script>
</head>
<body>
<a href="../share/{{.}}">share</a>
`

var tailText = `
</body>
</html>
`

var tableText = `
<h2>{{.Title}}</h2>
<table width="100%"><tr><td valign="top">
<textarea autofocus="true" id="{{.Name}}src" spellcheck="false" onkeydown="keyHandler{{.Name}}(event);" onkeyup="auto{{.Name}}();">{{.Content}}</textarea>
<div class="hints">
(Shift-Enter to compile and run.)&nbsp;&nbsp;&nbsp;&nbsp;
<input type="checkbox" id="auto{{.Name}}" value="checked" /> Compile and run after each keystroke
</div>
</tr>
<tr>
</td>
<td valign="top">
<div id="{{.Name}}dst"></div>
<div id="{{.Name}}err"></div>
</td>
</tr>
</table>
`
