fetch("/list")
	.then((response) => {
		return response.json();
	})
	.then((data) => {
		console.log(JSON.stringify(data));
	});
fetch("/search?name=bash")
	.then((response) => {
		return response.json();
	})
	.then((data) => {
		console.log(JSON.stringify(data));
	});

fetch("/show?name=ht")
	.then((response) => {
		return response.json();
	})
	.then((data) => {
		console.log(JSON.stringify(data));
	});
