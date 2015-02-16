var poll = function() {
	$.ajax({
  		url:  '/polling',
  		data: {
  			token: token
  		},
  		type: 'POST',
  		timeout: 100*1000
	}).done(function(data) {
  		console.log(data);
  		poll();
	});
};
poll();