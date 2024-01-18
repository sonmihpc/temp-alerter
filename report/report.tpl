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
					Temperature 1
				</td>
				<td>
					{{.Temp1}} °C
				</td>
			</tr>
			<tr>
            	<td>
            		Temperature 2
            	</td>
            	<td>
            		{{.Temp2}} °C
            	</td>
            </tr>
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