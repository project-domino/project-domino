export default class SaveManager {
	constructor(parent) {
		this.parent = parent;
		this.saving = false;

		this.parent.on("save", document => this.save(document));
	}
	save(document) {
		if(this.saving) return;
		this.saving = true;
		this.saveAsync(document, () => {
			this.saving = false;
			this.parent.emit("saved");
		});
	}
	saveAsync(document, callback) {
		this.saveSync(document);
		window.setTimeout(callback, 0);
	}
	saveSync(document) {
		console.log("TODO Saving");
		console.log("document =", JSON.stringify(document));
	}
}
