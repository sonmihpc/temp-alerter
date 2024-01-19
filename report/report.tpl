{{define "report"}}
<html>
	<head>
		<style type="text/css">
			body {
				font-size: 10pt;
			}
			th, td {
				padding: 5px;
				text-align: left;
			}
		</style>
	</head>
	<body>
		<p>Warning: The temperatures has exceeded the threshold:</p>
		<table>
		    <tr>
		        <td>
		            Position
		        <td>
		            {{.Position}}
		        <td>
		        <td>
		    </tr>
		    {{range $index, $value := .Temps}}
			<tr>
				<td>
					Temperature {{$index}}
				</td>
				<td>
					{{$value}} °C
				</td>
			</tr>
		    {{end}}
		</table>
		</center>
		<p>Regards,</p>

		<p>Admin</p>

		<p>
			<i>Note: This is an automated e-mail.</i>
		</p>
	</body>
</html>
{{- end}}