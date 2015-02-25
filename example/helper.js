var getParamByName = function(name) {
	var pattern =  new RegExp(name + '=([A-Za-z0-9]+)')
	var arr = document.URL.match(pattern);
	return arr[1];
}