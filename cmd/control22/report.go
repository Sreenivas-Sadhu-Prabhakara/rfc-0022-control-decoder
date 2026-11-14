package main

import "html/template"

var page = template.Must(template.New("report").Parse(`<!doctype html>
<html lang="en"><head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>RFC 22 · Binary control messages</title>
<style>
:root { background: #eef4f3; color: #132e34; font: 17px/1.55 system-ui, sans-serif; }
main { max-width: 880px; margin: auto; padding: 38px 22px; }
h1 { font-size: clamp(2.1rem, 6vw, 3.6rem); line-height: 1.08; margin: .25em 0; }
p { max-width: 74ch; }
code { overflow-wrap: anywhere; }
a { color: #115e67; }
.table { overflow-x: auto; background: white; border: 1px solid #a5c0bb; border-radius: 10px; }
table { text-align: left; width: 100%; border-collapse: collapse; }
th, td { padding: 8px 14px; border-top: 1px solid #d5e4e1; }
thead { background: #dbece8; }
caption { padding: 10px 14px; text-align: left; font-weight: 700; }
</style></head>
<body><main>
<h1>RFC 22 · Binary control messages</h1>
<p>Each opcode determines how many link bytes follow it. The table shows the commands in the supplied teaching payload.</p>
<p><strong>Payload ({{.Bytes}} bytes)</strong><br><code>{{.Hex}}</code></p>
<div class="table" role="region" aria-label="Decoded commands" tabindex="0">
<table><caption>{{len .Commands}} decoded commands</caption>
<thead><tr><th scope="col">Byte offset</th><th scope="col">Opcode</th><th scope="col">Command</th><th scope="col">Link bytes</th></tr></thead>
<tbody>{{range .Commands}}
<tr><th scope="row">{{.Offset}}</th><td><code>{{printf "%02X" .Opcode}}</code></td><td>{{.Name}}</td>
<td><code>{{range $i, $link := .Links}}{{if $i}} {{end}}{{printf "%02X" $link}}{{end}}</code></td></tr>
{{end}}</tbody></table>
</div>
<p>RFC 22 observed that CLOSING could arrive before the last data message. This decoder displays the directive and its link byte.</p>
<p>GitHub Pages serves this saved output from the Go CLI. <a href="https://www.rfc-editor.org/info/rfc22/">Official RFC 22</a> · <a href="https://github.com/Sreenivas-Sadhu-Prabhakara/rfc-0022-control-decoder">Source and run instructions</a></p>
</main></body></html>
`))
