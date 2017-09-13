package view

var INDEX = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Goud-Conifg</title>
	</head>
	<body>
		<h1>Goud-Conifg</h1>
		<h3>Parameters</h3>
		<form action="/update" method="post">
		{{range $k,$v:=.}}
			<h5>{{$k}}</h5>
			<input name="{{$k}}" value="{{$v}}"/><br/>
		{{end}}
			<input type="submit" value="更新"/>
		</form>
	</body>
	</html>`
