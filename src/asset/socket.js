var poll = function() {
	$.ajax({
  		url: "/polling?time=" + new Date().getTime(),
	}).done(function(data) {
  		console.log(data);
  		poll();
	});
};
poll();