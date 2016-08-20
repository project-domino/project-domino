import getModal from "./modal.js";

const modal = getModal();

const errorHandler = err => {
	console.log(err);

	var responseJSON;
	if(err.responseJSON)
		responseJSON = err.responseJSON;
	else if(err.responseText)
		responseJSON = JSON.parse(err.responseText);

	if(err.readyState === 0)
		modal.alert("Could not connect to server");
	else
		modal.alert(responseJSON.Errors[0]);
};

export {errorHandler};
