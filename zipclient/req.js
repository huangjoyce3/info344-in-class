//document.getElementById("firstname").value = "Joyce";
document.getElementById("form")
	.addEventListener("submit", function(evt) {
		evt.preventDefault();
		var name = document.getElementById("firstname").value;

		fetch("http://localhost:4000/hello?name=" + name)
	    	.then(function(response) {

	        	return response.text();
	    })
	    .then(function(data) {
	    	console.log(data);
	    	document.getElementById("response").textContent = data;
	    })
	})
