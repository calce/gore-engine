function() {
	
	var engineCode = "RE2"
	window.gore.engines[engineCode] = {

		function process(res) {
			// Ignore if response is not for the latest request
			if (res.requestId != window.gore.lastRequest.id) return;
			// otherwise, update the results
		}
	}

}()