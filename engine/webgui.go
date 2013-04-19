package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
	"code.google.com/p/mx3/util"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var (
	ui      = &guistate{Steps: 1000, Runtime: 1e-9, Cond: sync.NewCond(new(sync.Mutex))}
	uitempl = template.Must(template.New("gui").Parse(templText))
)

func gui(w http.ResponseWriter, r *http.Request) {
	ui.Lock()
	defer ui.Unlock()
	util.FatalErr(uitempl.Execute(w, ui))
}

type guistate struct {
	Msg                 string
	Steps               int
	Runtime             float64
	running, pleaseStop bool // todo: mv out of struct
	*sync.Cond               // todo: mv out of struct
}

func (s *guistate) Time() float32    { return float32(Time) }
func (s *guistate) Lock()            { ui.L.Lock() }
func (s *guistate) Unlock()          { ui.L.Unlock() }
func (s *guistate) ImWidth() int     { return ui.Mesh().Size()[2] }
func (s *guistate) ImHeight() int    { return ui.Mesh().Size()[1] }
func (s *guistate) Mesh() *data.Mesh { return &mesh }
func (s *guistate) Uname() string    { return Uname }
func (s *guistate) Version() string  { return VERSION }
func (s *guistate) Pwd() string      { pwd, _ := os.Getwd(); return pwd }
func (s *guistate) Solver() *cuda.Heun {
	if Solver == nil {
		return &zeroSolver
	} else {
		return Solver
	}
}

// surrogate solver if no real one is set, provides zero values for time step etc to template.
var zeroSolver cuda.Heun

const templText = `
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>mx3</title>
	<style media="screen" type="text/css">
		body { margin: 40px; font-family: Helvetica, Arial, sans-serif; font-size: 15px; }
		img  { margin: 10px; }
		h1   { font-size: 28px; font-color: gray; }
		h2   { font-size: 20px; }
		hr   { border-style: none; border-top: 1px solid gray; }
		a    { color: #375EAB; text-decoration: none; }
		table{ border:"20"; }
		input#text{ border:solid; border-color:grey; border-width:1px; padding-left:4px;}
		div#header{ color:gray; font-size:16px; }
		div#footer{ color:gray; font-size:14px; }
	</style>
</head>

<body>


<div id="header"> <h1> {{.Version}} </h1> <hr/> </div>




<div> <h2> solver </h2>

<table><tr><td>  

	<form action=/ctl/run method="POST">
        <input id=text size=8 name="value" value="{{.Runtime}}"> s <input type="submit" value="Run"/>
	</form>
	<form  action=/ctl/steps method="POST">
        <input id=text size=8 name="value" value="{{.Steps}}"> <input type="submit" value="Steps"/>
	</form>

	<form id=text action=/ctl/pause method="POST"> 
		<input type="submit" value="Pause"/>
	</form>

	<br/>

</td><td>  
 &nbsp; &nbsp; &nbsp;
</td><td>  

	<span id="running"><font color=red><b>Paused</b></font></span> 
	<span id="dash"> </span>

</td></tr></table>


<script>
	function httpGet(url){
    	var xmlHttp = new XMLHttpRequest();
    	xmlHttp.open("GET", url, false);
    	xmlHttp.send(null);
    	return xmlHttp.responseText;
    }
	var running = false
	function updateRunning(){
		running = (httpGet("/running/") === "true")
		if(running){
			document.getElementById("running").innerHTML = "<font color=green><b>Running</b></font>"
		}else{
			document.getElementById("running").innerHTML = "<font color=red><b>Paused</b></font>"
		}
	}
	setInterval(updateRunning, 200)
</script>

<script>
	function updateDash(){
		document.getElementById("dash").innerHTML = httpGet("/dash/")
	}
	function updateDashIfRunning(){
		if(running){
			updateDash();
		}
	}
	updateDash();
	setInterval(updateDashIfRunning, 200);
</script>

<hr/> </div>




<div> <h2> magnetization </h2> 
<img id="magnetization" src="/render/m" width={{.ImWidth}} height={{.ImHeight}} alt="m"/>

<form  action=/setm/ method="POST">
	<b>From file:</b> <input id="text" size=60 name="value" value="{{.Pwd}}"> <input type="submit" value="Submit"/>
</form>

<script>
	var img = new Image();
	img.src = "/render/m";
	function updateImg(){
		if(running && img.complete){
			document.getElementById("magnetization").src = img.src;
			img = new Image();
			img.src = "/render/m?" + new Date();
		}
	}
	setInterval(updateImg, 500);
</script>

<hr/></div>




<div> <h2> parameters </h2> 
	<form action=/setparam/ method="POST">
	<table>
	{{range $k, $v := .Params}}
		<tr><td> {{$k}}: </td><td> 
		{{range $v.Comp}} 
        	<input id=text size=8 name="{{$k}}{{.}}" value="{{$v.GetComp .}}"> 
		{{end}} {{$v.Unit}} <font color=grey>&nbsp;({{$v.Descr}})</font> </td></tr>
	{{end}}
	</table>
	<input type="submit" value="Submit"/>
	</form>
<hr/></div>




<div><h2> mesh </h2> 
<form action=/setmesh/ method="POST"><table> 
	<tr>
		<td> grid size: </td>
		<td> <input id=text size=8 name="gridsizex" value="{{index .Mesh.Size 2}}"> </td> <td> x </td>
		<td> <input id=text size=8 name="gridsizey" value="{{index .Mesh.Size 1}}"> </td> <td> x </td>
		<td> <input id=text size=8 name="gridsizez" value="{{index .Mesh.Size 0}}"> </td> <td>   </td>
	</tr>

	<tr>
		<td> cell size: </td>
		<td> <input id=text size=8 name="cellsizex" value="{{index .Mesh.CellSize 2}}"> </td> <td> x  </td>
		<td> <input id=text size=8 name="cellsizey" value="{{index .Mesh.CellSize 1}}"> </td> <td> x  </td>
		<td> <input id=text size=8 name="cellsizez" value="{{index .Mesh.CellSize 0}}"> </td> <td> m3 </td>
	</tr>

	<tr>
		<td> world size: &nbsp;&nbsp; </td>
		<td> {{index .Mesh.WorldSize 2}} </td> <td> x  </td>
		<td> {{index .Mesh.WorldSize 1}} </td> <td> x  </td>
		<td> {{index .Mesh.WorldSize 0}} </td> <td> m3 </td>
	</tr>
</table>
	<input type="submit" value=" Submit"/> <b> Changing the mesh requires some re-initialization time</b>
</form>

<hr/></div>




<font color=red> <div> <h2> Danger Zone </h2></font>
	<form action=/ctl/kill  method="POST"> <b> Kill process:</b> <input type="submit" value="Kill"/> </form>
<hr/></div>

<div id="footer">
<center>
{{.Uname}}
</center>
</div>

</body>
</html>
`
