document.getElementById("form")
	.addEventListener("submit", function(evt) {
		evt.preventDefault();
		var name = document.getElementById("inputBox").value;

		fetch("http://localhost:4000/zips/" + name)
	    	.then(function(response) {
	        	return response.text();
	    })
	    .then(function(data) {
	    	console.log(data);
	    	document.getElementById("response").textContent = data;
	    })
	})